<script lang="typescript">
  import { onMount } from "svelte";
  import { dispatch, subscribeComponent } from "@app/event/EventBus";
  import type { IFolder } from "@app/models/IFolder";
  import type { IForm } from "@app/models/IForm";
  import { getApi } from "@app/services/ApiService";
  import type { User } from "@app/models/User";
  import { me } from "@app/services/AuthService";
  import { LoadState } from "@app/models/LoadState";
  import { debounce } from "@app/util/Debounce";
  import FormList from "@app/features/folders/FormList.svelte";
  import Loader from "@app/components/Loader.svelte";
  import { cacheClear } from "@app/util/Cache";
  import FolderSettings from "./FolderSettings.svelte";
  import FolderNoForms from "./FolderNoForms.svelte";

  let forms: IForm[] = [];
  let user: User;
  let folder: IFolder;
  let state: LoadState = LoadState.NotStarted;
  let editing = false;

  subscribeComponent("folder_edit", () => {
    editing = true;
  });

  subscribeComponent("forms_moved", (newFolder) => {
    forms = [];
    state = LoadState.Loading;
    cacheClear(`api-request-form?folderId=${newFolder}`);
    setForms(false);
  });

  subscribeComponent(
    "folder_selected",
    async (e: { folder: IFolder; showForms: any }) => {
      forms = [];
      state = LoadState.Loading;
      folder = e.folder;
      debounceSetForms();
    }
  );

  async function setForms(cache: boolean = true) {
    forms = await getApi(`form?folderId=${folder.id}`);
    folder.forms = forms;
    dispatch("folder_loaded", folder);
    state = LoadState.Finished;
  }

  const debounceSetForms = debounce(async () => {
    setForms();
  }, 300);

  onMount(async () => {
    user = await me();
    dispatch("folder_content_loaded", {});
  });
</script>

{#if editing}
  <FolderSettings
    {folder}
    onClose={() => {
      editing = false;
    }} />
{/if}

{#if state === LoadState.Loading}
  <Loader />
{/if}

{#if state === LoadState.Finished && forms && forms.length === 0}
  <FolderNoForms folderId={folder?.id} />
{/if}

{#if state === LoadState.Finished && forms && forms.length > 0}
  <FormList {forms} />
{/if}
