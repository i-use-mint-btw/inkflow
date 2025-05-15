<script setup lang="ts">
import { getTextColor } from '../utils';


defineProps({
  modelValue: String
})

const emit = defineEmits(['update:modelValue'])

</script>


<template>
  <div class="editor-window">
    <textarea class="editor-window-textarea" name="Markdown editor" :style="{color: getTextColor()}" :value="modelValue" @input="emit('update:modelValue', $event.target.value)"></textarea>
  </div>
</template>

<style scoped>
  .editor-window-textarea {
    width: 100%;
    height: 100%;
    background-color: transparent;
    resize: none;
  }

  .editor-window-textarea:focus {
    outline: none;
  }
</style>