<script lang="typescript">
  import Dialog from "@app/components/layout/Dialog.svelte";
  import type { ButtonAction } from "@app/components/models/ComponentProps";
  import type { IFolder } from "@app/models/IFolder";
  import type { IForm } from "@app/models/IForm";
  import type { Dictionary } from "@app/models/Utility";
  import { getApi } from "@app/services/ApiService";
  import { onMount } from "svelte";
  import { saveForm } from "../form/edit/services/SaveForm";
  import { getFolders } from "./FolderService";
  import MoveForm from "./MoveForm.svelte";

  export let onClose: () => any;
  export let forms: IForm[];
  let folders: Dictionary<IFolder> = {};
  let selected: string = "";
  let name = "";

  function onSelected(folder: IFolder) {
    selected = folder.id;
    name = folder.name;
  }

  async function moveForms() {
    const promises = forms.map((f) => {
      return getApi<IForm>(`form/${f.id}`);
    });
    const fullForms = await Promise.all(promises);
    const saves = fullForms.map((f) => {
      f.folder = selected;
      return saveForm({ initiator: "user" }, f);
    });
    await Promise.all(saves);
  }

  function getActions() {
    let actions: ButtonAction[] = [
      {
        type: "primary",
        onClick: moveForms,
        label: "Move Form(s) to " + name + " Folder",
      },
    ];
    return actions;
  }

  onMount(async () => {
    folders = await getFolders(true);
    onSelected(folders[Object.keys(folders)[0]]);
  });
</script>

<Dialog
  isOpen={true}
  {onClose}
  title={'Select Folder To Move Form(s)'}
  {getActions}>
  <MoveForm {folders} {selected} {onSelected} />
</Dialog>
