<!-- eslint-disable -->
<template>
  <el-row>
    <!--    1. 展示 桌位 部分-->
    <el-col v-for="(desk,deskIndex) in desks" :key="deskIndex"
            style="height: 300px; width: 202px; margin-left: 10px; border: none">
      <desk-card :desk="desk"
                 @click.native="handleDeskCardClick(deskIndex)"></desk-card>
    </el-col>


    <!--    2. 添加 桌位 部分-->
    <el-col style="height: 300px; width: 202px; margin-left: 10px; border: none">
      <desk-spacial-card @click.native="openDeskEditorOfAdmin"></desk-spacial-card>
    </el-col>
  </el-row>
</template>

<script>
/* eslint-disable */
import DeskCard from "../card/DeskCard";
import DeskSpacialCard from "../card/DeskSpacialCard";
import test from "../../common/test/test";
import utils from "../../common/utils"

export default {
  components: {DeskSpacialCard, DeskCard},
  props: {
    className: String,
  },
  data() {
    return {
      desks: [],
    };
  },
  methods: {
    async handleDeskCardClick(deskIndex) {
      // todo: this.desks[deskIndex].id === undefined 时，此时 this.desks[deskIndex].id !== 0 也成立..
      if (this.desks[deskIndex].id !== 0) {
        this.$emit("turnToClassListMode", deskIndex)
        return
      }
      // 进行点单
      await utils.OrderDesk(this, {
        desk: this.desks[deskIndex],
      }, (res) => {
        this.desks[deskIndex].id = res.data.data.deskID;
        this.$emit("attachOrderID", res.data.data.orderID)
      })
    },
    openDeskEditorOfAdmin() {
      this.$emit("openDeskEditorOfAdmin", test.blankDesk, this.className)
    },
  }
}
</script>
