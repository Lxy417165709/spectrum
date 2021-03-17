<!-- eslint-disable -->
<template>
  <el-row>

    <!--    1. 展示 商品类 部分-->
    <el-col v-for="(goodClass,goodClassIndex) in goodClasses" :key="goodClassIndex"
            style="height: 300px; width: 202px; margin-left: 10px; border: none">

      <good-class-card :name="goodClass.name"
                       :pictureStorePath="goodClass.pictureStorePath"
                       @click.native="handleGoodClassCardClick(goodClassIndex)"></good-class-card>

    </el-col>

    <!--    2. 添加 商品类 部分-->
    <el-col v-if="props_isAdminView" style="height: 300px; width: 202px; margin-left: 10px; border: none">
      <good-class-spacial-card @click.native="tryToAddGoodClass"></good-class-spacial-card>
    </el-col>

    <!--    3. 商品类添加、编辑-->
    <el-dialog
      title="商品添加/编辑"
      :visible.sync="GoodClassEditorVisible"
      width="30%">
      <good-class-editor ref="GoodClassEditor"></good-class-editor>
    </el-dialog>
  </el-row>
</template>

<script>
/* eslint-disable */
import GoodClassCard from "../card/GoodClassCard";
import GoodClassSpacialCard from "../card/GoodClassSpacialCard";
import test from "../../common/test/test";
import GoodClassEditor from "../editor/GoodClassEditor";

export default {
  components: {GoodClassEditor, GoodClassSpacialCard, GoodClassCard},
  props: {
    goodClasses: Array,
    props_isAdminView: Boolean,
  },
  data() {
    return {
      GoodClassEditorVisible: false
    };
  },
  methods: {
    handleGoodClassCardClick(goodClassIndex) {
      this.$emit("turnToGoodListMode", goodClassIndex)
    },
    tryToAddGoodClass() {
      this.openGoodClassEditor(test.blankGoodClass)
    },
    openGoodClassEditor(goodClass) {
      this.GoodClassEditorVisible = true
      this.$nextTick(() => {
        this.$refs.GoodClassEditor.goodClass = goodClass
      })
    }
  }
}
</script>
