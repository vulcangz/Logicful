import type { User } from "@app/models/User";
import { postApi } from "./ApiService";

let memoryToken: string = "";
let refreshing = false;

export const emptyUser : User = {
  creationDate: "",
  lastActive: "",
  fullName: "",
  displayName: "",
  email: "",
  id: "",
  teamId: ""
};
export interface UserToken {
  token: string;
  expiration: number;
}

export function setToken(token: UserToken, remember: boolean = true) {
  if (!localStorage) {
    return;
  }
  localStorage.removeItem("token");
  if (remember) {
    localStorage.setItem("token", JSON.stringify(token));
  } else {
    memoryToken = JSON.stringify(token);
  }
}

export async function getToken(): Promise<string | undefined> {
  if (!localStorage) {
    return undefined;
  }
  const token = localStorage.getItem("token") ?? memoryToken;
  if (!token) {
    return undefined;
  }
  try {
    let parsed: UserToken = JSON.parse(token);
    const left = parsed.expiration - Date.now();

    if (left <= 0) {
      await refreshToken();
      return await getToken();
    }
    // 24 hours
    else if (left <= 8.64e7) {
      refreshToken();
    }

    return parsed.token;
  } catch (ex) {
    console.error(ex);
    localStorage.removeItem("token");
    memoryToken = "";
    return undefined;
  }
}

export async function me(): Promise<User> {
  const token = await getToken();
  if (!token) {
    return emptyUser;
  }
  try {
    const user: User = parseJwt(token);
    return user;
  } catch {
    return emptyUser;
  }
}

function parseJwt(token: string): any {
  const base64Url = token.split(".")[1];
  const base64 = base64Url.replace(/-/g, "+").replace(/_/g, "/");
  const jsonPayload = decodeURIComponent(
      atob(base64)
          .split("")
          .map(function (c) {
            return "%" + ("00" + c.charCodeAt(0).toString(16)).slice(-2);
          })
          .join("")
  );
  return JSON.parse(jsonPayload);
}

async function refreshToken() {
  if (refreshing) {
    return;
  }
  refreshing = true;
  try {
    const token = await postApi<UserToken>("user/refresh", {});
    setToken(token, true);
  } finally {
    refreshing = false;
  }
}
