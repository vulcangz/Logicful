<script lang="typescript">
  import type {
    LogicConditional,
    LogicRuleOptions,
  } from "@app/models/LogicBuilder";
  import type { IField } from "@app/models/IField";
  import type { LabelValue } from "@app/models/IField";
  import { onMount } from "svelte";
  import { dispatchSingle } from "@app/event/EventBus";
  import { randomString } from "@app/util/Generate";
  import { subscribeFieldChange } from "@app/event/FieldEvent";
  import formStore from "@app/store/FormStore";
  import { firstNotEmpty } from "@app/util/Format";
  import { fastClone} from "@app/util/Compare";
  import ConfigField from "@app/features/form/edit/ConfigField.svelte";
  import Button from "@app/components/Button.svelte";
  import JavascriptExpressionViewer from "@app/components/JavascriptExpressionViewer.svelte";
  import {getFieldLogicConditions} from "../../../services/FieldLogicConditions";

  export let helperText: string = "";
  export let field: IField;
  let options: LogicRuleOptions[] = [];

  subscribeFieldChange(onMount, (newField, change) => {
    if (change.field === "value") {
      return;
    }
    if (field.id === newField.id) {
      field = newField;
      if (field.validations) {
        field.validations.forEach((f, i) => {
          options[i] = getOptions(i);
        });
      }
    }
  });

  function getFields(): IField[] {
    let fields = dispatchSingle<IField[]>("get_form_fields", {});
    fields = fields.filter((w) => w.id !== field.id && w.type !== "spacer");
    return fields;
  }

  function remove(option: number) {
    const temp = fastClone(field.validations);
    temp.splice(option, 1);
    field.validations = temp;
    formStore.set(field);
  }

  function addNew() {
    field.validations = field.validations ?? [];
    field.validations = field.validations.concat([
      {
        errorIfTrue: {
          "Field is valid if the condition passes.":
            "Field is valid if the condition passes.",
        },
        value: "",
        condition: "",
        message: "",
      },
    ]);
    formStore.set(field);
  }

  function shouldShowValue(index: number) {
    const condition = field.validations?.[index]?.condition ?? "";
    const toNotShow = ["hasValue", "isTrue", "isFalse", "notHaveValue", "isEmail"];
    if (toNotShow.includes(condition)) {
      return false;
    }
    return true;
  }

  function getOptions(index: number): LogicRuleOptions {
    const value = field.validations?.[index]?.condition;
    const condition = conditions().find((w) => w.value === value);
    if (!condition) {
      return {
        valueType: "text",
        showValue: false,
        options: () => Promise.resolve([]),
        helperText: "",
        placeholder: "",
      };
    }
    return {
      valueType: condition.valueInput ?? "text",
      showValue: shouldShowValue(index),
      options: () => loadOptions(index),
      helperText: condition.helper ?? "",
      placeholder: condition.placeholder ?? "",
    };
  }

  function loadOptions(index: number): Promise<LabelValue[]> {
    const id = field.id;
    if (!id) {
      return Promise.resolve([]);
    }
    const options = dispatchSingle("get_options", {
      id,
    });
    return options as any;
  }

  function conditions(): LogicConditional[] {
    return getFieldLogicConditions({
      field,
      loadOptions
    });
  }

  function fieldsTransformer(fields: IField[]): LabelValue[] {
    return fields.map((f) => {
      return {
        label: `${firstNotEmpty(f.label, f.name)} - ${f.type}`,
        value: f.id,
      };
    });
  }
</script>

