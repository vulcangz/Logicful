<script lang="typescript">
  import Dialog from "@app/components/layout/Dialog.svelte";
  import type { ButtonAction } from "@app/components/models/ComponentProps";
  import { dispatch } from "@app/event/EventBus";
  import type { IFolder } from "@app/models/IFolder";
  import { deleteApi, putApi } from "@app/services/ApiService";
  export let folder: IFolder;
  export let onClose: () => any;
  let deleting = false;

  let deleteActions: ButtonAction[] = [
    {
      label: "Confirm",
      onClick: deleteFolder,
      type: "danger",
    },
    {
      label: "Cancel",
      onClick: () => (deleting = false),
      type: "secondary",
      onClose: false,
    },
  ];

  let actions: ButtonAction[] = [
    {
      label: "Delete",
      onClick: () => (deleting = true),
      type: "danger",
      onClose: false,
    },
    {
      label: "Save",
      onClick: save,
      type: "primary",
    },
  ];

  async function deleteFolder() {
    await deleteApi(`folder/${folder.id}`, {});
    dispatch("folder_deleted", folder);
  }

  async function save() {
    await putApi(`folder/${folder.id}`, folder);
    dispatch("folder_updated", folder);
  }
</script>

{#if deleting}
  <Dialog
    isOpen={true}
    actions={deleteActions}
    title={`Confirm Deletion For ${folder.name}`}
    {onClose}>
    <p>
      Are you sure you want to delete this folder? All forms in this folder will
      be moved to uncategorized.
    </p>
  </Dialog>
{/if}

{#if folder && !deleting}
  <Dialog isOpen={true} {actions} title={`Editing ${folder.name}`} {onClose}>
    <label for="name">Name</label>
    <input
      type="text"
      class="form-control"
      name="name"
      bind:value={folder.name} />
  </Dialog>
{/if}
