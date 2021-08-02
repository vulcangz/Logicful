<script lang="typescript">
  import Button from "@app/components/Button.svelte";
  import { dispatch } from "@app/event/EventBus";
  import { subscribeFieldChange } from "@app/event/FieldEvent";
  import { isUrl } from "@app/guards/Guard";
  import type { IForm, Integration } from "@app/models/IForm";
  import { randomString } from "@app/util/Generate";
  import { set } from "@app/util/Selection";
  import { onMount } from "svelte";
  import ConfigField from "../ConfigField.svelte";

  export let form: IForm;
  export let integration: Integration;
  export let index: number;
  let disabled: boolean = true;

  onMount(() => {
    set(form, `workflow.integrations[${index}].name`, integration.name);
  });

  subscribeFieldChange(onMount, (newField) => {
    const url = form.workflow?.integrations?.[index]?.config?.url;
    if (!url) {
      disabled = false;
    }
    if (url) {
      disabled = isUrl(url) ? false : true;
    }
    dispatch("integration_configured", url && isUrl(url));
  });
</script>

<ConfigField
  field={{ id: randomString(), required: true, type: 'string', label: 'Webhook Url', value: form.workflow?.integrations?.[index]?.config?.url, helperText: 'Specify the url you would like submission info to be sent to. Please include https:// or http:// in your url.', placeholder: 'https://e6e86c95a45ac19aada60a1a2003b385.m.pipedream.net', configTarget: 'form', configFieldTarget: `workflow.integrations[${index}].config.url` }} />

<ConfigField
  field={{ id: randomString(), required: true, type: 'string', label: 'Verification Token', value: form.workflow?.integrations?.[index]?.config?.token ?? randomString(), helperText: 'This token will be included as the request header X-Logicful-Token. You can use it to verify the webhook request came from our servers.', configTarget: 'form', configFieldTarget: `workflow.integrations[${index}].config.token` }} />

<ConfigField
  field={{ id: randomString(), required: true, type: 'switch', label: 'Enabled', value: form.workflow?.integrations?.[index]?.enabled, configTarget: 'form', configFieldTarget: `workflow.integrations[${index}].enabled` }} />

<div class="mt-2 border-b border-gray-200" />

<div class="mt-5 ml-3">
  <Button
    {disabled}
    type="primary"
    onClick={() => dispatch('save_form', { status: 'draft' })}>
    Save Form
  </Button>
  <p class="mt-2 text-sm text-gray-500">
    The form must be saved to apply integration changes.
  </p>
</div>
