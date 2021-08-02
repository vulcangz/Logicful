export function nameToInitials(name: string) {
  if (!name) {
    return undefined;
  }
  const split = name.split(" ");
  let result = "";
  split.forEach((s) => {
    result += s[0];
  });
  return result;
}
