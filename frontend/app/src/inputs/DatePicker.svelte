<script lang="typescript">
  import { subscribeFieldChange } from "@app/event/FieldEvent";
  import { DateTime } from "luxon";
  import { isMinDate } from "@app/util/Date";

  import type { IField } from "@app/models/IField";
  import formStore from "@app/store/FormStore";
  import { debounce } from "@app/util/Debounce";
  import { onMount } from "svelte";
  import Label from "../inputs/Label.svelte";
  import InputErrorText from "./validations/InputErrorText.svelte";
  import InputErrorIcon from "./validations/InputErrorIcon.svelte";
  import FieldHelperText from "@app/inputs/FieldHelperText.svelte";
  import { validateField } from "@app/services/FieldValidator";
  import { isDateObject, isNumber } from "@app/guards/Guard";
  import type { DateObject } from "@app/models/Date";
  import { fastClone } from "@app/util/Compare";
  import { toDateTime } from "@app/models/Date";
  export let field: IField;

  let value: DateObject | undefined;
  let isoDate: string;

  let debouncedOnChange: any;

  subscribeFieldChange(onMount, async (newField) => {
    if (newField.id === field.id) {
      if (newField.value) {
        value = await formStore.valueFromLoader(
          newField,
          undefined,
          newField.value
        );
      } else {
        value = undefined;
      }
      setIsoDate();
    }
  });

  function setIsoDate() {
    if (isDateObject(value)) {
      if (isMinDate(value)) {
        isoDate = "";
      } else {
        isoDate = toDateTime(value).toISODate();
      }
    } else {
      isoDate = "";
    }
  }

  onMount(async () => {
    debouncedOnChange = debounce((date: DateTime) => {
      const obj: any = date.toObject();
      const result: DateObject = {
        ...obj,
        formatted: date.toLocaleString(DateTime.DATETIME_MED_WITH_SECONDS),
      };
      field.value = result;
      formStore.set(field, {
        fromUser: true,
        field: "value",
        value: field.value,
      });
      validateField(field, field.value);
      field.onChange?.(date);
    }, 400);

    value = await formStore.valueFromLoader(field, undefined);

    setIsoDate();
  });

  function formatTime(hours: number, minutes: number) {
    function pad(value: number) {
      return `${value < 10 ? "0" : ""}${value}`;
    }
    return pad(hours) + ":" + pad(minutes);
  }

  function currentDateObject() {
    if (isDateObject(value)) {
      return value;
    }
    return fastClone({
      day: 0,
      minute: 0,
      second: 0,
      millisecond: 0,
      year: 0,
      month: 0,
    });
  }

  function onTimeChange(hour: number, minute: number) {
    try {
      const current = currentDateObject();
      current.hour = hour;
      current.minute = minute;
      current.month = current.month === 0 ? 12 : current.month;
      current.year = current.year === 0 ? 1969 : current.year;
      current.day = current.day === 0 ? 31 : current.day;
      const newDate = toDateTime(current);
      console.log("TIME CHANGE", current);
      debouncedOnChange(newDate);
    } catch (ex) {
      console.error(ex);
    }
  }

  function onDateChange(date: string) {
    try {
      const current = currentDateObject();
      const datetime = DateTime.fromISO(date);
      current.day = datetime.day;
      current.month = datetime.month;
      current.year = datetime.year;
      const newDate = toDateTime(current);
      debouncedOnChange(newDate);
    } catch (ex) {
      console.error(ex);
    }
  }
</script>

<Label {field} />
<div class="mt-1 relative rounded-md shadow-sm">
  <div
    class="grid"
    class:grid-cols-6={field.showDate && field.showTime}
    class:gap-2={field.showDate && field.showTime}
  >
    {#if field.showDate}
      <div
        class:col-span-6={field.showDate && field.showTime}
        class:sm:col-span-3={field.showDate && field.showTime}
      >
        <input
          type="date"
          required
          value={isoDate}
          data-test-id={"date-input-" +
            field.label +
            "-" +
            (field.required ? "required" : "optional")}
          on:click|stopPropagation
          class:input-error={field.error}
          on:input={(e) => {
            onDateChange(e.target.value);
          }}
          class="shadow-sm focus:ring-indigo-500 focus:border-indigo-500 block w-full sm:text-sm border-gray-300 rounded-md"
        />
      </div>
    {/if}
    {#if field.showTime}
      <div
        class:col-span-6={field.showDate && field.showTime}
        class:sm:col-span-3={field.showDate && field.showTime}
      >
        <input
          type="time"
          required
          value={value ? formatTime(value.hour, value.minute) : ""}
          on:click|stopPropagation
          data-test-id={"time-input-" +
            field.label +
            "-" +
            (field.required ? "required" : "optional")}
          class:input-error={field.error}
          on:input={(e) => {
            const split = e.target.value.split(":");
            const hour = parseInt(split[0], 10);
            const minute = parseInt(split[1], 10);
            onTimeChange(
              isNumber(hour) ? hour : 0,
              isNumber(minute) ? minute : 0
            );
          }}
          class="shadow-sm focus:ring-indigo-500 focus:border-indigo-500 block w-full sm:text-sm border-gray-300 rounded-md"
        />
      </div>
    {/if}
  </div>
  <InputErrorIcon {field} />
</div>
<InputErrorText {field} />
<FieldHelperText {field} />
