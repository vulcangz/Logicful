<script lang="typescript">
  import type { IField } from "@app/models/IField";
  import { onMount } from "svelte";
  import formStore from "@app/store/FormStore";
  import { subscribeFieldChange } from "@app/event/FieldEvent";
  import Label from "./Label.svelte";

  export let config: any;
  export let field: IField;
  let checked: boolean = false;

  export let value: boolean | undefined = undefined;
  let defaultValue = false;

  subscribeFieldChange(onMount, (newField, change) => {
    if (newField.id === field.id) {
      if (change.field === "defaultValue") {
        value = newField.defaultValue;
      } else {
        value = newField.value;
      }
      checked = value ?? false;
    }
  });

  onMount(() => {
    value = formStore.getValue(field.configTarget ?? field.id);
    checked = value ?? false;
  });
</script>

<div class="flex items-center">
  <span
    role="checkbox"
    tabindex="0"
    aria-checked={checked}
    on:click|stopPropagation={() => {
      checked = !checked;
      field.value = checked;
      formStore.set(field, {
        fromUser: true,
        value: field.value,
        field: "value",
      });
      field.onChange?.(field.value);
    }}
    class:bg-gray-200={!checked}
    class:bg-indigo-600={checked}
    class="relative inline-flex flex-shrink-0 h-6 w-11 border-2
      border-transparent rounded-full cursor-pointer transition-colors
      ease-in-out duration-200 focus:outline-none focus:ring mr-2"
  >
    <span
      aria-hidden="true"
      class:translate-x-0={checked}
      class:translate-x-5={checked}
      class="inline-block h-5 w-5 rounded-full bg-white shadow transform
        transition ease-in-out duration-200"
    />
  </span>
  <Label {field} showOptional={false} />
</div>
