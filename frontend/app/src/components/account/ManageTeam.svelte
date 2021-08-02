<script lang="typescript">
  import type { User } from "@app/models/User";

  import { me } from "@app/services/AuthService";
  import { onMount } from "svelte";

  let user: User;
  let teamId: string = "";

  onMount(async () => {
    user = await me();
    // teamId = user.teamId
  });
</script>

{#if user}
  <div>
    <div class="card border-light p-0 p-md-4 mb-4">
      <div class="card-body">
        <h3 class="h5 mb-0">Manage Team</h3>
        <form class="form mt-5" autocomplete="off">
          <div class="mb-4">
            <label for="inputPasswordOld">Current Password</label>
            <input
              type="password"
              class="form-control"
              id="inputPasswordOld"
              required />
          </div>
          <div class="mb-4">
            <label for="inputPasswordNew">New Password</label>
            <input
              type="password"
              class="form-control"
              id="inputPasswordNew"
              required />
            <span class="form-text small text-muted">
              The password must be 8-20 characters, and must <em>not</em> contain
              spaces.
            </span>
          </div>
          <div class="mb-4">
            <label for="inputPasswordNewVerify">Verify</label>
            <input
              type="password"
              class="form-control"
              id="inputPasswordNewVerify"
              required />
            <span class="form-text small text-muted">
              To confirm, type the new password again.
            </span>
          </div>
          <div class="form-group">
            <button type="submit" class="btn btn-dark">Save</button>
          </div>
        </form>
      </div>
    </div>
  </div>
{/if}
