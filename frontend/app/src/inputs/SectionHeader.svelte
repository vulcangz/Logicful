<script lang="typescript">
  import type { IField } from "@app/models/IField";
  export let field: IField;
</script>

<div>
  <h3 class="text-lg leading-6 font-medium text-gray-900">
    {field.header ?? ""}
  </h3>
  {#if field.helperText}
    <p class="mt-1 max-w-2xl text-sm leading-5 text-gray-500">
      {field.helperText}
    </p>
  {/if}
</div>
