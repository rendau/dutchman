<template>
  <div class="flex items-center">
    <ac-table-pagination v-if="showPagination"
                         :std-params="stdParams"
                         :data="data"
                         @page="onPage"/>

    <q-space/>

    <div class="q-pr-xs q-mr-sm">
      <ac-table-total-counter :value="totalCount"/>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  stdParams: { type: Object, default: () => ({}) },
  data: { type: Object, default: () => ({}) },
  noPagination: { type: Boolean, default: false },
})
const emits = defineEmits()

const pageSize = computed(() => (props.stdParams.page_size || 0))
const totalCount = computed(() => (props.data.total_count || 0))
const showPagination = computed(() => {
  return pageSize.value && !props.noPagination && (totalCount.value > pageSize.value)
})

const onPage = page => emits('page', page)
</script>
