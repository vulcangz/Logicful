<script lang="typescript">
  import { onMount } from "svelte";
  import { me } from "@app/services/AuthService";
  import { navigate } from "svelte-routing";

  onMount(async () => {
    const user = await me();
    if (!user.id) {
      sessionStorage.setItem(
        "redirect_after_login",
        window.location.pathname + window.location.search
      );
      navigate("/account/login");
    }
  });
</script>

<div>
  <slot />
</div>
