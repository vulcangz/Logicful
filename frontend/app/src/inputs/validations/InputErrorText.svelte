<script lang="typescript">
  import type { IField } from "@app/models/IField";

  export let field: IField;
</script>

{#if field.error}
  <p class="mt-2 text-sm text-red-600">{field.error}</p>
{/if}
