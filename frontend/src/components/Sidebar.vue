<script setup lang="ts">
import type { AppDocument } from "@/types";
import Document from "./Document.vue";

const documents = JSON.parse(localStorage.getItem("documents")!) as AppDocument[];
</script>

<template>
  <div
    class="bg-gray-100 w-full h-20 flex items-center pl-4 border-b-1 border-gray-400"
  >
    <h3 class="text-gray-700">My documents</h3>
  </div>
  <div class="w-full overflow-auto">
    <div v-for="document of documents" >
      <Document :data="document" />
    </div>
  </div>
</template>
