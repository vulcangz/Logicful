<script lang="typescript">
  import Link from "./Link.svelte";
  import type { SidebarOptions } from "@app/components/props/Props";

  export let options: SidebarOptions = [];

  let active = `group flex items-center px-3 py-2 text-sm leading-5 font-medium
            text-indigo-600 bg-indigo-50 border-l-4 border-indigo-600
            focus:outline-none focus:bg-indigo-100 transition ease-in-out
            duration-150`;

  let nonActive = `group flex items-center px-3 py-2 text-sm leading-5 font-medium
          text-gray-600 border-l-4 border-transparent hover:text-gray-900
          hover:bg-gray-50 focus:outline-none focus:text-gray-900
          focus:bg-gray-50 transition ease-in-out duration-150`;

  let activeIcon = `mr-3 h-6 w-6 text-indigo-500 transition ease-in-out
            duration-150`;

  let nonActiveIcon = `mr-3 h-6 w-6 text-gray-400 group-hover:text-gray-600
            group-focus:text-gray-600 transition ease-in-out duration-150`;

  function getClass(path: string) {
    return active;
  }

  function getIconClass(path: string) {
    return activeIcon;
  }
</script>

<div
  class="w-60 flex flex-col flex-grow border-r border-gray-200 pt-3 bg-white
    overflow-y-auto"
>
  <div class="flex-grow flex flex-col">
    <nav class="flex-1 bg-white space-y-1">
      {#each options as option}
        <Link href={option.href} class={getClass(option.href)}>
          <!-- Heroicon name: home -->
          {#if option.icon}
            <svg
              class={getIconClass(option.href)}
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d={option.icon}
              />
            </svg>
          {/if}
          {option.name}
        </Link>
      {/each}
    </nav>
  </div>
</div>
