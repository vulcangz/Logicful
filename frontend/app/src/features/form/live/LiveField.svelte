<script lang="typescript">
  import type { IField } from "@app/models/IField";
  export let field: IField;
  export let hidden: boolean = false;
  import Field from "../edit/Field.svelte";
</script>

<div class:hidden class="pt-3 pb-3">
  <Field {field} mode="live" />
</div>
