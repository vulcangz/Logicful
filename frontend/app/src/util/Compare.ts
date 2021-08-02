import { isObject, isString } from "../guards/Guard";

export function stringEquals(str1: any, str2: any) {
  if (str1 == null && str2 == null) {
    return true;
  }
  if (isString(str1) && isString(str2)) {
    return str1.toLowerCase().trim() === str2.toLowerCase().trim();
  }
  return str1 === str2;
}

export function toNumberOrDefault(value: any) {
  const parsed = parseFloat(value);
  return isNaN(parsed) ? 0 : parsed;
}

export function isEmptyOrNull(value: any[] | undefined | null): boolean {
  if (value == null) {
    return true;
  }
  if (!Array.isArray(value)) {
    return true;
  }
  return value.length === 0;
}

export function isNullString(value: any) {
  return (
    value == null ||
    value == "null" ||
    value == "undefined" ||
    value == undefined
  );
}

export function nullOrEmpty(str: any) {
  return str == null || str === "";
}

export function fastEquals(any: any, any2: any) {
  return JSON.stringify(any) === JSON.stringify(any2);
}

export function fastClone(any: any) {
  if (any == null || !isObject(any)) {
    return any;
  }
  return JSON.parse(JSON.stringify(any));
}