<div>
  {#each field.validations ?? [] as option, i}
    <div class="bg-gray-100 overflow-visible mt-3 relative rounded-md">
      <div class="absolute top-2 right-2" on:click={() => remove(i)}>
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
          class="w-6 h-6">
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
        </svg>
      </div>
      <div class="sm:p-2">
        <div>
          <div>
            <ConfigField
              config={{ search: true }}
              field={{ id: randomString(), label: 'Select Your Validation', value: { type: 'local', value: field.validations?.[i]?.condition }, type: 'combobox', required: true, configFieldTarget: `validations[${i}].condition`, configTarget: field.id, options: { type: 'local', value: conditions() } }} />
          </div>
          <div>
            <ConfigField
              config={{ search: true }}
              field={{ id: randomString(), label: 'Error Message To Show When Validation Fails', value: { type: 'local', value: field.validations?.[i]?.message ?? '' }, type: 'string', required: true, configFieldTarget: `validations[${i}].message`, configTarget: field.id, options: { type: 'local', value: conditions() } }} />
          </div>
          <ConfigField
            field={{ id: randomString(), required: true, value: field.validations?.[i]?.errorIfTrue, label: 'Select whether the field is valid or invalid if the condition passes.', type: 'radio-group', options: ['Field is valid if the condition passes.', 'Field is invalid if the condition passes.'], configTarget: field.id, configFieldTarget: `validations[${i}].errorIfTrue` }} />
          <div>
            {#if field.validations?.[i]?.condition && options[i]?.showValue}
              {#if options[i].valueType === 'text'}
                <ConfigField
                  field={{ id: randomString(), helperText: options[i].helperText, placeholder: options[i].placeholder, label: 'Provide Value For Conditional', value: { type: 'local', value: field.validations?.[i]?.value }, type: 'string', required: true, configFieldTarget: `validations[${i}].value`, configTarget: field.id }} />
              {:else if options[i].valueType === 'combobox'}
                <ConfigField
                  field={{ id: randomString(), helperText: options[i].helperText, placeholder: options[i].placeholder, label: 'Provide Value For Conditional', value: { type: 'local', value: field.validations?.[i]?.value }, type: 'combobox', required: true, configFieldTarget: `validations[${i}].value`, configTarget: field.id, options: { type: 'local', value: options[i].options } }} />
              {:else if options[i].valueType === 'checkbox-group'}
                <ConfigField
                  field={{ id: randomString(), helperText: options[i].helperText, placeholder: options[i].placeholder, label: 'Provide Value For Conditional', value: { type: 'local', value: field.validations?.[i]?.value }, type: 'combobox', required: true, configFieldTarget: `validations[${i}].value`, configTarget: field.id, options: { type: 'local', value: options[i].options } }} />
              {:else if options[i].valueType === 'radio-group'}
                <ConfigField
                  field={{ id: randomString(), helperText: options[i].helperText, placeholder: options[i].placeholder, label: 'Provide Value For Conditional', value: { type: 'local', value: field.validations?.[i]?.value }, type: 'combobox', required: true, configFieldTarget: `validations[${i}].value`, configTarget: field.id, options: { type: 'local', value: options[i].options } }} />
              {:else if options[i].valueType === 'date'}
                <ConfigField
                  field={{ id: randomString(), helperText: options[i].helperText, showDate: true, showTime: true, placeholder: options[i].placeholder, label: 'Afer Date', value: { type: 'local', value: field.validations?.[i]?.value }, type: 'date', required: true, configFieldTarget: `validations[${i}].value`, configTarget: field.id }} />
              {:else if options[i].valueType === 'date-between'}
                <ConfigField
                  field={{ id: randomString(), helperText: options[i].helperText, showDate: true, showTime: true, placeholder: options[i].placeholder, label: 'Minimum Date', value: { type: 'local', value: field.validations?.[i]?.value.min }, type: 'date', required: true, configFieldTarget: `validations[${i}].value.min`, configTarget: field.id }} />
                <ConfigField
                  field={{ id: randomString(), helperText: options[i].helperText, showDate: true, showTime: true, placeholder: options[i].placeholder, label: 'Maximum Date', value: { type: 'local', value: field.validations?.[i]?.value.max }, type: 'date', required: true, configFieldTarget: `validations[${i}].value.max`, configTarget: field.id }} />
              {:else if options[i].valueType === 'number'}
                <ConfigField
                  field={{ id: randomString(), helperText: options[i].helperText, placeholder: options[i].placeholder, label: 'Number of Days', value: { type: 'local', value: field.validations?.[i]?.value }, type: 'number', required: true, configFieldTarget: `validations[${i}].value`, configTarget: field.id }} />
              {:else if options[i].valueType === 'js-expression'}
                <ConfigField
                  field={{ id: randomString(), helperText: options[i].helperText, placeholder: options[i].placeholder, label: 'Custom Javascript Expression', value: { type: 'local', value: field.validations?.[i]?.value }, type: 'string', rows: 3, required: true, configFieldTarget: `validations[${i}].value`, configTarget: field.id }} />
                <div class="ml-3">
                  <JavascriptExpressionViewer fieldId={field.id} />
                </div>
              {/if}
            {/if}
          </div>
        </div>
      </div>
    </div>
  {/each}
  {#if helperText}
    <div class="helper-text">
      {@html helperText ?? ''}
    </div>
  {/if}
  <div class="ml-1 mt-4 mb-6">
    <Button type="primary" onClick={addNew}>Add New Rule</Button>
  </div>
</div>
