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
      <GoodClass :goodClasses="goodClasses" :isEditMode="true"/>
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


export default {

  name: "panel",
  components: {
    GoodClass,
    GoodAdder,
    OptionClassManager,
    GoodOrder,
    GoodWanner,
    GoodClassManager,
    DeskSet,
  },
  data() {
    return {
      activeName: 'second',
      goodClasses: [
        {
          name: "奶茶系列",
          goods: [
            {
              component: GoodCard,
              name: "波霸奶茶",
            },
            {
              component: GoodCard,
              name: "红豆奶茶",
            }
          ]
        },
        {
          name: "水果茶系列",
          goods: [
            {
              component: GoodCard,
              name: "超神水果茶",
            },
            {
              component: GoodCard,
              name: "我爱水果茶",
            }
          ]
        },
        {
          name: "小吃系列",
          goods: [
            {
              component: GoodCard,
              name: "炸鸡",
            },
            {
              component: GoodCard,
              name: "薯条",
            }
          ]
        }
      ],
    };
  },
  methods: {
    handleClick(tab, event) {
      console.log(tab, event);
    }
  }
};
</script>
