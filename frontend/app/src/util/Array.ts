export function shiftArray(arr: any[], oldIndex: number, newIndex: number) {
  if (newIndex >= arr.length) {
    let k = newIndex - arr.length + 1;
    while (k--) {
      arr.push(undefined);
    }
  }
  arr.splice(newIndex, 0, arr.splice(oldIndex, 1)[0]);
  return arr;
}

export function orderBy(arr: any[], arr2: any[], property: string) {
  arr.sort((a, b) => {
    return arr2.indexOf(a[property]) - arr2.indexOf(b[property]);
  });
  return arr;
}
