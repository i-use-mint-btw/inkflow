<script setup lang="ts">
import { useGlobalStore } from "../stores/global";

const props = defineProps<{ action: Function }>();
const store = useGlobalStore();
</script>

<template>
  <button
    class="container"
    @click="action()"
    :style="{
      backgroundColor: store.darkMode
        ? store.colorPallete.accentDark
        : store.colorPallete.accentLight,
    }"
  >
    <div
      :class="{ circle: true, active: store.darkMode, inactive: !store.darkMode }"
      :style="{
        backgroundColor: store.darkMode
          ? store.colorPallete.backgroundLight
          : store.colorPallete.backgroundDark,
      }"
    ></div>
  </button>
</template>

<style>
.container {
  width: 50px;
  height: 22px;
  border: none;
  border-radius: 10px;
}

.circle {
  width: 42%;
  height: 100%;
  border-radius: 13px;
  transition: all 400ms ease-in-out;
}

.active {
  transform: translateX(30px);
}

.inactive {
  transform: translateX(0);
}
</style>
