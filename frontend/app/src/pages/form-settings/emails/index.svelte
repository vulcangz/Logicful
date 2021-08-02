<script context="module" lang="ts">
  import formStore from "../../../store/FormStore";
  import {IForm} from "../../../models/IForm";
  import {getApi} from "../../../services/ApiService";

  export async function preload(this: any, page: any, session: any) {
    const formId = page.query.formId;
    if (!formId || formId == "undefined" || formId == "null") {
      return this.error(400, "Invalid Form Id");
    }
    const form: IForm = await getApi(`form/${formId}`);
    form.id = formId;
    formStore.setForm(form);

    return { form: form };
  }
</script>

<script lang="typescript">
  import { onMount } from "svelte";
  import type { IForm } from "@app/models/IForm";
  import FormSettings from "@app/components/form_settings/FormSettings.svelte";

  export let form: IForm;
  let selected: string = "emails";

  onMount(() => {});
</script>

<div>
  <FormSettings {form} {selected} />
</div>
