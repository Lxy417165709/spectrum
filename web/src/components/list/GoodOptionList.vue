<!-- eslint-disable -->
<template>
  <el-row>
    <!--    1. 展示 商品选项 部分-->
    <el-col v-for="(goodOption,goodOptionIndex) in goodOptions" :key="goodOptionIndex"
            style="height: 300px; width: 202px; margin-left: 10px; border: none">
      <good-option-card :goodOption="goodOption"></good-option-card>
    </el-col>
    <el-col v-if="props_isAdminView" style="height: 300px; width: 202px; margin-left: 10px; border: none">
      <good-option-special-card></good-option-special-card>
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
    goodOptions: Array,
    props_isAdminView: Boolean,
  },
  data() {
    return {};
  },
  methods: {}
}
</script>
