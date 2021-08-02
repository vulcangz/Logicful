<script lang="typescript">
  import type { IForm } from "@app/models/IForm";

  import { randomString } from "@app/util/Generate";
  import { subscribeComponent } from "@app/event/EventBus";
  import ConfigField from "./ConfigField.svelte";
  import SectionHeader from "@app/inputs/SectionHeader.svelte";
  import Button from "@app/components/Button.svelte";
  import ManageIntegrations from "./ManageIntegrations.svelte";
  import ManageSubmissionBehavior from "./ManageSubmissionBehavior.svelte";
  import FlyoutPanel from "./FlyoutPanel.svelte";

  export let form: IForm;
  let dialog: "" | "integrations" | "submission-behavior" = "";

  subscribeComponent("form_loaded", (updatedForm) => {
    form = updatedForm;
  });

  subscribeComponent("form_updated", (updatedForm) => {
    form = updatedForm;
  });
</script>

{#if dialog === 'integrations'}
  <FlyoutPanel title="Manage Integrations" onClose={() => (dialog = '')}>
    <ManageIntegrations {form} />
  </FlyoutPanel>
{/if}

{#if dialog === 'submission-behavior'}
  <FlyoutPanel
    title="Manage Submission Behavior"
    description="Redirect or show a message your form has been submitted."
    onClose={() => (dialog = '')}>
    <ManageSubmissionBehavior {form} />
  </FlyoutPanel>
{/if}

<ConfigField
  field={{ id: randomString(), required: true, label: 'Form Title', value: form.title, type: 'string', configFieldTarget: 'title', configTarget: 'form' }} />
<ConfigField
  field={{ id: randomString(), required: true, label: 'Form Description', value: form.description, type: 'string', configFieldTarget: 'description', configTarget: 'form' }} />

<div class="mt-3 ml-3 pl-1">
  <SectionHeader
    field={{ id: randomString(), type: 'section-header', header: 'Manage Integrations', helperText: 'Configure integrations to run upon form submission, such as sending an email.' }} />
  <div class="mt-3">
    <Button type="primary" onClick={() => (dialog = 'integrations')}>
      Manage Integrations
    </Button>
  </div>
</div>

<div class="mt-6 ml-3 pl-1">
  <SectionHeader
    field={{ id: randomString(), type: 'section-header', header: 'Submission Behavior', helperText: 'Configure what should happen for the end-user when the form is submitted.' }} />
  <div class="mt-3">
    <Button type="primary" onClick={() => (dialog = 'submission-behavior')}>
      Manage Submission Behavior
    </Button>
  </div>
</div>
