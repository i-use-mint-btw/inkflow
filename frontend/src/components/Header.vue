<script setup lang="ts">
import { API_URL } from "@/constants";
import { useGlobalStore } from "@/stores/global";
import { ref } from "vue";
import { AiFillFileAdd } from "vue-icons-plus/ai";
import { FaShare } from "vue-icons-plus/fa";
import CreateDocumentModal from "./CreateDocumentModal.vue";
import { GiHamburgerMenu } from "vue-icons-plus/gi";

const showModal = ref(false);
const store = useGlobalStore();
const shouldRerender = ref(false);

function toggleModal() {
  showModal.value = !showModal.value;
}

async function createDocument(title: string) {
  showModal.value = false;
  const res = await fetch(API_URL + "/document/create", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ Title: title }),
  });
  const document = await res.json();

  let localStorageDocuments = localStorage.getItem("documents");

  if (!localStorageDocuments) {
    localStorage.setItem("documents", JSON.stringify([document.data]));
    return;
  }

  const documents = JSON.parse(localStorageDocuments!) as any[];
  documents.push(document.data);
  localStorage.setItem("documents", JSON.stringify(documents));
  shouldRerender.value = !shouldRerender.value;
}

async function fetchExistingDocument(id: string) {
  showModal.value = false;
  const res = await fetch(API_URL + "/document/read/" + id);
  const document = await res.json();

  let localStorageDocuments = localStorage.getItem("documents");

  if (!localStorageDocuments) {
    localStorage.setItem("documents", JSON.stringify([document.data]));
    return;
  }

  const documents = JSON.parse(localStorageDocuments) as any[];
  documents.push(document.data);
  localStorage.setItem("documents", JSON.stringify(documents));
  shouldRerender.value = !shouldRerender.value;
}

function copyDocumentLink() {
  const link = location.href + store.selectedDocument.id;
  console.log(link);

  if (navigator.clipboard && window.isSecureContext) {
    navigator.clipboard
      .writeText(link)
      .then(() => {
        console.log("Link copied to clipboard!");
      })
      .catch((err) => {
        console.error("Could not copy link: ", err);
      });
  } else {
    const tempInput = document.createElement("input");
    tempInput.value = link;
    document.body.appendChild(tempInput);
    tempInput.select();
    document.execCommand("copy");
    document.body.removeChild(tempInput);
    console.log("Link copied to clipboard (fallback method)");
  }
  alert("Link copied to clipboard");
}
</script>

<template>
  <CreateDocumentModal
    v-if="showModal"
    :toggleModal="toggleModal"
    :createDocument="createDocument"
    :fetchExistingDocument="fetchExistingDocument"
  />
  <header
    class="flex w-screen h-1/12 border-b-1 px-4 md:px-16 justify-between border-gray-400 items-center"
  >
    <h1 class="font-medium" style="font-weight: 400; margin: 0">Inkflow</h1>
    <h2 class="hidden md:block">
      {{
        Object.keys(store.selectedDocument).length === 0
          ? "Choose a document"
          : store.selectedDocument.title
      }}
    </h2>
    <div class="flex items-center gap-4">
      <button
        @click="toggleModal"
        class="bg-blue-400 text-lg active:opacity-85 p-1.5 rounded-full"
      >
        <AiFillFileAdd class="text-white" />
      </button>
      <button
        class="bg-blue-400 flex items-center gap-2 text-lg active:opacity-85 p-1.5 rounded-full"
        @click="copyDocumentLink"
      >
        <FaShare class="text-white" />
      </button>
      <!-- <button class="block md:hidden" @click="store.toggleSideMenu">
        <GiHamburgerMenu />
      </button> -->
      <!-- <button
        class="bg-blue-300 dark:bg-transparent p-2 rounded-full ml-4 text-blue-600 dark:text-yellow-200"
        @click="store.toggleDarkMode"
      >
        <BsMoonStarsFill v-if="store.darkMode === false" />
        <PiSunDuotone v-else />
      </button> -->
    </div>
  </header>
</template>
