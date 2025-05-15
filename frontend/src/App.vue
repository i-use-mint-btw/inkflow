<script setup lang="ts">
import { debounce } from "lodash-es";
import Sidebar from "./components/Sidebar.vue";
import EditorWindow from "./components/EditorWindow.vue";
import PreviewWindow from "./components/PreviewWindow.vue";
import { ref, onUpdated } from "vue";
import { useGlobalStore } from "./stores/global";

const input = ref("# hello");

onUpdated()

const store = useGlobalStore();
</script>

<template>
  <div
    :style="{
      backgroundColor: store.darkMode
        ? store.colorPallete.backgroundDark
        : store.colorPallete.backgroundLight,
    }"
    class="main-container"
  >
    <Sidebar />
    <EditorWindow v-model="input" />
    <PreviewWindow :content="input" />
  </div>
</template>

<style scoped>
.main-container {
  display: grid;
  grid-template-columns: 1fr 2fr 2fr;
}

.main-container > * {
  height: 100vh;
}
</style>
