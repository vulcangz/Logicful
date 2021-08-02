import _set from "lodash.set";
import _get from "lodash.get";
import _has from "lodash.has";

export function select(o: any, s: string): any {
  return _get(o, s);
}

export function set(o: any, s: string, value: any): any {
  _set(o, s, value);
}

export function assertExists(o: any, s: string, value: any): boolean {
  if (!_has(o, s)) {
    _set(o, s, value);
    return true;
  }
  return false;
}
