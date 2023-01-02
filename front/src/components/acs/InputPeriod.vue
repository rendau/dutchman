<template>
  <div>
    <div class="row no-wrap items-center">
      <div>
        <ac-input-date :withTime="withTime"
                       :model-value="tsGte" @update:model-value="$emit('update:ts-gte', $event)"/>
      </div>

      <div class="text-color5 q-px-sm">
        --
      </div>

      <div>
        <ac-input-date :withTime="withTime" end-of-day
                       :model-value="tsLte" @update:model-value="$emit('update:ts-lte', $event)"/>
      </div>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  tsGte: { type: String },
  tsLte: { type: String },
  withTime: { type: Boolean },
})
</script>
