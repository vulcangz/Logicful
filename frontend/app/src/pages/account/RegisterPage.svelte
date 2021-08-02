<script lang="typescript">
  import { postApi } from "@app/services/ApiService";
  import { setToken } from "@app/services/AuthService";
  import { afterUpdate } from "svelte";
  import type { UserToken } from "@app/services/AuthService";
  import { navigate } from "svelte-routing";
  import Button from "@app/components/Button.svelte";

  let email = "";
  let password = "";
  let name = "";
  let displayName = "";
  let agree = false;
  let valid = false;
  let showPassword = false;
  let creating = false;
  let error = "";
  let lastName = "";

  async function onRegister() {
    error = "";
    creating = true;
    if (!name && !displayName) {
      name = email.split("@")[0].trim();
      displayName = name;
    }
    try {
      const result = await postApi<UserToken>("user/register", {
        email,
        fullName: name,
        displayName,
        password,
      });
      if (result.token) {
        setToken(result);
        navigate("/");
      } else {
        error = "Failed to register, unknown response from server.";
      }
    } catch (ex) {
      if (ex.message === "email already exists") {
        error = "Email address is already in use.";
      } else {
        error = "Failed to register, something went wrong.";
      }
    } finally {
      creating = false;
    }
  }

  afterUpdate(() => {
    valid = isValid();
  });

  function onNameBlur() {
    if (name && lastName !== name) {
      displayName = name.trim().split(" ")[0] ?? "";
      lastName = name;
    }
  }

  function isValid(): boolean {
    return email != "" && password != "" && agree;
  }

  function togglePassword() {
    showPassword = !showPassword;
  }
</script>

<svelte:head>
  <meta
    name="google-signin-client_id"
    content="807768276065-c8b2jvlks20jgm0mk3t2akfm7pc3jomu.apps.googleusercontent.com" />
  <script src="https://apis.google.com/js/platform.js" async defer>
  </script>
</svelte:head>

