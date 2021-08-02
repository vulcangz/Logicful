<script lang="typescript">
  import Button from "@app/components/Button.svelte";
  import type { IForm, Integration } from "@app/models/IForm";
  import SendEmailIntegration from "./integrations/SendEmailIntegration.svelte";
  import IntegrationStepper from "./integrations/IntegrationStepper.svelte";
  import { dispatch, subscribeComponent } from "@app/event/EventBus";
  import SendWebhookIntegration from "./integrations/SendWebhookIntegration.svelte";

  export let form: IForm;
  let configuring: Integration;
  let index: number;

  function isEnabled(index: number): boolean {
    return form.workflow?.integrations?.[index]?.enabled ?? false;
  }

  const integrations: Integration[] = [
    {
      name: "integration-email",
      label: "Send Email On Submission",
      description:
        "When a submission is received, an email will be sent out to your desired email address containing the submission data.",
      config: {},
      editor: SendEmailIntegration,
    },
    {
      name: "integration-webhook",
      label: "Send Webhook On Submission",
      description:
        "When a submission is received, an HTTP POST request will be sent to the desired url.",
      config: {},
      editor: SendWebhookIntegration,
    },
  ];

  subscribeComponent("save_form", () => {
    setTimeout(() => {
      dispatch("show_toast", {
        message: `Successfully updated integration.`,
      });
      steps[0].active = true;
      steps[1].active = false;
    }, 250);
  });

  subscribeComponent("integration_configured", (value: boolean) => {
    steps[1].completed = value;
  });

  function configure(integration: Integration, i: number) {
    configuring = integration;
    index = i;
    steps[0].completed = true;
    steps[0].active = false;
    steps[1].active = true;
  }

  function onStepClick(name: string) {
    if (name === "Configure" && !configuring) {
      return;
    }
    steps = steps.map((s) => {
      s.active = s.name === name;
      return s;
    });
  }

  let steps = [
    {
      name: "Select",
      active: true,
      completed: false,
    },
    {
      name: "Configure",
      active: false,
      completed: false,
    },
  ];
</script>

<IntegrationStepper {steps} onClick={onStepClick} />

{#if steps[0].active}
  {#each integrations as integration, i}
    <div class="bg-white py-5 border-b border-gray-200">
      <div
        class="-ml-4 -mt-4 flex justify-between items-center flex-wrap
          sm:flex-nowrap">
        <div class="ml-4 mt-4 w-10/12">
          <h3 class="text-lg leading-6 font-medium text-gray-900">
            {integration.label}
            {#if isEnabled(i)}
              <span
                class="inline-flex items-center px-3 py-0.5 rounded-full text-sm
                  font-medium leading-5 bg-indigo-100 text-indigo-800">
                Enabled
              </span>
            {/if}
          </h3>
          <p class="mt-1 text-sm leading-5 text-gray-500">
            {integration.description}
          </p>
        </div>
        <div class="mt-4 flex-shrink-0">
          <span class="inline-flex rounded-md shadow-sm">
            <Button type="primary" onClick={() => configure(integration, i)}>
              Configure
            </Button>
          </span>
        </div>
      </div>
    </div>
  {/each}
{/if}

{#if steps[1].active}
  <div class="mt-4">
    <svelte:component
      this={configuring.editor}
      {form}
      {index}
      integration={configuring} />
  </div>
{/if}
