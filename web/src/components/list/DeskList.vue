<!-- eslint-disable -->
<template>
  <el-row>
    <!--    1. 展示 桌位 部分-->
    <el-col v-for="(desk,deskIndex) in desks" :key="deskIndex"
            style="height: 300px; width: 202px; margin-left: 10px; border: none">
      <span @click="handleDeskCardClick(deskIndex)">
        <desk-card :num="desk.space.num"
                   :pictureStorePath="desk.space.pictureStorePath"></desk-card>
      </span>
    </el-col>
  </el-row>
</template>

<script>
/* eslint-disable */
import DeskCard from "../card/DeskCard";

export default {
  components: {DeskCard},
  props: {
    desks: Array
  },
  data() {
    return {};
  },
  methods: {
    handleDeskCardClick(deskIndex) {
      this.$emit("turnToClassListMode", deskIndex)
    },
  }
}
</script>
