<script lang="typescript">
  import "dragula/dist/dragula.min.css";
  import FormBuilder from "@app/features/form/edit/FormBuilder.svelte";
  import FormSidebar from "@app/features/form/edit/FormSidebar.svelte";
  import FieldEditSidebar from "@app/features/form/edit/FieldEditSidebar.svelte";
  import RightSidebar from "@app/features/form/edit/RightSidebar.svelte";
  import AlphaBanner from "@app/components/AlphaBanner.svelte";
</script>

<AlphaBanner />
<div class="h-screen flex overflow-hidden bg-white">
  <!-- Static sidebar for desktop -->
  <FormSidebar />
  <div class="flex flex-col w-0 flex-1 overflow-hidden">
    <main
      class="flex-1 relative z-0 overflow-y-auto focus:outline-none"
      tabindex="0">
      <div class="pt-2 pb-6 md:py-2">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 md:px-8">
          <div class="pb-2">
            <FormBuilder />
          </div>
        </div>
      </div>
    </main>
  </div>
  <RightSidebar>
    <FieldEditSidebar />
  </RightSidebar>
</div>
