import { isString, isObject } from "../guards/Guard";
import type { FormValue, IField } from "../models/IField";
import { select } from "@app/util/Selection";

export class FieldValueLoader {
  public async load(field: IField): Promise<any | undefined> {
    return await this.loadValue(field.value ?? field.defaultValue);
  }

  private async loadValue(value: FormValue): Promise<any | undefined> {
    if (value == null) {
      return;
    }
    if (isString(value)) {
      return value;
    }
    if (value.type === "remote") {
      return await this.loadRemote(value);
    }
    if (value.type === "local") {
      const localValue = value.value;
      if (isObject(localValue) && localValue.type === "remote") {
        return await this.loadChildren(localValue);
      }
      return localValue;
    }
    return value;
  }

  private async loadRemote(value: FormValue) {
    if (isString(value)) {
      return value;
    }
    const response = await fetch(value.value);
    const result = await response.json();
    return value.selector ? select(result, value.selector) : result;
  }

  private async loadChildren(value: FormValue | any) {
    const keys = Object.keys(value);
    const promises = await keys.map((w) => {
      return this.loadValue(value[w]);
    });
    const results = await Promise.all(promises);
    const response: any = {};
    for (let i = 0; i < keys.length; i++) {
      response[keys[i]] = results[i];
    }
    return response;
  }
}
