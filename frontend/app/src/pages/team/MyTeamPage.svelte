<script lang="typescript">
  import Shell from "@app/components/Shell.svelte";
  import type { User } from "@app/models/User";
  import { emptyUser, me } from "@app/services/AuthService";
  import { onMount } from "svelte";
  import NoTeam from "./NoTeam.svelte";
  import TeamDetails from "./TeamDetails.svelte";

  let user: User = emptyUser;

  onMount(async () => {
    user = await me();
  });
</script>

<Shell>
  {#if user.teamId === user.id}
    <NoTeam />
  {:else}
    <TeamDetails />
  {/if}
</Shell>
