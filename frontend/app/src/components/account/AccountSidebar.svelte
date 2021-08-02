<script lang="typescript">
  import type { User } from "@app/models/User";

  import { me } from "@app/services/AuthService";

  import { onMount } from "svelte";
  import Link from "../Link.svelte";

  export let page: string = "";
  let user: User;

  onMount(async () => {
    user = await me();
  });
</script>

{#if user}
  <div class="card border-light p-2">
    <div class="card-header bg-white border-0" style="display: flex;">
      <!-- <span><i class="fas fa-user"></i></span> -->
      <h2 class="h5" style="padding-left: 0em;">Hello, {user.displayName}!</h2>
    </div>
    <div class="card-body p-2">
      <div class="list-group dashboard-menu list-group-sm">
        <div style="display: flex;">
          <Link
            href="/account/settings"
            class="d-flex list-group-item border-0 list-group-item-action {page === 'settings' ? 'active' : ''}">
            <i class="fas fa-cogs" /><span style="padding-left: 0.5em;">
              Profile Settings</span><span class="icon icon-xs ml-auto"><span
                class="fas fa-chevron-right" /></span>
          </Link>
        </div>
        <div style="display: flex;">
          <Link
            href="/account/settings/team"
            class="d-flex list-group-item border-0 list-group-item-action {page === 'manage-team' ? 'active' : ''}">
            <i class="fas fa-users" />
            <span style="padding-left: 0.5em;"> Manage Team</span><span class="icon
                icon-xs ml-auto"><span class="fas fa-chevron-right" /></span>
          </Link>
        </div>
        <div style="display: flex;">
          <Link
            href="/account/settings/billing"
            class="d-flex list-group-item border-0 list-group-item-action {page === 'billing' ? 'active' : ''}">
            <i class="fas fa-wallet" />
            <span style="padding-left: 0.5em;"> Billing</span><span class="icon
                icon-xs ml-auto"><span class="fas fa-chevron-right" /></span>
          </Link>
        </div>
      </div>
    </div>
  </div>
{/if}
