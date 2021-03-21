<!-- eslint-disable -->
<template>
  <el-row>
    <el-divider content-position="left">{{ className }}</el-divider>
    <!--    1. 展示 商品 部分-->
    <el-col v-for="(good,goodIndex) in goods" :key="goodIndex"
            style="height: 300px; width: 202px; margin-left: 10px; border: none">
      <good-card :good="good" @dblclick.native="openGoodEditor(good)"></good-card>
    </el-col>

    <!--    2. 添加 商品 部分-->
    <el-col v-if="props_isAdminView" style="height: 300px; width: 202px; margin-left: 10px; border: none">
      <good-spacial-card @click.native="tryToAddGood()"></good-spacial-card>
    </el-col>

    <!--    3. 商品下单框-->
    <el-dialog
      title="商品点单"
      :visible.sync="GoodEditorOfUserVisible"
      width="30%">
      <good-editor-of-user ref="GoodEditorOfUser" :orderID="orderID"></good-editor-of-user>
    </el-dialog>

    <!--    4. 商品添加、编辑框-->
    <el-dialog
      title="商品添加/编辑"
      :visible.sync="GoodEditorOfAdminVisible"
      width="30%">
      <good-editor-of-admin ref="GoodEditorOfAdmin" :className="className"></good-editor-of-admin>
    </el-dialog>

  </el-row>
</template>

<script>
/* eslint-disable */
import GoodCard from "../card/GoodCard";
import GoodSpacialCard from "../card/GoodSpacialCard";
import test from "../../common/test/test"
import GoodEditorOfUser from "../editor/GoodEditorOfUser";
import GoodEditorOfAdmin from "../editor/GoodEditorOfAdmin";

export default {
  components: {GoodEditorOfAdmin, GoodEditorOfUser, GoodSpacialCard, GoodCard},
  props: {
    props_isAdminView: Boolean,
    className: String,
    orderID:Number,
  },
  data() {
    return {
      goods: [],
      GoodEditorOfUserVisible: false,
      GoodEditorOfAdminVisible: false,
    };
  },
  methods: {
    openGoodEditor(good) {
      if (this.props_isAdminView) {
        this.openGoodEditorOfAdmin(good, false)
      } else {
        this.openGoodEditorOfUser(good)
      }
    },
    tryToAddGood() {
      this.openGoodEditorOfAdmin(test.blankGood, true)
    },
    openGoodEditorOfAdmin(good, canModifyGoodName) {
      // 商品信息编辑(管理人员视图)
      this.GoodEditorOfAdminVisible = true
      this.$nextTick(() => {
        this.$refs.GoodEditorOfAdmin.good = good
        this.$refs.GoodEditorOfAdmin.canModifyGoodName = canModifyGoodName
      })
    },
    openGoodEditorOfUser(good) {
      // 商品信息编辑(普通用户视图)
      this.GoodEditorOfUserVisible = true
      this.$nextTick(() => {
        this.$refs.GoodEditorOfUser.good = good
      })
    },
  }
}
</script>
