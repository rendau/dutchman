<template>
  <q-layout view="HHH Lpr lFf" class="app-bg-color relative-position">
    <Header/>

    <Drawer/>

    <q-page-container>
      <q-page class="q-pt-md q-pl-lg q-pr-md q-pb-lg" style="padding-left: 20px">
        <router-view/>
      </q-page>
    </q-page-container>
  </q-layout>
</template>

<script setup>
import Header from './Header'
import Drawer from './Drawer'</script>
