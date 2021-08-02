<script lang="typescript">
  import type { IForm } from "@app/models/IForm";

  import { onMount } from "svelte";

  let recentSubmissions: string[] = [];
  let forms: IForm[] = [];

  onMount(() => {
    // forms = getRecentlyViewedSubmissions()
  });

  function getRecentlyViewedSubmissions() {
    return ["submission1"];
  }

  function getRecentFroms() {
    let form1: IForm = {
      title: "Form 1",
      id: "yeet",
      changeTime: "Oct 2 3:30 PM",
      fields: [],
    };
    let form2: IForm = {
      title: "Form 2",
      id: "yeet",
      changeTime: "Oct 12 3:30 PM",
      fields: [],
    };
    let form3: IForm = {
      title: "Form 3",
      id: "yeet",
      changeTime: "Oct 22 3:30 PM",
      fields: [],
    };

    return [form1, form2, form3];
  }
</script>

<div style="width: 100%">
  <h5 class="h5">Recently Viewed Form Submissions</h5>
  <div class="card-body px-0 pt-0">
    <ul class="list-group list-group-flush">
      {#each forms as form}{/each}
    </ul>
  </div>
</div>

<style>
  .h5 {
    padding-bottom: 0.2em;
  }
</style>
