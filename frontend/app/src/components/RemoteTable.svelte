<script lang="typescript">
  import type {
    TableRow,
    TableButtonAction,
  } from "@app/components/models/RemoteTableProps";
  import { afterUpdate, onMount } from "svelte";
  import Fuse from "fuse.js";
  import { LoadState } from "@app/models/LoadState";
  import { randomString } from "@app/util/Generate";
  import Pagination from "@app/components/Pagination.svelte";
  import { dispatch, dispatchPrivate } from "@app/event/EventBus";
  import { fastClone, fastEquals } from "@app/util/Compare";
  import Dialog from "@app/components/layout/Dialog.svelte";
  import { isObject } from "@app/guards/Guard";
  import type { Dictionary } from "@app/models/Utility";
  import Button from "./Button.svelte";
  import { tableRows } from "@app/store/RemoteTableStore";
  export let getRows: () => Promise<TableRow[]>;
  export let defaultSortColumn = "";
  export let searchPlaceHolder: string = "Search";
  export let columnMeta: { [key: string]: { type: any } } = {};

  let id: string = "";
  let rows: TableRow[] = [];
  let filtered: TableRow[] = [];
  let columns: string[] = [];
  let filteredColumns: string[] = [];
  let query = "";
  let fuse: Fuse<{}>;
  let state = LoadState.Loading;
  let range: { min: number; max: number } = { min: 1, max: 1 };
  let widths: { [key: string]: number } = {};
  let canvasContext: any;
  let sort = "";
  let sortDirection = "";
  let allRowsSelected = false;
  let selectedCount = 0;
  let modal: "delete" | "toggle_column" | "preview" | "" | "filter" = "";
  let filters: Dictionary<any> = {
    onlyUnread: false,
  };
  let lastFilters: Dictionary<any> = {};
  let appliedFilters = 0;
  let showFilter: boolean;

  export let headerActions: TableButtonAction[] = [];
  export let onEdit: ((row: any) => any) | undefined = undefined;
  export let onDelete: ((rows: any[]) => any) | undefined = undefined;
  export let onRead:
    | ((rows: any[], value: boolean) => any)
    | undefined = undefined;
  export let hidden: Set<string> = new Set<string>();
  export let sortColumns:
    | ((columns: string[]) => string[])
    | undefined = undefined;
  export let onFormat: (column: string, row: any) => any = () => undefined;
  export let onRowClick: (row: any) => any = () => {};

  afterUpdate(() => {
    if (!fastEquals(filters, lastFilters)) {
      lastFilters = fastClone(filters);
      applyFilters();
    }
  });

  function applyFilters() {
    appliedFilters = 0;
    filtered = rows;
    if (filters.onlyUnread) {
      appliedFilters++;
      filtered = filtered.filter((f) => {
        return isUnread(f);
      });
    }
    tableRows.set(filtered);
  }

  function createFuse(): Fuse<{}> {
    const list = rows.map((r) => {
      const result: any = {};
      Object.keys(r).forEach((key) => {
        result[key] = isObject(r[key]) ? JSON.stringify(r[key]) : r[key];
      });
      return result;
    });
    return new Fuse(list, {
      keys: Object.keys(rows[0]),
    });
  }

  onMount(() => {
    id = randomString();
    hidden.add("table_meta_id");
    const element = document.createElement("canvas");
    canvasContext = element.getContext("2d");
    load();
  });

  $: {
    if (rows.length === 0) {
      filtered = rows;
    } else if (query === "") {
      filtered = rows;
    } else {
      const result = fuse.search(query);
      filtered = result.map((r) => r.item);
    }
    tableRows.set(filtered);
  }

  function selectAllRows() {
    for (let i = 0; i < filtered.length; i++) {
      if (i >= range.min && i <= range.max) {
        filtered[i].meta_selected = allRowsSelected ? false : true;
      }
    }
    allRowsSelected = !allRowsSelected;
    let count = 0;
    for (let i = 0; i < filtered.length; i++) {
      if (filtered[i].meta_selected) {
        count++;
      }
    }
    selectedCount = count;
  }

  async function load() {
    try {
      state = LoadState.Loading;
      rows = await getRows();
      if (rows.length === 0) {
        state = LoadState.Finished;
        return;
      }
      rows.map((w) => {
        w.table_meta_id = randomString();
        return w;
      });
      fuse = createFuse();
      filtered = rows;
      tableRows.set(filtered);
      let allColumns = new Set<string>();
      rows.forEach((r) => {
        Object.keys(r).forEach((c) => allColumns.add(c));
      });
      columns = Array.from(allColumns);
      columns = sortColumns?.(columns) ?? columns;
      filteredColumns = columns.filter((w) => !hidden.has(w));
      state = LoadState.Finished;
      if (defaultSortColumn) {
        sortColumn(defaultSortColumn);
      }
      setWidths();
    } catch (ex) {
      state = LoadState.Failed;
    }
  }

  function sortColumn(column: string) {
    if (sort === column) {
      sort = column;
      sortDirection = sortDirection === "asc" ? "desc" : "asc";
    } else {
      sort = column;
      sortDirection = "desc";
    }
    let isDate = columnMeta[column]?.type === "date";
    dispatchPrivate(id, "on_sort", { sort, sortDirection });
    filtered = filtered.sort(function (a, b) {
      var nameA = a[sort]?.toString()?.toUpperCase();
      var nameB = b[sort]?.toString()?.toUpperCase();
      if (nameA == null && nameB == null) {
        return 0;
      }
      if (nameA == null) {
        return 1;
      }
      if (nameB == null) {
        return 1;
      }
      if (isDate) {
        if (new Date(nameA).getTime() < new Date(nameB).getTime()) {
          return 1;
        }
        if (new Date(nameA).getTime() > new Date(nameB).getTime()) {
          return -1;
        }
      }
      if (nameA < nameB) {
        return 1;
      }
      if (nameA > nameB) {
        return -1;
      }
      return 0;
    });
    if (sortDirection === "asc") {
      filtered = filtered.reverse();
    }
    tableRows.set(filtered);
  }

  function headerStyle(column: string) {
    if (widths[column]) {
      return "min-width: " + widths[column] + "px;";
    }
  }

  function setWidths() {
    let values = filtered.slice(range.min, range.max);
    widths = {};
    values.forEach((value) => {
      columns.forEach((c) => {
        const v = value[c];
        let width = getTextWidth(v, "");

        if (width < 200) {
          width = 200;
        }
        if (width > 600) {
          width = 600;
        }
        if ((widths[c] ?? 0) < width) {
          widths[c] = width;
        }
      });
    });
  }

  function isUnread(row: any) {
    return row["meta_unread"] === true;
  }

  function renderValue(row: any, column: string) {
    let value = row[column] ?? "";
    value = onFormat(column, row[column]) ?? value;
    return isObject(value) || Array.isArray(value)
      ? JSON.stringify(value)
      : value;
  }

  function getTextWidth(text: string, font: string) {
    canvasContext.font = "bold 1em arial";
    return canvasContext.measureText(text).width;
  }

  function toggleColumn(checked: boolean, column: string) {
    checked ? hidden.delete(column) : hidden.add(column);
    filteredColumns = columns.filter((w) => !hidden.has(w));
  }

  function onRowSelected(row: any) {
    const index = filtered.findIndex(
      (w) => w.table_meta_id === row.table_meta_id
    );
    if (filtered[index].meta_selected) {
      selectedCount--;
      filtered[index].meta_selected = false;
    } else {
      selectedCount++;
      filtered[index].meta_selected = true;
    }
  }

  async function markRead(value: boolean) {
    const selected = filtered.filter((w) => w.meta_selected);
    const ids = new Set(selected.map((s) => s.table_meta_id));
    await onRead?.(selected, value);
    filtered = filtered.map((f) => {
      const id = f.table_meta_id;
      if (ids.has(id)) {
        f.meta_unread = !value;
      }
      return f;
    });
    tableRows.set(filtered);
    dispatch("show_toast", {
      message: `Item(s) marked as ${value ? "read" : "unread"}.`,
    });
  }

  async function deleteEntries() {
    const selected = filtered.filter((w) => w.meta_selected);
    if (selected.length !== selectedCount) {
      throw new Error(
        "Selection count did not match actual selected, please try reloading the page."
      );
    }
    await onDelete?.(selected);
    dispatch("show_toast", {
      title: "Deletion Started",
      message: "Your entries have been queued for deletion.",
    });
    modal = "";
    const toRemove = new Set(selected.map((w) => w.table_meta_id));
    filtered = filtered.filter((w) => {
      return !toRemove.has(w.table_meta_id);
    });
    filtered = filtered.map((f) => {
      f.meta_selected = false;
      return f;
    });
    rows = rows.filter((w) => {
      return !toRemove.has(w.table_meta_id);
    });
    tableRows.set(filtered);
    dispatch("show_toast", {
      message: `Item(s) deleted.'}.`,
    });
  }
