<template>
    <input type="text"
            v-model="keyWord"
           placeholder="搜索..."
           class="h-14 w-2/3 text-left border border-gray-400 rounded-full py-3 px-6 shadow-inner focus:outline-none focus:ring focus:border-blue-300 text-lg">
    <button @click="search"
            class="bg-gray-500 shadow-inner rounded-full p-2 px-10 focus:outline-none focus:ring text-white w-52 h-14 m-10 text-xl">搜索</button>
</template>

<script>
export default {
    data() {
        return {
            keyWord:"",
        }
    },
  methods: {
    search () {
      this.$router.push({ path: "search", query: { q: this.keyWord } })
    }
  }
}
</script>