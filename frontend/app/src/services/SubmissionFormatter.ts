import { isDateObject, isObject } from "@app/guards/Guard";
import { DateTime } from "luxon";
import { toDateTime } from "@app/models/Date";

export let types: { [key: string]: string } = {};

export function formatSubmissionItem(column: string, value: any) {
  const type = types[column];
  if (type === "address" && isObject(value)) {
    const results = [
      value?.address1?.value,
      value?.address2?.value,
      value?.state?.value,
      value?.city?.value,
      value?.zip?.value,
    ];
    return results.filter((r) => r).join(" ");
  }
  if (type === "checkbox-group" && isObject(value)) {
    return Object.values(value)
      .filter((v) => v != null)
      .join(", ");
  }
  if (type === "radio-group" && isObject(value)) {
    return Object.values(value).find((v) => v != null);
  }
  if (type === "full-name" && isObject(value)) {
    const results = [
      value?.prefix?.value,
      value?.first?.value,
      value?.middle?.value,
      value?.last?.value,
    ];
    return results.filter((r) => r).join(" ");
  }
  if (type === "file" && isObject(value)) {
    if (!value) {
      return "No file submitted";
    }
    return `${value.name ?? value.id}, ${value.type}`;
  }
  if (type === "date" && isDateObject(value)) {
    return toDateTime(value).toLocaleString(DateTime.DATETIME_MED_WITH_SECONDS);
  }
  return undefined;
}
