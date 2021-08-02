<script lang="typescript">
  import type { IField } from "@app/models/IField";
  import Field from "./Field.svelte";

  export let field: IField;
  export let config: any = {};
</script>

<div class="p-3">
  <Field {field} {config} />
</div>
