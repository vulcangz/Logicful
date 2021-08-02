<script lang="ts">
  import Link from "../Link.svelte";
  import LogoFull from "../LogoFull.svelte";
  import { navigate } from "svelte-routing";
  import { onMount } from "svelte";
  import { emptyUser, me } from "@app/services/AuthService";
  import { nameToInitials } from "@app/util/Nav";

  let open = false;
  let profileOpen = false;
  let user = emptyUser;

  onMount(async () => {
    user = await me();
  });

  let activeClass =
    "ml-8 inline-flex items-center px-1 pt-1 border-b-2 border-indigo-500 text-sm font-medium leading-5 text-gray-900 focus:outline-none focus:border-indigo-700 transition duration-150 ease-in-out";

  let nonActiveClass =
    "ml-8 inline-flex items-center px-1 pt-1 border-b-2 border-transparent text-sm font-medium leading-5 text-gray-500 hover:text-gray-700 hover:border-gray-300 focus:outline-none focus:text-gray-700 focus:border-gray-300 transition duration-150 ease-in-out";

  let activeClassMobile = `block pl-3 pr-4 py-2 border-l-4 border-indigo-500 text-base
            font-medium text-indigo-700 bg-indigo-50 focus:outline-none
            focus:text-indigo-800 focus:bg-indigo-100 focus:border-indigo-700
            transition duration-150 ease-in-out`;

  let nonActiveClassMobile = `mt-1 block pl-3 pr-4 py-2 border-l-4 border-transparent
            text-base font-medium text-gray-600 hover:text-gray-800
            hover:bg-gray-50 hover:border-gray-300 focus:outline-none
            focus:text-gray-800 focus:bg-gray-50 focus:border-gray-300
            transition duration-150 ease-in-out`;

  function getClass(path: string, mobile: boolean = false): string {
    const active = window.location.pathname === path;
    if (mobile) {
      return active ? activeClassMobile : nonActiveClassMobile;
    }
    return active ? activeClass : nonActiveClass;
  }

  function onLogoClick() {
    navigate("/dashboard");
  }
</script>

