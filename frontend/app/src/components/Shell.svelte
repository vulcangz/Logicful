<script lang="ts">
  import AlphaBanner from "./AlphaBanner.svelte";

  import Navbar from "./layout/Navbar.svelte";
  import type { SidebarOptions } from "./props/Props";
  import ShellSidebar from "./ShellSidebar.svelte";

  export let header: string = "";
  export let sidebar: SidebarOptions = [];
  export let showSidebar: boolean = false;
  let open = false;
</script>

<AlphaBanner />
<Navbar />

<div class="h-screen flex overflow-hidden bg-white">
  {#if sidebar.length > 0 || showSidebar}
    <div class:sm:block={open} class:hidden={!open} class="md:block">
      <slot name="sidebar">
        <ShellSidebar options={sidebar} />
      </slot>
    </div>
  {/if}
  <div class="flex flex-col w-0 flex-1 overflow-hidden">
    <main
      class="flex-1 relative z-0 overflow-y-auto focus:outline-none"
      tabindex="0"
    >
      <div class="pt-2 pb-6 md:py-6">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 md:px-8">
          <slot name="header">
            <h1 class="text-2xl font-semibold text-gray-900">{header}</h1>
          </slot>
        </div>
        <div class="max-w-7xl mx-auto px-4 sm:px-6 md:px-8">
          <div class="py-4">
            <slot />
          </div>
        </div>
      </div>
    </main>
  </div>
</div>
