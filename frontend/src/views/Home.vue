<template>
  <div class="flex flex-col h-screen">
    <NavBar />
    <div class="flex flex-col">
      <div class="flex flex-col justify-center items-center m-10 h-full">
        <!-- <img src="https://img3.doubanio.com/dae/accounts/resources/d3e2921/shire/assets/nav_doubanapp_6.png"
             class="h-40 w-40 m-10"> -->
             <h1 class="text-9xl m-8">🎒</h1>
        <Search />
      </div>

    </div>
  </div>

</template>

<script lang="ts">
import { ref, defineComponent } from "vue";
import Search from "../components/Search.vue";
import NavBar from "../components/NavBar.vue";

export default defineComponent({
  name: "Home",
  components: {
    Search,
    NavBar,
  },
});
</script>