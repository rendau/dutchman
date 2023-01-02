<template>
  <div class="row items-center q-pr-md relative-position q-pb-lg">
    <div v-if="bb" class="q-pr-lg">
      <q-btn dense outline round color="grey-4" @click="$router.back()">
        <q-icon name="chevron_left" color="grey-7"/>
      </q-btn>
    </div>

    <slot/>
  </div>
</template>

<script setup>
const props = defineProps({
  bb: { type: Boolean },
})
</script>
