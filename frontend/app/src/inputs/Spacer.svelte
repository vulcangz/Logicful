<script lang="typescript">
  import type { IField } from "@app/models/IField";
  import { subscribeFieldChange } from "@app/event/FieldEvent";
  import { onMount } from "svelte";
  import { toNumberOrDefault } from "@app/util/Compare";
  import formStore from "@app/store/FormStore";

  export let field: IField;
  export let value = 0;

  subscribeFieldChange(onMount, (newField) => {
    if (newField.id === field.id) {
      value = toNumberOrDefault(newField.options?.spacer ?? 1);
    }
  });

  onMount(() => {
    value = toNumberOrDefault(formStore.get(field.id)?.options?.spacer ?? 1);
  });
</script>

<div style={`margin-bottom:${value}em`} />
