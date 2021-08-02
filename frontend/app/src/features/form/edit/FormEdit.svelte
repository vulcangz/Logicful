<script lang="typescript">
  import type { IForm } from "@app/models/IForm";

  let form: IForm;
  import { randomString } from "@app/util/Generate";
  import { onMount } from "svelte";
  import { subscribeComponent } from "@app/event/EventBus";
  import formStore from "@app/store/FormStore";
  import FormEditSettings from "./FormEditSettings.svelte";
  import ConfigField from "./ConfigField.svelte";

  subscribeComponent("form_updated", (props) => {
    form = props;
  });

  subscribeComponent("form_loaded", (props) => {
    form = props.form;
  });

  onMount(() => {
    const f = formStore.getForm();
    if (f.loaded) {
      form = f;
    }
  });
</script>

{#if form}
  <div>
    <ConfigField
      field={{ id: randomString(), type: 'switch', label: 'Enable Logic For Preview', value: { type: 'local', value: form.enableLogic ?? true }, configFieldTarget: 'enableLogic', configTarget: 'form' }} />
    <FormEditSettings {form} />
  </div>
{:else}
  <div class="spinner-border" style="width: 2rem; height: 2rem;" role="status">
    <span class="sr-only">Loading...</span>
  </div>
{/if}
