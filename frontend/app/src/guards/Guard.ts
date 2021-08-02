import type { DateObject } from "@app/models/Date";
import type { LabelValue } from "@app/models/IField";
import { nullOrEmpty } from "@app/util/Compare";

export function isString(value: any): value is string {
  return typeof value === "string";
}

export function isValueSet(value: any): boolean {
  if (value == null) {
    return false;
  }
  if (isString(value)) {
    return !nullOrEmpty(value);
  }
  if (Array.isArray(value)) {
    return value.length > 0;
  }
  if (isObject(value)) {
    return Object.keys(value).length > 0;
  }
  return true;
}

export function isNumber(value: any): value is number {
  if (value == null) {
    return false;
  }
  return !isNaN(parseInt(value));
}

export function isDateObject(value: any): value is DateObject {
  return (
    isObject(value) &&
    value.hour != null &&
    value.day != null &&
    value.month != null &&
    value.year != null
  );
}

export function isObject(value: any): value is any {
  return value != null && !isString(value) && typeof value === "object";
}

export function isFunction(value: any): value is () => any {
  return typeof value === "function" && !isString(value) && !nullOrEmpty(value);
}

export function isLabelValue(value: any): value is LabelValue {
  return !isString(value) && value?.label != null && value?.value != null;
}

export function isEmail(value: any): boolean {
  return value && value.toString().includes("@");
}

export function isUrl(value: any): boolean {
  if (!value) {
    return false;
  }
  return (
    value.toString().includes("http://") ||
    value.toString().includes("https://")
  );
}