</script>

<div class="bg-white px-4 py-5 border-b border-gray-200 sm:px-6">
  <div
    class="-ml-4 -mt-4 flex justify-between items-center flex-wrap
      sm:flex-nowrap"
  >
    <div>
      <div class="flex items-center">
        <div class="flex-shrink-0">
          <div class="mt-1 relative rounded-md shadow-sm">
            <div
              class="absolute inset-y-0 left-0 pl-3 flex items-center
                pointer-events-none"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-5 w-5"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
                />
              </svg>
            </div>
            <input
              class="form-input block w-full pl-10 sm:text-sm sm:leading-5"
              placeholder={searchPlaceHolder}
              bind:value={query}
            />
          </div>
        </div>
        <!-- Filter -->
        {#if showFilter}
          <div>
            <div class="text-gray-700 text-center px-4 py-2 m-2">
              <div on:click={() => (modal = "filter")}>
                <svg
                  class="w-5 h-5"
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.293A1 1 0 013 6.586V4z"
                  />
                </svg>
              </div>
            </div>
          </div>
        {/if}
        {#if headerActions && headerActions.length > 0}
          {#each headerActions as action}
            <div>
              <div class="text-gray-700 text-center px-4 py-2 m-2">
                <Button type="primary" onClick={action.onClick}>
                  {action.label}
                </Button>
              </div>
            </div>
          {/each}
        {/if}
        <!-- Applied Filters -->
        {#if appliedFilters > 0}
          <div>
            <div class="text-gray-700 text-center px-4 py-2 m-2">
              <span
                class="inline-flex items-center px-3 py-0.5 rounded-full text-sm
                  font-medium leading-5 bg-indigo-100 text-indigo-800"
              >
                {appliedFilters}
                Filter(s) Applied
                <button
                  on:click={() => {
                    filters = {};
                  }}
                  type="button"
                  class="flex-shrink-0 -mr-0.5 ml-1.5 inline-flex
                    text-indigo-500 focus:outline-none focus:text-indigo-700"
                  aria-label="Remove large badge"
                >
                  <svg
                    class="h-2 w-2"
                    stroke="currentColor"
                    fill="none"
                    viewBox="0 0 8 8"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-width="1.5"
                      d="M1 1l6 6m0-6L1 7"
                    />
                  </svg>
                </button>
              </span>
            </div>
          </div>
        {/if}
      </div>
    </div>
    {#if selectedCount > 0}
      <div class="ml-4 mt-4 flex-shrink-0 flex">
        <span class="inline-flex rounded-md">
          <div>
            Selected:
            <strong>{selectedCount} of {filtered.length}</strong>
          </div>
        </span>

        <slot name="selected_actions">
          <span
            class="ml-3 pl-2 inline-flex rounded-md cursor-pointer"
            on:click={() => markRead(true)}
          >
            <svg
              class="w-5 h-5"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
              />
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"
              />
            </svg>
          </span>

          <span
            class="ml-3 pl-2 inline-flex rounded-md cursor-pointer"
            on:click={() => markRead(false)}
          >
            <svg
              class="w-5 h-5"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.88 9.88l-3.29-3.29m7.532 7.532l3.29 3.29M3 3l3.59 3.59m0 0A9.953 9.953 0 0112 5c4.478 0 8.268 2.943 9.543 7a10.025 10.025 0 01-4.132 5.411m0 0L21 21"
              />
            </svg>
          </span>

          <span
            class="ml-3 pl-2 inline-flex rounded-md cursor-pointer"
            on:click={() => (modal = "delete")}
          >
            <svg
              class="w-5 h-5"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"
              />
            </svg>
          </span>
        </slot>
      </div>
    {/if}
    <span
      class="ml-3 pl-2 inline-flex rounded-md cursor-pointer"
      on:click={() => (modal = "toggle_column")}
    >
      <svg
        class="w-5 h-5"
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10"
        />
      </svg>
    </span>
  </div>
</div>

<div class="flex flex-col">
  <div class="-my-2 overflow-x-scroll sm:-mx-6 lg:-mx-8">
    <div class="py-2 align-middle inline-block sm:px-6 lg:px-8 min-w-full">
      <div
        class="shadow overflow-hidden border-b border-gray-200 sm:rounded-lg"
      >
        <table class="min-w-full divide-y divide-gray-200">
          <thead>
            <tr>
              <th
                class="px-6 py-3 bg-gray-50 text-left text-xs leading-4
                  font-medium text-gray-500 uppercase tracking-wider w-2"
              >
                <div>
                  <input
                    class="form-checkbox h-4 w-4 text-indigo-600 transition
                      duration-150 ease-in-out"
                    type="checkbox"
                    value=""
                    checked={allRowsSelected}
                    on:change={selectAllRows}
                    id={"row-toggle-all"}
                  />
                </div>
              </th>
              {#each filteredColumns as column (column)}
                <th
                  style={headerStyle(column)}
                  class="px-6 py-3 bg-gray-50 text-left text-xs leading-4
                    font-medium text-gray-500 uppercase tracking-wider"
                  on:click={() => sortColumn(column)}
                >
                  {column}
                  <span>
                    {#if sort === column && sortDirection === "asc"}
                      <span>
                        <svg
                          class="w-5 h-5"
                          xmlns="http://www.w3.org/2000/svg"
                          fill="none"
                          viewBox="0 0 24 24"
                          stroke="currentColor"
                        >
                          <path
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2"
                            d="M5 15l7-7 7 7"
                          />
                        </svg>
                      </span>
                    {:else if sort === column && sortDirection === "desc"}
                      <span>
                        <svg
                          class="w-5 h-5"
                          xmlns="http://www.w3.org/2000/svg"
                          fill="none"
                          viewBox="0 0 24 24"
                          stroke="currentColor"
                        >
                          <path
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2"
                            d="M19 9l-7 7-7-7"
                          />
                        </svg>
                      </span>
                    {/if}
                  </span>
                </th>
              {/each}
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            {#each filtered as row, index}
              {#if index >= range.min && index <= range.max}
                <tr
                  class:active={row.meta_selected}
                  style="vertical-align: middle; cursor: pointer;"
                  on:click={() => {
                    onRowClick(row);
                    row["meta_unread"] = false;
                  }}
                >
                  <td
                    class:bg-cool-gray-50={!isUnread(row)}
                    class="px-6 py-3 text-sm leading-5 font-medium text-gray-900"
                  >
                    <div class="flex items-center space-x-3">
                      <a href="#" class="truncate hover:text-gray-600">
                        <input
                          class="form-checkbox h-4 w-4 text-indigo-600
                            transition duration-150 ease-in-out"
                          type="checkbox"
                          value=""
                          checked={row.meta_selected}
                          on:click|stopPropagation
                          on:change|stopPropagation={(e) => {
                            onRowSelected(row);
                          }}
                          id={"row-toggle-" + index}
                        />
                      </a>
                      {#if isUnread(row)}
                        <div
                          class="flex-shrink-0 w-2.5 h-2.5 rounded-full
                            bg-indigo-400"
                        />
                      {/if}
                    </div>
                  </td>
                  {#each filteredColumns as column}
                    <td
                      class:bg-cool-gray-50={!isUnread(row)}
                      class:font-bold={isUnread(row)}
                      class:text-gray-800={isUnread(row)}
                      class="px-6 py-4 text-sm leading-5 text-gray-500"
                    >
                      <div class="text" class:text-unread={isUnread(row)}>
                        {renderValue(row, column)}
                      </div>
                    </td>
                  {/each}
                </tr>
              {/if}
            {/each}
          </tbody>
        </table>
      </div>
    </div>
  </div>
</div>
<Pagination
  {id}
  count={filtered.length}
  onRangeChange={(r) => {
    if (fastEquals(r, range)) {
      return;
    }
    range = r;
    setWidths();
    columns = columns;
  }}
/>

{#if modal === "toggle_column"}
  <Dialog
    title={"Toggle Column Visibility"}
    isOpen={true}
    onClose={() => {
      modal = "";
    }}
  >
    {#each columns as column}
      {#if column !== "table_meta_id"}
        <div class="mt-4">
          <div class="relative flex items-start">
            <div class="flex items-center h-5">
              <input
                class="form-checkbox h-4 w-4 text-indigo-600 transition
                  duration-150 ease-in-out"
                type="checkbox"
                on:click|stopPropagation
                checked={!hidden.has(column)}
                on:change={(e) => {
                  toggleColumn(e.target.checked, column);
                }}
              />
            </div>
            <div class="ml-3 text-sm leading-5">
              <label for={"toggle-" + column} class="font-medium text-gray-700"
                >{column}</label
              >
            </div>
          </div>
        </div>
      {/if}
    {/each}
  </Dialog>
{:else if modal === "delete"}
  <slot name="delete_entries">
    <Dialog
      title={"Confirm Deletion"}
      isOpen={true}
      actions={[
        {
          label: `Delete ${selectedCount} Entries`,
          type: "danger",
          onClick: deleteEntries,
        },
        { label: "Cancel", type: "secondary" },
      ]}
      onClose={() => {
        modal = "";
      }}
    >
      <p>Are you sure you want to delete {selectedCount} entries?</p>
    </Dialog>
  </slot>
{:else if modal === "filter"}
  <Dialog
    title={"Manage Filters"}
    isOpen={true}
    onClose={() => {
      modal = "";
    }}
  >
    <div class="form-check">
      <input
        class="form-check-input"
        type="checkbox"
        bind:checked={filters.onlyUnread}
        on:click|stopPropagation
      />
      <label class="form-check-label">Only Show Unread Items</label>
    </div>
  </Dialog>
{/if}

<style>
  .text {
    overflow: hidden;
    text-overflow: ellipsis;
    display: -webkit-box;
    -webkit-line-clamp: 2; /* number of lines to show */
    -webkit-box-orient: vertical;
  }
</style>
