<!-- eslint-disable -->
<template>
  <div>

    <!--    1. 顶栏-->
    <el-row style="height: 40px;margin-left:10px;margin-bottom: 10px">
      <el-button v-show="viewMode === 1 && hasFather || viewMode === 2" style="height: 40px" @click="handleButtonClick"
                 type="primary">
        退回
      </el-button>
    </el-row>

    <!--    2. 商品类展示-->
    <div v-show="viewMode === 1">
      <el-divider content-position="left">普通商品类</el-divider>
      <good-class-list :goodClasses="goodClasses" :isAdminView="isAdminView"
                       @turnToGoodListMode="turnToGoodListMode" @openGoodClassEditor="openGoodClassEditor"
                       :classType="1"></good-class-list>

      <el-divider content-position="left" v-if="isAdminView">附属类</el-divider>
      <good-class-list :goodClasses="goodClasses" :isAdminView="isAdminView"
                       @turnToGoodListMode="turnToGoodListMode" @openGoodClassEditor="openGoodClassEditor"
                       :classType="-1" v-if="isAdminView"></good-class-list>
    </div>

    <!--    3. 商品类内的商品展示-->
    <div v-show="viewMode === 2">
      <el-divider content-position="left">元素</el-divider>
      <good-list v-if="curGoodClassIndex!==-1"
                 :isAdminView="isAdminView"
                 :goods="goodClasses[curGoodClassIndex].goods"
                 :className="goodClasses[curGoodClassIndex].name"
                 @turnToGoodClassListMode="turnToGoodClassListMode"
                 @openGoodEditorOfAdmin="openGoodEditorOfAdmin"
                 @openGoodEditorOfUser="openGoodEditorOfUser"></good-list>
    </div>

    <!--    4. 商品添加、编辑框-->
    <el-dialog
      title="商品添加/编辑"
      :visible.sync="GoodEditorOfAdminVisible"
      width="30%">
      <good-editor-of-admin ref="GoodEditorOfAdmin"></good-editor-of-admin>
    </el-dialog>

    <!--    5. 商品下单框-->
    <el-dialog
      title="商品点单"
      :visible.sync="GoodEditorOfUserVisible"
      width="30%">
      <good-editor-of-user ref="GoodEditorOfUser"></good-editor-of-user>
    </el-dialog>

    <!--    6. 商品类添加、编辑框-->
    <el-dialog
      title="商品类添加/编辑"
      :visible.sync="GoodClassEditorVisible"
      width="30%">
      <good-class-editor ref="GoodClassEditor"></good-class-editor>
    </el-dialog>
  </div>
</template>

<script>
/* eslint-disable */
import GoodClassList from "../list/GoodClassList";
import GoodList from "../list/GoodList";
import GoodEditorOfAdmin from "../editor/GoodEditorOfAdmin";
import GoodEditorOfUser from "../editor/GoodEditorOfUser";
import GoodClassEditor from "../editor/GoodClassEditor";
import test from "../../common/test/test";

export default {
  name: "GoodClass",
  components: {GoodClassEditor, GoodEditorOfAdmin, GoodList, GoodClassList, GoodEditorOfUser},
  props: {
    isAdminView: Boolean,
    hasFather: Boolean,
  },
  mounted() {
    this.goodClasses = test.goodClasses
  },
  data() {
    return {
      viewMode: 1,
      GoodEditorOfAdminVisible: false,
      GoodClassEditorVisible: false,
      GoodEditorOfUserVisible: false,
      curGoodClassIndex: -1,

      // 数据库读取属性
      goodClasses: [],
    }
  },
  methods: {
    turnToGoodListMode(mode, goodClassIndex) {
      this.viewMode = mode
      this.curGoodClassIndex = goodClassIndex
    },
    turnToGoodClassListMode() {
      this.viewMode = 1
    },
    handleButtonClick() {
      if (this.viewMode === 1) {
        this.$emit("turnToFatherMode")
      }
      if (this.viewMode === 2) {
        this.viewMode = 1
      }
    },
    openGoodEditorOfAdmin(good, className) {
      this.GoodEditorOfAdminVisible = true
      this.$nextTick(() => {
        this.$refs.GoodEditorOfAdmin.good = good
        this.$refs.GoodEditorOfAdmin.className = className
      })
    },
    openGoodEditorOfUser(good) {
      this.GoodEditorOfUserVisible = true
      this.$nextTick(() => {
        this.$refs.GoodEditorOfUser.good = good
      })
    },
    openGoodClassEditor(goodClass) {
      this.GoodClassEditorVisible = true
      this.$nextTick(() => {
        console.log("openGoodClassEditor-this.$nextTick", goodClass)
        this.$refs.GoodClassEditor.goodClass = goodClass
      })
    }
  }
}
</script>

<style scoped>

</style>
