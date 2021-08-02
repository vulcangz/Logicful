<script lang="typescript">
  import type { IFolder } from "@app/models/IFolder";
  import type { Dictionary } from "@app/models/Utility";

  export let folders: Dictionary<IFolder>;
  export let selected: string = "";
  export let onSelected: (folder: IFolder) => any;
  let expanded: { [key: string]: boolean } = {};
  export let padding: boolean = false;
</script>

{#each Object.values(folders) as folder}
  <div class="pt-1" class:pl-2={padding}>
    <div>
      <button
        on:click={() => {
          expanded[folder.id] = !expanded[folder.id] ?? true;
          if (selected === folder.id) {
            return;
          }
          onSelected(folder);
        }}
        class="group w-full flex items-center pl-2 pr-1 py-2 text-base leading-5
          font-medium rounded-md bg-white text-gray-600 hover:text-gray-900
          hover:bg-gray-50 focus:outline-none focus:text-gray-900
          focus:bg-gray-50 transition ease-in-out duration-150">
        {#if folder.id === selected}
          <div>
            <svg
              class="w-6 h-6"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor">
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M5 19a2 2 0 01-2-2V7a2 2 0 012-2h4l2 2h4a2 2 0 012 2v1M5 19h14a2 2 0 002-2v-5a2 2 0 00-2-2H9a2 2 0 00-2 2v5a2 2 0 01-2 2z" />
            </svg>
          </div>
        {:else}
          <div>
            <svg
              class="w-6 h-6"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor">
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6l-2-2H5a2 2 0 00-2 2z" />
            </svg>
          </div>
        {/if}
        <p class="pl-1 pr-1 font-light">{folder.name}</p>
      </button>
      <div>
        {#if expanded[folder.id]}
          <svelte:self
            padding={true}
            folders={folder.children ?? {}}
            {selected}
            {onSelected} />
        {/if}
      </div>
    </div>
  </div>
{/each}
