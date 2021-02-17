<template>
  <NavBar />
  <!-- <Search /> -->
  <Table />
</template>

<script lang="ts">
import { ref, defineComponent } from "vue";
import Table from "../components/Table.vue";
import NavBar from "../components/NavBar.vue";

export default defineComponent({
  name: "Home",
  components: {
    Table,
    NavBar,
  },
});
</script>