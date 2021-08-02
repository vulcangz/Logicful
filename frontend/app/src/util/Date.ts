import { isDateObject, isString } from "@app/guards/Guard";
import type { DateObject } from "@app/models/Date";
import { DateTime } from "luxon";

export function isMinDate(date: DateObject | string): boolean {
  if (isDateObject(date)) {
    return date.year === 1969 && date.month === 12;
  }
  if (isString(date)) {
    const parsed = DateTime.fromISO(date);
    return isMinDate(parsed.toObject() as DateObject);
  }
  return false;
}
