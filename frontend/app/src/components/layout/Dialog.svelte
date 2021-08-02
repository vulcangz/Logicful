<script lang="typescript">
  import { afterUpdate, onMount } from "svelte";
  import type { ButtonAction } from "@app/components/models/ComponentProps";
  import Button from "../Button.svelte";
  import { dispatch } from "@app/event/EventBus";

  export let title: string = "";
  export let isOpen: boolean = false;
  export let onClose = () => {};
  export let actions: ButtonAction[] = [];
  export let getActions: (() => ButtonAction[]) | undefined = undefined;

  let loaded = false;
  let processing = -1;
  let failed = false;
  let error = "";

  onMount(() => {
    dispatch("modal_open", {});
    actions = getActions?.() ?? actions;
    setTimeout(() => {
      loaded = true;
    }, 500);
    return () => {
      isOpen = false;
    };
  });

  afterUpdate(() => {
    actions = getActions?.() ?? actions;
  });

  async function close() {
    if (!isOpen || !loaded) {
      return;
    }
    dispatch("modal_close", {});
    isOpen = false;
    loaded = false;
    await onClose?.();
  }

  async function runAction(action: ButtonAction, index: number) {
    try {
      failed = false;
      error = "";
      processing = index;
      if (action && action.onClick) {
        await action.onClick();
      }
      processing = -1;
      if (action.onClose == null || action.onClose) {
        close();
      }
    } catch (ex) {
      console.error(ex);
      error = ex.message;
      failed = true;
    }
  }

  function buttonClass(isLast: boolean) {
    return !isLast
      ? "flex w-full rounded-md shadow-sm sm:w-auto"
      : "mt-3 flex w-full rounded-md shadow-sm sm:mt-0 sm:ml-3 sm:w-auto";
  }
</script>

{#if isOpen}
  <div class="fixed z-50 inset-0 overflow-y-auto w-auto min-w-96">
    <div
      class="flex items-end justify-center min-h-screen pt-4 px-4 pb-20
        text-center sm:block sm:p-0">
      <div class="fixed inset-0 transition-opacity">
        <div class="absolute inset-0 bg-gray-500 opacity-75" />
      </div>

      <span class="hidden sm:inline-block sm:align-middle sm:h-screen min-w-" />&#8203;
      <div
        class="inline-block align-bottom bg-white rounded-lg px-4 pt-5 pb-4
          text-left overflow-hidden shadow-xl transform transition-all sm:my-8
          sm:align-middle w-auto sm:p-6 max-w-screen-lg min-w-96"
        role="dialog"
        on:click|stopPropagation
        aria-modal="true"
        aria-labelledby="modal-headline">
        <div class="hidden sm:block absolute top-0 right-0 pt-4 pr-4">
          <button
            on:click={close}
            type="button"
            class="text-gray-400 hover:text-gray-500 focus:outline-none
              focus:text-gray-500 transition ease-in-out duration-150"
            aria-label="Close">
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
        <div class="sm:flex sm:items-start w-full">
          <div class="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left w-full">
            <h3
              class="text-lg leading-6 font-medium text-gray-900"
              id="modal-headline">
              {title}
            </h3>
            <div class="mt-3 w-full z-50">
              <slot />
            </div>
          </div>
        </div>
        <div class="mt-5 sm:mt-4 sm:pl-4 sm:flex">
          {#if actions.length > 0}
            {#each actions as action, index}
              <span class={buttonClass(index === actions.length - 1)}>
                {#if processing === index}
                  {#if failed}
                    <Button
                      type={action.type}
                      width="full"
                      onClick={() => runAction(action, index)}>
                      Failed To Run, Click To Try Again
                    </Button>
                  {:else}
                    <Button type={action.type} width="full" disabled={true}>
                      Processing...
                    </Button>
                  {/if}
                {:else if action.focus}
                  <Button
                    type={action.type}
                    focus={true}
                    width="full"
                    onClick={() => runAction(action, index)}>
                    {action.label}
                  </Button>
                {:else}
                  <Button
                    type={action.type}
                    width="full"
                    onClick={() => runAction(action, index)}>
                    {action.label}
                  </Button>
                {/if}
              </span>
            {/each}
          {/if}
        </div>
        <div class="mt-5 sm:mt-4 sm:pl-4 sm:flex">
          {#if error}
            <div class="rounded-md bg-red-50 p-4">
              <div class="flex">
                <div class="flex-shrink-0">
                  <svg
                    class="h-5 w-5 text-red-400"
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 20 20"
                    fill="currentColor">
                    <path
                      fill-rule="evenodd"
                      d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z"
                      clip-rule="evenodd" />
                  </svg>
                </div>
                <div class="ml-3">
                  <h3 class="text-sm leading-5 font-medium text-red-800">
                    An error has occured.
                  </h3>
                  <div class="mt-2 text-sm leading-5 text-red-700">
                    <ul class="list-disc pl-5">
                      <li>{error}</li>
                    </ul>
                  </div>
                </div>
              </div>
            </div>
          {/if}
        </div>
      </div>
    </div>
  </div>
{/if}

<svelte:body on:click={close} />
