<script lang="typescript">
  // @ts-nocheck
  import type { IForm } from "@app/models/IForm";
  import type { IField } from "@app/models/IField";
  import { subscribeFieldChange } from "@app/event/FieldEvent";
  import { DynamicFormMode } from "@app/components/models/ComponentProps";
  import { dispatch, subscribeComponent } from "@app/event/EventBus";
  import { LogicBuilder } from "@app/services/LogicBuilder";
  import { fastClone } from "@app/util/Compare";
  import { onMount } from "svelte";
  import Dialog from "@app/components/layout/Dialog.svelte";
  import EditableField from "./EditableField.svelte";
  import Button from "@app/components/Button.svelte";
  import { loadScripts } from "@app/util/Script";

  export let form: IForm;
  export let mode: DynamicFormMode = DynamicFormMode.Live;
  let saving = false;
  let deleting = false;

  onMount(async () => {
    await loadScripts([
      "https://cdnjs.cloudflare.com/ajax/libs/acorn/6.4.2/acorn.min.js",
      "https://cdn.jsdelivr.net/npm/js-interpreter@2.2.0/lib/js-interpreter.min.js",
    ]);
  });

  subscribeComponent("save_form", () => {
    saving = true;
  });

  subscribeComponent("form_saved", () => {
    saving = false;
  });

  subscribeComponent("confirm_field_deletion", () => {
    deleting = true;
  });

  subscribeFieldChange(onMount, (updatedField: IField) => {
    if (!form || !form.fields) {
      return;
    }
    const index = form.fields.findIndex((w) => w.id === updatedField.id);
    if (index === -1) {
      return;
    }
    form.fields[index].updated = !form.fields[index].updated;
    const fieldsWithRules = form.fields.filter((w) => {
      if (!w.logic || !w.logic.rules) {
        return false;
      }
      const hasRule = w.logic.rules.find(
        (rule) => rule.field === updatedField.id
      );
      return hasRule != null;
    });
    for (let fieldWithRule of fieldsWithRules) {
      let ruleIndex = form.fields.findIndex((w) => w.id === fieldWithRule.id);
      form.fields[ruleIndex].updated = !form.fields[ruleIndex].updated;
    }
  });

  function display(field: IField): boolean {
    if (!form.enableLogic) {
      return true;
    }
    if (!field.logic) {
      return true;
    }
    const builder = new LogicBuilder();
    return builder.evaluate(field);
  }

  function onDelete() {
    const selected = form.fields.find((w) => w.selected);
    if (selected) {
      dispatch("field_delete", {
        field: selected,
      });
    }
  }
</script>

{#if deleting}
  <Dialog
    title={'Confirm Deletion'}
    isOpen={true}
    actions={[{ label: `Delete Field`, type: 'danger', onClick: onDelete, focus: true }, { label: 'Cancel', type: 'secondary' }]}
    onClose={() => {
      deleting = false;
    }}>
    <p class="text-sm leading-5 text-gray-500">
      Are you sure you want to delete this field?
    </p>
  </Dialog>
{/if}
<div
  class="border-b border-gray-200 px-4 py-4 sm:flex sm:items-center
    sm:justify-between">
  <div class="flex-1 min-w-0">
    <h1 class="text-lg font-medium leading-6 text-gray-900 sm:truncate" data-test-id="form-title">
      {form.title}
    </h1>
    <p class="max-w-4xl text-sm leading-5 text-gray-500">{form.description}</p>
    <span
      class="hidden lg:inline-flex items-center mt-2 px-3 py-0.5 rounded-full
        text-sm font-medium leading-5 bg-indigo-100 text-indigo-800">
      {#if saving}
        Saving...
      {:else}
        <div class="hidden xl:inline-flex">All Changes Saved</div>
        <div class="xl:hidden">Saved</div>
      {/if}
    </span>
  </div>
  <div class="mt-4 flex sm:mt-0 sm:ml-4">
    <span class="order-1 ml-3 shadow-sm rounded-md sm:order-0 sm:ml-0">
      <Button
        type={'secondary'}
        href={`/form/submissions?formId=${form.id}&mode=local`}>
        Submissions
      </Button>
    </span>
    <span class="order-0 sm:order-1 sm:ml-3 shadow-sm rounded-md">
      <Button
        type={'primary'}
        href={`/form/preview?formId=${form.id}&mode=local`}>
        Preview
      </Button>
    </span>
  </div>
</div>

<div class="sm:mx-auto sm:w-full">
  <div class="py-5 px-4 sm:rounded-lg sm:px-2">
    <form action="#" method="POST" id="form-preview">
      <div id="form-preview-fields">
        {#each form.fields as field (field.id)}
          {#if display(field)}
            <div id={`form-field-${field.id}`}>
              <EditableField field={fastClone(field)} />
            </div>
          {:else}
            <div id={`form-field-${field.id}`}>
              <EditableField field={fastClone(field)} hidden={true} />
            </div>
          {/if}
        {/each}
      </div>
    </form>
  </div>
</div>

<style>
  :global(.ex-over) {
    background-color: #f5f5f5;
    height: 100%;
    min-height: 25vh;
    margin-top: 1em;
    margin-bottom: 1em;
  }
</style>
