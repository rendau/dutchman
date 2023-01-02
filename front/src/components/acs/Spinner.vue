<template>
  <q-inner-loading :showing="showing">
    <ac-spn color="primary" size="3em"/>
  </q-inner-loading>
</template>

<script setup>
const props = defineProps({
  showing: { type: Boolean },
})
</script>
