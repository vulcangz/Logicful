<script lang="typescript">
  import { dispatch } from "@app/event/EventBus";

  export let title: string;
  export let description: string;
  function onClose() {
    dispatch("unselect_field", {});
  }
</script>

<header class="space-y-1 pb-6 pt-4 mb-4 px-4 bg-indigo-700 sm:px-6">
  <div class="flex items-center justify-between space-x-3">
    <h2 class="text-lg leading-7 font-medium text-white">{title}</h2>
    <div class="h-7 flex items-center">
      <button
        on:click={onClose}
        aria-label="Close panel"
        class="text-indigo-200 hover:text-white transition ease-in-out
          duration-150">
        <svg
          class="h-6 w-6"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor">
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M6 18L18 6M6 6l12 12" />
        </svg>
      </button>
    </div>
  </div>
  <div>
    <p class="text-sm leading-5 text-indigo-300">{description}</p>
  </div>
</header>
