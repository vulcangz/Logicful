<script lang="typescript">
  import { subscribeComponent } from "@app/event/EventBus";
  import Field from "@app/features/form/edit/Field.svelte";

  import type { IForm } from "@app/models/IForm";

  import { randomString } from "@app/util/Generate";

  export let form: IForm;

  subscribeComponent("form_loaded", (updatedForm) => {
    form = updatedForm;
  });
  subscribeComponent("form_updated", (updatedForm) => {
    form = updatedForm;
  });
</script>

<p style="padding-top: 0.6em; margin-bottom: 0em;">
  Configure what emails are sent when the form is submitted.
</p>

<Field
  field={{ id: randomString(), type: 'switch', label: 'Send me an email summary on submission', value: { type: 'local', value: form.emailOnSubmission ?? false }, configFieldTarget: 'emailOnSubmission', configTarget: 'form' }} />

<h2 class="h5" style="padding-top: 0.5em;">Custom Emails</h2>

<Field
  field={{ id: randomString(), required: true, label: 'Form Title', value: { type: 'local', value: form.title }, type: 'string', configFieldTarget: 'title', configTarget: 'form' }} />

<Field
  field={{ id: randomString(), type: 'switch', label: 'Disable after a maximum number of submissions', value: { type: 'local', value: form.maxSubmissions ?? false }, configFieldTarget: 'maxSubmissions', configTarget: 'form' }} />

<Field
  field={{ id: randomString(), type: 'date', required: true, label: 'Submissions close after date/time ', value: { type: 'local', value: form.closeDateTime ?? '' }, configFieldTarget: 'closeDateTime', configTarget: 'form' }} />
