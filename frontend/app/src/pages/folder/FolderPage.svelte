<script lang="typescript">
  import Folders from "@app/features/folders/Folders.svelte";
  import FolderContent from "@app/features/folders/FolderContent.svelte";
  import Shell from "@app/components/Shell.svelte";
  import { dispatch, subscribeComponent } from "@app/event/EventBus";
  import type { IFolder } from "@app/models/IFolder";

  let folder: IFolder = { id: "", name: "My Forms", forms: [] };

  subscribeComponent(
    "folder_selected",
    async (e: { folder: IFolder; showForms: any }) => {
      folder = e.folder;
    }
  );

  subscribeComponent("folder_loaded", async (f: IFolder) => {
    folder = f;
  });
</script>

<Shell showSidebar={true}>
  <div slot="header">
    <div class="lg:flex lg:items-center lg:justify-between">
      <div class="flex-1 min-w-0">
        <div class="flex">
          <div class="flex-initial text-gray-700 text-center">
            <h1
              class="font-bold leading-4 text-gray-900 sm:text-2xl sm:leading-9
                sm:truncate">
              {folder.name}
            </h1>
          </div>
          <div
            class="flex-initial text-gray-700 text-center pl-2 md:mt-2
              cursor-pointer"
            on:click={() => dispatch('folder_edit', folder)}>
            <svg
              class="h-5 w-5"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor">
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
            </svg>
          </div>
        </div>
        <div class="mt-1 flex flex-col sm:mt-0 sm:flex-row sm:flex-wrap">
          <div
            class="mt-2 flex items-center text-sm leading-5 text-gray-500
              sm:mr-6">
            <svg
              class="flex-shrink-0 mr-1.5 h-5 w-5"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor">
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
            </svg>
            {#if folder.forms}{folder.forms.length} Forms{:else}0 Forms{/if}
          </div>
        </div>
      </div>
    </div>
  </div>
  <div slot="sidebar" class="mt-4 hidden md:block">
    <Folders />
  </div>
  <div class="mt-8">
    <div class="sm:block md:hidden -ml-4 -mt-6 pb-3">
      <Folders />
    </div>
    <FolderContent />
  </div>
</Shell>
