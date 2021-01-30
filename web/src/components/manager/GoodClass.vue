<!-- eslint-disable -->
<template>
  <div>
    <el-row style="height: 40px;margin-left:10px;margin-bottom: 10px">
      <el-button v-show="viewMode === 1 && hasFather || viewMode === 2" style="height: 40px" @click="handleButtonClick"
                 type="primary">
        退回
      </el-button>
    </el-row>
    <div v-show="viewMode === 1">
      <el-divider content-position="left">普通商品类</el-divider>
      <good-class-list :goodClasses="goodClasses" :isEditMode="isEditMode"
                       @turnToGoodListMode="turnToGoodListMode" @openGoodClassEditor="openGoodClassEditor"
                       :classType="1"></good-class-list>

      <el-divider content-position="left" v-if="isEditMode">附属类</el-divider>
      <good-class-list :goodClasses="goodClasses" :isEditMode="isEditMode"
                       @turnToGoodListMode="turnToGoodListMode" @openGoodClassEditor="openGoodClassEditor"
                       :classType="-1" v-if="isEditMode"></good-class-list>
    </div>
    <div v-show="viewMode === 2">
      <el-divider content-position="left">元素</el-divider>
      <good-list v-if="curGoodClassIndex!==-1" :isEditMode="isEditMode" :goods="goodClasses[curGoodClassIndex].goods"
                 @turnToGoodClassListMode="turnToGoodClassListMode" @openGoodEditor="openGoodEditor"></good-list>
    </div>

    <!--    :before-close="handleClose"-->
    <el-dialog
      title="商品添加/编辑"
      :visible.sync="goodEditorVisible"
      width="30%">
      <good-editor ref="goodEditor"></good-editor>
    </el-dialog>

    <el-dialog
      title="商品类添加/编辑"
      :visible.sync="goodClassEditorVisible"

      width="30%">
      <good-class-editor></good-class-editor>
    </el-dialog>
  </div>
</template>

<script>
/* eslint-disable */
import GoodClassList from "../list/GoodClassList";
import GoodList from "../list/GoodList";
import GoodEditor from "../editor/GoodEditor";
import GoodClassEditor from "../editor/GoodClassEditor";
import test from "../../common/test/test";

export default {
  name: "GoodClass",
  components: {GoodClassEditor, GoodEditor, GoodList, GoodClassList},
  props: {
    goodClasses: Array,
    isEditMode: Boolean,
    hasFather: Boolean,
  },
  data() {
    return {
      viewMode: 1,
      goodEditorVisible: false,
      goodClassEditorVisible: false,
      curGoodClassIndex: -1,
    }
  },
  mounted() {
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
    openGoodEditor(good) {
      this.goodEditorVisible = true
      this.$nextTick(() => {
        this.$refs.goodEditor.good = good
      })
    },
    openGoodClassEditor() {
      this.goodClassEditorVisible = true
    }
  }
}
</script>

<style scoped>

</style>
