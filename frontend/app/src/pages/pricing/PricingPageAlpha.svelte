<script lang="ts">
		import AlphaBanner from "@app/components/AlphaBanner.svelte";
		import MarketingNavbar from "@app/components/layout/MarketingNavbar.svelte";
				import MarketingFooter from "../../components/layout/MarketingFooter.svelte";
</script>

<AlphaBanner />

<div class="bg-white">
    <div class="relative bg-white">
        <div class="max-w-7xl mx-auto px-4 sm:px-6">
            <div class="relative pt-6 pb-12 lg:pb-20">
                <MarketingNavbar />
            </div>
        </div>
    </div>
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <div class="max-w-3xl mx-auto">
            <h3>Pricing Coming Soon</h3>
            <p>Due to Logicful being in early alpha, pricing is still being decided.</p>
            <p>See our <a class="underline text-blue-500" href="/faq">FAQ page for more information.</a></p>
        </div>
    </div>
    <MarketingFooter/>
</div>
