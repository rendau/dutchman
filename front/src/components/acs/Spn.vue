<template>
  <q-spinner-puff :color="color" :size="size"/>
</template>

<script setup>
const props = defineProps({
  color: { type: String, default: 'primary' },
  size: { type: String, default: '2rem' },
})
</script>
