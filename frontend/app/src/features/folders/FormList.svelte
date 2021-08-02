<script lang="typescript">
  import type { IForm } from "@app/models/IForm";
  import MoveFormManager from "./MoveFormManager.svelte";
  import { dispatch } from "@app/event/EventBus";
  import FormListItem from "./FormListItem.svelte";

  export let forms: IForm[] = [];
  let moving: IForm | undefined;

  function moveFolders(form: IForm) {
    moving = form;
  }

  function onMoveComplete() {
    dispatch("forms_moved", moving!.id);
    moving = undefined;
  }

  function onDelete(formId: string | undefined) {}
</script>

{#if moving}
  <div>
    <MoveFormManager onClose={onMoveComplete} forms={[moving]} />
  </div>
{/if}

{#each forms as form}
  <div class="-mt-2 mb-4">
    <FormListItem
      {form}
      onMove={moveFolders}
      onDelete={() => onDelete(form.id)} />
    <div class="border-b border-gray-200 sm:px-6 pt-3" />
  </div>
{/each}
