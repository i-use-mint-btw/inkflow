<script setup lang="ts">
import "vue3-markdown/dist/vue3-markdown.css";
import { VMarkdownView } from "vue3-markdown";

defineProps<{ content: string }>();

</script>

<template>
  <section class="w-1/2 overflow-y-auto">
    <div
      class="w-full h-20 flex items-center pl-4 border-b-1 border-gray-400 bg-gray-100"
    >
      <h3 class="text-gray-700">Markdown Preview</h3>
    </div>
    <VMarkdownView  class="p-4" :content="content"></VMarkdownView>
  </section>
</template>
