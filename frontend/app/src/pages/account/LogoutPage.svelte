<script lang="typescript">
  import Loader from "@app/components/Loader.svelte";
  import Shell from "@app/components/Shell.svelte";
  import { onMount } from "svelte";

  let error = false;

  function signOut(attempt: number) {
    //@ts-ignore
    if (!window.gapi) {
      setTimeout(signOut, 100);
      return;
    }
    //@ts-ignore
    window.gapi.load("auth2", function () {
      try {
        //@ts-ignore
        window.gapi.auth2.init();
        //@ts-ignore
        var auth2 = window.gapi.auth2.getAuthInstance();
        setTimeout(() => {
          auth2
            .signOut()
            .then(function () {
              localStorage.removeItem("token");
              window.location.replace("/account/login");
            })
            .catch((ex: any) => {
              console.error(ex);
              if (attempt > 1) {
                error = true;
                return;
              }
              setTimeout(() => {
                signOut(attempt + 1);
              }, 100);
            });
        }, 500);
      } catch (ex) {
        console.error(ex);
        error = true;
      }
    });
  }

  onMount(() => {
    setTimeout(() => {
      signOut(1);
    }, 500);
  });
</script>

<svelte:head>
  <meta
    name="google-signin-client_id"
    content="807768276065-c8b2jvlks20jgm0mk3t2akfm7pc3jomu.apps.googleusercontent.com" />
  <script src="https://apis.google.com/js/platform.js" async defer>
  </script>
</svelte:head>

<Shell sidebar={false}>
  {#if error}
    <p>Failed to sign out, please try refreshing the page.</p>
  {:else}
    <Loader />
    <h3>Signing out... please wait.</h3>
  {/if}
</Shell>
