<script lang="ts">
  import Button from "@app/components/Button.svelte";
  export let folderId: string;
</script>

<div class="flex">
  <div>
    <p class="mt-2 text-base leading-4 font-extrabold text-gray-900">
      Create New Form
    </p>
    <p class="mt-5 text-base leading-4 text-gray-500">
      You have no forms in this folder, click below to create one now.
    </p>
    <div class="mt-5">
      <Button
        type="primary"
        size="regular"
        href={`/form/create?folder=${folderId}`}>
        Create New Form
      </Button>
    </div>
  </div>
</div>
