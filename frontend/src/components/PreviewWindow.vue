<script setup lang="ts">
import { marked } from 'marked';
import { computed } from 'vue';


const props = defineProps<{content: string}>()
const output = computed(() => marked.parse(props.content, {breaks: true}))

</script>

<template>
  <section v-html="output" class="w-1/2 p-4 overflow-y-auto"></section>
</template>