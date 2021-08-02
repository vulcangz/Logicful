<script lang="typescript">
  import { onMount } from "svelte";
  import FieldEdit from "./FieldEdit.svelte";
  import type { IField } from "@app/models/IField";
  import FormEdit from "./FormEdit.svelte";
  import { subscribeFieldChange } from "@app/event/FieldEvent";
  import { fastClone } from "@app/util/Compare";
  import FieldConfigSidebarHeader from "./FieldConfigSidebarHeader.svelte";
  import { subscribeComponent } from "@app/event/EventBus";

  let field: IField | undefined;
  let fieldId: string | undefined;

  subscribeComponent("field_delete", () => {
    field = undefined;
    fieldId = undefined;
  });

  subscribeFieldChange(onMount, (newField: IField) => {
    if (newField.id === fieldId && !newField.selected) {
      field = undefined;
      fieldId = undefined;
      return;
    }
    if (newField.selected) {
      field = fastClone(newField);
      fieldId = field!.id;
    }
  });
</script>

{#if field}
  {#each [field] as f (fieldId)}
    <div>
      <FieldConfigSidebarHeader
        title={`Editing ${f.label}`}
        description={`Configure settings for your ${f.type} field.`} />
      <div class="sm:mx-auto sm:w-full">
        <div class="px-4 sm:rounded-lg sm:px-2">
          <FieldEdit field={f} />
        </div>
      </div>
    </div>
  {/each}
{:else}
  <FieldConfigSidebarHeader
    title={`Form Settings`}
    description={`Configure settings for your form.`} />
  <div class="sm:mx-auto sm:w-full">
    <div class="px-4 sm:rounded-lg sm:px-2">
      <FormEdit />
    </div>
  </div>
{/if}
