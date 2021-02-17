<template>
  <Search />
  <Table />
</template>

<script lang="ts">
import { ref, defineComponent } from "vue";
import Table from "../components/Table.vue";
import Search from "../components/Search.vue";

export default defineComponent({
  name: "Home",
  components: {
    Table,
    Search,
  },
});
</script>