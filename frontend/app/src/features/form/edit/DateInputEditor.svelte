<script lang="typescript">
  import SectionHeader from "@app/inputs/SectionHeader.svelte";

  import type { IField } from "@app/models/IField";

  import { randomString } from "@app/util/Generate";

  import ConfigField from "./ConfigField.svelte";

  export let field: IField;
</script>

<ConfigField
  field={{ id: randomString(), type: 'switch', label: 'Include Date', value: field.showDate, configFieldTarget: 'showDate', configTarget: field.id }} />

<ConfigField
  field={{ id: randomString(), type: 'switch', label: 'Include Time', value: field.showTime, configFieldTarget: 'showTime', configTarget: field.id }} />

<div class="pt-4">
  <div class="pl-3">
    <SectionHeader
      field={{ id: randomString(), header: 'Date/Time Range', helperText: 'Restrict selection to date range', type: 'section-header' }} />
    <p class="mt-2 text-sm text-gray-500">
      To restrict only time, leave date as blank
    </p>
  </div>
  <ConfigField
    field={{ id: randomString(), type: 'date', showTime: field.showTime, showDate: field.showDate, label: 'From', value: field.rangeFrom, configFieldTarget: 'rangeFrom', configTarget: field.id }} />
  <div class="-mt-2">
    <ConfigField
      field={{ id: randomString(), type: 'date', showTime: field.showTime, showDate: field.showDate, label: 'To', value: field.rangeTo, configFieldTarget: 'rangeTo', configTarget: field.id }} />
  </div>
</div>
