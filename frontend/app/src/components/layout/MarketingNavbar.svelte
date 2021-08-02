<script lang="typescript">
  import LogoMarketing from "@app/components/LogoMarketing.svelte";
  import Button from "../Button.svelte";
  import LogoFullAlternative from "../LogoFullAlternative.svelte";
  let open = false;
</script>

<div class="max-w-screen-xl mx-auto px-4 sm:px-6">
  <nav
    class="relative flex items-center justify-between sm:h-10 md:justify-center">
    <div class="flex items-center flex-1 md:absolute md:inset-y-0 md:left-0">
      <div class="flex items-center justify-between w-full md:w-auto">
        <a href="/" aria-label="Home" class="h-10">
          <LogoMarketing />
        </a>
        <div
          class="-mr-2 flex items-center md:hidden"
          on:click={() => (open = true)}>
          <button
            type="button"
            class="inline-flex items-center justify-center p-2 rounded-md
              text-gray-400 hover:text-gray-500 hover:bg-gray-100
              focus:outline-none focus:bg-gray-100 focus:text-gray-500
              transition duration-150 ease-in-out"
            id="main-menu"
            aria-label="Main menu"
            aria-haspopup="true">
            <svg
              class="h-6 w-6"
              stroke="currentColor"
              fill="none"
              viewBox="0 0 24 24">
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M4 6h16M4 12h16M4 18h16" />
            </svg>
          </button>
        </div>
      </div>
    </div>
    <div class="hidden md:flex md:space-x-10">
      <a
        href="/features"
        class="font-medium text-gray-500 hover:text-gray-900 transition
          duration-150 ease-in-out">Features</a>
      <a
        href="/pricing"
        class="font-medium text-gray-500 hover:text-gray-900 transition
          duration-150 ease-in-out">Pricing</a>
      <a
        href="/faq"
        class="font-medium text-gray-500 hover:text-gray-900 transition
          duration-150 ease-in-out">FAQ</a>
        <a
                href="https://trello.com/b/T0c1HOPI"
                target="_blank"
                rel="noopener noreferrer"
                class="font-medium text-gray-500 hover:text-gray-900 transition
          duration-150 ease-in-out">Roadmap</a>
    </div>
    <div
      class="hidden md:absolute md:flex md:items-center md:justify-end
        md:inset-y-0 md:right-0">
      <span class="inline-flex rounded-md shadow">
        <Button href="/form/builder?formId=demo" type="white">Try Demo</Button>
      </span>
      <span class="inline-flex rounded-md shadow ml-3">
        <Button href="/account/login" type="primary">Log in</Button>
      </span>
    </div>
  </nav>
</div>

<!--
    Mobile menu, show/hide based on menu open state.

    Entering: "duration-150 ease-out"
      From: "opacity-0 scale-95"
      To: "opacity-100 scale-100"
    Leaving: "duration-100 ease-in"
      From: "opacity-100 scale-100"
      To: "opacity-0 scale-95"
  -->
{#if open}
  <div
    class="absolute top-0 inset-x-0 p-2 transition transform origin-top-right
      md:hidden z-50 overflow-visible">
    <div class="rounded-lg shadow-md">
      <div
        class="rounded-lg bg-white ring-1 ring-black ring-opacity-5 overflow-hidden"
        role="menu"
        aria-orientation="vertical"
        aria-labelledby="main-menu">
        <div class="px-5 pt-4 flex items-center justify-between">
          <a class="h-5" href="/">
            <LogoFullAlternative />
          </a>
          <div class="-mr-2">
            <button
              on:click={() => (open = false)}
              type="button"
              class="inline-flex items-center justify-center p-2 rounded-md
                text-gray-400 hover:text-gray-500 hover:bg-gray-100
                focus:outline-none focus:bg-gray-100 focus:text-gray-500
                transition duration-150 ease-in-out"
              aria-label="Close menu">
              <svg
                class="h-6 w-6"
                stroke="currentColor"
                fill="none"
                viewBox="0 0 24 24">
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
        </div>
        <div class="px-2 pt-2 pb-3">
          <a
            href="/features"
            class="block px-3 py-2 rounded-md text-base font-medium
              text-gray-700 hover:text-gray-900 hover:bg-gray-50
              focus:outline-none focus:text-gray-900 focus:bg-gray-50 transition
              duration-150 ease-in-out"
            role="menuitem">Features</a>
          <a
            href="/pricing"
            class="mt-1 block px-3 py-2 rounded-md text-base font-medium
              text-gray-700 hover:text-gray-900 hover:bg-gray-50
              focus:outline-none focus:text-gray-900 focus:bg-gray-50 transition
              duration-150 ease-in-out"
            role="menuitem">Pricing</a>
          <a
            href="/faq"
            class="mt-1 block px-3 py-2 rounded-md text-base font-medium
              text-gray-700 hover:text-gray-900 hover:bg-gray-50
              focus:outline-none focus:text-gray-900 focus:bg-gray-50 transition
              duration-150 ease-in-out"
            role="menuitem">FAQ</a>
          <a
                  href="https://trello.com/b/T0c1HOPI"
                  target="_blank"
                  rel="noopener noreferrer"
                  class="mt-1 block px-3 py-2 rounded-md text-base font-medium
              text-gray-700 hover:text-gray-900 hover:bg-gray-50
              focus:outline-none focus:text-gray-900 focus:bg-gray-50 transition
              duration-150 ease-in-out"
                  role="menuitem">Roadmap</a>
        </div>
        <div>
          <a
            href="/account/login"
            class="block w-full px-5 py-3 text-center font-medium
              text-indigo-600 bg-gray-50 hover:bg-gray-100 hover:text-indigo-700
              focus:outline-none focus:bg-gray-100 focus:text-indigo-700
              transition duration-150 ease-in-out"
            role="menuitem">
            Log in
          </a>
        </div>
      </div>
    </div>
  </div>
{/if}
