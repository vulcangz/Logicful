import type { LogicRule } from "../models/LogicBuilder";
import type { IField } from "../models/IField";
import formStore from "../store/FormStore";
import { nullOrEmpty } from "@app/util/Compare";
import {
  isDateObject,
  isLabelValue,
  isObject,
  isString,
} from "@app/guards/Guard";
import type { Dictionary } from "@app/models/Utility";
import { DateTime } from "luxon";
import { dispatch } from "@app/event/EventBus";
import { compileCode } from "./JsEvaluator";
import type { DateObject } from "@app/models/Date";
import { isMinDate } from "@app/util/Date";
import { toDateTime } from "@app/models/Date";

export class LogicBuilder {
  evaluate(field: IField): boolean {
    if (!field.logic) {
      return true;
    }
    if (!field.logic.rules) {
      return true;
    }

    const rules = field.logic.rules.filter((r) => {
      if (nullOrEmpty(r.field) || nullOrEmpty(r.condition)) {
        return false;
      }
      return true;
    });

    if (rules.length === 0) {
      return true;
    }

    if (field.logic.action === "show-any-match") {
      for (let rule of rules) {
        const fieldTargetValue = formStore.getValue(rule.field);
        if (fieldTargetValue == null) {
          continue;
        }
        if (this.evaluateCondition(rule, fieldTargetValue)) {
          return true;
        }
      }
      return false;
    }

    if (field.logic.action === "show-all-match") {
      for (let rule of rules) {
        const fieldTargetValue = formStore.getValue(rule.field);
        if (fieldTargetValue == null) {
          return false;
        }
        if (!this.evaluateCondition(rule, fieldTargetValue)) {
          return false;
        }
      }
      return true;
    }

    if (field.logic.action === "hide-all-match") {
      for (let rule of rules) {
        const fieldTargetValue = formStore.getValue(rule.field);
        if (fieldTargetValue == null) {
          return true;
        }
        if (!this.evaluateCondition(rule, fieldTargetValue)) {
          return true;
        }
      }
      return false;
    }

    if (field.logic.action === "hide-any-match") {
      for (let rule of rules) {
        const fieldTargetValue = formStore.getValue(rule.field);
        if (fieldTargetValue == null) {
          continue;
        }
        if (this.evaluateCondition(rule, fieldTargetValue)) {
          return false;
        }
      }
      return true;
    }

    return false;
  }

  public evaluateCondition(rule: LogicRule, value: any): boolean {
    console.log(
      "eval",
      JSON.stringify(rule.value, null, 4),
      JSON.stringify(value, null, 4)
    );
    if (isLabelValue(value)) {
      return (
        this.evaluateCondition(rule, value.value) ||
        this.evaluateCondition(rule, value.label)
      );
    }
    switch (rule.condition) {
      case "contains":
        return this.toLowerCase(value).includes(this.toLowerCase(rule.value));
      case "startsWith":
        return this.toLowerCase(value).startsWith(this.toLowerCase(rule.value));
      case "endsWith":
        return this.toLowerCase(value).endsWith(this.toLowerCase(rule.value));
      case "eq":
        return this.toLowerCase(value) == this.toLowerCase(rule.value);
      case "gt":
        return parseFloat(value) > parseFloat(rule.value);
      case "lt":
        return parseFloat(value) < parseFloat(rule.value);
      case "lte":
        return parseFloat(value) <= parseFloat(rule.value);
      case "gte":
        return parseFloat(value) >= parseFloat(rule.value);
      case "hasValue":
        return this.hasValue(value);
      case "notHaveValue":
        return !this.hasValue(value);
      case "isTrue":
        return value != null && value == true;
      case "isFalse":
        return value != null && value == false;
      case "isFileExtension":
        return this.isFileExtension(value, rule);
      case "isNotFileExtension":
        return !this.isFileExtension(value, rule);
      case "hasValueChecked":
        return this.hasValueChecked(value, rule);
      case "notHasValueChecked":
        return !this.hasValueChecked(value, rule);
      case "noValuesChecked":
        return this.hasNoValuesChecked(value, rule);
      case "dateIsBetween":
        return this.dateIsBetween(value, rule);
      case "dateIsNotBetween":
        return !this.dateIsBetween(value, rule);
      case "dateIsBefore":
        return this.dateIsBefore(value, rule);
      case "dateIsAfter":
        return this.dateIsAfter(value, rule);
      case "isWithinXDays":
        return this.dateIsWithinXDays(value, rule);
      case "js":
        return this.javascriptExpression(value, rule);
      case "isEmail":
        return this.isEmail(value);
      default:
        return false;
    }
  }

