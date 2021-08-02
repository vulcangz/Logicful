<script lang="typescript">
  import type { IField } from "@app/models/IField";
  import { subscribeFieldChange } from "@app/event/FieldEvent";
  import { onMount } from "svelte";
  import formStore from "@app/store/FormStore";
  import { isString } from "@app/guards/Guard";
  import { LoadState } from "@app/models/LoadState";

  export let field: IField;
  export let isPreview: boolean = false;
  let value = "";
  let lastUrl = "";
  let state = LoadState.NotStarted;
  let url;

  subscribeFieldChange(onMount, (newField) => {
    if (newField.id === field.id && lastUrl !== newField.value) {
      url = newField.value;
      load(url);
    }
  });

  onMount(async () => {
    url = formStore.getValue(field.configTarget ?? field.id);
    load(url);
  });

  async function load(url: string) {
    state = LoadState.Loading;
    try {
      if (!url) {
        return;
      }
      if (!isString(url)) {
        return;
      }
      if (url.startsWith("http")) {
        lastUrl = url;
        const response = await fetch(url);
        const html = await response.text();
        value = html ?? "";
      } else {
        value = url;
      }
    } catch (ex) {
      state = LoadState.Failed;
    } finally {
      if (state !== LoadState.Failed) {
        state = LoadState.Finished;
      }
    }
  }
</script>

<div>
  {#if isPreview && (value === "" || value == null)}
    <h5>Content Placeholder</h5>
    <p>
      From the field configuration settings, select a content block to display.
    </p>
  {:else if state === LoadState.Finished}
    {@html value}
  {:else if state === LoadState.Failed}
    <p>Failed to load content.</p>
  {:else}
    <div class="d-flex justify-content-center">
      <div class="spinner-border text-dark" role="status">
        <span class="sr-only">Loading...</span>
      </div>
    </div>
  {/if}
</div>
