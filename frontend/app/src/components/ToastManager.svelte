<script lang="typescript">
  import { subscribeComponent } from "@app/event/EventBus";
  import { randomString } from "@app/util/Generate";
  import Toast from "./Toast.svelte";
  let toasts: { id: string; message: string; title: string }[] = [];

  subscribeComponent("show_toast", (props) => {
    const id = randomString();
    toasts = toasts.concat([
      { id, message: props.message, title: props.title },
    ]);
    setTimeout(() => {
      removeById(id);
    }, props.timeout || 4000);
  });

  subscribeComponent("toast_closed", (props) => {
    removeById(props.id);
  });

  function removeById(id: string) {
    toasts.splice(
      toasts.findIndex((t) => t.id === id),
      1
    );
    toasts = toasts;
  }
</script>

{#each toasts as toast}
  <Toast message={toast.message} title={toast.title} />
{/each}
