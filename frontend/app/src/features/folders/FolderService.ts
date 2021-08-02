import type { IFolder } from "@app/models/IFolder";
import type { Dictionary } from "@app/models/Utility";
import { getApi } from "@app/services/ApiService";
import { me } from "@app/services/AuthService";
import { fastClone } from "@app/util/Compare";
import { set } from "@app/util/Selection";

export async function getFolders(
  cache: boolean = true,
  each?: (folder: IFolder) => any
): Promise<Dictionary<IFolder>> {
  const current = await getApi<IFolder[]>("folder");
  const user = await me();
  current.push({
    name: "Uncategorized",
    id: `${user.teamId}:uncategorized`,
    isUncategorized: true,
  });

  let results: { [key: string]: IFolder } = {};
  let parentMap: { [key: string]: string | undefined } = {};
  let folderIdMap: { [key: string]: IFolder } = {};

  current.forEach((f) => {
    each?.(f);
    parentMap[f.id] = f.parent;
    folderIdMap[f.id] = fastClone(f);
    if (!f.parent) {
      results[f.id] = f;
    }
  });

  current.forEach((f) => {
    if (f.parent) {
      let root: string | undefined = f.parent;
      let paths: Set<string> = new Set([f.id, root]);
      while (true) {
        if (!root) {
          break;
        }
        root = parentMap[root];
        if (!root) {
          break;
        }
        paths.add(root);
      }
      const fullPath = Array.from(paths)
        .map((p) => folderIdMap[p])
        .reverse();

      let path = "";
      fullPath.forEach((folder, i) => {
        path += folder.id;
        set(results, path, folder);
        path += `.children.`;
      });
    }
  });
  return results;
}
