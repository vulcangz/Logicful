export function debounce(
  func: any,
  wait: number,
  immediate: number | null = null
) {
  let timeout: number | null;
  return function () {
    //@ts-ignore
    var context = this,
      args = arguments;
    var later = function () {
      timeout = null;
      //@ts-ignore
      if (!immediate) func.apply(context, args);
    };
    var callNow = immediate && !timeout;
    //@ts-ignore
    clearTimeout(timeout);
    //@ts-ignore
    timeout = setTimeout(later, wait);
    //@ts-ignore
    if (callNow) func.apply(context, args);
  };
}
