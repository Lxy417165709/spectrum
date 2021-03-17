<!--eslint-disable-->
<template>
  <el-tabs v-model="activeName" @tab-click="handleClick">
    <el-tab-pane label="桌位管理" name="1">
      <OrderPage/>
    </el-tab-pane>
    <el-tab-pane label="商品管理" name="2">
      <ManageGoodPage :props_isAdminView="true" :props_haveParentComponent="false"/>
    </el-tab-pane>
    <el-tab-pane label="订单" name="3">
      <ManageOrderPage/>
    </el-tab-pane>
  </el-tabs>
</template>
<script>

/* eslint-disable */
import OrderPage from "./page/OrderPage";
import ManageGoodPage from "./page/ManageGoodPage";
import ManageOrderPage from "./page/ManageOrderPage";

export default {
  name: "panel",
  components: {
    ManageOrderPage,
    ManageGoodPage,
    OrderPage,
  },
  mounted() {
  },
  data() {
    return {
      activeName: 'second',
    }
  },
  methods: {
    handleClick(tab, event) {
      console.log(tab, event);
    }
  }
};
</script>
