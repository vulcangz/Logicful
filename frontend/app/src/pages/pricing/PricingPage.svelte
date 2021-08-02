<script lang="ts">
  import AlphaBanner from "@app/components/AlphaBanner.svelte";

  import MarketingNavbar from "@app/components/layout/MarketingNavbar.svelte";
</script>

<!--
  Tailwind UI components require Tailwind CSS v1.8 and the @tailwindcss/ui plugin.
  Read the documentation to get started: https://tailwindui.com/documentation
-->
<AlphaBanner />

<div class="bg-white">
  <div class="relative bg-white">
    <div class="max-w-7xl mx-auto px-4 sm:px-6">
      <div class="relative pt-6 pb-12 lg:pb-20">
        <MarketingNavbar />
      </div>
    </div>

    <!--
        Mobile menu, show/hide based on mobile menu state.

        Entering: "duration-200 ease-out"
          From: "opacity-0 scale-95"
          To: "opacity-100 scale-100"
        Leaving: "duration-100 ease-in"
          From: "opacity-100 scale-100"
          To: "opacity-0 scale-95"
      -->
    <div
      class="absolute z-10 top-0 inset-x-0 p-2 transition transform
        origin-top-right md:hidden">
      <div class="rounded-lg shadow-lg">
        <div class="bg-white rounded-lg ring-1 ring-black ring-opacity-5 divide-y-2 divide-gray-50">
          <div class="pt-5 pb-6 px-5 space-y-6">
            <div class="flex items-center justify-between">
              <div>
                <img
                  class="h-8 w-auto"
                  src="https://tailwindui.com/img/logos/workflow-mark-purple-600.svg"
                  alt="Workflow" />
              </div>
              <div class="-mr-2">
                <button
                  type="button"
                  class="rounded-md p-2 inline-flex items-center justify-center
                    text-gray-400 hover:text-gray-500 hover:bg-gray-100
                    focus:outline-none focus:bg-gray-100 focus:text-gray-500
                    transition duration-150 ease-in-out">
                  <!-- Heroicon name: x -->
                  <svg
                    class="h-6 w-6"
                    aria-hidden="true"
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
            <div>
              <nav class="grid gap-y-8">
                <a
                  href="#"
                  class="-m-3 rounded-md p-3 flex items-center space-x-3
                    hover:bg-gray-50 transition ease-in-out duration-150">
                  <!-- Heroicon name: chart-bar -->
                  <svg
                    class="flex-shrink-0 h-6 w-6 text-purple-600"
                    aria-hidden="true"
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor">
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
                  </svg>
                  <div class="text-base leading-6 font-medium text-gray-900">
                    Analytics
                  </div>
                </a>
                <a
                  href="#"
                  class="-m-3 rounded-md p-3 flex items-center space-x-3
                    hover:bg-gray-50 transition ease-in-out duration-150">
                  <!-- Heroicon name: cursor-click -->
                  <svg
                    class="flex-shrink-0 h-6 w-6 text-purple-600"
                    aria-hidden="true"
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor">
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M15 15l-2 5L9 9l11 4-5 2zm0 0l5 5M7.188 2.239l.777 2.897M5.136 7.965l-2.898-.777M13.95 4.05l-2.122 2.122m-5.657 5.656l-2.12 2.122" />
                  </svg>
                  <div class="text-base leading-6 font-medium text-gray-900">
                    Engagement
                  </div>
                </a>
                <a
                  href="#"
                  class="-m-3 rounded-md p-3 flex items-center space-x-3
                    hover:bg-gray-50 transition ease-in-out duration-150">
                  <!-- Heroicon name: shield-check -->
                  <svg
                    class="flex-shrink-0 h-6 w-6 text-purple-600"
                    aria-hidden="true"
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor">
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
                  </svg>
                  <div class="text-base leading-6 font-medium text-gray-900">
                    Security
                  </div>
                </a>
                <a
                  href="#"
                  class="-m-3 rounded-md p-3 flex items-center space-x-3
                    hover:bg-gray-50 transition ease-in-out duration-150">
                  <!-- Heroicon name: view-grid -->
                  <svg
                    class="flex-shrink-0 h-6 w-6 text-purple-600"
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor">
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z" />
                  </svg>
                  <div class="text-base leading-6 font-medium text-gray-900">
                    Integrations
                  </div>
                </a>
                <a
                  href="#"
                  class="-m-3 rounded-md p-3 flex items-center space-x-3
                    hover:bg-gray-50 transition ease-in-out duration-150">
                  <!-- Heroicon name: refresh -->
                  <svg
                    class="flex-shrink-0 h-6 w-6 text-purple-600"
                    aria-hidden="true"
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor">
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                  </svg>
                  <div class="text-base leading-6 font-medium text-gray-900">
                    Automations
                  </div>
                </a>
              </nav>
            </div>
          </div>
          <div class="py-6 px-5 space-y-6">
            <div class="grid grid-cols-2 gap-y-4 gap-x-8">
              <a
                href="#"
                class="text-base leading-6 font-medium text-gray-900
                  hover:text-gray-700 transition ease-in-out duration-150">
                Pricing
              </a>
              <a
                href="#"
                class="text-base leading-6 font-medium text-gray-900
                  hover:text-gray-700 transition ease-in-out duration-150">
                Docs
              </a>
              <a
                href="#"
                class="text-base leading-6 font-medium text-gray-900
                  hover:text-gray-700 transition ease-in-out duration-150">
                Enterprise
              </a>
              <a
                href="#"
                class="text-base leading-6 font-medium text-gray-900
                  hover:text-gray-700 transition ease-in-out duration-150">
                Blog
              </a>
              <a
                href="#"
                class="text-base leading-6 font-medium text-gray-900
                  hover:text-gray-700 transition ease-in-out duration-150">
                Help Center
              </a>
              <a
                href="#"
                class="text-base leading-6 font-medium text-gray-900
                  hover:text-gray-700 transition ease-in-out duration-150">
                Guides
              </a>
              <a
                href="#"
                class="text-base leading-6 font-medium text-gray-900
                  hover:text-gray-700 transition ease-in-out duration-150">
                Security
              </a>
              <a
                href="#"
                class="text-base leading-6 font-medium text-gray-900
                  hover:text-gray-700 transition ease-in-out duration-150">
                Events
              </a>
            </div>
            <div class="space-y-6">
              <a
                href="#"
                class="w-full bg-purple-100 border border-transparent rounded-md
                  py-2 px-4 flex items-center justify-center text-base leading-6
                  font-medium text-purple-600 hover:bg-purple-50
                  focus:outline-none focus:border-purple-200
                  focus:ring-purple active:bg-purple-200 transition
                  ease-in-out duration-150">
                Sign up
              </a>
              <p
                class="text-center text-base leading-6 font-medium text-gray-500">
                Existing customer? <a href="#" class="text-purple-600
                    hover:text-purple-500 transition ease-in-out duration-150"> Sign
                  in </a>
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- Pricing with four tiers and toggle -->
  <div class="bg-gradient-to-b from-white to-gray-50">
    <div class="max-w-screen-xl mx-auto pt-10 px-4 sm:px-6 lg:px-8">
      <div class="sm:flex sm:flex-col sm:align-center">
        <h1
          class="text-5xl leading-none font-extrabold text-gray-900
            sm:text-center">
          Pricing Plans
        </h1>
        <p class="mt-5 text-xl leading-7 text-gray-500 sm:text-center">
          Start building for free, then as you need more form features, upgrade.
          Account plans unlock additional features.
        </p>
        =
      </div>
      <div
        class="mt-12 space-y-4 sm:mt-16 sm:space-y-0 sm:grid sm:grid-cols-2
          sm:gap-6 md:grid-cols-3 lg:max-w-3xl lg:mx-auto xl:max-w-none xl:mx-0
          xl:grid-cols-3">
        <div
          class="border border-gray-200 rounded-lg shadow-sm divide-y
            divide-gray-200">
          <div class="p-6">
            <h2 class="text-lg leading-6 font-medium text-gray-900">Free</h2>
            <p class="mt-4 text-sm leading-5 text-gray-500">
              Basic form features and limited workflows
            </p>
            <p class="mt-8">
              <span
                class="text-4xl leading-10 font-extrabold text-gray-900">$0</span>
              <span
                class="text-base leading-6 font-medium text-gray-500">/mo</span>
            </p>
            <button
              type="button"
              class="mt-8 w-full bg-purple-600 border border-transparent
                rounded-md py-2 text-sm leading-5 font-semibold text-white
                hover:bg-purple-500 focus:outline-none focus:border-purple-700
                focus:ring-purple active:bg-purple-700 transition
                duration-150 ease-in-out">Sign up</button>
          </div>
          <div class="pt-6 pb-8 px-6">
            <h3
              class="text-xs leading-4 font-medium text-gray-900 tracking-wide
                uppercase">
              What's included
            </h3>
            <ul class="mt-6 space-y-4">
              <li class="flex space-x-3">
                <!-- Heroicon name: check -->
                <svg
                  class="flex-shrink-0 h-5 w-5 text-green-500"
                  aria-hidden="true"
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 20 20"
                  fill="currentColor">
                  <path
                    fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd" />
                </svg>
                <span class="text-sm leading-5 text-gray-500">5 forms per month</span>
              </li>

              <li class="flex space-x-3">
                <!-- Heroicon name: check -->
                <svg
                  class="flex-shrink-0 h-5 w-5 text-green-500"
                  aria-hidden="true"
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 20 20"
                  fill="currentColor">
                  <path
                    fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd" />
                </svg>
                <span class="text-sm leading-5 text-gray-500">Unlimited
                  submissions</span>
              </li>

              <li class="flex space-x-3">
                <!-- Heroicon name: check -->
                <svg
                  class="flex-shrink-0 h-5 w-5 text-green-500"
                  aria-hidden="true"
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 20 20"
                  fill="currentColor">
                  <path
                    fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd" />
                </svg>
                <span class="text-sm leading-5 text-gray-500">1 team member</span>
              </li>
              <li class="flex space-x-3">
                <!-- Heroicon name: check -->
                <svg
                  class="flex-shrink-0 h-5 w-5 text-green-500"
                  aria-hidden="true"
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 20 20"
                  fill="currentColor">
                  <path
                    fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd" />
                </svg>
                <span class="text-sm leading-5 text-gray-500">Mobile friendly
                  forms</span>
              </li>
            </ul>
          </div>
        </div>

        <div
          class="border border-gray-200 rounded-lg shadow-sm divide-y
            divide-gray-200">
          <div class="p-6">
            <h2 class="text-lg leading-6 font-medium text-gray-900">Silver</h2>
            <p class="mt-4 text-sm leading-5 text-gray-500">
              All the form basics plus additional workflows and the ability to
              add team members
            </p>
            <p class="mt-8">
              <span
                class="text-4xl leading-10 font-extrabold text-gray-900">$24</span>
              <span
                class="text-base leading-6 font-medium text-gray-500">/mo</span>
            </p>
            <button
              type="button"
              class="mt-8 w-full bg-purple-600 border border-transparent
                rounded-md py-2 text-sm leading-5 font-semibold text-white
                hover:bg-purple-500 focus:outline-none focus:border-purple-700
                focus:ring-purple active:bg-purple-700 transition
                duration-150 ease-in-out">Buy Silver</button>
          </div>
          <div class="pt-6 pb-8 px-6">
            <h3
              class="text-xs leading-4 font-medium text-gray-900 tracking-wide
                uppercase">
              What's included
            </h3>
            <ul class="mt-6 space-y-4">
              <li class="flex space-x-3">
                <!-- Heroicon name: check -->
                <svg
                  class="flex-shrink-0 h-5 w-5 text-green-500"
                  aria-hidden="true"
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 20 20"
                  fill="currentColor">
                  <path
                    fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd" />
                </svg>
                <span class="text-sm leading-5 text-gray-500">100 forms per
                  month</span>
              </li>

              <li class="flex space-x-3">
                <!-- Heroicon name: check -->
                <svg
                  class="flex-shrink-0 h-5 w-5 text-green-500"
                  aria-hidden="true"
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 20 20"
                  fill="currentColor">
                  <path
                    fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd" />
                </svg>
                <span class="text-sm leading-5 text-gray-500">Unlimited
                  submissions
                </span>
              </li>

              <li class="flex space-x-3">
                <!-- Heroicon name: check -->
                <svg
                  class="flex-shrink-0 h-5 w-5 text-green-500"
                  aria-hidden="true"
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 20 20"
                  fill="currentColor">
                  <path
                    fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd" />
                </svg>
                <span class="text-sm leading-5 text-gray-500"><b>2</b> team members</span>
              </li>

              <li class="flex space-x-3">
                <!-- Heroicon name: check -->
                <svg
                  class="flex-shrink-0 h-5 w-5 text-green-500"
                  aria-hidden="true"
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 20 20"
                  fill="currentColor">
                  <path
                    fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd" />
                </svg>
                <span class="text-sm leading-5 text-gray-500">Mobile friendly
                  forms
                </span>
              </li>

              <li class="flex space-x-3">
                <!-- Heroicon name: check -->
                <svg
                  class="flex-shrink-0 h-5 w-5 text-green-500"
                  aria-hidden="true"
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 20 20"
                  fill="currentColor">
                  <path
                    fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd" />
                </svg>
                <span class="text-sm leading-5 text-gray-500">Email
                  Notifications
                </span>
              </li>
            </ul>
          </div>
        </div>

        <div
          class="border border-gray-200 rounded-lg shadow-sm divide-y
            divide-gray-200">
          <div class="p-6">
            <h2 class="text-lg leading-6 font-medium text-gray-900">Gold</h2>
            <p class="mt-4 text-sm leading-5 text-gray-500">
              All the basics as well as additional workflows and advanced
              integrations
            </p>
            <p class="mt-8">
              <span
                class="text-4xl leading-10 font-extrabold text-gray-900">$32</span>
              <span
                class="text-base leading-6 font-medium text-gray-500">/mo</span>
            </p>
            <button
              type="button"
              class="mt-8 w-full bg-purple-600 border border-transparent
                rounded-md py-2 text-sm leading-5 font-semibold text-white
                hover:bg-purple-500 focus:outline-none focus:border-purple-700
                focus:ring-purple active:bg-purple-700 transition
                duration-150 ease-in-out">Buy Gold</button>
          </div>
          <div class="pt-6 pb-8 px-6">
            <h3
              class="text-xs leading-4 font-medium text-gray-900 tracking-wide
                uppercase">
              What's included
            </h3>
            <ul class="mt-6 space-y-4">
              <li class="flex space-x-3">
                <!-- Heroicon name: check -->
                <svg
                  class="flex-shrink-0 h-5 w-5 text-green-500"
                  aria-hidden="true"
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 20 20"
                  fill="currentColor">
                  <path
                    fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd" />
                </svg>
                <span class="text-sm leading-5 text-gray-500">1,000 forms per
                  month
                </span>
              </li>

              <li class="flex space-x-3">
                <!-- Heroicon name: check -->
                <svg
                  class="flex-shrink-0 h-5 w-5 text-green-500"
                  aria-hidden="true"
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 20 20"
                  fill="currentColor">
                  <path
                    fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd" />
                </svg>
                <span class="text-sm leading-5 text-gray-500">Unlimited
                  Submissions</span>
              </li>

              <li class="flex space-x-3">
                <!-- Heroicon name: check -->
                <svg
                  class="flex-shrink-0 h-5 w-5 text-green-500"
                  aria-hidden="true"
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 20 20"
                  fill="currentColor">
                  <path
                    fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd" />
                </svg>
                <span class="text-sm leading-5 text-gray-500"><b>10</b> team members</span>
              </li>

              <li class="flex space-x-3">
                <!-- Heroicon name: check -->
                <svg
                  class="flex-shrink-0 h-5 w-5 text-green-500"
                  aria-hidden="true"
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 20 20"
                  fill="currentColor">
                  <path
                    fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd" />
                </svg>
                <span class="text-sm leading-5 text-gray-500">Mobile friendly
                  forms</span>
              </li>

              <li class="flex space-x-3">
                <!-- Heroicon name: check -->
                <svg
                  class="flex-shrink-0 h-5 w-5 text-green-500"
                  aria-hidden="true"
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 20 20"
                  fill="currentColor">
                  <path
                    fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd" />
                </svg>
                <span class="text-sm leading-5 text-gray-500">Customizable
                  Notifications</span>
              </li>

              <li class="flex space-x-3">
                <!-- Heroicon name: check -->
                <svg
                  class="flex-shrink-0 h-5 w-5 text-green-500"
                  aria-hidden="true"
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 20 20"
                  fill="currentColor">
                  <path
                    fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd" />
                </svg>
                <span class="text-sm leading-5 text-gray-500">Advanced
                  Integreations, read more here</span>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>

    <!-- Feature list -->
    <div class="max-w-screen-xl mx-auto py-16 px-4 sm:px-6 lg:py-24 lg:px-8">
      <div class="max-w-3xl mx-auto text-center">
        <h2 class="text-3xl leading-9 font-extrabold text-gray-900">
          All-in-one platform
        </h2>
        <p class="mt-4 text-lg leading-7 text-gray-500">
          Everything you need for your form needs from start to finish. From
          creating and editting to setting up workflows and integrations to
          viewing submissions and easily interating results into your own
          workflows.
        </p>
      </div>
      <dl
        class="mt-12 space-y-10 sm:space-y-0 sm:grid sm:grid-cols-2 sm:gap-x-6
          sm:gap-y-12 lg:grid-cols-4 lg:gap-x-8">
        <div class="flex space-x-3">
          <!-- Heroicon name: check -->
          <svg
            class="flex-shrink-0 h-6 w-6 text-green-500"
            aria-hidden="true"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M5 13l4 4L19 7" />
          </svg>
          <div class="space-y-2">
            <dt class="text-lg leading-6 font-medium text-gray-900">
              Invite team members
            </dt>
            <dd class="flex space-x-3 lg:py-0 lg:pb-4">
              <span class="text-base leading-6 text-gray-500">Share you forms
                with team members so forms can be created and edited by more
                than one person</span>
            </dd>
          </div>
        </div>
        <div class="flex space-x-3">
          <!-- Heroicon name: check -->
          <svg
            class="flex-shrink-0 h-6 w-6 text-green-500"
            aria-hidden="true"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M5 13l4 4L19 7" />
          </svg>
          <div class="space-y-2">
            <dt class="text-lg leading-6 font-medium text-gray-900">
              Notifications
            </dt>
            <dd class="flex space-x-3">
              <span class="text-base leading-6 text-gray-500">Workflows include
                recieving emails upon form submission</span>
            </dd>
          </div>
        </div>
        <div class="flex space-x-3">
          <!-- Heroicon name: check -->
          <svg
            class="flex-shrink-0 h-6 w-6 text-green-500"
            aria-hidden="true"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M5 13l4 4L19 7" />
          </svg>
          <div class="space-y-2">
            <dt class="text-lg leading-6 font-medium text-gray-900">
              List view
            </dt>
            <dd class="flex space-x-3">
              <span class="text-base leading-6 text-gray-500">Etiam cras augue
                ornare pretium sit malesuada morbi orci, venenatis. Dictum
                lacus.</span>
            </dd>
          </div>
        </div>
        <div class="flex space-x-3">
          <!-- Heroicon name: check -->
          <svg
            class="flex-shrink-0 h-6 w-6 text-green-500"
            aria-hidden="true"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M5 13l4 4L19 7" />
          </svg>
          <div class="space-y-2">
            <dt class="text-lg leading-6 font-medium text-gray-900">
              Mobile friendly forms
            </dt>
            <dd class="flex space-x-3">
              <span class="text-base leading-6 text-gray-500">Anyone from any
                devie can easily view and submit forms</span>
            </dd>
          </div>
        </div>
      </dl>
    </div>
  </div>

  <!-- Logo cloud on brand -->
  <div class="bg-purple-600">
    <div class="max-w-screen-xl mx-auto py-16 px-4 sm:py-20 sm:px-6 lg:px-8">
      <div class="lg:space-y-10">
        <h2 class="text-3xl leading-9 font-extrabold text-white">
          The world's most innovative companies use Workflow
        </h2>
        <div class="mt-8 flow-root lg:mt-0">
          <div class="-mt-4 -ml-8 flex flex-wrap justify-between lg:-ml-4">
            <div
              class="mt-4 ml-8 flex flex-grow flex-shrink-0 lg:flex-grow-0
                lg:ml-4">
              <img
                class="h-12"
                src="https://tailwindui.com/img/logos/tuple-logo-purple-200.svg"
                alt="Tuple" />
            </div>
            <div
              class="mt-4 ml-8 flex flex-grow flex-shrink-0 lg:flex-grow-0
                lg:ml-4">
              <img
                class="h-12"
                src="https://tailwindui.com/img/logos/mirage-logo-purple-200.svg"
                alt="Mirage" />
            </div>
            <div
              class="mt-4 ml-8 flex flex-grow flex-shrink-0 lg:flex-grow-0
                lg:ml-4">
              <img
                class="h-12"
                src="https://tailwindui.com/img/logos/statickit-logo-purple-200.svg"
                alt="StaticKit" />
            </div>
            <div
              class="mt-4 ml-8 flex flex-grow flex-shrink-0 lg:flex-grow-0
                lg:ml-4">
              <img
                class="h-12"
                src="https://tailwindui.com/img/logos/transistor-logo-purple-200.svg"
                alt="Transistor" />
            </div>
            <div
              class="mt-4 ml-8 flex flex-grow flex-shrink-0 lg:flex-grow-0
                lg:ml-4">
              <img
                class="h-12"
                src="https://tailwindui.com/img/logos/workcation-logo-purple-200.svg"
                alt="Workcation" />
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- FAQ offset -->
  <div class="max-w-screen-xl mx-auto py-16 px-4 sm:px-6 lg:py-20 lg:px-8">
    <div class="lg:grid lg:grid-cols-3 lg:gap-8">
      <div class="space-y-4">
        <h2 class="text-3xl leading-9 font-extrabold text-gray-900">
          Frequently asked questions
        </h2>
        <p class="text-lg leading-7 text-gray-500">
          Can’t find the answer you’re looking for? Reach out to our <a
            href="#"
            class="font-medium text-purple-600 hover:text-purple-500 transition
              duration-150 ease-in-out">customer support</a> team.
        </p>
      </div>
      <div class="mt-12 lg:mt-0 lg:col-span-2">
        <dl class="space-y-12">
          <div class="space-y-2">
            <dt class="text-lg leading-6 font-medium text-gray-900">
              Why do you never see elephants hiding in trees?
            </dt>
            <dd class="text-base leading-6 text-gray-500">
              Because they're so good at it. Lorem ipsum dolor sit amet
              consectetur adipisicing elit. Quas cupiditate laboriosam fugiat.
            </dd>
          </div>
          <div class="space-y-2">
            <dt class="text-lg leading-6 font-medium text-gray-900">
              What's the best thing about Switzerland?
            </dt>
            <dd class="text-base leading-6 text-gray-500">
              I don't know, but the flag is a big plus. Lorem ipsum dolor sit
              amet consectetur adipisicing elit. Quas cupiditate laboriosam
              fugiat.
            </dd>
          </div>
          <div class="space-y-2">
            <dt class="text-lg leading-6 font-medium text-gray-900">
              How do you make holy water?
            </dt>
            <dd class="text-base leading-6 text-gray-500">
              You boil the hell out of it. Lorem ipsum dolor sit amet
              consectetur adipisicing elit. Quas cupiditate laboriosam fugiat.
            </dd>
          </div>
          <div class="space-y-2">
            <dt class="text-lg leading-6 font-medium text-gray-900">
              What do you call someone with no body and no nose?
            </dt>
            <dd class="text-base leading-6 text-gray-500">
              Nobody knows. Lorem ipsum dolor sit amet consectetur adipisicing
              elit. Quas cupiditate laboriosam fugiat.
            </dd>
          </div>
        </dl>
      </div>
    </div>
  </div>
</div>
