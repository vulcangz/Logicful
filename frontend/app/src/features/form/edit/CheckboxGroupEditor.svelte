<script lang="typescript">
  import type { IField, LabelValue } from "@app/models/IField";
  import { randomString } from "@app/util/Generate";
  import Repeater from "@app/components/Repeater.svelte";
  import formStore from "@app/store/FormStore";
  import { isEmptyOrNull } from "@app/util/Compare";
  import ConfigField from "./ConfigField.svelte";

  export let field: IField;

  function onOptionsChange(options: string[] | LabelValue[]) {
    if (options.length === 0) {
      options = ["Checkbox Item 1"];
    }
    field.options = options;
    formStore.set(field, {
      fromUser: true,
      field: "options",
      value: options,
    });
  }

  function options(): LabelValue[] {
    if (isEmptyOrNull(field.options)) {
      return [{ label: "Checkbox Item 1", value: "Checkbox Item 1" }];
    }
    return field.options?.map((w: string) => {
      return { label: w, value: w };
    });
  }
</script>

<div class="ml-3">
  <Repeater
    options={options()}
    onlyLabel={true}
    label={'Checkbox Options'}
    onChange={onOptionsChange} />
</div>
<ConfigField
  field={{ id: randomString(), type: 'switch', label: "Include 'Other' Option", value: { type: 'local', value: field.includeOther || false }, configFieldTarget: 'includeOther', configTarget: field.id }} />
