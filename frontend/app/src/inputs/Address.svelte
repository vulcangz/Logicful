<script lang="typescript">
  import { subscribeFieldChange } from "@app/event/FieldEvent";

  import type { IField } from "@app/models/IField";
  import formStore from "@app/store/FormStore";
  import { onMount } from "svelte";
  import Field from "@app/features/form/edit/Field.svelte";
  export let field: IField;
  export let value: { [key: string]: any };

  let address1: IField;
  let address2: IField;
  let city: IField;
  let state: IField;
  let zip: IField;

  subscribeFieldChange(onMount, (newField) => {
    if (newField.id === field.id) {
      value = newField.value;
    }
  });

  onMount(() => {
    value = formStore.getValue(field.id) ?? {};
  });
</script>

<div>
  <div class="grid grid-cols-1 gap-y-6 gap-x-4 sm:grid-cols-6">
    <div class="sm:col-span-6">
      <Field
        field={{
          required: true,
          name: `${field.name}.address1`,
          id: `${field.id}.value.address1`,
          label: "Street Address",
          rows: 3,
          value: value?.address1?.value ?? "",
          type: "string",
        }}
      />
    </div>
    <div class="sm:col-span-2">
      <Field
        field={{
          required: true,
          name: `${field.name}.city`,
          id: `${field.id}.value.city`,
          label: "City",
          value: value?.city?.value ?? "",
          type: "string",
        }}
      />
    </div>

    <div class="sm:col-span-2">
      <Field
        field={{
          name: `${field.name}.state`,
          id: `${field.id}.value.state`,
          label: "State",
          value: value?.state?.value,
          required: true,
          type: "combobox",
          options: {
            type: "remote",
            value:
              "https://gist.githubusercontent.com/mshafrir/2646763/raw/8b0dbb93521f5d6889502305335104218454c2bf/states_hash.json",
          },
        }}
      />
    </div>

    <div class="sm:col-span-2">
      <Field
        field={{
          required: true,
          name: `${field.name}.zip`,
          id: `${field.id}.value.zip`,
          label: "Zip / Postal Code",
          value: value?.zip?.value ?? "",
          type: "string",
          properties: { pattern: "[d]{5}(-[d]{4})?" },
        }}
      />
    </div>
  </div>
</div>
