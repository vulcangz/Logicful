<script lang="typescript">
  import type { IField } from "@app/models/IField";
  import OptionSetsList from "./OptionSetsList.svelte";
  import { randomString } from "@app/util/Generate";
  import ConfigField from "./ConfigField.svelte";
  import Button from "@app/components/Button.svelte";
  import FlyoutPanel from "./FlyoutPanel.svelte";

  export let field: IField;
  let dialog: "option_sets" | "" = "";

  function loadTransformer(value: any[]) {
    return value.map((v) => {
      return {
        label: v.name,
        value: v.value,
      };
    });
  }
</script>

{#if dialog === 'option_sets'}
  <FlyoutPanel title={'Manage Option Sets'} onClose={() => (dialog = '')}>
    <OptionSetsList />
  </FlyoutPanel>
{:else}
  <ConfigField
    field={{ id: randomString(), loadTransformer: loadTransformer, required: true, label: 'Option Set', value: field.options, name: `${field.id}-builder-config-field-field_editor-options`, type: 'combobox', options: { type: 'remote', value: `option-set`, isOurApi: true }, configFieldTarget: 'options', configTarget: field.id }} />

  <div class="ml-3 mt-3">
    <Button type="primary" onClick={() => (dialog = 'option_sets')}>
      Manage Option Sets
    </Button>
  </div>
{/if}
