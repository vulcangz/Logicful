<script lang="typescript">
  import type { IField } from "@app/models/IField";

  export let field: IField;
  import TextInput from "@app/inputs/TextInput.svelte";
  import { afterUpdate, onMount } from "svelte";
  import ComboBox from "@app/inputs/ComboBox.svelte";
  import { LoadState } from "@app/models/LoadState";
  import { FieldValueLoader } from "@app/loader/FieldValueLoader";
  import Address from "@app/inputs/Address.svelte";
  import FullName from "@app/inputs/FullName.svelte";

  import TextArea from "@app/inputs/TextArea.svelte";
  import CheckboxGroup from "@app/inputs/CheckboxGroup.svelte";
  import Spacer from "@app/inputs/Spacer.svelte";
  import RadioGroup from "@app/inputs/RadioGroup.svelte";

  import formStore from "@app/store/FormStore";
  import RichTextDisplay from "@app/inputs/RichTextDisplay.svelte";
  import Switch from "../../../inputs/Switch.svelte";
  import DatePicker from "@app/inputs/DatePicker.svelte";
  import { firstNotEmpty } from "@app/util/Format";
  import FileUpload from "@app/inputs/FileUpload.svelte";
  import SectionHeader from "@app/inputs/SectionHeader.svelte";
  import { fastEquals } from "@app/util/Compare";

  let state = LoadState.NotStarted;
  let value: any;
  let lastValue: any;
  export let config: any = {};
  export let hidden: boolean = false;
  export let padding: boolean = true;
  export let mode: "builder" | "live" = "builder";

  onMount(load);

  afterUpdate(async () => {
    if (field.value && !fastEquals(field.value, lastValue)) {
      await load();
    }
  });

  async function load() {
    lastValue = field.value;
    if ((field.value ?? field.defaultValue) != null) {
      state = LoadState.Loading;
      try {
        const loader = new FieldValueLoader();
        const result = await loader.load(field);
        value = result;
        field.value = result;
        formStore.set(field, {
          value: result,
          field: "value",
          fromUser: false,
        });
        state = LoadState.Finished;
      } catch (e) {
        console.error(e);
        state = LoadState.Failed;
      }
    }
  }
</script>

<div>
  {#if hidden}
    {#if mode === 'builder'}
      <p class="text-gray-700">
        {firstNotEmpty(field.label, field.name)} is hidden by rules defined in logic.
        This message is only displayed on this preview.
      </p>
    {:else}
      <div class="hidden" />
    {/if}
  {:else if field.type === 'address'}
    <Address {field} {value} />
  {:else if field.type === 'string'}
    <TextInput {field} />
  {:else if field.type === 'email'}
    <TextInput {field} type="email" />
  {:else if field.type === 'number'}
    <TextInput {field} type={'number'} />
  {:else if field.type === 'combobox'}
    <ComboBox {field} {...config} />
  {:else if field.type === 'block'}
    <RichTextDisplay {field} isPreview={true} />
  {:else if field.type === 'block-editor'}
    <TextArea {field} {...config} isPreview={true} />
  {:else if field.type === 'spacer'}
    <Spacer {field} />
  {:else if field.type === 'switch'}
    <Switch {field} {...config} />
  {:else if field.type === 'date'}
    <DatePicker {field} {...config} />
  {:else if mode === 'builder' && field.type === 'placeholder'}
    <div class="placeholder">
      <i class="fas fa-grip-horizontal" />
      <p>You have no fields, drag one from the left sidebar to get started.</p>
    </div>
  {:else if field.type === 'file'}
    <FileUpload {field} />
  {:else if field.type === 'checkbox-group'}
    <CheckboxGroup {field} />
  {:else if field.type === 'radio-group'}
    <RadioGroup {field} />
  {:else if field.type === 'full-name'}
    <FullName {field} {value} />
  {:else if field.type === 'section-header'}
    <SectionHeader {field} />
  {:else if mode === 'builder'}
    <p>No field found for field. {JSON.stringify(field, null, 2)}</p>
  {:else}
    <div class="hidden" />
  {/if}
</div>
