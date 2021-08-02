<script lang="typescript">
  import { subscribeComponent } from "@app/event/EventBus";

  import { subscribeFieldChange } from "@app/event/FieldEvent";
  import type { IField } from "@app/models/IField";
  import formStore from "@app/store/FormStore";
  import { onMount } from "svelte";
  let value: any;
  let error: any;
  export let fieldId: string | undefined;

  subscribeFieldChange(onMount, (newField: IField) => {
    if (newField.id === fieldId) {
      value = newField.value;
    }
  });

  subscribeComponent("logic_rule_javascript_error", (err) => {
    error = err;
  });

  onMount(() => {
    if (!fieldId) {
      return;
    }
    value = formStore.getValue(fieldId);
  });
</script>

{#if error}
  <div class="rounded-lg bg-red-50 p-4 sm:rounded-lg">
    <div class="flex">
      <div class="flex-shrink-0">
        <!-- Heroicon name: x-circle -->
        <svg
          class="h-5 w-5 text-red-400"
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 20 20"
          fill="currentColor"
        >
          <path
            fill-rule="evenodd"
            d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z"
            clip-rule="evenodd"
          />
        </svg>
      </div>
      <div class="ml-3">
        <h3 class="text-sm leading-5 font-medium text-red-800">
          Javascript Expression Error
        </h3>
        <p class="text-sm text-red-800">
          Your javascript expression has been evaluated against the selected
          input and returned this error:
        </p>
        <div class="mt-2 text-sm leading-5 text-red-700">
          <ul class="list-disc pl-5">
            <li class="mt-1">{error?.message ?? ""}</li>
          </ul>
        </div>
      </div>
    </div>
  </div>
{/if}

<div class="bg-white shadow overflow-hidden sm:rounded-lg mt-3 mb-3">
  <div class="px-4 py-5 border-b border-gray-200 sm:px-6">
    <h3 class="text-lg leading-6 font-medium text-gray-900">
      Javascript Expression Guide
    </h3>
    <p class="mt-1 max-w-2xl text-sm leading-5 text-gray-500" />
  </div>
  <div>
    <dl>
      <div class="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
        <dt class="text-sm leading-5 font-medium text-gray-500">
          Instructions
        </dt>
        <dd
          class="mt-1 text-sm leading-5 text-gray-900 sm:mt-0 sm:col-span-2
            overflow-scroll"
        >
          <p>
            Write any valid Javascript above, your expression must evalulate to
            a boolean value. Your expression must conform to
            <strong>ES5</strong>, ES6 methods are not supported.
          </p>
          <br />
          <p>
            The current value of the input you are targeting will be supplied
            for your expression as the variable:
            <code><strong>value</strong></code>
          </p>
          <br />
          <p>
            You can see its current value below in the 'Current Value' column
          </p>
        </dd>
      </div>
      <div class="bg-white px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
        <dt class="text-sm leading-5 font-medium text-gray-500">
          Current Value
        </dt>
        <dd class="mt-1 text-sm leading-5 text-gray-900 sm:mt-0 sm:col-span-2">
          {#if !value}
            <p>Enter a value on your selected input to see its value here.</p>
          {:else}
            <textarea
              type="text"
              class="form-textarea block w-full transition duration-150
                ease-in-out sm:text-sm sm:leading-5"
              disabled
              readonly
              rows={5}
              value={JSON.stringify(value, null, 4)}
            />
          {/if}
        </dd>
      </div>
      <div class="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
        <dt class="text-sm leading-5 font-medium text-gray-500">
          String Example Expression
        </dt>
        <dd class="mt-1 text-sm leading-5 text-gray-900 sm:mt-0 sm:col-span-2">
          <p><code>value.indexOf('hello') !== -1</code></p>
        </dd>
      </div>
      <div class="bg-white px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
        <dt class="text-sm leading-5 font-medium text-gray-500">
          Checkbox Example Expression
        </dt>
        <dd class="mt-1 text-sm leading-5 text-gray-900 sm:mt-0 sm:col-span-2">
          <p><code>value['My Option'] != null</code></p>
        </dd>
      </div>
    </dl>
  </div>
</div>
