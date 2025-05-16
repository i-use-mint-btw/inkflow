import { useGlobalStore } from "./stores/global"

export function getTextColor() {
  const store = useGlobalStore()
  return store.darkMode ? store.colorPallete.textDark : store.colorPallete.textLight
}

export function debounce(callback: () => void, delayInMs: number) {
  let timer
  return function() {
    clearTimeout(timer)
    setTimeout(callback, delayInMs)
  }
}