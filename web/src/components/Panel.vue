<!--eslint-disable-->
<template>
  <el-tabs v-model="activeName" @tab-click="handleClick">
    <el-tab-pane label="桌位管理" name="1">
      <ManageDeskPage/>
    </el-tab-pane>
    <el-tab-pane label="商品管理" name="2">
      <ManageGoodPage :props_isAdminView="true" :props_haveParentComponent="false"/>
    </el-tab-pane>
    <el-tab-pane label="订单管理" name="3">
      <ManageOrderPage/>
    </el-tab-pane>
  </el-tabs>
</template>
<script>

/* eslint-disable */
import ManageDeskPage from "./page/ManageDeskPage";
import ManageGoodPage from "./page/ManageGoodPage";
import ManageOrderPage from "./page/ManageOrderPage";

export default {
  name: "panel",
  components: {
    ManageOrderPage,
    ManageGoodPage,
    ManageDeskPage,
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
