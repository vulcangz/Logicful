<script lang="ts">
  import Dialog from "@app/components/layout/Dialog.svelte";
  import type { ButtonAction } from "@app/components/models/ComponentProps";
  import { postApi } from "@app/services/ApiService";
  import { getUrlParameter } from "@app/util/Http";
  import { onMount } from "svelte";
  import MyTeamPage from "./MyTeamPage.svelte";

  let invalid = false;
  let showLogoutPrompt = false;
  let team: string = "";
  let name: string = "";
  let actions: ButtonAction[] = [];
  let logoutActions: ButtonAction[] = [
    {
      label: "Logout Now",
      onClick: () => {
        sessionStorage.setItem("redirect_after_login", "/team");
        window.location.replace("/account/logout");
      },
      type: "primary",
    },
  ];

  async function join() {
    await postApi<any>(`team/members/invite/accept?team=${team}`, {});
    showLogoutPrompt = true;
    actions = [];
  }

  onMount(() => {
    team = getUrlParameter("team") ?? "";
    name = getUrlParameter("name") ?? "";
    if (!team || !name) {
      invalid = true;
    }
    actions = [
      {
        label: "Join Team",
        onClick: join,
        type: "primary",
      },
      {
        label: "Cancel",
        type: "secondary",
      },
    ];
  });
</script>

<MyTeamPage />

{#if actions.length > 0}
  <Dialog isOpen={true} title="Accept Team Invite" {actions}>
    {#if invalid}
      <p>
        Invalid invite, please try clicking Join Team from your invite email
        again.
      </p>
    {:else}
      <p>Click below to join <strong>{name}.</strong></p>
      <p class="mt-3">
        If you are currently on a team, you will be <strong>removed</strong> from
        that team when joining this one.
      </p>
    {/if}
  </Dialog>
{/if}

{#if showLogoutPrompt}
  <Dialog isOpen={true} title="Succesfully Joined Team" actions={logoutActions}>
    <p>
      Successfully joined team, to see changes, you must <strong>log out</strong>
      and log back in.
    </p>
  </Dialog>
{/if}
