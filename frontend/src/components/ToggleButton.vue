<script setup lang="ts">
import { useGlobalStore } from "../stores/global";

defineProps<{ action: Function }>();
const store = useGlobalStore();
</script>

<template>
  <button
    class="container"
    @click="action()"
  >
    <div
      :class="{ circle: true, active: store.darkMode, inactive: !store.darkMode }"
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
