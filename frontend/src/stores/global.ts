import { defineStore } from "pinia";

export const useGlobalStore = defineStore("global", {
  state: () => ({
    darkMode: false,
    colorPallete: {
      backgroundLight: "#FAFAFA", // Tailwind gray-50
      backgroundDark: "#1F2937", // Tailwind gray-800

      textLight: "#1A1A1A", // Near black
      textDark: "#F9FAFB", // Tailwind gray-100

      primaryLight: "#3B82F6", // Tailwind blue-500
      primaryDark: "#60A5FA", // Tailwind blue-400

      accentLight: "#10B981", // Tailwind green-500
      accentDark: "#34D399", // Tailwind green-400

      borderLight: "#E5E7EB", // Tailwind gray-200
      borderDark: "#374151", // Tailwind gray-700

      errorLight: "#EF4444", // Tailwind red-500
      errorDark: "#F87171", // Tailwind red-400
    },
  }),
  actions: {
    toggleDarkMode() {
      this.darkMode = !this.darkMode;
    },
  },
});
