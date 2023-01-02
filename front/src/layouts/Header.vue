<template>
  <q-header bordered class="bg-grey-2 text-primary">
    <q-toolbar class="q-px-sm">
      <div v-if="$store.getters['app/screenSmall']" class="q-pr-xs">
        <q-btn flat dense round icon="menu" color="grey-8" @click="$store.commit('app/toggleDrawer')"/>
      </div>

      <q-avatar>
        <q-icon name="rocket_launch" size="40px"/>
<!--        <q-img no-spinner no-transition src="~assets/icon.png"
               width="40px"/>-->
      </q-avatar>

      <q-toolbar-title class="text-weight-medium">
        Temp
      </q-toolbar-title>

      <q-space/>

      <div v-if="$store.getters['profile/isAuthed']">
        <q-btn flat dense round color="primary" @click="onLogout">
          <q-icon name="logout" size="1.2rem"/>
        </q-btn>
      </div>
    </q-toolbar>
  </q-header>
</template>

<script setup>
import { useQuasar } from 'quasar'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import { cns } from 'boot/cns'

const store = useStore()
const router = useRouter()
const $q = useQuasar()

const onLogout = () => {
  $q.dialog({
    message: 'Вы действительно хотите выйти?',
    ok: { label: 'Да', noCaps: true },
    cancel: { label: 'Отмена', flat: true, noCaps: true },
  }).onOk(() => {
    store.dispatch('profile/logout').then(() => {
      router.push({ name: cns.authRouteName })
    })
  })
}
</script>
