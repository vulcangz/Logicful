import type { IField } from "@app/models/IField";

export function setFieldDefaults(field: IField): IField {
  if (field.type === "checkbox-group") {
    field.value = { "Option 1": "Option 1" };
    field.options = ["Option 1", "Option 2"];
  }
  if (field.type === "radio-group") {
    field.value = { "Option 1": "Option 1" };
    field.options = ["Option 1", "Option 2"];
  }
  if (field.type === "section-header") {
    field.helperText =
      "Use this text to display helper text that will assist a user in filling out the following fields.";
    field.header = "Example Section Header";
  }
  if (field.type === "date") {
    field.showDate = true;
    field.showTime = true;
  }
  return field;
}
