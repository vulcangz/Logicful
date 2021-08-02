import { isDateObject } from "@app/guards/Guard";
import type { DateObject } from "@app/models/Date";
import type { IField } from "@app/models/IField";
import { LogicBuilder } from "@app/services/LogicBuilder";
import { isMinDate } from "@app/util/Date";
import { DateTime } from "luxon";
import {toDateTime} from "@app/models/Date";

export function runBuiltInValidations(
  field: IField,
  value: any
): string | undefined {
  if (field.type === "date") {
    return validateDateField(field, value);
  }
  return undefined;
}

function validateDateField(field: IField, value: any): string | undefined {
  const evaluator = new LogicBuilder();
  const result = evaluator.evaluateCondition(
    {
      condition: "dateIsBetween",
      value: {
        min: field.rangeFrom,
        max: field.rangeTo,
      },
    },
    value
  );
  if (!result) {
    const from: DateObject | undefined = isDateObject(field.rangeFrom)
      ? field.rangeFrom
      : undefined;
    const to: DateObject | undefined = isDateObject(field.rangeTo)
      ? field.rangeTo
      : undefined;

    function formatTime(d: DateObject) {
      return toDateTime(d).toLocaleString(DateTime.TIME_SIMPLE);
    }

    function formatDate(d: DateObject) {
      return toDateTime(d).toLocaleString(DateTime.DATETIME_MED);
    }

    if (
      from &&
      (from.minute || from.hour) &&
      to &&
      isMinDate(to) &&
      (to.minute || to.hour)
    ) {
      return `Time must be between ${formatTime(from)} - ${formatTime(to)}.`;
    }

    if (from && isMinDate(from) && (from.minute || from.hour)) {
      return `Time must be after ${formatTime(from)}.`;
    }

    if (to && isMinDate(to) && (to.minute || to.hour)) {
      return `Time must be before ${formatTime(to)}.`;
    }

    if (from && to) {
      return `Date must be between ${formatDate(from)} - ${formatDate(to)}.`;
    }

    if (from) {
      return `Date must be after ${formatDate(from)}.`;
    }
    if (to) {
      return `Date must be before ${formatDate(to)}.`;
    }
  }
  return undefined;
}
