<script lang="typescript">
  import type { IField } from "@app/models/IField";
  import { subscribeFieldChange } from "@app/event/FieldEvent";
  import Label from "./Label.svelte";
  import { onMount } from "svelte";
  import formStore from "@app/store/FormStore";
  import { debounce } from "@app/util/Debounce";
  import { subscribeComponent } from "@app/event/EventBus";
  import InputErrorText from "./validations/InputErrorText.svelte";
  import FieldHelperText from "./FieldHelperText.svelte";
  import { validateField } from "../services/FieldValidator";

  export let field: IField;
  export let value: { [key: string]: string } = {};
  let otherText: string = "";
  let debouncedOnChange: any;

  subscribeFieldChange(onMount, (newField) => {
    if (newField.id === field.id) {
      value = newField.value ?? {};
      otherText = value.other ?? "";
    }
  });

  subscribeComponent("get_options", (props) => {
    if (props.id === field.id) {
      return field.options;
    }
  });

  onMount(() => {
    debouncedOnChange = debounce((field: IField) => {
      formStore.set(field, {
        fromUser: true,
        field: "value",
        value: field.value,
      });
      validateField(field, field.value);
    }, 500);

    value = formStore.getValue(field.configTarget ?? field.id) ?? {};
    otherText = value.other ?? "";
  });

  function onOtherChange(e: any) {
    otherText = e.target.value;
    if (otherText === "" || otherText == null) {
      delete value.other;
    } else {
      value.other = otherText ?? "";
    }
    field.value = value;
    debouncedOnChange(field);
  }

  function onChange(e: any, option: any) {
    e.stopPropagation();
    if (e.target.checked) {
      value[option] = option;
    } else {
      delete value[option];
    }
    field.value = value;
    formStore.set(field, {
      fromUser: true,
      field: "value",
      value: field.value,
    });
    validateField(field, field.value);
  }

  function isChecked(option: any) {
    return value[option] != null && value[option] != "";
  }
</script>

{#if !field.hideLabel}
  <Label {field} type="inline" />
{/if}
{#if field.options}
  {#each field.options as option}
    <div class="mt-4">
      <div class="relative flex items-start">
        <div class="flex items-center h-5">
          <input
            class="focus:ring-indigo-500 h-4 w-4 text-indigo-600 border-gray-300 rounded"
            type="checkbox"
            class:border-red-300={field.error}
            on:click|stopPropagation
            data-test-id={"checkbox-input-" +
              `${field.label}-${option}` +
              "-" +
              (field.required ? "required" : "optional")}
            on:change={(e) => onChange(e, option)}
            value=""
            checked={isChecked(option)}
            id={`${field.id}-${option}`}
          />
        </div>
        <div class="ml-3 text-sm leading-5">
          <label
            for={`${field.id}-${option}`}
            class:text-red-900={field.error}
            class="font-medium text-gray-700">{option}</label
          >
        </div>
      </div>
    </div>
  {/each}
  {#if !field.includeOther}
    <InputErrorText {field} />
    <FieldHelperText {field} />
  {/if}
  {#if field.includeOther}
    <div class="mt-4">
      <div class="relative flex items-start">
        <div class="flex items-center h-5">
          <input
            class="focus:ring-indigo-500 h-4 w-4 text-indigo-600 border-gray-300 rounded"
            data-test-id={"checkbox-input-" +
              `${field.label}-other` +
              "-" +
              (field.required ? "required" : "optional")}
            on:click|stopPropagation
            class:border-red-300={field.error}
            type="checkbox"
            value=""
            checked={otherText !== ""}
            id={`${field.id}-other`}
          />
        </div>
        <div class="ml-3 text-sm leading-5">
          <label
            for={`${field.id}-other`}
            class:text-red-900={field.error}
            class="font-medium text-gray-700"
            >Other:
          </label>
        </div>
      </div>
    </div>
    <div class="mt-4">
      <input
        class="shadow-sm focus:ring-indigo-500 focus:border-indigo-500 block w-full sm:text-sm border-gray-300 rounded-md"
        type="text"
        class:input-error={field.error}
        on:click|stopPropagation
        data-test-id={"checkbox-input-" +
          `${field.label}-other-input` +
          "-" +
          (field.required ? "required" : "optional")}
        on:input={(e) => onOtherChange(e)}
        value={otherText}
        id={`${field.id}-other-input`}
      />
      <InputErrorText {field} />
      <FieldHelperText {field} />
    </div>
  {/if}
{/if}
