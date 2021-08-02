<script lang="typescript">
    import type { IField } from "@app/models/IField";
  import FlyoutPanel from "./FlyoutPanel.svelte";
  import Button from "@app/components/Button.svelte";
  import ValidationsBuilder from "./ValidationsBuilder.svelte";
  export let field: IField;
  let visible = false;
</script>

<Button type="primary" onClick={() => (visible = true)}>
  Manage Validations
</Button>

{#if visible}
  <FlyoutPanel
    title="Manage Validations"
    description={'Add validations to limit allowed values in your inputs.'}
    onClose={() => (visible = false)}>
    <ValidationsBuilder {field} />
  </FlyoutPanel>
{/if}
