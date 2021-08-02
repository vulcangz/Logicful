<script lang="typescript">
  import { subscribePrivate } from "@app/event/EventBus";

  import { afterUpdate, onMount } from "svelte";

  export let id = "";
  export let rowsPerPage = 10;
  export let page = 1;
  export let count = 0;
  export let onRangeChange: ({ min, max }: { min: number; max: number }) => any;
  let pages = 1;
  let hasNext = true;
  let hasPrevious = false;
  let lastRowsPerPage = 0;
  let rowsPerPageEntries = [10, 25, 50, 100];
  let showingCount = 0;
  let showing = "";

  let classes = `-ml-px relative inline-flex items-center px-4 py-2 border
            border-gray-300 bg-white text-sm leading-5 font-medium text-gray-700
            hover:text-gray-500 focus:z-10 focus:outline-none
            focus:border-indigo-300 focus:ring-indigo active:bg-gray-100
            active:text-gray-700 transition ease-in-out duration-150`;

  afterUpdate(() => {
    onChange();
  });

  onMount(() => {
    subscribePrivate(id, "on_sort", () => {
      page = 1;
    });
  });

  function onChange() {
    pages = Math.ceil(count / rowsPerPage);
    if (page > pages) {
      page = pages || 1;
    }
    hasNext = page < pages;
    hasPrevious = page > 1;
    showingCount = Math.floor(page * rowsPerPage);
    showingCount = showingCount >= count ? Math.floor(count) : showingCount;
    showing = `Showing ${showingCount} / ${Math.floor(count)} Entries`;
    onRangeChange(range());
  }

  $: {
    if (rowsPerPage !== lastRowsPerPage) {
      lastRowsPerPage = rowsPerPage;
      page = 1;
    }
  }

  function setPage(newPage: number) {
    page = newPage;
  }

  function range(): { min: number; max: number } {
    const max = rowsPerPage * page;
    const min = max - rowsPerPage;
    return { min, max };
  }
</script>

<div class="bg-white px-4 py-3 flex border-t border-gray-200 sm:px-6">
  <div class="flex-1 flex just- sm:hidden">
    <a
      href="#"
      class="relative inline-flex items-center px-4 py-2 border border-gray-300
        text-sm leading-5 font-medium rounded-md text-gray-700 bg-white
        hover:text-gray-500 focus:outline-none focus:ring-indigo
        focus:border-indigo-300 active:bg-gray-100 active:text-gray-700
        transition ease-in-out duration-150"
    >
      Previous
    </a>
    <a
      href="#"
      class="ml-3 relative inline-flex items-center px-4 py-2 border
        border-gray-300 text-sm leading-5 font-medium rounded-md text-gray-700
        bg-white hover:text-gray-500 focus:outline-none
        focus:ring-indigo focus:border-indigo-300 active:bg-gray-100
        active:text-gray-700 transition ease-in-out duration-150"
    >
      Next
    </a>
  </div>
  <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
    <div>
      <p class="text-sm leading-5 text-gray-700">
        Showing
        <span class="font-medium"
          >{showingCount === count
            ? showingCount
            : showingCount - rowsPerPage}</span
        >
        to
        <span class="font-medium">{showingCount}</span>
        of
        <span class="font-medium">{count}</span>
        results
      </p>
    </div>
    <div class="flex-initial">
      <div />
    </div>
    <div class="flex-initial">
      <nav class="relative z-0 inline-flex shadow-sm">
        <div class="relative px-2">
          <select
            bind:value={rowsPerPage}
            class="mt-1 form-select block w-full pl-3 pr-10 py-2 text-base
              leading-6 border-gray-300 focus:outline-none
              focus:ring-indigo focus:border-indigo-300 sm:text-sm
              sm:leading-5"
          >
            {#each rowsPerPageEntries as entry}
              <option value={entry}>Show {entry} Entries</option>
            {/each}
          </select>
        </div>
        <a
          href="javascript:void(0)"
          class:disabled={!hasPrevious}
          on:click={() => setPage(page - 1)}
          class="relative inline-flex items-center px-2 py-2 rounded-l-md border
            border-gray-300 bg-white text-sm leading-5 font-medium text-gray-500
            hover:text-gray-400 focus:z-10 focus:outline-none
            focus:border-indigo-300 focus:ring-indigo
            active:bg-gray-100 active:text-gray-500 transition ease-in-out
            duration-150"
          aria-label="Previous"
        >
          <!-- Heroicon name: chevron-left -->
          <svg class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
            <path
              fill-rule="evenodd"
              d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z"
              clip-rule="evenodd"
            />
          </svg>
        </a>
        {#if hasPrevious && page - 2 >= 1}
          <a
            class={classes}
            on:click={() => setPage(page - 2)}
            href="javascript:void(0)">{page - 2}</a
          >
        {/if}
        {#if page > 1}
          <a
            class={classes}
            on:click={() => setPage(page - 1)}
            href="javascript:void(0)">{page - 1}</a
          >
        {/if}
        <a
          class="-ml-px relative inline-flex items-center px-4 py-2 border
            border-gray-300 bg-indigo-700 text-sm leading-5 font-medium
            text-white hover:text-gray-500 focus:z-10 focus:outline-none
            focus:border-indigo-300 focus:ring-indigo
            active:bg-indigo-700 active:text-white transition ease-in-out
            duration-150"
          href="javascript:void(0)">{page}</a
        >
        {#if hasNext}
          <a
            class={classes}
            on:click={() => setPage(page + 1)}
            href="javascript:void(0)">{page + 1}</a
          >
        {/if}
        {#if page + 2 <= pages}
          <a
            class={classes}
            on:click={() => setPage(page + 2)}
            href="javascript:void(0)">{page + 2}</a
          >
        {/if}
        <a
          href="javascript:void(0)"
          class:disabled={!hasNext}
          on:click={() => setPage(page + 1)}
          class="-ml-px relative inline-flex items-center px-2 py-2 rounded-r-md
            border border-gray-300 bg-white text-sm leading-5 font-medium
            text-gray-500 hover:text-gray-400 focus:z-10 focus:outline-none
            focus:border-indigo-300 focus:ring-indigo
            active:bg-gray-100 active:text-gray-500 transition ease-in-out
            duration-150"
          aria-label="Next"
        >
          <!-- Heroicon name: chevron-right -->
          <svg class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
            <path
              fill-rule="evenodd"
              d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z"
              clip-rule="evenodd"
            />
          </svg>
        </a>
      </nav>
    </div>
  </div>
</div>

<!-- <li>
      <div class="dropdown">
        <button
          class="btn btn-secondary dropdown-toggle rows-page-button"
          type="button"
          id="dropdownMenuButton"
          data-toggle="dropdown"
          aria-expanded="false">{showing}</button>
        <ul class="dropdown-menu" aria-labelledby="dropdownMenuButton">
          {#each rowsPerPageEntries as entry}
            <li>
              <a
                class="dropdown-item"
                href="javascript:void(0)"
                on:click={() => {
                  setRowsPerPage(entry);
                }}>
                Show {entry} Entries
              </a>
            </li>
          {/each}
        </ul>
      </div>
    </li> -->
