import type { IField } from "@app/models/IField";
import type { LogicConditional } from "@app/models/LogicBuilder";

export function getFieldLogicConditions({
  field,
  loadOptions,
}: {
  field: IField;
  loadOptions: any;
}): LogicConditional[] {
  const customExpression = {
    label: "Custom Javascript Expression",
    value: "js",
    valueInput: "js-expression",
    placeholder: "value === 'my value'",
  };
  if (field.type === "string") {
    return [
      customExpression,
      {
        label: "Is Email Address",
        value: "isEmail",
      },
      {
        label: "Contains",
        value: "contains",
      },
      {
        label: "Starts With",
        value: "startsWith",
      },
      {
        label: "Ends With",
        value: "endsWith",
      },
      {
        label: "Equals",
        value: "eq",
      },
      {
        label: "Has Value",
        value: "hasValue",
      },
    ];
  }
  if (field.type === "combobox") {
    return [
      customExpression,
      {
        label: "Equals",
        value: "eq",
        valueInput: "combobox",
        options: loadOptions,
      },
      {
        label: "Not Equals",
        value: "notEq",
        valueInput: "combobox",
        options: loadOptions,
      },
      {
        label: "Has Selected Option",
        value: "hasValue",
      },
    ];
  }
  if (field.type === "date") {
    return [
      customExpression,
      {
        label: "Is Within N Days Of Current Date",
        value: "isWithinXDays",
        valueInput: "number",
      },
      {
        label: "Is After Date",
        value: "dateIsAfter",
        valueInput: "date",
      },
      {
        label: "Is Before Date",
        value: "dateIsBefore",
        valueInput: "date",
      },
      {
        label: "Is Between Dates",
        value: "dateIsBetween",
        valueInput: "date-between",
      },
      {
        label: "Is Not Between Dates",
        value: "dateIsNotBetween",
        valueInput: "date-between",
      },
    ];
  }
  if (field.type === "switch") {
    return [
      {
        label: "Is Toggled",
        value: "isTrue",
      },
      {
        label: "Is Not Toggled",
        value: "isFalse",
      },
    ];
  }
  if (field.type === "checkbox-group") {
    return [
      customExpression,
      {
        label: "Has Value Checked",
        value: "hasValueChecked",
        valueInput: "checkbox-group",
        options: loadOptions,
      },
      {
        label: "Does Not Have Value Checked",
        value: "notHasValueChecked",
        valueInput: "checkbox-group",
        options: loadOptions,
      },
      {
        label: "Has No Values Checked",
        value: "noValuesChecked",
      },
    ];
  }
  if (field.type === "radio-group") {
    return [
      customExpression,
      {
        label: "Has Value Selected",
        value: "hasValueChecked",
        valueInput: "radio-group",
        options: loadOptions,
      },
      {
        label: "Does Not Have Value Selected",
        value: "notHasValueChecked",
        valueInput: "radio-group",
        options: loadOptions,
      },
      {
        label: "Has No Values Selected",
        value: "noValuesChecked",
      },
    ];
  }
  if (field.type === "file") {
    return [
      customExpression,
      {
        label: "Has Chosen File",
        value: "hasValue",
      },
      {
        label: "Has Not Chosen File",
        value: "notHaveValue",
      },
      {
        label: "Is File Extension",
        value: "isFileExtension",
        helper:
          "You can include multiple extensions by seperating with a comma.",
        placeholder: "pdf, txt, png",
      },
      {
        label: "Is Not File Extension",
        helper:
          "You can include multiple extensions by seperating with a comma.",
        value: "isNotFileExtension",
        placeholder: "pdf, txt, png",
      },
    ];
  }
  if (field.type === "number") {
    return [
      customExpression,
      {
        label: "Greater Than",
        value: "gt",
      },
      {
        label: "Less Than",
        value: "lt",
      },
      {
        label: "Less Than or Equal To",
        value: "lte",
      },
      {
        label: "Greater Than or Equal To",
        value: "gte",
      },
      {
        label: "Equal To",
        value: "eq",
      },
      {
        label: "Has Value",
        value: "hasValue",
      },
    ];
  }

  return [customExpression];
}
