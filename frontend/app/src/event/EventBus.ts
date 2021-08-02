import { randomString } from "@app/util/Generate";
import { onMount } from "svelte";

const map = new Map<string, { [key: string]: any }>();

export function subscribe(
  event: string,
  subscriber: ((payload: any) => any) | ((payload: any) => Promise<any>)
): string {
  let id = randomString();
  if (!map.has(event)) {
    const subscribers = {
      [id]: subscriber,
    };
    map.set(event, subscribers);
  } else {
    const subscribers = map.get(event);
    subscribers![id] = subscriber;
    map.set(event, subscribers!);
  }
  return id;
}

export function unsubscribe(event: string, id: string) {
  if (!map.has(event) || !map.get(event)) {
    return;
  }
  const result = map.get(event)!;
  delete result[id];
  map.set(event, result);
}

export function subscribeComponent(
  event: string,
  subscriber: ((payload: any) => any) | ((payload: any) => Promise<any>)
) {
  onMount(() => {
    const id = subscribe(event, subscriber);
    return () => {
      unsubscribe(event, id);
    };
  });
}

export function subscribePrivateComponent(
  id: string,
  event: string,
  subscriber: ((payload: any) => any) | ((payload: any) => Promise<any>)
) {
  onMount(() => {
    const unsubscribeId = subscribePrivate(id, event, subscriber);
    return () => {
      unsubscribe(event, unsubscribeId);
    };
  });
}

export function subscribePrivate(
  id: string,
  event: string,
  subscriber: ((payload: any) => any) | ((payload: any) => Promise<any>)
): string {
  const e = `${id}-${event}`;
  return subscribe(e, subscriber);
}

export async function dispatchPrivate(id: string, event: string, payload: any) {
  const e = `${id}-${event}`;
  dispatch(e, payload);
}

export function dispatchSingle<T>(event: string, payload: any): T {
  const result = dispatchSync(event, payload)?.filter((r) => r)[0] as T;
  return result;
}

export async function dispatch(event: string, payload: any) {
  console.debug("dispatch", event, payload);
  if (map.has(event)) {
    const subscribers = map.get(event);
    if (!subscribers) {
      return;
    }
    const promises = Object.values(subscribers).map((s) => {
      return s(payload);
    });
    await Promise.all(promises);
  }
}

export function dispatchSync(event: string, payload: any) {
  console.debug("dispatch_sync", event, payload);
  if (map.has(event)) {
    const subscribers = map.get(event);
    if (!subscribers) {
      return [];
    }
    return Object.values(subscribers).map((subscriber) => {
      return subscriber(payload);
    });
  }
  return [];
}
