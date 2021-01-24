<!-- eslint-disable -->
<template>
  <el-row>
    <el-col v-for="(desk,deskIndex) in desks" :key="deskIndex"
            style="height: 300px; width: 202px; margin-left: 10px; border: none">

      <!-- todo: 不知为何 在 span 中才能触发响应事件 -->
      <span @click="handleDeskCardClick(deskIndex)">
        <component :is="desk.component" :num="desk.space.num"
                   :pictureStorePath="desk.space.pictureStorePath"></component>
      </span>
    </el-col>
  </el-row>
</template>

<script>
/* eslint-disable */
export default {
  props: {
    desks: Array
  },
  data() {
    return {};
  },
  methods: {
    handleDeskCardClick(deskIndex) {
       console.log("handleDeskCardClick")
      this.$emit("passViewMode", 1, deskIndex)
    },
  }
}
</script>
