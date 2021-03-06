<!-- eslint-disable -->
<template>
  <el-row>
    <el-col v-for="(good,goodIndex) in goods" :key="goodIndex"
            style="height: 300px; width: 202px; margin-left: 10px; border: none">
      <good-card :good="good" @dblclick.native="openEditor(good)"></good-card>
    </el-col>
    <el-col v-if="isEditMode" style="height: 300px; width: 202px; margin-left: 10px; border: none">
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
    isEditMode: Boolean
  },
  data() {
    return {
    };
  },
  methods: {
    handleButtonClick() {
      this.$emit("turnToGoodClassListMode")
    },
    openEditor(good) {
      if (this.isEditMode) {
        this.openGoodInfoEditor(good)
      } else {
        this.openGoodSellEditor(good)
      }
    },
    tryToAddGood() {
      this.openGoodInfoEditor(test.blankGood)
    },
    openGoodInfoEditor(good) {
      // 商品信息编辑(管理人员视图)
      this.$emit("openGoodInfoEditor", good)
    },
    openGoodSellEditor(good) {
      // 商品信息编辑(普通用户视图)
      this.$emit("openGoodSellEditor", good)
    }
  }
}

// todo: 菜单显示不太正常！！！
</script>
