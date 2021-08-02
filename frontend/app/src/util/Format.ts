import { nullOrEmpty } from "./Compare";

export function firstNotEmpty(...values: (string | any)[]) {
  for (let v of values) {
    if (!nullOrEmpty(v)) {
      return v;
    }
  }
  return "";
}
