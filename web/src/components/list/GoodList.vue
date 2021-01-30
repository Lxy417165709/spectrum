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

    <vue-context-menu :contextMenuData="contextMenuData" @openEditor="openEditor"></vue-context-menu>
  </el-row>

</template>

<script>
/* eslint-disable */
import GoodCard from "../card/GoodCard";
import GoodSpacialCard from "../card/GoodSpacialCard";

export default {
  components: {GoodSpacialCard, GoodCard},
  props: {
    goods: Array,
    isEditMode: Boolean
  },
  data() {
    return {
      contextMenuData: {
        menuName: 'demo',
        good: {},
        //菜单显示的位置
        axis: {
          x: null,
          y: null
        },
        //菜单选项
        menulists: [{
          fnHandler: 'openEditor', //绑定事件
          // icoName: 'fa fa-home fa-fw', //icon图标
          btnName: '编辑' //菜单名称
        }]
      }
    };
  },
  methods: {
    handleButtonClick() {
      this.$emit("turnToGoodClassListMode")
    },
    handleSpecialCardClick() {
      console.log("handleSpecialCardClick")
      this.$emit("openGoodEditor", {})
    },
    showMenu(good, e) {
      e.preventDefault()
      this.contextMenuData.good = good
      this.contextMenuData.axis = {
        x: e.clientX,
        y: e.clientY
      }
    },
    openEditor() {
      console.log("this.contextMenuData.good", this.contextMenuData.good)
      this.$emit("openGoodEditor", this.contextMenuData.good)
    },
  }
}
</script>
