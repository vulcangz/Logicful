import { dispatch } from "@app/event/EventBus";
import { dispatchFieldChange } from "@app/event/FieldEvent";
import get from "lodash.get";
import type { IField } from "@app/models/IField";
import type { IForm } from "@app/models/IForm";
import { fastClone, fastEquals } from "@app/util/Compare";
import { set } from "@app/util/Selection";
import { FieldValueLoader } from "@app/loader/FieldValueLoader";

let configStore: { [key: string]: IField } = {};

let files: { [key: string]: File } = {};

let store: {
  [key: string]: any;
  fields: { [key: string]: IField };
} = {
  fields: {},
};

export type FieldChange = {
  field: string;
  value: any;
  fromUser: boolean;
};

export class FormStore {
  initialized(): boolean {
    return this.getForm().initialized === true;
  }
  setForm(form: IForm) {
    const copy: IForm = fastClone(form);
    store = { fields: {} };
    copy.fields.forEach((f) => {
      formStore.set(f, {
        fromUser: false,
        field: "",
        value: "",
      });
    });
    Object.keys(copy).forEach((f) => {
      if (f === "fields") {
        return;
      }
      //@ts-ignore
      store[f] = copy[f];
    });
    store.initialized = true;
    dispatch("form_updated", this.getForm());
  }
  set(
    field: IField,
    change: FieldChange = { field: "", value: "", fromUser: false }
  ) {
    if (field.configTarget === "form") {
      const isSame = fastEquals(configStore[field.id], field);

      if (isSame) {
        return;
      }
      set(store, field.configFieldTarget, field.value);
      dispatch("form_updated", this.getForm());
      dispatchFieldChange(fastClone(field), {
        field: field.configFieldTarget,
        value: field.value,
        fromUser: change.fromUser,
      });
      return;
    }

    if (field.configTarget) {
      const isSame = fastEquals(configStore[field.id], field);

      if (isSame) {
        return;
      }

      set(
        store.fields[field.configTarget],
        field.configFieldTarget,
        field.value
      );
      const copy = fastClone(field);
      configStore[field.id] = copy;
      dispatchFieldChange(copy, change);
      const newField = get(store.fields, field.configTarget);
      dispatchFieldChange(fastClone(newField), {
        field: field.configFieldTarget,
        value: field.value,
        fromUser: change.fromUser,
      });
      return;
    }

    const isSame = fastEquals(field, get(store.fields, field.id));

    if (isSame) {
      return;
    }

    const copy = fastClone(field);
    set(store.fields, field.id, copy);
    dispatchFieldChange(copy, change);
  }
  get(fieldId: string) {
    const field = store.fields[fieldId];
    if (!field) {
      return undefined;
    }
    const copy = fastClone(field);
    return copy;
  }
  getValue(fieldId: string) {
    const copy = this.get(fieldId);
    return copy?.value ?? undefined;
  }
  async valueFromLoader(
    field: IField,
    defaultValue: any,
    value: any = undefined
  ): Promise<any> {
    const clone = fastClone(field);
    value =
      value ??
      clone.value ??
      formStore.getValue(clone.configTarget ?? clone.id) ??
      defaultValue;
    const loader = new FieldValueLoader();
    clone.value = value;
    const result = await loader.load(field);
    return result;
  }
  getForm(): IForm {
    const form: any = { fields: [] };
    Object.keys(store).forEach((k) => {
      if (k === "fields") {
        return;
      }
      form[k] = store[k];
    });
    Object.keys(store.fields).forEach((fieldId) => {
      const field = store.fields[fieldId];
      if (field.configTarget || field.isConfigField) {
        return;
      }
      form.fields.push(fastClone(field));
    });
    return form;
  }
  setFile(id: string, file: File) {
    files[id] = file;
  }
  clearFile(id: string) {
    delete files[id];
  }
  getFile(id: string): File | undefined {
    return files[id];
  }
}

const formStore = new FormStore();

//@ts-ignore
window.getForm = () => {
  return formStore.getForm();
};

export default formStore;
