<script setup lang="ts">
defineProps<{ createDocument: (title: string) => {}, toggleModal: () => void }>();

//createDocument(documentTitle.value)
</script>

<template>
  <div
    @click="toggleModal"
    class="absolute top-0 bottom-0 right-0 left-0 flex items-center justify-center"
  >
    <div class="w-1/6 bg-gray-100 p-4 rounded-md" @click="e => e.stopPropagation()">
      <label for="title input" class="block mb-4 focus">Enter a title: </label>
      <!-- @vue-ignore -->
      <input
        name="title input"
        :autofocus="true"
        @keypress="(e) => e.key === 'Enter' && createDocument(e.target.value)"
        type="text"
        placeholder="Quarterly report"
        class="border-black border-1 w-full outline-none p-2"
      />
    </div>
  </div>
</template>
