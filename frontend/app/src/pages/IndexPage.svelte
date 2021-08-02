<script lang="typescript">
  import AlphaBanner from "@app/components/AlphaBanner.svelte";
  import MarketingNavbar from "@app/components/layout/MarketingNavbar.svelte";
  import MarketingFooter from "../components/layout/MarketingFooter.svelte";
  import TheTeam from "../components/layout/TheTeam.svelte";
</script>

<AlphaBanner />
<div class="bg-gray-50">
  <div class="relative overflow-hidden">
    <div class="block absolute inset-y-0 h-full w-full">
      <div class="relative h-full">
        <svg
          class="absolute right-full transform translate-y-1/3 translate-x-1/4
            md:translate-y-1/2 sm:translate-x-1/2 lg:translate-x-full"
          width="404"
          height="784"
          fill="none"
          viewBox="0 0 404 784"
        >
          <defs>
            <pattern
              id="e229dbec-10e9-49ee-8ec3-0286ca089edf"
              x="0"
              y="0"
              width="20"
              height="20"
              patternUnits="userSpaceOnUse"
            >
              <rect
                x="0"
                y="0"
                width="4"
                height="4"
                class="text-gray-200"
                fill="currentColor"
              />
            </pattern>
          </defs>
          <rect
            width="404"
            height="784"
            fill="url(#ad9a0a02-b58e-4a1d-8c36-1b649889af63)"
          />
        </svg>
        <svg
          class="absolute left-full transform -translate-y-3/4 -translate-x-1/4
            sm:-translate-x-1/2 md:-translate-y-1/2 lg:-translate-x-3/4"
          width="404"
          height="784"
          fill="none"
          viewBox="0 0 404 784"
        >
          <defs>
            <pattern
              id="d2a68204-c383-44b1-b99f-42ccff4e5365"
              x="0"
              y="0"
              width="20"
              height="20"
              patternUnits="userSpaceOnUse"
            >
              <rect
                x="0"
                y="0"
                width="4"
                height="4"
                class="text-gray-200"
                fill="currentColor"
              />
            </pattern>
          </defs>
          <rect
            width="404"
            height="784"
            fill="url(#478e97d6-90df-4a89-8d63-30fdbb3c7e57)"
          />
        </svg>
      </div>
    </div>

    <div class="relative pt-6 pb-12 lg:pb-20">
      <MarketingNavbar />
      <div class="mt-5 mx-auto max-w-screen-xl px-4 sm:px-6 md:mt-16 lg:mt-20">
        <div class="text-center">
          <h2
            class="text-4xl tracking-tight leading-10 font-extrabold
              text-gray-900 sm:text-5xl sm:leading-none md:text-6xl"
          >
            simple,
            <br />
            <span class="text-indigo-600">dynamic forms</span>
          </h2>
          <p
            class="mt-3 max-w-md mx-auto text-base text-gray-500 sm:text-lg
              md:mt-5 md:text-xl md:max-w-3xl"
          >
            Powerful easy to use dynamic forms to enhance your business
          </p>
        </div>
      </div>
    </div>
    <div class="relative">
      <div class="absolute inset-0 flex flex-col">
        <div class="flex-1" />
        <div class="flex-1 w-full bg-gray-800" />
      </div>
      <div class="max-w-screen-xl mx-auto px-4 sm:px-6">
        <img
          class="relative rounded-lg shadow-lg"
          src="/assets/builder.png"
          alt="App screenshot"
        />
      </div>
    </div>
  </div>
  <div class="bg-gray-800">
    <div
      class="max-w-screen-xl mx-auto pt-8 pb-8 px-4 sm:px-6 md:pb-24 lg:px-8"
    />
  </div>
  <div class="bg-gray-50 overflow-hidden">
    <div class="relative max-w-screen-xl mx-auto py-12 px-4 sm:px-6 lg:px-8">
      <svg
        class="absolute top-0 left-full transform -translate-x-1/2
          -translate-y-3/4 lg:left-auto lg:right-full lg:translate-x-2/3
          lg:translate-y-1/4"
        width="404"
        height="784"
        fill="none"
        viewBox="0 0 404 784"
      >
        <defs>
          <pattern
            id="8b1b5f72-e944-4457-af67-0c6d15a99f38"
            x="0"
            y="0"
            width="20"
            height="20"
            patternUnits="userSpaceOnUse"
          >
            <rect
              x="0"
              y="0"
              width="4"
              height="4"
              class="text-gray-200"
              fill="currentColor"
            />
          </pattern>
        </defs>
        <rect
          width="404"
          height="784"
          fill="url(#8b1b5f72-e944-4457-af67-0c6d15a99f38)"
        />
      </svg>

      <div class="relative lg:grid lg:grid-cols-3 lg:gap-x-8">
        <div class="lg:col-span-1">
          <h3
            class="text-3xl leading-9 font-extrabold tracking-tight
              text-gray-900 sm:text-4xl sm:leading-10"
          >
            A better way to collect data.
          </h3>
        </div>
        <div
          class="mt-10 sm:grid sm:grid-cols-2 sm:gap-x-8 sm:gap-y-10
            lg:col-span-2 lg:mt-0"
        >
          <div>
            <div
              class="flex items-center justify-center h-12 w-12 rounded-md
                bg-indigo-500 text-white"
            >
              <!-- Heroicon name: globe-alt -->
              <svg
                class="h-6 w-6"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M21 12a9 9 0 01-9 9m9-9a9 9 0 00-9-9m9 9H3m9 9a9 9 0 01-9-9m9 9c1.657 0 3-4.03 3-9s-1.343-9-3-9m0 18c-1.657 0-3-4.03-3-9s1.343-9 3-9m-9 9a9 9 0 019-9"
                />
              </svg>
            </div>
            <div class="mt-5">
              <h4 class="text-lg leading-6 font-medium text-gray-900">
                Advanced Integrations
              </h4>
              <p class="mt-2 text-base leading-6 text-gray-500">
                Integrate your forms with our various prebuilt integrations at
                the click of a button.
              </p>
              <p class="mt-2 text-base leading-6 text-gray-500">
                <strong>Examples</strong>
                of integrations are sending an email when a submission is received,
                posting to a webhook, or uploading to an S3 bucket.
              </p>
            </div>
          </div>
          <div class="mt-10 sm:mt-0">
            <div
              class="flex items-center justify-center h-12 w-12 rounded-md
                bg-indigo-500 text-white"
            >
              <svg
                class="h-6 w-6"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.88 9.88l-3.29-3.29m7.532 7.532l3.29 3.29M3 3l3.59 3.59m0 0A9.953 9.953 0 0112 5c4.478 0 8.268 2.943 9.543 7a10.025 10.025 0 01-4.132 5.411m0 0L21 21"
                />
              </svg>
            </div>
            <div class="mt-5">
              <h4 class="text-lg leading-6 font-medium text-gray-900">
                Powerful Form Logic
              </h4>
              <p class="mt-2 text-base leading-6 text-gray-500">
                Hide or show fields based on the input from other fields. Easy
                to use logic builder to create
                <strong>interactive forms</strong>, no coding required.
              </p>
            </div>
          </div>
          <div class="mt-10 sm:mt-0">
            <div
              class="flex items-center justify-center h-12 w-12 rounded-md
                bg-indigo-500 text-white"
            >
              <!-- Heroicon name: lightning-bolt -->
              <svg
                class="h-6 w-6"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M13 10V3L4 14h7v7l9-11h-7z"
                />
              </svg>
            </div>
            <div class="mt-5">
              <h4 class="text-lg leading-6 font-medium text-gray-900">
                Easy to use builder
              </h4>
              <p class="mt-2 text-base leading-6 text-gray-500">
                Simple, easy to use builder to quickly create your advanced
                forms.
              </p>
              <p class="mt-2 text-base leading-6 text-gray-500">
                All controls are available in one place in the same consistent
                interface.
              </p>
            </div>
          </div>
          <div class="mt-10 sm:mt-0">
            <div
              class="flex items-center justify-center h-12 w-12 rounded-md
                bg-indigo-500 text-white"
            >
              <svg
                class="h-6 w-6"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M3 15a4 4 0 004 4h9a5 5 0 10-.1-9.999 5.002 5.002 0 10-9.78 2.096A4.001 4.001 0 003 15z"
                />
              </svg>
            </div>
            <div class="mt-5">
              <h4 class="text-lg leading-6 font-medium text-gray-900">
                Dynamic, Reusable Content
              </h4>
              <p class="mt-2 text-base leading-6 text-gray-500">
                Reuse values for dropdowns, radio buttons, checkboxes, and
                content by specifying an option set.
              </p>
              <p class="mt-2 text-base leading-6 text-gray-500">
                Option sets can be declared in-line or pull from a
                <strong>remote HTTP source</strong>.
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <TheTeam />

  <div class="bg-indigo-700">
    <div
      class="max-w-2xl mx-auto text-center py-16 px-4 sm:py-20 sm:px-6 lg:px-8"
    >
      <h2
        class="text-3xl leading-9 font-extrabold text-white sm:text-4xl sm:leading-10"
      >
        <span class="block">Boost your productivity.</span>
        <span class="block">Sign up for Logicful.</span>
      </h2>
      <p class="mt-4 text-lg leading-6 text-indigo-200">
        Get notified when Logicful is ready for production use.
      </p>
      <a
        href="/form/preview?formId=9abcd1fb-dc6e-4574-9235-47600aa53f1a"
        class="mt-8 w-full inline-flex items-center justify-center px-5 py-3 border border-transparent text-base leading-6 font-medium rounded-md text-indigo-600 bg-white hover:text-indigo-500 hover:bg-indigo-50 transition duration-150 ease-in-out sm:w-auto"
      >
        Sign up for updates
      </a>
    </div>
  </div>
  <MarketingFooter />
</div>
