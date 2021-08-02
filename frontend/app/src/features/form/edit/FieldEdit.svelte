<script lang="typescript">
  import type { IField } from "@app/models/IField";
  import { randomString } from "@app/util/Generate";
  import FieldTypeEditor from "./FieldTypeEditor.svelte";
  import LogicAccordion from "./LogicAccordion.svelte";
  import ContentBlockEditor from "./ContentBlockEditor.svelte";
  import type { FieldEditConfig } from "./models/FieldEditConfig";
  import SectionHeaderEditor from "./SectionHeaderEditor.svelte";
  import ConfigField from "./ConfigField.svelte";
  import ManageValidations from "./ManageValidations.svelte";

  export let field: IField;
  export let config: FieldEditConfig = {};
  let cantBeRequired = ["switch"];
</script>

<form>
  {#if field.type === 'spacer'}
    <ConfigField
      field={{ id: randomString(), label: 'Increase value to add more spacing between the previous and next field.', required: true, value: field.options?.spacer ?? 1, type: 'number', configFieldTarget: 'options.spacer', configTarget: field.id }} />
  {:else if field.type === 'block'}
    <ContentBlockEditor {field} expanded={field.expanded} />
  {:else if field.type === 'section-header'}
    <SectionHeaderEditor {field} />
  {:else}
    <div id={`field-button-${field.id}`}>
      {#if !cantBeRequired.includes(field.type)}
        <ConfigField
          config={{ search: false }}
          field={{ id: randomString(), customCss: 'padding-bottom: 0em;', label: 'Required', value: { type: 'local', value: field.required }, type: 'switch', configFieldTarget: 'required', configTarget: field.id, options: { type: 'local', value: [{ label: 'Yes', value: true }, { label: 'No', value: false }] } }} />
      {/if}
      <ConfigField
        field={{ id: randomString(), label: 'Label', value: field.label, type: 'string', configFieldTarget: 'label', configTarget: field.id }} />
      <ConfigField
        field={{ id: randomString(), label: 'Helper Text', value: field.helperText, type: 'string', configFieldTarget: 'helperText', configTarget: field.id }} />
      <ConfigField
        field={{ id: randomString(), label: 'Field Type', value: { type: 'local', value: field.type }, type: 'combobox', required: true, configFieldTarget: 'type', configTarget: field.id, options: { type: 'local', value: [{ label: 'Email', value: 'email' }] } }} />
      <FieldTypeEditor {field} />
    </div>
  {/if}
  <div class="ml-3 mt-2">
    <LogicAccordion {field} />
    <ManageValidations {field} />
  </div>
</form>
