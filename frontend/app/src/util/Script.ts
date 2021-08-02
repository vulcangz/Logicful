let loaded = new Set<string>();

export async function loadScripts(urls: string[]) {
  urls = urls.filter((u) => !loaded.has(u));
  if (urls.length === 0) {
    return;
  }
  const promises: Promise<any>[] = [];
  urls.forEach((u) => {
    promises.push(
      new Promise((r) => {
        let script = document.createElement("script");
        script.onload = function () {
          loaded.add(u);
          r(undefined);
        };
        script.src = u;
        document.head.appendChild(script);
      })
    );
  });
  await Promise.all(promises);
}
