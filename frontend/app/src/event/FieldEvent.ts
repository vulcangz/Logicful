import type { IField } from "../models/IField";
import { dispatchSync, subscribeComponent } from "@app/event/EventBus";

export function dispatchFieldChange(
  field: IField,
  change: { field: string; value: any; fromUser: boolean }
) {
  dispatchSync("field_changed", {
    field,
    change,
  });
}

export function subscribeFieldChange(
  onMount: any,
  callback: (
    field: IField,
    change: { field: string; value: any; fromUser: boolean }
  ) => any
) {
  subscribeComponent("field_changed", (payload) => {
    if (!payload.field) {
      console.error("Field change was undefined.", payload);
      return;
    }
    callback(payload.field, payload.change);
  });
}
