<script lang="typescript">
  import type { IField } from "@app/models/IField";
  import { subscribeFieldChange } from "@app/event/FieldEvent";
  import Label from "./Label.svelte";
  import { onMount } from "svelte";
  import formStore from "@app/store/FormStore";
  import { debounce } from "@app/util/Debounce";
  import { subscribeComponent } from "@app/event/EventBus";
  import { validateField } from "../services/FieldValidator";
  import InputErrorText from "./validations/InputErrorText.svelte";
  import FieldHelperText from "./FieldHelperText.svelte";

  export let field: IField;
  export let value: { [key: string]: string } = {};
  let otherText: string = "";
  let debouncedOnChange: any;
  let otherSelected = false;

  subscribeFieldChange(onMount, (newField) => {
    if (newField.id === field.id) {
      value = newField.value ?? {};
      if (otherText && !value.other) {
        return;
      }
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
    otherSelected = otherText != null && otherText != "";
  });

  function onOtherChange(e: any) {
    otherText = e.target.value;
    value = {};
    if (otherText != "" && otherText != null) {
      value.other = otherText;
    }
    field.value = value;
    debouncedOnChange(field);
  }

  function onOtherRadioChange() {
    otherSelected = true;
    value = { other: otherText };
    field.value = value;
    formStore.set(field, {
      fromUser: true,
      field: "value",
      value: field.value,
    });
    validateField(field, field.value);
  }

  function onChange(e: any, option: any) {
    e.stopPropagation();
    value = {};
    value[option] = option;
    otherSelected = false;
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

<Label {field} type="inline" />
{#if field.options}
  {#each field.options as option}
    <div class="mt-4">
      <div class="relative flex items-start">
        <div class="flex items-center h-5">
          <input
            class="focus:ring-indigo-500 h-4 w-4 text-indigo-600 border-gray-300"
            type="radio"
            data-test-id={"radio-input-" +
              `${field.label}-${option}` +
              "-" +
              (field.required ? "required" : "optional")}
            on:click|stopPropagation
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
            class="focus:ring-indigo-500 h-4 w-4 text-indigo-600 border-gray-300"
            data-test-id={"radio-input-" +
              `${field.label}-other` +
              "-" +
              (field.required ? "required" : "optional")}
            on:click|stopPropagation
            on:click={onOtherRadioChange}
            type="radio"
            value=""
            checked={otherSelected}
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
        on:input={(e) => onOtherChange(e)}
        value={otherText}
        data-test-id={"radio-input-" +
          `${field.label}-other-input` +
          "-" +
          (field.required ? "required" : "optional")}
        id={`${field.id}-other-input`}
      />
      <InputErrorText {field} />
      <FieldHelperText {field} />
    </div>
  {/if}
{/if}
