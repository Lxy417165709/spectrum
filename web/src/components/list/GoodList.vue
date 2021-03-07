<!-- eslint-disable -->
<template>
  <el-row>
    <!--    1. 展示 商品 部分-->
    <el-col v-for="(good,goodIndex) in goods" :key="goodIndex"
            style="height: 300px; width: 202px; margin-left: 10px; border: none">
      <good-card :good="good" @dblclick.native="openGoodEditor(good)"></good-card>
    </el-col>

    <!--    2. 添加 商品 部分-->
    <el-col v-if="props_isAdminView" style="height: 300px; width: 202px; margin-left: 10px; border: none">
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
    props_isAdminView: Boolean,
    className: String
  },
  data() {
    return {};
  },
  methods: {
    openGoodEditor(good) {
      if (this.props_isAdminView) {
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
</script>
