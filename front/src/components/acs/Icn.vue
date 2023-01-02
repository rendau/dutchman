<template>
  <div class="inline-block" style="line-height: 1px">
    <!-- logo -->
    <svg v-if="name === 'logo'"></svg>
  </div>
</template>

<script setup>
const props = defineProps({
  name: { type: String, required: true },
})
</script>