<nav class="bg-white shadow">
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
    <div class="flex justify-between h-16">
      <div class="flex">
        <div
          class="flex-shrink-0 flex items-center cursor-pointer"
          on:click={onLogoClick}>
          <div class="block lg:hidden h-8 w-auto">
            <LogoFull />
          </div>
          <div class="hidden lg:block h-8 w-auto">
            <LogoFull />
          </div>
        </div>
        <div class="hidden sm:ml-6 sm:flex">
          <Link href="/dashboard" class={getClass('/dashboard')}>
            Dashboard
          </Link>
          <Link href="/folder" class={getClass('/folder')}>My Forms</Link>
          <Link href="/team" class={getClass('/team')}>My Team</Link>
        </div>
      </div>
      <div class="hidden sm:ml-6 sm:flex sm:items-center">
        <button
          class="p-1 border-2 border-transparent text-gray-400 rounded-full
            hover:text-gray-500 focus:outline-none focus:text-gray-500
            focus:bg-gray-100 transition duration-150 ease-in-out"
          aria-label="Notifications">
          <!-- Heroicon name: bell -->
          <svg
            class="h-6 w-6"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
          </svg>
        </button>

        <!-- Profile dropdown -->
        <div class="ml-3 relative">
          <span
            on:click={() => (profileOpen = !profileOpen)}
            class="inline-flex items-center justify-center h-10 w-10
              rounded-full bg-indigo-500 cursor-pointer">
            <span
              class="font-medium leading-none text-white">{nameToInitials(user.fullName)}</span>
          </span>
          {#if profileOpen}
            <div class="relative inline-block text-left">
              <div
                class="origin-top-right absolute right-0 mt-2 w-56 rounded-md
                  shadow-lg z-50">
                <div
                  class="rounded-md bg-white ring-1 ring-black ring-opacity-5"
                  role="menu"
                  aria-orientation="vertical"
                  aria-labelledby="options-menu">
                  <div class="px-4 py-3">
                    <p class="text-sm leading-5">Signed in as</p>
                    <p
                      class="text-sm leading-5 font-medium text-gray-900
                        truncate">
                      {user.displayName} ({user.email || 'Loading...'})
                    </p>
                  </div>
                  <div class="border-t border-gray-100" />
                  <div class="py-1">
                    <a
                      href="#"
                      class="block px-4 py-2 text-sm leading-5 text-gray-700
                        hover:bg-gray-100 hover:text-gray-900 focus:outline-none
                        focus:bg-gray-100 focus:text-gray-900"
                      role="menuitem">Account settings</a>
                    <a
                      href="#"
                      class="block px-4 py-2 text-sm leading-5 text-gray-700
                        hover:bg-gray-100 hover:text-gray-900 focus:outline-none
                        focus:bg-gray-100 focus:text-gray-900"
                      role="menuitem">Support</a>
                    <a
                      href="#"
                      class="block px-4 py-2 text-sm leading-5 text-gray-700
                        hover:bg-gray-100 hover:text-gray-900 focus:outline-none
                        focus:bg-gray-100 focus:text-gray-900"
                      role="menuitem">License</a>
                  </div>
                  <div class="border-t border-gray-100" />
                  <div class="py-1">
                    <a
                      href="/account/logout"
                      class="block w-full text-left px-4 py-2 text-sm leading-5
                        text-gray-700 hover:bg-gray-100 hover:text-gray-900
                        focus:outline-none focus:bg-gray-100 focus:text-gray-900"
                      role="menuitem">
                      Sign out
                    </a>
                  </div>
                </div>
              </div>
            </div>
          {/if}
        </div>
      </div>
      <div class="-mr-2 flex items-center sm:hidden">
        <!-- Mobile menu button -->
        <button
          on:click={() => (open = !open)}
          class="inline-flex items-center justify-center p-2 rounded-md
            text-gray-400 hover:text-gray-500 hover:bg-gray-100
            focus:outline-none focus:bg-gray-100 focus:text-gray-500 transition
            duration-150 ease-in-out"
          aria-label="Main menu"
          aria-expanded="false">
          <!-- Icon when menu is closed. -->
          <!--
            Heroicon name: menu

            Menu open: "hidden", Menu closed: "block"
          -->
          <svg
            class="block h-6 w-6"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M4 6h16M4 12h16M4 18h16" />
          </svg>
          <!-- Icon when menu is open. -->
          <!--
            Heroicon name: x

            Menu open: "block", Menu closed: "hidden"
          -->
          <svg
            class="hidden h-6 w-6"
            xmlns="http://www.w3.org/2000/svg"
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
  </div>

  <!--
    Mobile menu, toggle classes based on menu state.

    Menu open: "block", Menu closed: "hidden"
  -->
  {#if open}
    <div>
      <div class="pt-2 pb-3">
        <Link href="/dashboard" class={getClass('/dashboard', true)}>
          Dashboard
        </Link>
        <Link href="/folder" class={getClass('/folder', true)}>My Forms</Link>
      </div>
      <div class="pt-4 pb-3 border-t border-gray-200">
        <div class="flex items-center px-4">
          <div class="flex-shrink-0">
            <span
              class="inline-flex items-center justify-center h-10 w-10
                rounded-full bg-indigo-500 cursor-pointer">
              <span
                class="font-medium leading-none text-white">{nameToInitials(user.fullName)}</span>
            </span>
          </div>
          <div class="ml-3">
            <div class="text-base font-medium leading-6 text-gray-800">
              {user.displayName}
            </div>
            <div class="text-sm font-medium leading-5 text-gray-500">
              {user.email}
            </div>
          </div>
        </div>
        <div class="mt-3">
          <a
            href="#"
            class="block px-4 py-2 text-base font-medium text-gray-500
              hover:text-gray-800 hover:bg-gray-100 focus:outline-none
              focus:text-gray-800 focus:bg-gray-100 transition duration-150
              ease-in-out">Your Profile</a>
          <a
            href="#"
            class="mt-1 block px-4 py-2 text-base font-medium text-gray-500
              hover:text-gray-800 hover:bg-gray-100 focus:outline-none
              focus:text-gray-800 focus:bg-gray-100 transition duration-150
              ease-in-out">Settings</a>
          <a
            href="/account/logout"
            class="mt-1 block px-4 py-2 text-base font-medium text-gray-500
              hover:text-gray-800 hover:bg-gray-100 focus:outline-none
              focus:text-gray-800 focus:bg-gray-100 transition duration-150
              ease-in-out">Sign out</a>
        </div>
      </div>
    </div>
  {/if}
</nav>
