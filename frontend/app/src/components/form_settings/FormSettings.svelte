<script lang="ts">
  import FieldConfigSidebarHeader from "@app/features/form/edit/FieldConfigSidebarHeader.svelte";

  import type { IForm } from "@app/models/IForm";
	import FormSettingsSidebar from "./FormSettingsSidebar.svelte";

  export let form: IForm;
  export let selected: string;
</script>

<FieldConfigSidebarHeader
  title={`Form Settings`}
  description={`Configure settings for your form.`} />
<div class="sm:mx-auto sm:w-full">
  <div class="px-4 sm:rounded-lg sm:px-2">
    {#if form != null}
      <FormSettingsSidebar {form} {selected} />
    {/if}
  </div>
</div>
