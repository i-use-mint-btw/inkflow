<script setup lang="ts">
import { getTextColor } from '../utils';


defineProps({
  modelValue: String
})

const emit = defineEmits(['update:modelValue'])

</script>


<template>
  <section class="w-1/2 h-full bg-transparent border-r-2 border-gray-400 overflow-hidden">
    <textarea class="focus:outline-none w-full h-full p-4 resize-none" name="Markdown editor" :style="{color: getTextColor()}" :value="modelValue" @input="emit('update:modelValue', $event.target?.value)"></textarea>
  </section>
</template>