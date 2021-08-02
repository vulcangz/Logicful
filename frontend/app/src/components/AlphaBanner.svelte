<script lang="typescript">
  import { onMount } from "svelte";
  let show = false;
  onMount(() => {
    show = sessionStorage.getItem("hide_alpha_banner") == null;
  });

  function close() {
    sessionStorage.setItem("hide_alpha_banner", "true");
    show = false;
  }
</script>

{#if show}
  <div class="bg-indigo-600">
    <div class="max-w-screen-xl mx-auto py-3 px-3 sm:px-6 lg:px-8">
      <div class="flex items-center justify-between flex-wrap">
        <div class="w-0 flex-1 flex items-center">
          <span class="flex p-2 rounded-lg bg-indigo-800">
            <svg
              class="h-6 w-6 text-white"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M11 5.882V19.24a1.76 1.76 0 01-3.417.592l-2.147-6.15M18 13a3 3 0 100-6M5.436 13.683A4.001 4.001 0 017 6h1.832c4.1 0 7.625-1.234 9.168-3v14c-1.543-1.766-5.067-3-9.168-3H7a3.988 3.988 0 01-1.564-.317z"
              />
            </svg>
          </span>
          <p class="ml-3 font-medium text-white truncate">
            <span> Logicful is in very early alpha stage. </span>
            <span class="md:inline">
              Important features may not be implemented yet or work as expected.
            </span>
            <span class="md:inline">
              <a href="/faq" class="underline">More Info</a>
            </span>
          </p>
        </div>
        <div class="order-2 flex-shrink-0 sm:order-3 sm:ml-3">
          <button
            type="button"
            on:click={close}
            class="-mr-1 flex p-2 rounded-md hover:bg-indigo-500
              focus:outline-none focus:bg-indigo-500 sm:-mr-2 transition
              ease-in-out duration-150"
            aria-label="Dismiss"
          >
            <!-- Heroicon name: x -->
            <svg
              class="h-6 w-6 text-white"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M6 18L18 6M6 6l12 12"
              />
            </svg>
          </button>
        </div>
      </div>
    </div>
  </div>
{/if}
