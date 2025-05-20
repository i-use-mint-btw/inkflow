import { useGlobalStore } from "./stores/global"

export function getTextColor() {
  const store = useGlobalStore()
  return store.darkMode ? store.colorPallete.textDark : store.colorPallete.textLight
}

export function debounce(callback: () => void, delayInMs: number) {
  let timer: any
  return function() {
    clearTimeout(timer)
    timer = setTimeout(callback, delayInMs)
  }
}