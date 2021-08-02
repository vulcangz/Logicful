<script lang="typescript">
  import type { User } from "@app/models/User";

  import { me } from "@app/services/AuthService";
  import { onMount } from "svelte";

  let user: User;
  let displayName: string = "";
  let fullName: string = "";
  let email: string = "";

  onMount(async () => {
    user = await me();
    displayName = user.displayName;
    fullName = user.fullName;
    email = user.email;
  });
</script>

{#if user}
  <div>
    <div class="row">
      <div class="col-lg-12">
        <div
          class="card card-body bg-white border-light mb-4"
          style="padding-left: 2em; padding-right: 2em;">
          <h2 class="h5 mb-4 mt-2">General Information</h2>
          <form>
            <div class="row">
              <div class="col-md-6 mb-3">
                <div class="mb-3">
                  <label for="first_name">Full Name</label>
                  <input
                    bind:value={fullName}
                    class="form-control"
                    id="first_name"
                    type="text" />
                </div>
              </div>
              <div class="col-md-6 mb-3">
                <div class="mb-3">
                  <label for="last_name">Display Name</label>
                  <input
                    bind:value={displayName}
                    class="form-control"
                    id="last_name"
                    type="text" />
                </div>
              </div>
            </div>
            <div class="row">
              <div class="col-md-6 mb-3">
                <div class="mb-3">
                  <label for="email">Email</label>
                  <input
                    bind:value={email}
                    class="form-control"
                    id="email"
                    type="email" />
                </div>
              </div>
              <div class="col-md-6 mb-3" />
            </div>
            <div class="mt-1">
              <button type="submit" class="btn btn-primary">Save</button>
            </div>
          </form>
        </div>
      </div>
    </div>

    <div class="card bg-white border-light mb-4">
      <div class="card-body" style="padding-left: 2em; padding-right: 2em;">
        <h3 class="h5 mb-4 mt-2">Change Password</h3>
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
              Password must be 6 characters minimum.</span>
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