  private hasValue(value: any): boolean {
    return value != null && value != "";
  }

  private isEmail(value: any): boolean {
    return isString(value) && value.includes("@");
  }

  private javascriptExpression(value: any, rule: LogicRule): boolean {
    try {
      dispatch("logic_rule_javascript_error", undefined);
      const result = compileCode(rule.value, value);
      return Boolean(result);
    } catch (ex) {
      dispatch("logic_rule_javascript_error", ex);
      return false;
    }
  }

  private dateIsBetween(value: any, rule: LogicRule): boolean {
    const minObject: DateObject = isDateObject(rule.value.min)
      ? rule.value.min
      : undefined;
    const maxObject: DateObject = isDateObject(rule.value.max)
      ? rule.value.max
      : undefined;

    // Time comparison only
    if (
      minObject &&
      isMinDate(minObject) &&
      (minObject.minute || minObject.hour) &&
      maxObject &&
      isMinDate(maxObject) &&
      (maxObject.minute || maxObject.hour)
    ) {
      const date = value as DateObject;
      const current = date.hour * 3600 + date.minute * 60;
      const max = maxObject.hour * 3600 + maxObject.minute * 60;
      const min = minObject.hour * 3600 + minObject.minute * 60;
      return current >= min && current <= max;
    }

    // Time comparison only
    if (
      minObject &&
      isMinDate(minObject) &&
      (minObject.minute || minObject.hour)
    ) {
      const date = value as DateObject;
      const current = date.hour * 3600 + date.minute * 60;
      const min = minObject.hour * 3600 + minObject.minute * 60;
      return current >= min;
    }

    // Time comparison only
    if (
      maxObject &&
      isMinDate(maxObject) &&
      (maxObject.minute || maxObject.hour)
    ) {
      const date = value as DateObject;
      const current = date.hour * 3600 + date.minute * 60;
      const max = maxObject.hour * 3600 + maxObject.minute * 60;
      return current <= max;
    }

    const min = rule.value.min
      ? toDateTime(rule.value.min).toMillis()
      : undefined;

    const max = rule.value.max
      ? toDateTime(rule.value.max).toMillis()
      : undefined;

    const date = isDateObject(value) ? toDateTime(value).toMillis() : undefined;

    if (!date) {
      return false;
    }
    if (min && max) {
      return date >= min && date <= max;
    }
    if (min) {
      return date >= min;
    }
    if (max) {
      return date <= max;
    }
    return false;
  }

  private dateIsAfter(value: DateObject, rule: LogicRule): boolean {
    const diff = toDateTime(value).diff(toDateTime(rule.value)).milliseconds;
    return diff > 0;
  }

  private dateIsWithinXDays(value: DateObject, rule: LogicRule): boolean {
    try {
      const days = parseInt(rule.value, 10);
      const valueAsDate = toDateTime(value);
      const diff = parseInt(valueAsDate.diffNow("days").days.toFixed(0), 10);
      if (valueAsDate.startOf("day").equals(DateTime.local().startOf("day"))) {
        return true;
      }
      return diff <= days && diff >= 0;
    } catch (ex) {
      return false;
    }
  }

  private dateIsBefore(value: any, rule: LogicRule): boolean {
    const before = rule.value ? toDateTime(rule.value).toMillis() : undefined;
    const date = value ? toDateTime(value).toMillis() : undefined;
    if (!date || !before) {
      return false;
    }
    return value < before;
  }

  private hasValueChecked(value: Dictionary<string>, rule: LogicRule): boolean {
    return Object.keys(value).includes(rule.value);
  }

  private hasNoValuesChecked(
    value: Dictionary<string>,
    rule: LogicRule
  ): boolean {
    return Object.keys(value).length === 0;
  }

  private isFileExtension(value: any, rule: LogicRule): boolean {
    if (!this.hasValue(value)) {
      return false;
    }
    const file = formStore.getFile(isObject(value) ? value.id : value);
    if (!file) {
      return false;
    }
    const fileName = file.name;
    const split = fileName.split(".");
    if (split.length < 2) {
      return false;
    }
    const rules = rule.value.split(",").map((r: string) => {
      return r.replace(" ", "").replace(".", "");
    });
    for (let r of rules) {
      if (r === split[split.length - 1]) {
        return true;
      }
    }
    return false;
  }

  private toLowerCase(value: any) {
    if (!this.hasValue(value)) {
      return "";
    }
    return value.toString().toLowerCase();
  }
}
