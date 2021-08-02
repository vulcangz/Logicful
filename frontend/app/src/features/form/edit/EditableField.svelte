<script lang="typescript">
  import { dispatch } from "@app/event/EventBus";
  import type { IField } from "@app/models/IField";
  import formStore from "@app/store/FormStore";
  import Field from "./Field.svelte";

  export let editor: boolean = false;
  export let hidden: boolean = false;
  export let field: IField;

  function onClone() {
    dispatch("field_clone", {
      field,
    });
  }

  function select() {
    if (field.configTarget || editor) {
      return;
    }
    field.selected = !field.selected;
    formStore.set(field, {
      field: "selected",
      value: field.selected,
      fromUser: false,
    });
  }

  let isConfigInput =
    field.configTarget || editor || field.type === "placeholder";
</script>

<div
  on:click|stopPropagation={select}
  class:bg-gray-100={!isConfigInput && field.selected}
  class={isConfigInput ? 'p-3 mb-2' : `hover:bg-gray-100 p-3 relative mb-2 rounded-md`}>
  {#if field.selected}
    <div
      class="absolute top-1 right-6 h-8 w-8 cursor-pointer pt-1"
      on:click|stopPropagation={onClone}>
      <svg
        class="w-5 h-5 align-middle"
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor">
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z" />
      </svg>
    </div>
    <div
      class="absolute top-1 right-0 h-8 w-8 cursor-pointer pt-1"
      on:click|stopPropagation={() => dispatch('confirm_field_deletion', {})}>
      <svg
        class="w-5 h-5"
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor">
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
      </svg>
    </div>
  {/if}
  <Field {field} {hidden} mode="builder" />
</div>
