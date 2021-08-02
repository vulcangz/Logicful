<script lang="typescript">
  import { onMount } from "svelte";

  let show = true;
  onMount(() => {
    setTimeout(() => {
      show = false;
    }, 150);
  });
</script>

{#if show}
  <div
    class="preloader bg-dark flex-column justify-content-center
      align-items-center"
    id="preloader">
    <div class="d-flex justify-content-center">
      <div
        class="spinner-border"
        style="width: 3rem; height: 3rem;"
        role="status">
        <span class="sr-only">Loading...</span>
      </div>
    </div>
  </div>
{/if}
