<script lang="typescript">
  import Dialog from "@app/components/layout/Dialog.svelte";
  import type { ButtonAction } from "@app/components/models/ComponentProps";
  import type { Team } from "@app/models/IForm";
  import { deleteApi } from "@app/services/ApiService";
  import { tableRows } from "@app/store/RemoteTableStore";
  import { onMount } from "svelte";

  export let team: Team;
  export let onClose: () => any;
  let userIds: string[] = [];

  async function remove() {
    await deleteApi(`team/members/remove`, {
      userIds,
    });
  }

  let actions: ButtonAction[] = [
    {
      label: "Confirm Removal",
      type: "primary",
      focus: true,
      onClick: remove,
    },
    {
      label: "Cancel",
      type: "secondary",
    },
  ];

  onMount(() => {
    const unsubscribe = tableRows.subscribe((rows) => {
      userIds = rows.filter((w) => w.meta_selected).map((w) => w.userId);
    });
    return () => {
      unsubscribe();
    };
  });
</script>

{#if userIds.length > 0}
  <Dialog
    isOpen={true}
    {actions}
    title={`Are you sure you want to remove ${userIds.length} teammate(s)?`}
    {onClose}>
    <p>Teammates will need to be re-invited to join this team again.</p>
  </Dialog>
{/if}
