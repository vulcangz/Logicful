import { DateTime } from "luxon";

export type DateObject = {
  day: number;
  hour: number;
  millisecond: number;
  minute: number;
  month: number;
  second: number;
  year: number;
  formatted: string;
};

export function toDateTime(value: DateObject): DateTime {
  return DateTime.local(
    value.year,
    value.month,
    value.day,
    value.hour,
    value.minute,
    value.second
  );
}
