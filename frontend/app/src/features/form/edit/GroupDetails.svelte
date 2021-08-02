<script lang="typescript">
  import type { IField} from "@app/models/IField";
  import { randomString } from "@app/util/Generate";
  import { onMount } from "svelte";
  import formStore from "@app/store/FormStore";
  import ConfigField from "./ConfigField.svelte";
  import Button from "@app/components/Button.svelte";
  import type { Group } from "@app/models/Group";

  export let groups: Group[] = [];
  export let field: IField;

  onMount(() => {
    let form = formStore.getForm();
    groups = form?.groups ?? [];
  });
</script>

<div>
  {#if groups.length === 0}
    <p class="pl-3 mb-3">
      You have not created any groups, click <strong>Manage Groups</strong> to add
      one now.
    </p>
  {:else}
    <ConfigField
      config={{ search: true }}
      field={{ id: randomString(), label: 'Specify Group', helperText: 'Select a group to add this field to.', value: { type: 'local', value: field.groupId }, type: 'combobox', required: true, configFieldTarget: `groupId`, configTarget: field.id, options: { type: 'local', value: groups } }} />
  {/if}
  <div class="pl-3">
    <Button type="primary">Manage Groups</Button>
  </div>
</div>
