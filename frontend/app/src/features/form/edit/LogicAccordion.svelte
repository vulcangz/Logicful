<script lang="typescript">
  import LogicBuilder from "./LogicBuilder.svelte";
  import type { IField } from "@app/models/IField";
  import FlyoutPanel from "./FlyoutPanel.svelte";
  import Button from "@app/components/Button.svelte";
  export let field: IField;
  let visible = false;
</script>

<Button type="primary" onClick={() => (visible = true)}>Manage Logic</Button>

{#if visible}
  <FlyoutPanel
    title="Manage Field Logic"
    description={'Add rules to change the behavior of when your field should show or hide.'}
    onClose={() => (visible = false)}>
    <LogicBuilder {field} />
  </FlyoutPanel>
{/if}
