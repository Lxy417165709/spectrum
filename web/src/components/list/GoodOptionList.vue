<!-- eslint-disable -->
<template>
  <el-row>
    <!--    1. 展示 商品选项 部分-->
    <el-col v-for="(goodOption,goodOptionIndex) in goodOptions" :key="goodOptionIndex"
            style="height: 300px; width: 202px; margin-left: 10px; border: none">
      <good-option-card :goodOption="goodOption"></good-option-card>
    </el-col>

    <!--    1. 添加 商品选项 部分-->
    <el-col v-if="props_isAdminView" style="height: 300px; width: 202px; margin-left: 10px; border: none">
      <good-option-special-card @click.native="tryToAddGoodOption()"></good-option-special-card>
    </el-col>
  </el-row>
</template>

<script>
/* eslint-disable */
import GoodCard from "../card/GoodCard";
import GoodSpacialCard from "../card/GoodSpacialCard";
import test from "../../common/test/test"
import GoodOptionCard from "../card/GoodOptionCard";
import GoodOptionSpecialCard from "../card/GoodOptionSpecialCard";

export default {
  components: {GoodOptionSpecialCard, GoodOptionCard, GoodSpacialCard, GoodCard},
  props: {
    props_isAdminView: Boolean,
  },
  data() {
    return {
      goodOptions: [],
      className: "",
    };
  },
  methods: {
    tryToAddGoodOption() {
      if (this.className === "附属选项类") {
        this.openGoodOptionEditorOfAdmin(test.blankGoodOption)
      }
      if (this.className === "附属商品类") {
        this.openGoodOptionEditorOfAdmin(test.blankGoodIngredient)
      }
    },
    openGoodOptionEditorOfAdmin(option) {
      this.$emit("openGoodOptionEditorOfAdmin", option, this.className)
    },
  }
}
</script>
