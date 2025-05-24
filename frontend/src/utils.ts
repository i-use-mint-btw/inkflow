export function debounce(callback: () => void, delayInMs: number) {
  let timer: any;
  return function () {
    clearTimeout(timer);
    timer = setTimeout(callback, delayInMs);
  };
}
