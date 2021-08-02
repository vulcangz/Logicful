<script lang="typescript">
  import Button from "@app/components/Button.svelte";
  import { dispatch } from "@app/event/EventBus";
  import { subscribeFieldChange } from "@app/event/FieldEvent";
  import { isEmail } from "@app/guards/Guard";
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
    const email = form.workflow?.integrations?.[index]?.config?.email;
    if (!email) {
      disabled = false;
    }
    if (email) {
      disabled = !isEmail(email);
    }
    dispatch("integration_configured", email && isEmail(email));
  });
</script>

<ConfigField
  field={{ id: randomString(), required: true, type: 'email', label: 'Email Address', value: form.workflow?.integrations?.[index]?.config?.email, helperText: 'Specify which email you would like to recieve submission info.', configTarget: 'form', configFieldTarget: `workflow.integrations[${index}].config.email` }} />

<ConfigField
  field={{ id: randomString(), required: true, type: 'switch', label: 'Enabled', value: form.workflow?.integrations?.[index]?.enabled, configTarget: 'form', configFieldTarget: `workflow.integrations[${index}].enabled` }} />

<div class="mt-2 border-b border-gray-200"></div>

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
