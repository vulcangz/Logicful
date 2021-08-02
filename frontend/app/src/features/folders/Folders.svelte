<script lang="typescript">
  import { dispatch, subscribe, subscribeComponent } from "@app/event/EventBus";
  import type { IFolder } from "@app/models/IFolder";
  import { LoadState } from "@app/models/LoadState";
  import type { User } from "@app/models/User";
  import { postApi } from "@app/services/ApiService";
  import { me } from "@app/services/AuthService";
  import { onMount } from "svelte";
  import Dialog from "@app/components/layout/Dialog.svelte";
  import FolderList from "./FolderList.svelte";
  import { getFolders } from "@app/features/folders/FolderService";
  import Button from "@app/components/Button.svelte";

  let selected: IFolder | undefined = undefined;
  let folders: { [key: string]: IFolder } = {};
  let searchPlaceHolder = "Search for a form";
  let query = "";
  let creatingNewFolder = false;
  let newFolderName = "";
  let user: User;
  let state = LoadState.NotStarted;
  let contentLoaded = false;
  let parent = "";

  async function onFolderSelected() {
    if (contentLoaded && selected) {
      dispatch("folder_selected", { folder: selected });
    }
  }

  subscribeComponent("folder_deleted", async (folder) => {
    state = LoadState.Loading;
    await loadFolders();
  });

  subscribeComponent("folder_updated", async (folder) => {
    state = LoadState.Loading;
    await loadFolders();
    onSelected(folder);
  });

  onMount(async () => {
    user = await me();
    state = LoadState.Loading;
    subscribe("folder_content_loaded", () => {
      contentLoaded = true;
    });
    await loadFolders();
  });

  async function loadFolders(cache: boolean = true) {
    folders = await getFolders(cache);
    onSelected(folders[Object.keys(folders)[0]]);
    state = LoadState.Finished;
  }

  function onSelected(folder: IFolder) {
    selected = folder;
    onFolderSelected();
  }

  function onNewFolder(parentFolder: string = "") {
    parent = parentFolder;
    creatingNewFolder = true;
  }

  async function createNewFolder() {
    await postApi("folder", {
      name: newFolderName,
      teamId: user.teamId,
      parent: parent,
    });
    newFolderName = "";
    await loadFolders(false);
    dispatch("folder_created", parent);
  }
</script>

{#if creatingNewFolder}
  <Dialog
    title={'Create New Folder'}
    isOpen={true}
    actions={[{ label: `Create Folder`, type: 'primary', onClick: createNewFolder }, { label: 'Cancel', type: 'secondary' }]}
    onClose={() => {
      creatingNewFolder = false;
    }}>
    <label class="block text-gray-700 text-sm font-bold mb-2" for="folderName">
      Folder Name
    </label>
    <input
      bind:value={newFolderName}
      class="bg-white focus:outline-none focus:ring border
        border-gray-300 rounded-lg py-2 px-4 block appearance-none
        leading-normal w-full"
      type="text"
      id="folderName"
      name="folderName"
      placeholder="" />
  </Dialog>
{/if}

<div
  class="flex flex-col flex-grow border-r border-gray-200 pb-4 pt-2 bg-white
    overflow-y-auto sm:w-screen w-screen md:w-80">
  <div class="flex-grow flex flex-col">
    <nav class="flex-1 px-3 bg-white">
      <h3
        class="pl-2 mb-3 text-xs leading-4 font-semibold text-gray-500 uppercase
          tracking-wider">
        My Folders
      </h3>
      <div class="mb-2 flex flex-col space-y-2">
        <span class="inline-flex rounded-md shadow-sm">
          <Button type="primary" width="full" href="/form/create">
            + New Form
          </Button>
        </span>
        <span class="inline-flex rounded-md shadow-sm">
          <Button type="white" width="full" onClick={() => onNewFolder()}>
            + New Folder
          </Button>
        </span>
      </div>
      <div>
        <FolderList
          {onNewFolder}
          {onSelected}
          {folders}
          selected={selected?.id} />
      </div>
    </nav>
  </div>
</div>
