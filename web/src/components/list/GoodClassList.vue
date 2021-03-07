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
      <good-class-spacial-card @click.native="openGoodClassEditor"></good-class-spacial-card>
    </el-col>
  </el-row>
</template>

<script>
/* eslint-disable */
import GoodClassCard from "../card/GoodClassCard";
import GoodClassSpacialCard from "../card/GoodClassSpacialCard";
import test from "../../common/test/test";

export default {
  components: {GoodClassSpacialCard, GoodClassCard},
  props: {
    goodClasses: Array,
    props_isAdminView: Boolean,
  },
  data() {
    return {};
  },
  methods: {
    handleGoodClassCardClick(goodClassIndex) {
      this.$emit("turnToGoodListMode", goodClassIndex)
    },
    openGoodClassEditor() {
      this.$emit("openGoodClassEditor", test.blankGoodClass)
    }
  }
}
</script>
