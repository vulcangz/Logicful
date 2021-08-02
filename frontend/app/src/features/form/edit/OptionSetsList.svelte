<script lang="typescript">
  import RemoteTable from "@app/components/RemoteTable.svelte";
  import type { TableRow } from "@app/components/models/RemoteTableProps";
  import type { OptionSet } from "@app/models/OptionSet";
  import ManageOptionSets from "./ManageOptionSets.svelte";
  import { getApi } from "@app/services/ApiService";
  import FlyoutPanel from "./FlyoutPanel.svelte";

  let dialog = "";
  let editingId = "";

  async function getRows(): Promise<TableRow[]> {
    const data: OptionSet[] = await getApi("option-set");
    return data.map((d) => {
      return {
        id: d.id,
        Name: d.name,
        Value: d.value,
        Type: d.type === "local" ? "Inline" : "Remote",
        "Last Updated": new Date(d.changeTime).toLocaleString(),
        "Modified By": d.changeBy,
        "Forms Using": 3,
        Status: "Published",
      };
    });
  }

  function onClose() {
    editingId = "";
    dialog = "";
  }

  const hidden = new Set(["Value", "id"]);
</script>

{#if dialog === 'create'}
  <FlyoutPanel title="Create Option Set" {onClose}>
    <ManageOptionSets {editingId} isNew={!editingId} onSave={onClose} />
  </FlyoutPanel>
{:else}
  <RemoteTable
    headerActions={[{ label: '+ New Option Set', onClick: () => (dialog = 'create') }]}
    {getRows}
    {hidden}
    onRowClick={(row) => {
      editingId = row.id;
      dialog = 'create';
    }}
    onDelete={() => {
      alert('delete');
    }} />
{/if}
