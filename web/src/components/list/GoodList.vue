<!-- eslint-disable -->
<template>
  <el-row>
    <el-col v-for="(good,goodIndex) in goods" :key="goodIndex"
            style="height: 300px; width: 202px; margin-left: 10px; border: none">
      <good-card :good="good" @dblclick.native="openGoodEditor(good)"></good-card>
    </el-col>
    <el-col v-if="isAdminView" style="height: 300px; width: 202px; margin-left: 10px; border: none">
      <good-spacial-card @click.native="tryToAddGood()"></good-spacial-card>
    </el-col>
  </el-row>

</template>

<script>
/* eslint-disable */
import GoodCard from "../card/GoodCard";
import GoodSpacialCard from "../card/GoodSpacialCard";
import test from "../../common/test/test"

export default {
  components: {GoodSpacialCard, GoodCard},
  props: {
    goods: Array,
    isAdminView: Boolean,
    className: String
  },
  data() {
    return {};
  },
  methods: {
    handleButtonClick() {
      this.$emit("turnToGoodClassListMode")
    },
    openGoodEditor(good) {
      if (this.isAdminView) {
        this.openGoodEditorOfAdmin(good)
      } else {
        this.openGoodEditorOfUser(good)
      }
    },
    tryToAddGood() {
      this.openGoodEditorOfAdmin(test.blankGood)
    },
    openGoodEditorOfAdmin(good) {
      // 商品信息编辑(管理人员视图)
      this.$emit("openGoodEditorOfAdmin", good, this.className)
    },
    openGoodEditorOfUser(good) {
      // 商品信息编辑(普通用户视图)
      this.$emit("openGoodEditorOfUser", good, this.className)
    }
  }
}

// todo: 菜单显示不太正常！！！
</script>
