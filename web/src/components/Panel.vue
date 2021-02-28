<!--eslint-disable-->
<template>
  <el-tabs v-model="activeName" @tab-click="handleClick">
    <el-tab-pane label="桌位管理" name="1">
      <DeskSet/>
    </el-tab-pane>
    <el-tab-pane label="商品管理" name="2">
      <GoodClass :isEditMode="true" :hasFather="false"/>
    </el-tab-pane>
    <el-tab-pane label="订单" name="3">
      <Order></Order>
    </el-tab-pane>
  </el-tabs>
</template>
<script>

/* eslint-disable */
import GoodOrder from "./unuse/GoodOrder";
import DeskSet from "./page/DeskSet";
import GoodClass from "./manager/GoodClass";
import GoodSellEditor from "./editor/GoodSellEditor";
import Order from "./page/Order";
import test from "../common/test/test";

export default {

  name: "panel",
  components: {
    Order,
    GoodSellEditor,
    GoodClass,
    GoodOrder,
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
