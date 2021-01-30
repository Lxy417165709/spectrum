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
    <el-tab-pane label="商品点单" name="8">
      <GoodWanner/>
    </el-tab-pane>
    <el-tab-pane label="桌位管理" name="9">
      <DeskSet/>
    </el-tab-pane>
    <el-tab-pane label="商品管理" name="10">
      <!--      todo: 不要用 props 吧，用Props组件内部修改属性时，会报错-->
      <GoodClass :goodClasses="goodClasses" :isEditMode="true" :hasFather="false"/>
    </el-tab-pane>
    <el-tab-pane label="商品销售" name="11">
      <GoodSellEditor :good="goodClasses[0].goods[0]"></GoodSellEditor>
    </el-tab-pane>
    <el-tab-pane label="订单" name="12">
      <Order></Order>
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
import DeskSet from "./page/DeskSet";
import GoodClass from "./manager/GoodClass";
import GoodCard from "./card/GoodCard";
import GoodSellEditor from "./editor/GoodSellEditor";
import Order from "./page/Order";
import test from "../common/test/test";


export default {

  name: "panel",
  components: {
    Order,
    GoodSellEditor,
    GoodClass,
    GoodAdder,
    OptionClassManager,
    GoodOrder,
    GoodWanner,
    GoodClassManager,
    DeskSet,
  },
  mounted() {
    this.goodClasses = test.goodClasses
  },
  data() {
    return {
      activeName: 'second',
      goodClasses: [],
    }
  },
  methods: {
    handleClick(tab, event) {
      console.log(tab, event);
    }
  }
};
</script>
