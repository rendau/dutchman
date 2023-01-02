<template>
  <q-page class="fullscreen flex flex-center q-pa-md">
    <div>
      <div class="column no-wrap items-center">
        <div>
          <q-icon name="lock_person" color="grey-6" size="100px"/>
        </div>

        <div class="text-h5 text-center q-pt-xl">
          Для входа в приложение<br/>
          необходимо авторизоваться
        </div>

        <div class="q-pt-xl">
          <q-btn no-caps label="Авторизоваться" color="positive" @click="onAuthClick"/>
        </div>
      </div>
    </div>
  </q-page>
</template>

<script setup>
import { onMounted, onUnmounted } from 'vue'
import { useQuasar } from 'quasar'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n/index'
import { cns } from 'boot/cns'

const $q = useQuasar()
const { t } = useI18n()
const store = useStore()
const router = useRouter()

let authWindow = null

const messageH = event => {
  if (event.source === authWindow) {
    try {
      let data = JSON.parse(event.data || '{}')
      store.commit('profile/setToken', {
        access: data.access_token || '',
        refresh: data.refresh_token || '',
      })
      store.dispatch('profile/refresh', true).then(() => {
        if (store.getters['profile/isAuthed']) {
          router.push('/')
        } else {
          $q.notify({ type: 'negative', message: 'Не удалось авторизоваться' })
        }
      }, err => {
        $q.notify({ type: 'negative', message: err.data.desc })
      })
    } catch (e) {
      console.error(e)
      $q.notify({ type: 'negative', message: t(cns.ErrSystem) })
    }

    if (authWindow) {
      authWindow.close()
      authWindow = null
    }
  }
}

const onAuthClick = () => {
  if (authWindow) {
    authWindow.close()
  }
  authWindow = window.open(cns.AccountAuthUrl)
}

onMounted(() => {
  window.addEventListener('message', messageH)
})

onUnmounted(() => {
  window.removeEventListener('message', messageH)
})
</script>
