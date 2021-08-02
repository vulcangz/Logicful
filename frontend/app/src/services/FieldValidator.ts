import { dispatch } from "@app/event/EventBus";
import { runBuiltInValidations } from "@app/inputs/validations/FieldValidations";
import type { IField } from "@app/models/IField";
import formStore from "@app/store/FormStore";
import { debounce } from "@app/util/Debounce";
import { LogicBuilder } from "./LogicBuilder";
import { isValueSet } from "@app/guards/Guard";

export const debounceValidateField: any = debounce(validateField, 400);

export function validateField(field: IField, value: any): boolean {
  let prev = field.error;
  const validations = field.validations ?? [];
  field.error = undefined;
  const evaluator = new LogicBuilder();

  const isEmpty = !isValueSet(value);

  for (let v of validations) {
    if (!field.required && isEmpty) {
      field.error = undefined;
      continue;
    }

    const result = evaluator.evaluateCondition(v, value);

    const isValidIfTrue =
      v.errorIfTrue?.["Field is valid if the condition passes."];

    const isErrorIfTrue =
      v.errorIfTrue?.["Field is invalid if the condition passes."];

    if (isValidIfTrue && result) {
      field.error = undefined;
      continue;
    }

    if (isValidIfTrue && !result) {
      field.error = v.message;
      break;
    }

    if (isErrorIfTrue && result) {
      field.error = v.message;
      break;
    }

    if (isErrorIfTrue && !result) {
      field.error = undefined;
      continue;
    }
  }

  if (field.required && isEmpty) {
    field.error = "This field is required.";
  }

  if (!field.error && !isEmpty) {
    const error = runBuiltInValidations(field, value);
    if (error) {
      field.error = error;
    }
  }

  if (field.error !== prev) {
    dispatch("on_field_validate", field);
    field.value = value;
    formStore.set(field);
  }

  return field.error != null;
}