<div
  class="min-h-screen bg-gray-100 flex flex-col justify-center py-12 sm:px-6
    lg:px-8">
  <div class="sm:mx-auto sm:w-full sm:max-w-md">
    <h2
      class="mt-6 text-center text-3xl leading-9 font-extrabold text-gray-900">
      Sign up for our platform
    </h2>
    <p class="mt-2 text-center text-sm leading-5 text-gray-600 max-w">
      Or <a href="/account/login" class="font-medium text-indigo-600
          hover:text-indigo-500 focus:outline-none focus:underline transition
          ease-in-out duration-150"> Sign in </a>
    </p>
  </div>

  <div class="sm:mx-auto sm:w-full sm:max-w-md">
    {#if error}
      <div class="inset-x-0 pt-2 mb-3">
        <div class="max-w-screen-xl mx-auto px-2 sm:px-6 lg:px-8">
          <div class="p-2 rounded-md bg-red-600 shadow-lg sm:p-2">
            <div class="flex items-center justify-between flex-wrap">
              <div class="w-0 flex-1 flex items-center">
                <p class="ml-3 font-medium text-white truncate">
                  <span> {error} </span>
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    {/if}

    <div class="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10 mt-3">
      <div style="margin-bottom: 1em">
        <div
          class="g-signin2"
          data-longtitle="true"
          data-theme="dark"
          data-height="50"
          data-width="365"
          data-onsuccess="onGoogleSignIn" />
      </div>

      <div class="mt-6">
        <div class="relative">
          <div class="absolute inset-0 flex items-center">
            <div class="w-full border-t border-gray-300" />
          </div>
          <div class="relative flex justify-center text-sm leading-5">
            <span class="px-2 bg-white text-gray-500"> Or </span>
          </div>
        </div>

        <form on:submit|preventDefault|stopPropagation={onRegister}>
          <div class="mt-2">
            <label
              for="name"
              class="block text-sm font-medium leading-5 text-gray-700">Full
              Name</label>
            <div class="mt-1 relative rounded-md shadow-sm">
              <div
                class="absolute inset-y-0 left-0 pl-3 flex items-center
                  pointer-events-none">
                <svg
                  class="w-5 h-5"
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 20 20"
                  fill="currentColor">
                  <path
                    fill-rule="evenodd"
                    d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-6-3a2 2 0 11-4 0 2 2 0 014 0zm-2 4a5 5 0 00-4.546 2.916A5.986 5.986 0 0010 16a5.986 5.986 0 004.546-2.084A5 5 0 0010 11z"
                    clip-rule="evenodd" />
                </svg>
              </div>
              <input
                id="name"
                autocomplete="name"
                type="text"
                bind:value={name}
                on:blur={onNameBlur}
                class="form-input block w-full pl-10 sm:text-sm sm:leading-5" />
            </div>
          </div>
          <div class="mt-2">
            <label
              for="displayName"
              class="block text-sm font-medium leading-5 text-gray-700">What
              should we call you?</label>
            <div class="mt-1 relative rounded-md shadow-sm">
              <div
                class="absolute inset-y-0 left-0 pl-3 flex items-center
                  pointer-events-none">
                <svg
                  class="w-5 h-5"
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 20 20"
                  fill="currentColor">
                  <path
                    fill-rule="evenodd"
                    d="M18 10c0 3.866-3.582 7-8 7a8.841 8.841 0 01-4.083-.98L2 17l1.338-3.123C2.493 12.767 2 11.434 2 10c0-3.866 3.582-7 8-7s8 3.134 8 7zM7 9H5v2h2V9zm8 0h-2v2h2V9zM9 9h2v2H9V9z"
                    clip-rule="evenodd" />
                </svg>
              </div>
              <input
                id="displayName"
                autocomplete="off"
                type="text"
                bind:value={displayName}
                class="form-input block w-full pl-10 sm:text-sm sm:leading-5" />
            </div>
          </div>
          <div class="mt-2">
            <label
              for="email"
              class="block text-sm font-medium leading-5 text-gray-700">Email</label>
            <div class="mt-1 relative rounded-md shadow-sm">
              <div
                class="absolute inset-y-0 left-0 pl-3 flex items-center
                  pointer-events-none">
                <svg
                  class="w-5 h-5"
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 20 20"
                  fill="currentColor">
                  <path
                    d="M2.003 5.884L10 9.882l7.997-3.998A2 2 0 0016 4H4a2 2 0 00-1.997 1.884z" />
                  <path
                    d="M18 8.118l-8 4-8-4V14a2 2 0 002 2h12a2 2 0 002-2V8.118z" />
                </svg>
              </div>
              <input
                id="email"
                autocomplete="email"
                type="email"
                bind:value={email}
                class="form-input block w-full pl-10 sm:text-sm sm:leading-5" />
            </div>
          </div>
          <div class="mt-2">
            <label
              for="password"
              class="block text-sm font-medium leading-5 text-gray-700">Password</label>
            <div class="mt-1 relative rounded-md shadow-sm">
              <div
                class="absolute inset-y-0 left-0 pl-3 flex items-center
                  pointer-events-none">
                <svg
                  class="h-5 w-5"
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 20 20"
                  fill="currentColor">
                  <path
                    fill-rule="evenodd"
                    d="M5 9V7a5 5 0 0110 0v2a2 2 0 012 2v5a2 2 0 01-2 2H5a2 2 0 01-2-2v-5a2 2 0 012-2zm8-2v2H7V7a3 3 0 016 0z"
                    clip-rule="evenodd" />
                </svg>
              </div>
              {#if showPassword}
                <input
                  id="password"
                  type="text"
                  minlength="5"
                  required
                  autocomplete="current-password"
                  bind:value={password}
                  class="form-input block w-full pl-10 sm:text-sm sm:leading-5" />
                <div
                  class="absolute inset-y-0 right-0 pr-3 flex items-center
                    cursor-pointer"
                  on:click={togglePassword}>
                  <svg
                    class="w-5 h-5"
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor">
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.88 9.88l-3.29-3.29m7.532 7.532l3.29 3.29M3 3l3.59 3.59m0 0A9.953 9.953 0 0112 5c4.478 0 8.268 2.943 9.543 7a10.025 10.025 0 01-4.132 5.411m0 0L21 21" />
                  </svg>
                </div>
              {:else}
                <input
                  id="password"
                  type="password"
                  minlength="5"
                  required
                  autocomplete="current-password"
                  bind:value={password}
                  class="form-input block w-full pl-10 sm:text-sm sm:leading-5" />
                <div
                  class="absolute inset-y-0 right-0 pr-3 flex items-center
                    cursor-pointer"
                  on:click={togglePassword}>
                  <svg
                    class="w-5 h-5"
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
                  </svg>
                </div>
              {/if}
            </div>
          </div>
          <div class="mt-6 flex items-center justify-between">
            <div class="flex items-center">
              <input
                id="terms"
                bind:checked={agree}
                type="checkbox"
                class="form-checkbox h-4 w-4 text-indigo-600 transition
                  duration-150 ease-in-out" />
              <label
                for="terms"
                class="ml-2 block text-sm leading-5 text-gray-900">
                I agree to the <a href="/tos" target="_blank" class="text-indigo-700">terms
                  and conditions</a>
              </label>
            </div>
          </div>

          <div class="mt-5">
            {#if !creating}
              <Button type="primary" width={'full'} submit={true}>
                Sign Up With Email
              </Button>
            {:else}
              <Button type="primary" width={'full'} disabled={true}>
                Creating Account...
              </Button>
            {/if}
          </div>
        </form>
      </div>
    </div>
  </div>
</div>
