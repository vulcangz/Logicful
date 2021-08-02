<script lang="typescript">
  import type { IField } from "@app/models/IField";

  export let field: IField;
</script>

{#if field.helperText}
  <p class="mt-2 text-sm text-gray-500">
    {@html field.helperText ?? ""}
  </p>
{/if}
