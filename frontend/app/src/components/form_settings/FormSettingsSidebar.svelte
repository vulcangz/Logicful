<script lang="typescript">
  import { subscribeComponent } from "@app/event/EventBus";
  import type { IForm } from "@app/models/IForm";
  export let form: IForm;
  export let selected: string;

  subscribeComponent("form_loaded", (updatedForm) => {
    form = updatedForm;
  });
  subscribeComponent("form_updated", (updatedForm) => {
    form = updatedForm;
  });
</script>
