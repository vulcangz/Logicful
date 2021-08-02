<script lang="typescript">
  import Button from "@app/components/Button.svelte";
  import RemoteTable from "@app/components/RemoteTable.svelte";
  import type { Team } from "@app/models/IForm";
  import type { User } from "@app/models/User";
  import { getApi } from "@app/services/ApiService";
  import { onMount } from "svelte";
  import AddMembers from "./AddMembers.svelte";
  import RemoveMembers from "./RemoveMembers.svelte";
  import {DateTime} from "luxon";

  let team: Team = {
    name: "",
    id: "",
  };
  let dialog: "removing" | "adding" | "" = "";
  let hidden: Set<string>;

  onMount(async () => {
    team = await getApi("team");
    hidden = new Set(["userId", 'table_meta_id']);
  });

  async function getRows() {
    const users: User[] = await getApi("team/members");
    return users.map((u) => {
      return {
        Email: u.email,
        Name: u.fullName,
        "Register Date": DateTime.fromISO(u.creationDate).toLocaleString(DateTime.DATE_MED),
        "Last Online": u.lastActive ? DateTime.fromISO(u.lastActive).toLocaleString(DateTime.DATE_MED) : 'Never',
        userId: u.id,
      };
    });
  }
</script>

<div class="bg-white px-4 py-5 border-b border-gray-200 sm:px-6">
  <div
    class="-ml-4 -mt-4 flex justify-between items-center flex-wrap
      sm:flex-nowrap">
    <div class="ml-4 mt-4">
      <h3 class="text-lg leading-6 font-medium text-gray-900">{team?.name}</h3>
      <p class="mt-1 text-sm leading-5 text-gray-500">
        All resources created on this team are shared with all team members.
      </p>
    </div>
    <div class="ml-4 mt-4 flex-shrink-0">
      <span class="inline-flex rounded-md shadow-sm">
        <Button type="primary" onClick={() => (dialog = 'adding')}>
          Add Member
        </Button>
        <div class="ml-2">
          <Button type="secondary">Leave Team</Button>
        </div>
      </span>
    </div>
  </div>

  <div class="pt-8 w-full">
    <RemoteTable {getRows} {hidden}>
      <div slot="selected_actions" class="ml-3">
        <Button
          type="secondary"
          size="small"
          onClick={() => (dialog = 'removing')}>
          Remove Member(s)
        </Button>
      </div>
    </RemoteTable>
  </div>
</div>

{#if dialog === 'removing'}
  <RemoveMembers {team} onClose={() => (dialog = '')} />
{/if}

{#if dialog === 'adding'}
  <AddMembers {team} onClose={() => (dialog = '')} />
{/if}
