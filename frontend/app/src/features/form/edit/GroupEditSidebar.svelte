<script lang="typescript">
  import type { Group } from "@app/models/Group";

  import type { IForm } from "@app/models/IForm";

  import formStore from "@app/store/FormStore";

  import { onMount } from "svelte";
  import { randomString } from "@app/util/Generate";
  import ConfigField from "./ConfigField.svelte";

  export let groupId: string;
  let form: IForm;
  let group: Group;
  let index: number;

  onMount(() => {
    form = formStore.getForm();
    let groups = form.groups ?? [];
    index = groups.findIndex((group) => {
      return group.value === groupId;
    });
    group = groups[index];
  });
</script>

{#if group}
  <div style="padding-left: 0.5em;">
    <h5 style="padding-bottom: 0.2em;">Group Settings</h5>
    <hr />
  </div>
  <div style="padding: .75em 0.6em;">
    <ConfigField
      field={{ id: randomString(), required: true, label: 'Group Name', value: { type: 'local', value: group.label }, type: 'string', configFieldTarget: `groups[${index}].label`, configTarget: 'form' }} />

    <!-- Repeater for fields in group add and remove -->

    <!-- <LogicAccordion {field} /> -->
  </div>
{:else}
  <div class="spinner" />
{/if}
