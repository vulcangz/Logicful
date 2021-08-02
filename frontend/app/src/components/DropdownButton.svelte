<script lang="typescript">
  import type { DropdownButtonAction } from "./models/ComponentProps";

  export let label: string;
  export let actions: DropdownButtonAction[] = [];
  export let processing: boolean = false;
  export let processingLabel: string = "Processing...";

  let showing = false;

  async function executeAction(action: DropdownButtonAction) {
    try {
      processing = true;
      await action.onClick();
    } finally {
      processing = false;
    }
  }

  function show() {
    showing = true;
  }

  function hide() {
    showing = false;
  }
</script>

<div class="btn-group mr-2 mb-2">
  <button
    type="button"
    class={`btn btn-primary`}
    onclick={() => executeAction(actions[0])}
    >{processing ? `${processingLabel}` : actions[0].label}</button
  >
  <button
    type="button"
    class="btn btn-primary dropdown-toggle dropdown-toggle-split"
    data-toggle="dropdown"
    aria-haspopup="true"
    aria-expanded="false"
  >
    <span class="fas fa-angle-down dropdown-arrow" />
    <span class="sr-only">Toggle Dropdown</span>
  </button>
  <div class="dropdown-menu" style="">
    {#each actions as action, i}
      {#if i === 0}
        <span />
      {:else}
        <a class="dropdown-item" href="#" on:click={() => executeAction(action)}
          >{action.label}</a
        >
      {/if}
    {/each}
  </div>
</div>
