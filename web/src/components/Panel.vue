<!--eslint-disable-->
<template>
  <el-tabs v-model="activeName" @tab-click="handleClick">
    <el-tab-pane label="商品添加" name="1">
      <GoodAdder/>
    </el-tab-pane>
    <el-tab-pane label="商品类管理" name="2">
      <GoodClassManager/>
    </el-tab-pane>
    <el-tab-pane label="选项类管理" name="3">
      <OptionClassManager/>
    </el-tab-pane>
    <el-tab-pane label="商品点单"  name="8">
      <GoodWanner/>
    </el-tab-pane>
  </el-tabs>
</template>
<script>

/* eslint-disable */
import GoodAdder from "./good/GoodAdder";
import OptionClassManager from "./option/OptionClassManager";
import GoodOrder from "./unuse/GoodOrder";
import GoodWanner from "./good/GoodWanner";
import GoodClassManager from "./good/GoodClassManager";
import Discounter from "./discount/Discounter";
export default {
  name: "panel",
  components: {
    GoodAdder,
    OptionClassManager,
    GoodOrder,
    GoodWanner,
    GoodClassManager,
    Discounter,
  },
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
