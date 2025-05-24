<script setup lang="ts">
import { useGlobalStore } from '@/stores/global';
import type { AppDocument } from '@/types';

defineProps<{ data: AppDocument }>();

const store = useGlobalStore()
</script>

<template>
  <button @click="store.setSelectedDocument(data)" class="h-16 w-full flex items-center justify-center bg-gray-100 border-y-1 border-gray-300">
    <p class="mr-4 text-lg">{{ data.title }}</p>
  </button>
</template>
