<!--eslint-disable-->
<template>
  <el-tabs v-model="activeName" @tab-click="handleClick">
    <el-tab-pane label="桌位管理" name="9">
      <DeskSet/>
    </el-tab-pane>
    <el-tab-pane label="商品管理" name="10">
      <GoodClass :isEditMode="true" :hasFather="false"/>
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
import GoodOrder from "./unuse/GoodOrder";
import DeskSet from "./page/DeskSet";
import GoodClass from "./manager/GoodClass";
import GoodSellEditor from "./editor/GoodSellEditor";
import Order from "./page/Order";
import test from "../common/test/test";
import global from "../common/global_object/global";
import utils from "../common/utils";


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
      // let model = utils.getRequestModel("mvp", "GetAllGoodClasses", {})
      // utils.sendRequestModel(model).then(res => {
      //   if (!utils.hasRequestSuccess(res)) {
      //     console.log(res.data.err)
      //     return
      //   }
      //   // localStorage.setItem("global_good_classes",JSON.stringify({
      //   //   classes:global.goodClasses
      //   // }))
      //   if (utils.hasData(res.data.data.goodClasses)) {
      //     global.goodClasses = res.data.data.goodClasses
      //   }
      //   console.log("global_good_classes",global.goodClasses)
      // })
    }
  }
};
</script>
