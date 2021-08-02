<script lang="typescript">
  import type { IField } from "@app/models/IField";
  import ComboBoxOptionsEditor from "./ComboBoxOptionsEditor.svelte";
  import { randomString } from "@app/util/Generate";
  import AddressEditor from "./AddressEditor.svelte";
  import CheckboxGroupEditor from "./CheckboxGroupEditor.svelte";
  import RadioGroupEditor from "./RadioGroupEditor.svelte";
  import ConfigField from "./ConfigField.svelte";
  import DateInputEditor from "./DateInputEditor.svelte";

  export let field: IField;
</script>

<div>
  {#if field.type === 'string'}
    <ConfigField
      field={{ id: randomString(), type: 'number', label: 'Rows', value: { type: 'local', value: field.rows || 1 }, configFieldTarget: 'rows', configTarget: field.id }} />
  {:else if field.type === 'combobox'}
    <ComboBoxOptionsEditor {field} />
  {:else if field.type === 'address'}
    <AddressEditor {field} />
  {:else if field.type === 'checkbox-group'}
    <CheckboxGroupEditor {field} />
  {:else if field.type === 'radio-group'}
    <RadioGroupEditor {field} />
  {:else if field.type === 'switch'}
    <ConfigField
      field={{ id: randomString(), type: 'switch', label: 'Default Value', value: { type: 'local', value: field.defaultValue || false }, configFieldTarget: 'defaultValue', configTarget: field.id }} />
  {:else if field.type === 'date'}
    <DateInputEditor {field} />
  {/if}
</div>
