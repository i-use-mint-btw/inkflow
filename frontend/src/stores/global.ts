import type { AppDocument } from "@/types";
import { defineStore } from "pinia";

export const useGlobalStore = defineStore("global", {
  state: () => ({
    darkMode: false,
    selectedDocument: {} as AppDocument
  }),
  actions: {
    toggleDarkMode() {
      const root = document.getElementById("app")
      
      if (!this.darkMode) {
        root?.classList.add("dark")
        this.darkMode = true
        return
      }

      root?.classList.remove("dark")
      this.darkMode = false
    },
    setSelectedDocument(document: AppDocument) {
      this.selectedDocument = document
    }
  },
});
