<script lang="typescript">
  import type { IField } from "@app/models/IField";
  import { firstNotEmpty } from "@app/util/Format";
  export let field: IField;
  export let showOptional: boolean = true;
</script>

{#if !field.hideLabel}
  <label
    for={field.id}
    class:text-red-900={field.error}
    class="block text-sm font-medium leading-5 text-gray-700 mb-2"
  >
    {firstNotEmpty(field.label, field.name)}
    {#if !field.required && showOptional}
      <span class="text-sm leading-5 text-gray-500">(optional)</span>
    {/if}
  </label>
{/if}
