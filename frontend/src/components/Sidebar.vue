<script setup lang="ts">
import { useGlobalStore } from "../stores/global";
import { getTextColor } from "../utils";
import ToggleButton from "./ToggleButton.vue";
const store = useGlobalStore()

</script>

<template>
  <div class="sidebar" :style="{color: getTextColor()}">
    <h1 class="logo">Inflow</h1>
    <p>Documents:</p>
    <button class="bg-red-400 active:bg-red-800">Click me</button>
    <div class="toggle-dark-mode-button-container">
      <ToggleButton :action="store.toggleDarkMode" />
    </div>
  </div>
</template>

<style scoped>
  .sidebar {
    display: grid;
    grid-template-rows: 1fr 4fr 1fr;
  }

  .logo {
    margin: 10px 0;
  }

  .toggle-dark-mode-button-container {
    display: flex;
    justify-content: center;
    align-items: center;
  }
</style>
