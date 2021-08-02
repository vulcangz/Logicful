<script lang="typescript">
  export let onClose: () => void;
  export let title = "";
  export let description = "";
</script>

<div class="fixed inset-0 overflow-hidden" on:click={onClose}>
  <div class="absolute inset-0 overflow-hidden">
    <section class="absolute inset-y-0 right-0 pl-10 max-w-full flex">
      <div class="w-screen max-w-2xl">
        <div class="h-full flex flex-col bg-white shadow-xl overflow-y-scroll">
          <header class="space-y-1 py-6 px-4 bg-indigo-700 sm:px-6">
            <div class="flex items-center justify-between space-x-3">
              <h2 class="text-lg leading-7 font-medium text-white">{title}</h2>
              <div class="h-7 flex items-center">
                <button
                  aria-label="Close panel"
                  on:click={onClose}
                  class="text-indigo-200 hover:text-white transition ease-in-out
                    duration-150">
                  <!-- Heroicon name: x -->
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
          <div
            class="relative flex-1 py-6 px-4 sm:px-6"
            on:click|stopPropagation>
            <!-- Replace with your content -->
            <div class="absolute inset-0 py-6 px-4 sm:px-6">
              <slot />
            </div>
            <!-- /End replace -->
          </div>
        </div>
      </div>
    </section>
  </div>
</div>
