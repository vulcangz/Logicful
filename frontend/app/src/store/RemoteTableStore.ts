import type { TableRow } from "@app/components/models/RemoteTableProps";
import { writable } from "svelte/store";

export const tableRows = writable<TableRow[]>([]);
