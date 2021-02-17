<template>
<NavBar/>
    <router-view />
</template>

<script lang="ts">
import { ref,defineComponent, onMounted } from 'vue'
import NavBar from './components/NavBar.vue'
import axios from 'axios'

export default defineComponent({
  name: 'App',
    components: {
      NavBar
  },
})
</script>

<style>
</style>