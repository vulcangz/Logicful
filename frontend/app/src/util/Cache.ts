const cache: { [key: string]: any } = {};

export function cacheSet(key: string, value: any) {
  cache[key] = value;
}

export function cacheGet(key: string) {
  return cache[key];
}

export function cacheClear(key: string) {
  delete cache[key];
}
