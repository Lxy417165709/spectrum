<!-- eslint-disable -->
<template>
  <el-row>
    <el-col v-for="(good,goodIndex) in goods" :key="goodIndex"
            style="height: 300px; width: 202px; margin-left: 10px; border: none">
      <!--      todo: 为什么放在这里面才能右键呢-->
      <div @contextmenu="showMenu(good,$event)">
        <good-card :good="good"></good-card>
      </div>
    </el-col>
    <el-col v-if="isEditMode" style="height: 300px; width: 202px; margin-left: 10px; border: none">
      <good-spacial-card @click.native="handleSpecialCardClick"></good-spacial-card>
    </el-col>

    <vue-context-menu v-if='isEditMode' :contextMenuData="editContextMenuData"
                      @openGoodInfoEditor="openGoodInfoEditor"></vue-context-menu>
    <vue-context-menu v-if='!isEditMode' :contextMenuData="notEditContextMenuData"
                      @openGoodSellEditor="openGoodSellEditor"></vue-context-menu>
  </el-row>

</template>

<script>
/* eslint-disable */
import GoodCard from "../card/GoodCard";
import GoodSpacialCard from "../card/GoodSpacialCard";
import test from "../../common/test/test"

export default {
  components: {GoodSpacialCard, GoodCard},
  props: {
    goods: Array,
    isEditMode: Boolean
  },
  data() {
    return {
      editContextMenuData: {
        menuName: 'demo',
        good: {},
        //菜单显示的位置
        axis: {
          x: null,
          y: null
        },
        //菜单选项
        menulists: [
          {
            fnHandler: 'openGoodInfoEditor', //绑定事件
            // icoName: 'fa fa-home fa-fw', //icon图标
            btnName: '编辑' //菜单名称
          },
        ]
      },
      notEditContextMenuData: {
        menuName: 'demo',
        good: {},
        //菜单显示的位置
        axis: {
          x: null,
          y: null
        },
        //菜单选项
        menulists: [
          {
            fnHandler: 'openGoodSellEditor', //绑定事件
            // icoName: 'fa fa-home fa-fw', //icon图标
            btnName: '点单' //菜单名称
          },
        ]
      }
    };
  },
  methods: {
    handleButtonClick() {
      this.$emit("turnToGoodClassListMode")
    },
    handleSpecialCardClick() {
      console.log("handleSpecialCardClick")
      this.$emit("openGoodEditor", test.blankGood)
    },
    showMenu(good, e) {
      e.preventDefault()
      this.notEditContextMenuData.good = good
      this.editContextMenuData.good = good
      this.editContextMenuData.axis = {
        x: e.clientX,
        y: e.clientY
      }
      this.notEditContextMenuData.axis = {
        x: e.clientX,
        y: e.clientY
      }
    },
    openGoodInfoEditor() {
      console.log("this.editContextMenuData.good", this.editContextMenuData.good)
      this.$emit("openGoodInfoEditor", this.editContextMenuData.good)
    },
    openGoodSellEditor() {
      console.log("this.notEditContextMenuData.good", this.notEditContextMenuData.good)
      this.$emit("openGoodSellEditor", this.notEditContextMenuData.good)
    }
  }
}

// todo: 菜单显示不太正常！！！
</script>
