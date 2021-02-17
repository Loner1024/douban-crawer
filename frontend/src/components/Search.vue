<template>
    <div class="flex justify-center m-10">
        <input type="text" placeholder="搜索..." class="h-10 w-2/3 text-left border border-gray-400 rounded-full py-3 px-6 shadow-inner focus:outline-none focus:ring focus:border-blue-300">
        <button v-on:click="search" class="ml-10 bg-gray-500 shadow-inner rounded-full p-2 px-10 focus:outline-none focus:ring text-white">搜索</button>
    </div>
</template>

<script>
export default {
    search(){
        console.log("click")
    }
}
</script>