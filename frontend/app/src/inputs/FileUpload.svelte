<script lang="typescript">
  import type { IField } from "@app/models/IField";
  import { subscribeFieldChange } from "@app/event/FieldEvent";
  import Label from "./Label.svelte";
  import { afterUpdate, onMount } from "svelte";
  import formStore from "@app/store/FormStore";
  import { firstNotEmpty } from "@app/util/Format";
  import { validateField } from "../services/FieldValidator";
  import FieldHelperText from "./FieldHelperText.svelte";
  import InputErrorText from "./validations/InputErrorText.svelte";

  export let field: IField;
  let files: FileList | undefined;
  let placeholder = "Choose a file...";
  let hasFile = false;
  let fileId = "";
  let fileInput: any;

  onMount(() => {
    placeholder = firstNotEmpty(
      field.placeholder ?? field.value?.name,
      "Choose a file..."
    );
  });

  function openSelector(e: any) {
    e.preventDefault();
    fileInput.click();
  }

  subscribeFieldChange(onMount, (newField: IField) => {
    if (newField.id === field.id) {
      placeholder = firstNotEmpty(
        field.placeholder ?? field.value?.name,
        "Choose a file..."
      );
    }
  });

  function clear() {
    if (files) {
      hasFile = false;
      placeholder = firstNotEmpty(field.placeholder, "Choose a file...");
      files = undefined;
      formStore.clearFile(fileId);
      field.value = undefined;
      field.error = "";
      formStore.set(field, {
        field: "value",
        value: undefined,
        fromUser: true,
      });
    }
  }

  afterUpdate(() => {
    if (files && files[0] && !hasFile) {
      const file = files[0];
      hasFile = true;
      placeholder = file.name;
      fileId = field.id;
      formStore.setFile(fileId, file);
      field.value = {
        name: placeholder,
        id: fileId,
        size: file.size,
        type: file.type,
      };
      formStore.set(field, {
        field: "value",
        value: field.value,
        fromUser: true,
      });
      validateField(field, field.value);
    }
  });
</script>

<div>
  <input
    type="file"
    bind:this={fileInput}
    {placeholder}
    data-test-id={"file-input-hidden-" +
      field.label +
      "-" +
      (field.required ? "required" : "optional")}
    id={`${field.id}-file-input`}
    bind:files
    on:click|stopPropagation
    class="hidden"
  />
  <Label {field} />
  {#if !hasFile}
    <div class="mt-1 flex rounded-md shadow-sm" on:click|stopPropagation>
      <input
        type="text"
        on:click|stopPropagation|preventDefault={openSelector}
        {placeholder}
        data-test-id={"file-input-" +
          field.label +
          "-" +
          (field.required ? "required" : "optional")}
        class:input-error={field.error}
        id={`${field.id}-file-input`}
        readonly
        class="form-input block w-full rounded-none rounded-l-md transition
          ease-in-out duration-150 sm:text-sm sm:leading-5"
      />
      <button
        on:click|stopPropagation|preventDefault={openSelector}
        class="-ml-px relative inline-flex items-center px-4 py-2 border
          border-gray-300 text-sm leading-5 font-medium rounded-r-md
          text-gray-700 bg-gray-50 hover:text-gray-500 hover:bg-white
          focus:outline-none focus:ring-indigo focus:border-indigo-300
          active:bg-gray-100 active:text-gray-700 transition ease-in-out
          duration-150"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          class="w-5 h-5"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"
          />
        </svg>
        <span class="ml-2">Browse</span>
      </button>
    </div>
  {:else}
    <div class="mt-1 flex rounded-md shadow-sm">
      <input
        type="text"
        value={placeholder}
        on:click|stopPropagation|preventDefault
        readonly
        data-test-id={"file-input-" +
          field.label +
          "-" +
          (field.required ? "required" : "optional")}
        class:input-error={field.error}
        id={`${field.id}-file-input`}
        disabled
        class="form-input block w-full rounded-none rounded-l-md transition
          ease-in-out duration-150 sm:text-sm sm:leading-5"
      />
      <button
        data-test-id={"file-input-clear-" +
          field.label +
          "-" +
          (field.required ? "required" : "optional")}
        on:click|stopPropagation|preventDefault={clear}
        class="-ml-px relative inline-flex items-center px-4 py-2 border
          border-gray-300 text-sm leading-5 font-medium rounded-r-md
          text-gray-700 bg-gray-50 hover:text-gray-500 hover:bg-white
          focus:outline-none focus:ring-indigo focus:border-indigo-300
          active:bg-gray-100 active:text-gray-700 transition ease-in-out
          duration-150"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          class="w-5 h-5"
          stroke="currentColor"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M12 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2M3 12l6.414 6.414a2 2 0 001.414.586H19a2 2 0 002-2V7a2 2 0 00-2-2h-8.172a2 2 0 00-1.414.586L3 12z"
          />
        </svg>
        <span class="ml-2">Clear</span>
      </button>
    </div>
  {/if}
  <InputErrorText {field} />
  <FieldHelperText {field} />
</div>
