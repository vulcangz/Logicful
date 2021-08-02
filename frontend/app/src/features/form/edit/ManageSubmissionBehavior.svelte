<script lang="typescript">
  import type { IForm } from "@app/models/IForm";
  import { randomString } from "@app/util/Generate";
  import ConfigField from "./ConfigField.svelte";
  import SaveFormButton from "./SaveFormButton.svelte";

  export let form: IForm;
</script>

<ConfigField
  field={{ id: randomString(), required: true, value: form.submissionConfig?.afterSubmitAction, label: 'Choose action to do after form has been submitted.', type: 'radio-group', options: ['Redirect to URL', 'Show Message'], configTarget: 'form', configFieldTarget: 'submissionConfig.afterSubmitAction' }} />

{#if form.submissionConfig?.afterSubmitAction?.['Show Message']}
  <ConfigField
    field={{ id: randomString(), required: true, value: form.submissionConfig?.afterSubmitConfig?.message, label: 'Message to display', type: 'block-editor', configTarget: 'form', configFieldTarget: 'submissionConfig.afterSubmitConfig.message' }} />
{/if}

{#if form.submissionConfig?.afterSubmitAction?.['Redirect to URL']}
  <ConfigField
    field={{ id: randomString(), required: true, value: form.submissionConfig?.afterSubmitConfig?.url, label: 'URL to redirect to:', type: 'string', configTarget: 'form', configFieldTarget: 'submissionConfig.afterSubmitConfig.url' }} />
{/if}

<div class="ml-3 mt-3">
  <SaveFormButton />
</div>
