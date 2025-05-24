<script setup lang="ts">
import { VMarkdownView } from "vue3-markdown";
import "vue3-markdown/dist/vue3-markdown.css";

defineProps<{ content: string }>();
</script>

<template>
  <div
    class="w-full h-20 flex items-center pl-4 border-b-1 border-gray-400 bg-gray-100"
  >
    <h3 class="text-gray-700">Markdown Preview</h3>
  </div>
  <div class="overflow-y-auto h-full">
    <VMarkdownView class="p-4 h-full" :content="content"></VMarkdownView>
  </div>
</template>