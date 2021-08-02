<script lang="typescript">
  import Button from "@app/components/Button.svelte";
  import type { IForm } from "@app/models/IForm";

  export let form: IForm;
  export let onMove: (form: IForm) => any;
  export let onDelete: (form: IForm) => any;
  let showActions = false;
</script>

<div class="lg:flex lg:items-center lg:justify-between">
  <div class="flex-1 min-w-0">
    <h2
      class="font-bold leading-7 text-gray-900 sm:leading-9 sm:truncate
        cursor-pointer hover:text-indigo-500">
      <a
        href={`/form/builder?formId=${form.id}`}
        target="_blank">{form.title}</a>
    </h2>
    <div class="mt-1 flex flex-col sm:mt-0 sm:flex-row sm:flex-wrap">
      {#if form.unreadSubmissions > 0}
        <div
          class="mt-2 flex items-center text-sm leading-5 text-gray-500 sm:mr-6">
          <span
            class="inline-flex items-center px-3 py-0.5 rounded-full text-sm
              font-medium leading-5 bg-purple-100 text-purple-800
              hover:font-bold">
            <a href={'/form/submissions?formId=' + form.id} target="_blank">
              {form.unreadSubmissions} Unread</a>
          </span>
        </div>
      {/if}
      <div
        class="mt-2 flex items-center text-sm leading-5 text-gray-700 sm:mr-6
          cursor-pointer hover:text-indigo-500">
        <svg
          class="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400
            hover:text-indigo-500"
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor">
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4" />
        </svg>
        <a
          href={'/form/submissions?formId=' + form.id}
          target="_blank">{form.submissionCount} Submissions</a>
      </div>
      <div
        class="mt-2 flex items-center text-sm leading-5 text-gray-500 sm:mr-6">
        <svg
          class="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400"
          viewBox="0 0 20 20"
          fill="currentColor">
          <path
            fill-rule="evenodd"
            d="M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0 00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z"
            clip-rule="evenodd" />
        </svg> Updated {new Date(form.changeDate ?? new Date()).toLocaleDateString()}
      </div>
    </div>
  </div>
  <div class="mt-5 flex lg:mt-0 lg:ml-4">
    <span class="sm:block shadow-sm rounded-md">
      <Button
        href={`/form/preview?formId=${form.id}`}
        type="white"
        size="small"
        hasIcon={true}>
        <!-- Heroicon name: pencil -->
        <svg
          class="-ml-1 mr-1 h-5 w-5 text-gray-500"
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor">
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
        </svg> Preview
      </Button>
    </span>

    <span class="sm:block ml-3 shadow-sm rounded-md">
      <Button
        href={`/form/builder?formId=${form.id}`}
        type="white"
        size="small"
        hasIcon={true}>
        <!-- Heroicon name: link -->
        <svg
          class="-ml-1 mr-1 h-5 w-5 text-gray-500"
          viewBox="0 0 20 20"
          fill="currentColor">
          <path
            d="M13.586 3.586a2 2 0 112.828 2.828l-.793.793-2.828-2.828.793-.793zM11.379 5.793L3 14.172V17h2.828l8.38-8.379-2.83-2.828z" />
        </svg> Edit
      </Button>
    </span>

    <!-- Dropdown -->
    <span class="ml-3 relative shadow-sm rounded-md">
      <Button
        hasIcon={true}
        onClick={() => (showActions = !showActions)}
        type="white"
        size="small"
        aria-haspopup="true">
        Actions
        <!-- Heroicon name: chevron-down -->
        <svg
          class="-mr-1 ml-2 h-5 w-5 text-gray-500"
          viewBox="0 0 20 20"
          fill="currentColor">
          <path
            fill-rule="evenodd"
            d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z"
            clip-rule="evenodd" />
        </svg>
      </Button>

      <div
        class:hidden={!showActions}
        class="origin-top-right absolute right-0 mt-2 -mr-1 w-48 rounded-md
          shadow-lg"
        aria-labelledby="mobile-menu"
        role="menu">
        <div class="py-1 rounded-md bg-white ring-1 ring-black ring-opacity-5">
          <a
            href="javascript:void(0)"
            on:click={() => onMove(form)}
            class="block px-4 py-2 text-sm leading-5 text-gray-700
              hover:bg-gray-100 focus:outline-none focus:bg-gray-100 transition
              duration-150 ease-in-out"
            role="menuitem">Move To New Folder</a>
        </div>
        <div class="py-1 rounded-md bg-white ring-1 ring-black ring-opacity-5">
          <a
            href="javascript:void(0)"
            on:click={() => onDelete(form)}
            class="block px-4 py-2 text-sm leading-5 text-red-600
              hover:bg-gray-100 focus:outline-none focus:bg-gray-100 transition
              duration-150 ease-in-out"
            role="menuitem">Delete</a>
        </div>
      </div>
    </span>
  </div>
</div>

<svelte:body on:click={() => (showActions = false)} />
