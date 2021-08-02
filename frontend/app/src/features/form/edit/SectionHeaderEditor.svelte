<script lang="typescript">
  import type { IField } from "@app/models/IField";
  import { randomString } from "@app/util/Generate";
  import ConfigField from "./ConfigField.svelte";
  export let field: IField;
</script>

<div>
  <ConfigField
    field={{ id: randomString(), label: 'Header', type: 'string', value: field.header, configTarget: field.id, configFieldTarget: 'header' }} />
  <ConfigField
    field={{ id: randomString(), label: 'Help Text', type: 'string', value: field.helperText, configTarget: field.id, configFieldTarget: 'helperText' }} />
</div>
