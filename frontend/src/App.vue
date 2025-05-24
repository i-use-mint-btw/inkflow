<script setup lang="ts">
import { ref, watch } from "vue";
import { VMarkdownEditor } from "vue3-markdown";
import "vue3-markdown/dist/vue3-markdown.css";
import Header from "./components/Header.vue";
import PreviewWindow from "./components/PreviewWindow.vue";
import Sidebar from "./components/Sidebar.vue";
import { API_URL } from "./constants";
import { useGlobalStore } from "./stores/global";
import { debounce } from "./utils";

const input = ref("");
const store = useGlobalStore();
let sock: WebSocket | null;

watch(
  () => store.selectedDocument.id,
  () => {
    if (sock) {
      sock.close();
      return;
    }
    sock = new WebSocket(
      API_URL + `/document/edit/${store.selectedDocument.id}`
    );

    sock.addEventListener("message", (event) => {
      input.value = event.data;
    });

    const queueBroadcast = debounce(() => {
      sock!.send(input.value);
    }, 1500);

    watch(input, queueBroadcast);
  }
);
</script>

<template>
  <div class="h-screen w-screen">
    <Header />
    <main class="flex w-screen h-11/12">
      <section
        class="w-full md:w-1/2 h-full bg-transparent border-r-1 border-gray-400 overflow-hidden"
      >
        <VMarkdownEditor v-model="input" locale="en" mode="light" />
      </section>
      <section
        class="hidden md:block w-1/2 h-full overflow-auto border-r-1 border-gray-400"
      >
        <PreviewWindow :content="input" />
      </section>
      <aside class=" md:w-3/16 h-full">
        <Sidebar />
      </aside>
    </main>
  </div>
</template>
