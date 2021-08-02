<script lang="typescript">
  import type { IForm } from "@app/models/IForm";
  import { dispatch, subscribeComponent } from "@app/event/EventBus";

  import FormEditSettings from "@app/features/form/edit/FormEditSettings.svelte";
  import { saveForm } from "@app/features/form/edit/services/SaveForm";
  import { onMount } from "svelte";
  import EmailSettings from "./EmailSettings.svelte";
  let saving = false;

  export let form: IForm;
  export let selected: string;

  async function saveDraft() {
    saving = true;
    await dispatch("save_form", {
      status: "draft",
    });
    saving = false;
  }

  subscribeComponent("save_form", async (params) => {
    await saveForm();
  });

  onMount(() => {});
</script>

<div class="row mb-5">
  <div class="col-12 mb-4">
    <div
      class="d-flex justify-content-end"
      style="padding-bottom: 1em; padding-left: 0; display: flex; text-align: right;">
      <div style="text-align: right;">
        <a
          href={`/form/submissions?formId=${form.id}`}
          target="_blank"
          class="btn btn-outline-dark">View Submissions</a>
      </div>
      <div style="text-align: right; padding-left: 0.5em;">
        <a
          href={`/form/builder?formId=${form.id}`}
          target="_blank"
          class="btn btn-outline-dark">
          <span class="fas fa-pencil-alt" /> Edit Form
        </a>
      </div>
      <div style="text-align: right; padding-left: 0.5em;">
        <a
          href={`/form/preview?formId=${form.id}&mode=local`}
          target="_blank"
          class="btn btn-outline-dark">Preview Form</a>
      </div>
    </div>
    <form action="#" method="post" class="card border-light p-3 mb-4">
      {#if selected === 'general'}
        <div
          class="card-header bg-white border-light p-3 mb-4 mb-md-0"
          style="display: flex; padding-top: 0.2em !important;">
          <h3 class="h5 mb-0" style="padding-top: 0.4em;">General</h3>
        </div>
        <div
          class="card-body p-0 p-md-4"
          style="padding-top: 0.5em !important;">
          <div class="row justify-content-center">
            {#if form}
              <FormEditSettings {form} />
            {:else}
              <div class="spinner" />
            {/if}
          </div>
        </div>
      {/if}
      {#if selected === 'workflows'}
        <div
          class="card-header bg-white border-light p-3 mb-4 mb-md-0"
          style="display: flex; padding-top: 0.2em !important;">
          <h3 class="h5 mb-0" style="padding-top: 0.4em;">Workflows</h3>
        </div>
        <div
          class="card-body p-0 p-md-4"
          style="padding-top: 0.5em !important;">
          <div class="row justify-content-center">
            {#if form}
              <!-- <FormEditSettings {form} /> -->
            {:else}
              <div class="spinner" />
            {/if}
          </div>
        </div>
      {/if}
      {#if selected === 'emails'}
        <div
          class="card-header bg-white border-light p-3 mb-4 mb-md-0"
          style="display: flex; padding-top: 0.2em !important;">
          <h3 class="h5 mb-0" style="padding-top: 0.4em;">Configure Emails</h3>
        </div>
        <div
          class="card-body p-0 p-md-4"
          style="padding-top: 0.5em !important;">
          <div class="row justify-content-center">
            {#if form}
              <EmailSettings {form} />
            {:else}
              <div class="spinner" />
            {/if}
          </div>
        </div>
      {/if}
      <div
        class="d-flex justify-content-end ml-auto"
        style="padding-right: 1em; padding-bottom: 1em;">
        {#if saving}
          <button
            class="btn btn-primary"
            type="button"
            disabled>Saving...</button>
        {:else}
          <button
            class="btn btn-primary"
            type="button"
            on:click={saveDraft}>Save Changes</button>
        {/if}
      </div>
    </form>
  </div>
</div>
