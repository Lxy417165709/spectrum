<!-- eslint-disable -->
<template>
  <el-row>
    <!--    1. 展示 商品选项 部分-->
    <el-col v-for="(goodOption,goodOptionIndex) in goodOptions" :key="goodOptionIndex"
            style="height: 300px; width: 202px; margin-left: 10px; border: none">
      <good-option-card :goodOption="goodOption"></good-option-card>
    </el-col>

    <!--    2. 添加 商品选项 部分-->
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
import cst from "../../common/cst";

export default {
  components: {GoodOptionSpecialCard, GoodOptionCard, GoodSpacialCard, GoodCard},
  props: {
    props_isAdminView: Boolean,
  },
  data() {
    return {
      goodOptions: [],// todo: 这里的 goodOption 含义包括了 附属选项、附属商品，之后要加以区分呀！
      className: "",
    };
  },
  methods: {
    tryToAddGoodOption() {
      if (this.className === cst.ATTACH_CLASS_NAME.GOOD_OPTION_CLASS_NAME) {
        this.openGoodOptionEditorOfAdmin(test.blankGoodOption)
      }
      if (this.className === cst.ATTACH_CLASS_NAME.GOOD_INGREDIENT_CLASS_NAME) {
        this.openGoodOptionEditorOfAdmin(test.blankGoodIngredient)
      }
    },
    openGoodOptionEditorOfAdmin(option) {
      this.$emit("openGoodOptionEditorOfAdmin", option, this.className)
    },
  }
}
</script>
