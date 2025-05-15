<script setup lang="ts">
import { marked } from 'marked';
import { computed } from 'vue';
import { getTextColor } from '../utils';


const props = defineProps<{content: string}>()

const output = computed(() => marked(props.content))

</script>

<template>
  <div v-html="output" :style="{color: getTextColor()}"></div>
</template>

<style>

</style>