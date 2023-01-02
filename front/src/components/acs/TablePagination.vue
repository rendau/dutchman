<template>
  <q-pagination :model-value="page"
                :max="pageMax"
                :max-pages="5"
                :direction-links="true"
                size=".9rem"
                color="primary"
                @update:model-value="onInput"/>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  stdParams: { type: Object, default: () => ({}) },
  data: { type: Object, default: () => ({}) },
})
const emits = defineEmits()

const page = computed(() => ((props.stdParams.page || 0) + 1))
const pageSize = computed(() => (props.stdParams.page_size || 0))
const totalCount = computed(() => (props.data.total_count || 0))
const pageMax = computed(() => Math.ceil((totalCount.value || 1) / (pageSize.value || 1)))

const onInput = v => {
  if (v !== page.value) {
    emits('page', v - 1)
  }
}
</script>
