<!-- eslint-disable -->
<template>
  <el-row>
    <!--    1. 展示 商品选项 部分-->
    <el-col v-for="(goodOption,goodOptionIndex) in goodOptions" :key="goodOptionIndex"
            style="height: 300px; width: 202px; margin-left: 10px; border: none">
      <good-option-card :goodOption="goodOption"
                        @dblclick.native="openGoodOptionEditorOfAdmin(goodOption,false)"></good-option-card>
    </el-col>

    <!--    2. 添加 商品选项 部分-->
    <el-col v-if="props_isAdminView" style="height: 300px; width: 202px; margin-left: 10px; border: none">
      <good-option-special-card @click.native="tryToAddGoodOption()"></good-option-special-card>
    </el-col>

    <!--    3. 商品选项添加、编辑框-->
    <el-dialog
      title="商品选项添加/编辑"
      :visible.sync="GoodOptionEditorOfAdminVisible"
      width="30%">
      <good-option-editor-of-admin ref="GoodOptionEditorOfAdmin" :className="className"></good-option-editor-of-admin>
    </el-dialog>

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
import GoodOptionEditorOfAdmin from "../editor/GoodOptionEditorOfAdmin";

export default {
  components: {GoodOptionSpecialCard, GoodOptionCard, GoodSpacialCard, GoodCard, GoodOptionEditorOfAdmin},
  props: {
    props_isAdminView: Boolean,
    className: String,
  },
  data() {
    return {
      GoodOptionEditorOfAdminVisible: false,
      goodOptions: []
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
      this.GoodOptionEditorOfAdminVisible = true
      this.$nextTick(() => {
        this.$refs.GoodOptionEditorOfAdmin.option = option
        this.$refs.GoodOptionEditorOfAdmin.curSizeInfoIndexString = option.selectedIndex.toString()
      })
    },

  }
}
</script>
