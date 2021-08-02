<script lang="ts">
		import AlphaBanner from "@app/components/AlphaBanner.svelte";
		import MarketingNavbar from "@app/components/layout/MarketingNavbar.svelte";
				import MarketingFooter from "../../components/layout/MarketingFooter.svelte";
</script>

<AlphaBanner/>

<div class="bg-white">
    <div class="relative bg-white">
        <div class="max-w-7xl mx-auto px-4 sm:px-6">
            <div class="relative pt-6 pb-12 lg:pb-20">
                <MarketingNavbar/>
            </div>
        </div>
    </div>
    <div class="bg-white">
        <div class="max-w-screen-xl mx-auto py-16 px-4 sm:px-6 lg:py-20 lg:px-8">
            <div class="lg:grid lg:grid-cols-3 lg:gap-8">
                <div class="space-y-4">
                    <h2 class="text-3xl leading-9 font-extrabold text-gray-900">
                        Frequently asked questions
                    </h2>
                </div>
                <div class="mt-12 lg:mt-0 lg:col-span-2">
                    <dl class="space-y-12">
                        <div class="space-y-2">
                            <dt class="text-lg leading-6 font-medium text-gray-900">
                                What does 'Early Alpha Stage' mean?
                            </dt>
                            <dd class="text-base leading-6 text-gray-500">
                                This means that Logicful is still under active development and should not be used for production use cases.
                            </dd>
                        </div>
                        <div class="space-y-2">
                            <dt class="text-lg leading-6 font-medium text-gray-900">
                                Is Logicful free to use?
                            </dt>
                            <dd class="text-base leading-6 text-gray-500">
                                Logicful is free to use during Alpha period.
                            </dd>
                            <dd class="text-base leading-6 text-gray-500">
                                Once out of alpha and stable enough for production use, a monthly fee will be required.
                            </dd>
                        </div>
                        <div class="space-y-2">
                            <dt class="text-lg leading-6 font-medium text-gray-900">
                                What is Logicful?
                            </dt>
                            <dd class="text-base leading-6 text-gray-500">
                               <p>Logicful is a powerful fully managed dynamic form builder software.</p>
                                <p>See our <a href="/features" class="underline text-blue-500">features</a> page for more information.</p>
                            </dd>
                        </div>
                    </dl>
                </div>
            </div>
        </div>
    </div>
    <MarketingFooter/>
</div>
