<svelte:options immutable />

<script lang="ts">
  export let title: string;
  export let paragraph: string;
</script>

<h1>{title}</h1>
<p>{paragraph}</p>

<style>
  h1 {
    color: #3f83f8;
    font-size: 1.875rem;
    font-weight: 200;
    margin: 0;
  }

  p {
    color: #374151;
    margin: 0;
    margin-top: 1rem;
  }
</style>
