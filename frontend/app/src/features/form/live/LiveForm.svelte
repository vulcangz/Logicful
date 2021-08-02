<script lang="typescript">
    // @ts-nocheck
    import LiveField from "@app/features/form/live/LiveField.svelte";
    import type {IForm} from "@app/models/IForm";
    import type {IField} from "@app/models/IField";
    import {subscribeFieldChange} from "@app/event/FieldEvent";
    import {LogicBuilder} from "@app/services/LogicBuilder";
    import {fastClone} from "@app/util/Compare";
    import {onMount} from "svelte";
    import {submitForm} from "@app/features/form/live/service/SubmitForm";
    import {LoadState} from "@app/models/LoadState";
    import {subscribeComponent} from "@app/event/EventBus";
    import Button from "@app/components/Button.svelte";
    import formStore from "@app/store/FormStore";
    import {getUrlParameter} from "../../../util/Http";

    export let form: IForm;
    export let mode: string = "";
    let state = LoadState.NotStarted;
    let uploadingFiles = false;
    let message: string;
    let redirectUrl: string;
    let canSubmit: boolean = true;

    subscribeFieldChange(onMount, (updatedField: IField) => {
        if (!form || !form.fields) {
            return;
        }
        const index = form.fields.findIndex((w) => w.id === updatedField.id);
        if (index === -1) {
            return;
        }
        form.fields[index].updated = !form.fields[index].updated;
        const fieldsWithRules = form.fields.filter((w) => {
            if (!w.logic || !w.logic.rules) {
                return false;
            }
            const hasRule = w.logic.rules.find(
                (rule) => rule.field === updatedField.id
            );
            return hasRule != null;
        });
        for (let fieldWithRule of fieldsWithRules) {
            let ruleIndex = form.fields.findIndex((w) => w.id === fieldWithRule.id);
            form.fields[ruleIndex].updated = !form.fields[ruleIndex].updated;
        }
    });

    function display(field: IField): boolean {
        if (!field.logic) {
            return true;
        }
        const builder = new LogicBuilder();
        return builder.evaluate(field);
    }

    subscribeComponent("submission_uploading_files", () => {
        uploadingFiles = true;
    });

    subscribeComponent("submission_uploading_files_finished", () => {
        uploadingFiles = false;
    });

    subscribeComponent("on_field_validate", (field: IField) => {
        if (field.error) {
            canSubmit = false;
            return;
        }
        const form = formStore.getForm();
        canSubmit = form.fields.find((w) => w.id !== field.id && w.error) == null;
    });

    async function onSubmit() {
        state = LoadState.Loading;
        try {
            if (form.id !== "demo" && !getUrlParameter("cypress_test")) {
                await submitForm();
            }
            state = LoadState.Finished;
            afterSubmission();
        } catch (ex) {
            if (ex.message === "validation_error") {
                state = LoadState.NotStarted;
                return;
            }
            console.error(ex);
            state = LoadState.Failed;
        }
    }

    function afterSubmission() {
        if (form.submissionConfig?.afterSubmitAction?.["Redirect to URL"]) {
            const url = form.submissionConfig?.afterSubmitConfig?.url;
            if (url) {
                redirectUrl = url;
                message =
                    "Thank you for your submission, you will be redirected in 5 seconds...";
                setTimeout(() => {
                    window.location.replace(url);
                }, 5000);
            }
        }
        if (form.submissionConfig?.afterSubmitAction?.["Show Message"]) {
            message = form.submissionConfig?.afterSubmitConfig?.message ?? "";
        }
        if (!message) {
            message = `Thank you for your submission.`;
        }
    }
</script>

<div class="bg-indigo-600 -m-6">
    <div class="max-w-screen-xl mx-auto px-3 sm:px-6 lg:px-8 h-6">
        <div class="flex items-center justify-between flex-wrap">
        </div>
    </div>
</div>

<div class="px-4 overflow-hidden sm:px-6 lg:px-8 py-12 mt-6">
    <div class="relative max-w-xl mx-auto">
        <svg class="absolute left-full transform translate-x-1/2" width="404" height="404" fill="none"
             viewBox="0 0 404 404">
            <defs>
                <pattern id="85737c0e-0916-41d7-917f-596dc7edfa27" x="0" y="0" width="20" height="20"
                         patternUnits="userSpaceOnUse">
                    <rect x="0" y="0" width="4" height="4" class="text-gray-200" fill="currentColor"/>
                </pattern>
            </defs>
            <rect width="404" height="404" fill="url(#85737c0e-0916-41d7-917f-596dc7edfa27)"/>
        </svg>
        <svg class="absolute right-full bottom-0 transform -translate-x-1/2" width="404" height="404" fill="none"
             viewBox="0 0 404 404">
            <defs>
                <pattern id="85737c0e-0916-41d7-917f-596dc7edfa27" x="0" y="0" width="20" height="20"
                         patternUnits="userSpaceOnUse">
                    <rect x="0" y="0" width="4" height="4" class="text-gray-200" fill="currentColor"/>
                </pattern>
            </defs>
            <rect width="404" height="404" fill="url(#85737c0e-0916-41d7-917f-596dc7edfa27)"/>
        </svg>
        <div class="bg-white shadow overflow-hidden sm:rounded-md pl-5 pr-5 pb-5 mt-10">
            <div class="py-5 border-b border-gray-200">
                <h2 class="text-2xl leading-7 font-bold text-gray-900">
                    {form.title}
                </h2>
                <p class="mt-1 max-w-2xl text-sm leading-5 text-gray-500">
                    {form.description}
                </p>
            </div>
            <div class="px-4 py-5 sm:p-0">
                {#if message}
                    <div class="mt-3">
                        <p class="text-base text-gray-900">
                            {@html message}
                        </p>
                        {#if redirectUrl}
                            <p class="text-base text-gray-900">
                                Not redirecting? <a href={redirectUrl} class="text-indigo underline">Click
                                here.</a>
                            </p>
                        {/if}
                    </div>
                {:else}
                    <form on:submit|preventDefault={onSubmit}>
                        <div>
                            <div id="form-preview-fields">
                                {#each form.fields as field (field.id)}
                                    {#if display(field)}
                                        <div id={`form-field-${field.id}`}>
                                            <LiveField field={fastClone(field)}/>
                                        </div>
                                    {:else}
                                        <div id={`form-field-${field.id}`}>
                                            <LiveField field={fastClone(field)} hidden={true}/>
                                        </div>
                                    {/if}
                                {/each}
                            </div>
                        </div>
                        <div class="mt-3">
          <span class="inline-flex rounded-md shadow-sm">
              {#if state === LoadState.NotStarted}
              <Button submit={true} type="primary" size="regular" disabled={!canSubmit} testId="submit-form-button">
                Submit
              </Button>
            {:else if state === LoadState.Failed}
              <Button submit={true} type="primary" testId="submit-form-button">
                Failed to Submit, Click To Try Again
              </Button>
            {:else if state === LoadState.Loading}
              <Button type="primary" disabled testId="submit-form-button">
                {#if uploadingFiles}Uploading Files...{:else}Submitting...{/if}
              </Button>
            {:else if state === LoadState.Finished}
              <Button type="primary" disabled testId="submit-form-button">Submitted Successfully.</Button>
            {/if}
          </span>
                        </div>
                    </form>
                {/if}
            </div>
        </div>
    </div>
</div>

