import { getToken } from "./AuthService";

async function authHeaders() {
  const token = await getToken();
  if (!token) {
    return;
  }
  return {
    Authorization: "Bearer " + token,
  };
}

export function apiEndpoint() {
  if (window.location.host === "localhost:5000") {
    return "http://localhost:3000/api/";
  } else {
    return "/api/";
  }
}

export async function getApi<T>(path: string): Promise<T> {
  const endpoint = apiEndpoint();
  const response = await fetch(`${endpoint}${path}`, {
    headers: await authHeaders(),
  });
  if (!response.ok) {
    const body = await response.json();
    throw new Error(body.message);
  }
  const body = await response.json();
  return body as T;
}

export async function postApi<T>(path: string, body: any): Promise<T> {
  return await requestApiWithBody<T>("POST", path, body);
}

export async function putApi<T>(path: string, body: any): Promise<T> {
  return await requestApiWithBody<T>("PUT", path, body);
}

export async function deleteApi<T>(path: string, body: any): Promise<T> {
  return await requestApiWithBody<T>("DELETE", path, body);
}

async function requestApiWithBody<T>(
  method: string,
  path: string,
  body: any
): Promise<T> {
  const endpoint = apiEndpoint();
  try {
    const response = await fetch(`${endpoint}${path}`, {
      method: method,
      headers: {
        "Content-Type": "application/json",
        ...(await authHeaders()),
      },
      body: JSON.stringify(body),
    });
    if (!response.ok) {
      const body = await response.json();
      throw new Error(body.message);
    }

    const result = await response.text();
    if (result === null || result === "") {
      return {} as T;
    }
    return JSON.parse(result) as T;
  } catch (ex) {
    throw ex;
  }
}
