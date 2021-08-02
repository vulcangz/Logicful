<script lang="typescript">
  import type { IField } from "@app/models/IField";
  import FlyoutPanel from "./FlyoutPanel.svelte";
  import GroupDetails from "./GroupDetails.svelte";
  import Button from "@app/components/Button.svelte";
  export let field: IField;
  let visible = false;
</script>

<Button type="primary" onClick={() => (visible = true)}>Manage Grouping</Button>

{#if visible}
  <FlyoutPanel
    title="Manage Field Grouping"
    description={'Group fields together and apply actions to all of them at once, such as showing a series of fields via a condition.'}
    onClose={() => (visible = false)}>
    <GroupDetails {field} />
  </FlyoutPanel>
{/if}
