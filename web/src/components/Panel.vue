<!--eslint-disable-->
<template>
  <el-tabs v-model="activeName" @tab-click="handleClick">
    <el-tab-pane label="商品添加" name="1">
      <GoodAdder/>
    </el-tab-pane>
    <el-tab-pane label="选项类添加" name="2">
      <OptionClassAdder/>
    </el-tab-pane>
    <el-tab-pane label="选项类管理" name="3">
      <OptionClassManager/>
    </el-tab-pane>
    <el-tab-pane label="附属品管理" name="4">
      <AttachGoodManager/>
    </el-tab-pane>
    <el-tab-pane label="商品展示" name="5">
      <GoodShower/>
    </el-tab-pane>
    <el-tab-pane label="商品选择"  name="6">
      <GoodSelector/>
    </el-tab-pane>
  </el-tabs>
</template>
<script>

/* eslint-disable */
import GoodAdder from "../components/GoodAdder";
import OptionClassAdder from "./OptionClassAdder";
import OptionClassManager from "../components/OptionClassManager";
import GoodShower from "./GoodShower";
import GoodSelector from "./GoodSelector";
import GoodOrder from "./GoodOrder";
import AttachGoodManager from "./AttachGoodManager"

export default {
  name: "panel",
  components: {GoodSelector, GoodShower, GoodAdder,OptionClassAdder,OptionClassManager,GoodOrder,AttachGoodManager},
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
