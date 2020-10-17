<!--eslint-disable-->
<template>
  <el-tabs v-model="activeName" @tab-click="handleClick">
    <el-tab-pane label="商品添加" name="first">
      <GoodAdder/>
    </el-tab-pane>
    <el-tab-pane label="选项类添加" name="second">
      <OptionAdder/>
    </el-tab-pane>
    <el-tab-pane label="选项类管理" name="third">
      <OptionClassManager/>
    </el-tab-pane>
    <el-tab-pane label="定时任务补偿" name="fourth">定时任务补偿</el-tab-pane>
  </el-tabs>
</template>
<script>

/* eslint-disable */

import GoodAdder from "../components/GoodAdder";
import OptionAdder from "../components/OptionAdder";
import OptionClassManager from "../components/OptionClassManager";
export default {
  name: "panel",
  components: {GoodAdder,OptionAdder,OptionClassManager},
  data() {
    return {
      activeName: 'second'
    };
  },
  methods: {
    handleClick(tab, event) {
      console.log(tab, event);
    }
  }
};
</script>
