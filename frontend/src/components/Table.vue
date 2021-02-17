<template>
  <div class="flex justify-around m-10">
    <table class="table-fixed text-green-600 text-lg w-5/6">
      <thead>
        <tr class="text-left">
          <th class="w-1/6">书名</th>
          <th class="w-3/12">作者</th>
          <th class="w-3/12">出版时间</th>
          <th class="w-3/12">出版社</th>
          <th class="w-3/12">ISBN</th>
          <th class="w-1/12">详情</th>
        </tr>
      </thead>
      <tbody v-for="from in froms" key:="from.Payload.Id" >
        <tr class="text-left shadow-inner">
          <td class="p-3">{{ from.Payload.Name }}</td>
          <td>{{ from.Payload.Author }}</td>
          <td>{{ from.Payload.PublishYear }}</td>
          <td>{{ from.Payload.Publisher }}</td>
          <td>{{ from.Payload.ISBN }}</td>
          <td>
            <a :href="from.Url" target="_blank">详情</a>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      froms: [],
    };
  },
  async mounted() {
    const res  = await axios.get("http://localhost:9000/?q="+this.$route.query.q)
    this.froms = res.data.data.Items
  },
}
</script>