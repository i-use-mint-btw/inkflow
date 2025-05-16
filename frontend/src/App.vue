<script setup lang="ts">
import Sidebar from "./components/Sidebar.vue";
import EditorWindow from "./components/EditorWindow.vue";
import PreviewWindow from "./components/PreviewWindow.vue";
import { ref, watch } from "vue";
import { useGlobalStore } from "./stores/global";
import { API_URL } from "./constants";
import { debounce } from "./utils";

const input = ref("# hello");

const store = useGlobalStore();

const markdownChangeEvent = new CustomEvent("markdownChange")
const sock = new WebSocket(API_URL + "/api/document/edit/6fef16d4-7725-4c2f-9fec-44586b691bcd")

sock.addEventListener("message", event => {
  input.value = event.data
  console.log(input.value, event.data)
})

const syncMarkdown = debounce(() => {
  sock.send(input.value)
}, 1000)

sock.addEventListener("markdownChange", syncMarkdown)

watch(input, () => sock.dispatchEvent(markdownChangeEvent))

const documentTitle = ref("")

async function createDocument() {
  const res = await fetch(
    API_URL + "/api/document/create", 
    {
      method: "POST", 
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({title: documentTitle.value})
    }
  )
  console.log(await res.json())
}
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
    <input type="text" v-model="documentTitle" placeholder="Document title" >
    <button @click="createDocument">Create new document</button>
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
