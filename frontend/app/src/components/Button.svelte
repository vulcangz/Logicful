<script lang="typescript">
  import { afterUpdate, onMount } from "svelte";

  export let type: "primary" | "secondary" | "warn" | "danger" | "white";
  export let submit: boolean = false;
  export let size: "regular" | "large" | "small" = "regular";
  export let onClick: () => any = () => {};
  export let disabled: boolean = false;
  export let width: "" | "full" = "";
  export let focus: boolean = false;
  export let href = "";
  export let hrefTarget = "";
  export let hasIcon: boolean = false;
  export let testId: string = "";
  let focusable: any = null;

  let px = "px-2.5";
  let text = "text-sm";
  let leading = "leading-4";
  let sizeClass = "";

  function setClasses() {
    if (size === "small") {
      px = "px-2";
      text = "text-xs";
      leading = "leading-3";
    }
    if (size === "regular") {
      px = "px-4";
      text = "text-sm";
      leading = "leading-5";
    }
    if (size === "large") {
      px = "px-4";
      text = "text-base";
      leading = "leading-6";
    }
    sizeClass = `${px} ${text} ${leading}`;
    if (hasIcon) {
      sizeClass = `${sizeClass} pl-3 pr-3`;
    }
    if (width === "full") {
      sizeClass = `${sizeClass} pl-3 pr-3 w-full inline-flex items-center justify-center`;
    }

    if (disabled) {
      sizeClass = `${sizeClass} opacity-50 cursor-not-allowed`;
    }
  }

  onMount(() => {
    setClasses();
  });

  afterUpdate(() => {
    setClasses();
    if (focus && focusable) {
      setTimeout(() => {
        try {
          focusable.focus();
        } catch {}
      }, 100);
    }
  });
</script>

{#if type === "danger"}
  <button
    type={submit ? "submit" : "button"}
    {disabled}
    bind:this={focusable}
    data-test-id={testId}
    on:click|stopPropagation={onClick}
    class={`inline-flex justify-center rounded-md border
      border-transparent py-2 bg-red-600 font-medium
      text-white shadow-sm hover:bg-red-500 focus:outline-none
      focus:border-red-700 focus:ring-indigo transition ease-in-out
      duration-150 sm:text-sm sm:leading-5 ${sizeClass}`}
  >
    <slot />
  </button>
{:else if type === "primary"}
  {#if href}
    <a
      target={hrefTarget || "_blank"}
      {href}
      {disabled}
      data-test-id={testId}
      bind:this={focusable}
      class={`inline-flex items-center py-2 border border-transparent
        text-sm rounded-md text-white bg-indigo-600
        hover:bg-indigo-500 focus:outline-none focus:border-indigo-700
        focus:ring-indigo active:bg-indigo-700 transition ease-in-out
        duration-150 ${sizeClass}`}
    >
      <slot />
    </a>
  {:else}
    <button
      type={submit ? "submit" : "button"}
      on:click|stopPropagation={onClick}
      {disabled}
      data-test-id={testId}
      bind:this={focusable}
      class={`inline-flex items-center px-4 py-2 border border-transparent
        text-sm leading-5 font-medium rounded-md text-white bg-indigo-600
        hover:bg-indigo-500 focus:outline-none focus:border-indigo-700
        focus:ring-indigo active:bg-indigo-700 transition ease-in-out
        duration-150 ${sizeClass}`}
    >
      <slot />
    </button>
  {/if}
{:else if type === "secondary"}
  {#if href}
    <a
      target={hrefTarget || "_blank"}
      {href}
      {disabled}
      data-test-id={testId}
      bind:this={focusable}
      class={`inline-flex items-center py-2 border border-transparent
        font-medium rounded-md text-indigo-700 bg-indigo-100
        hover:bg-indigo-50 focus:outline-none focus:border-indigo-300
        focus:ring-indigo active:bg-indigo-200 transition ease-in-out
        duration-150 ${sizeClass}`}
    >
      <slot />
    </a>
  {:else}
    <button
      type={submit ? "submit" : "button"}
      on:click|stopPropagation={onClick}
      {disabled}
      bind:this={focusable}
      data-test-id={testId}
      class={`inline-flex items-center py-2 border border-transparent
        font-medium rounded-md text-indigo-700 bg-indigo-100
        hover:bg-indigo-50 focus:outline-none focus:border-indigo-300
        focus:ring-indigo active:bg-indigo-200 transition ease-in-out
        duration-150 ${sizeClass}`}
    >
      <slot />
    </button>
  {/if}
{:else if type === "white"}
  {#if href}
    <a
      target={hrefTarget || "_blank"}
      {href}
      {disabled}
      bind:this={focusable}
      data-test-id={testId}
      class={`inline-flex items-center py-2 border border-gray-300
        font-medium rounded-md text-gray-700 bg-white
        hover:text-gray-500 focus:outline-none focus:border-indigo-300
        focus:ring-indigo active:text-gray-800 active:bg-gray-50
        transition ease-in-out duration-150 ${sizeClass}`}
    >
      <slot />
    </a>
  {:else}
    <button
      type={submit ? "submit" : "button"}
      on:click|stopPropagation={onClick}
      {disabled}
      bind:this={focusable}
      data-test-id={testId}
      class={`inline-flex items-center py-2 border border-gray-300
        font-medium rounded-md text-gray-700 bg-white
        hover:text-gray-500 focus:outline-none focus:border-indigo-300
        focus:ring-indigo active:text-gray-800 active:bg-gray-50
        transition ease-in-out duration-150 ${sizeClass}`}
    >
      <slot />
    </button>
  {/if}
{/if}
