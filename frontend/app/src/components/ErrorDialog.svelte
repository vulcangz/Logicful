<script lang="typescript">
  import { dispatch, subscribeComponent } from "@app/event/EventBus";

  import { onMount } from "svelte";
  import Dialog from "./layout/Dialog.svelte";
  let error: string = "";

  subscribeComponent("unhandled_error", (err) => {
    error = err;
  });

  onMount(() => {
    window.onunhandledrejection = (e: any) => {
      console.error(e);
      const message = e.reason?.message ?? e.stack ?? e.message;
      dispatch("unhandled_error", message);
      //alert(`An unhandled exception has occured.\n\n${message}\n\nYour page may not work as expected.\n\nPlease try reloading the page.\n\nThis issue has been reported.`)
    };
  });
</script>

{#if error !== ""}
  <Dialog
    isOpen={true}
    title={"An error has occured"}
    actions={[
      {
        label: "Reload Page",
        onClick: () => {
          window.location.reload();
        },
        type: "primary",
      },
    ]}
  >
    <p>An uncaught exception has occured.</p>
    <p style="color: indigo">{error}</p>
    <p>
      Your page may not continue working as expected, please reload the page.
    </p>
  </Dialog>
{/if}
