<script lang="typescript">
  import Button from "@app/components/Button.svelte";
  import { dispatch, subscribeComponent } from "@app/event/EventBus";
  let saving = false;
  subscribeComponent("form_saved", () => {
    dispatch("show_toast", {
      message: "Successfully saved your form.",
    });
    setTimeout(() => {
      saving = false;
    }, 1000);
  });
</script>

{#if saving}
  <Button type="primary" disabled={true}>Saving...</Button>
{:else}
  <Button
    type="primary"
    onClick={() => {
      saving = true;
      dispatch('save_form', { status: 'draft' });
    }}>
    Save Form
  </Button>
{/if}
