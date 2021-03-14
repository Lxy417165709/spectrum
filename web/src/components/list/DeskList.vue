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


    <!--    2. 添加 桌位 部分-->
<!--    <el-col style="height: 300px; width: 202px; margin-left: 10px; border: none">-->
<!--      <desk-spacial-card @click.native="openDeskEditorOfAdmin"></desk-spacial-card>-->
<!--    </el-col>-->
  </el-row>
</template>

<script>
/* eslint-disable */
import DeskCard from "../card/DeskCard";
import DeskSpacialCard from "../card/DeskSpacialCard";
import test from "../../common/test/test";

export default {
  components: {DeskSpacialCard, DeskCard},
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
    // openDeskEditorOfAdmin() {
    //   this.$emit("openDeskEditorOfAdmin", test.blankDesk)
    // },
  }
}
</script>
