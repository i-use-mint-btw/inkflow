<script setup lang="ts">
import { onMounted, ref, watch } from "vue";
import { VMarkdownEditor } from "vue3-markdown";
import "vue3-markdown/dist/vue3-markdown.css";
import Header from "./components/Header.vue";
import PreviewWindow from "./components/PreviewWindow.vue";
import { API_URL } from "./constants";
import { debounce } from "./utils";

const input = ref("# Hello");

onMounted(() => {
  const markdownChangeEvent = new CustomEvent("markdownChange");

  const sock = new WebSocket(
    API_URL + "/api/document/edit/6fef16d4-7725-4c2f-9fec-44586b691bcd"
  );

  sock.addEventListener("message", (event) => {
    input.value = event.data;
  });

  const syncMarkdown = debounce(() => {
    sock.send(input.value);
  }, 800);

  sock.addEventListener("markdownChange", syncMarkdown);
  watch(input, () => sock.dispatchEvent(markdownChangeEvent));
});

//const documentTitle = ref("");
//createDocument(documentTitle.value)
</script>

<template>
  <div class="h-screen w-screen">
    <Header />
    <main class="flex w-screen h-11/12">
      <section
        class="w-1/2 h-full bg-transparent border-r-1 border-gray-400 overflow-hidden"
      >
        <VMarkdownEditor v-model="input" locale="en" />
      </section>
      <PreviewWindow :content="input" />
    </main>
  </div>
</template>
