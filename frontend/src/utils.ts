import { useGlobalStore } from "./stores/global"

export function getTextColor() {
  const store = useGlobalStore()
  return store.darkMode ? store.colorPallete.textDark : store.colorPallete.textLight
}