<template>
  <div class="row items-start q-col-gutter-sm cursor-pointer">
    <div style="width: 170px">
      <q-field outlined dense :label="label" :stack-label="!!date" bg-color="white" class="cursor-pointer">
        <template v-slot:control>
          <div class="self-center full-width no-outline">{{ date }}</div>
        </template>

        <template v-slot:append>
          <q-icon name="event" class="cursor-pointer label-color" size="1.4rem"/>
        </template>

        <q-popup-proxy ref="qDateProxyRef">
          <div>
            <q-date :model-value="date" today-btn mask="YYYY-MM-DD" @update:model-value="onDateValueInput"/>
          </div>
        </q-popup-proxy>
      </q-field>
    </div>

    <div style="width: 100px">
      <q-field outlined dense :label="timeLabel" :stack-label="!!time" bg-color="white" class="cursor-pointer">
        <template v-slot:control>
          <div class="self-center full-width no-outline">{{ time }}</div>
        </template>

        <template v-slot:append>
          <q-icon name="schedule" class="cursor-pointer label-color" size="1.4rem"/>
        </template>

        <q-popup-proxy ref="qTimeProxyRef">
          <div>
            <q-time :model-value="time" now-btn format24h mask="HH:mm" @update:model-value="onTimeValueInput"/>
          </div>
        </q-popup-proxy>
      </q-field>
    </div>
  </div>
</template>

<script setup>
import { ref, watch, onMounted } from 'vue'
import moment from 'moment'

const props = defineProps({
  modelValue: {},
  label: { type: String },
  timeLabel: { type: String },
  withTime: { type: Boolean, default: false },
  endOfDay: { type: Boolean, default: false },
})

const emits = defineEmits()

const date = ref('')
const time = ref('')
const qDateProxyRef = ref()
const qTimeProxyRef = ref()

const onValueChange = () => {

  if (!props.modelValue) {
    date.value = null
    time.value = null
  } else {
    let v = moment(props.modelValue)
    date.value = v.format('YYYY-MM-DD')
    time.value = v.format('HH:mm')
  }
}

const onDateValueInput = v => {
  if (qDateProxyRef.value) {
    qDateProxyRef.value.hide()
  }
  date.value = v
  if (!date.value) {
    time.value = null
  }
  emitInput()
}

const onTimeValueInput = v => {
  if (qTimeProxyRef.value) {
    qTimeProxyRef.value.hide()
  }
  time.value = v
  emitInput()
}

const emitInput = () => {
  let v = null

  if (date.value) {
    let t = time.value

    if (!t) {
      t = props.endOfDay ? '23:59' : '00:00'
    }

    v = moment(date.value + ' ' + t, 'YYYY-MM-DD HH:mm').toJSON()
  } else {
    time.value = null
  }

  emits('update:model-value', v)
}

watch(() => props.modelValue, () => {
  onValueChange()
})

onMounted(() => {
  onValueChange()
})
</script>
