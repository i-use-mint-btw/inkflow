<script setup lang="ts">
defineProps<{ fetchExistingDocument: (id:string) => {}, createDocument: (title: string) => {}, toggleModal: () => void }>();

function extractID(url: string) {
  if (url.length === 0) return
  const tokens = url.split("/")
  return tokens[tokens.length - 1]
}

//createDocument(documentTitle.value)
</script>

<template>
  <div
    @click="toggleModal"
    class="absolute top-0 bottom-0 right-0 left-0 flex items-center justify-center"
  >
    <div class=" w-5/6 md:w-1/6 bg-gray-100 p-4 rounded-md" @click="e => e.stopPropagation()">
      <label for="document-title" class="block mb-4 focus">Enter a title: </label>
      <!-- @vue-ignore -->
      <input
        name="document-title"
        :autofocus="true"
        @keypress="(e) => e.key === 'Enter' && createDocument(e.target.value)"
        type="text"
        placeholder="Quarterly report"
        class="border-black border-1 w-full outline-none p-2 mb-6"
      />
      <label for="shared-document">Or add a shared document:</label>
      <!-- @vue-ignore -->
      <input
        name="shared-document"
        :autofocus="true"
        @keypress="(e) => e.key === 'Enter' && fetchExistingDocument(extractID(e.target.value))"
        type="text"
        placeholder="Quarterly report"
        class="border-black border-1 w-full outline-none p-2"
      />
      
    </div>
  </div>
</template>
