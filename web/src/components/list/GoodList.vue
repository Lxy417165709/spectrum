<!-- eslint-disable -->
<template>
  <el-row>
    <el-col v-for="(good,goodIndex) in goods" :key="goodIndex"
            style="height: 300px; width: 202px; margin-left: 10px; border: none">
      <!--      todo: 为什么放在这里面才能右键呢-->
      <div @contextmenu="showMenu">
        <good-card :name="good.name" :pictureStorePath="good.pictureStorePath"></good-card>
      </div>

    </el-col>
    <el-col v-if="isEditMode" style="height: 300px; width: 202px; margin-left: 10px; border: none">
      <good-spacial-card @click.native="handleSpecialCardClick"></good-spacial-card>
    </el-col>


    <vue-context-menu :contextMenuData="contextMenuData" @home="home" @deletedata="deletedata"></vue-context-menu>

  </el-row>

</template>

<script>
/* eslint-disable */
import GoodCard from "../card/GoodCard";
import GoodSpacialCard from "../card/GoodSpacialCard";

export default {
  components: {GoodSpacialCard, GoodCard},
  watch: {
    visible(value) {
      if (value) {
        document.body.addEventListener('click', this.closeMenu)
      } else {
        document.body.removeEventListener('click', this.closeMenu)
      }
    }
  },
  props: {
    goods: Array,
    isEditMode: Boolean
  },
  data() {
    return {
      contextMenuData: {
        menuName: 'demo',
        //菜单显示的位置
        axis: {
          x: null,
          y: null
        },
        //菜单选项
        menulists: [{
          fnHandler: 'home', //绑定事件
          icoName: 'fa fa-home fa-fw', //icon图标
          btnName: '回到主页' //菜单名称
        }, {
          fnHandler: 'deletedata',
          icoName: 'fa fa-minus-square-o  fa-fw',
          btnName: '删除布局'
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
      this.$emit("openGoodEditor")
    },
    showMenu() {
      event.preventDefault()
      let x = event.clientX
      let y = event.clientY
      // Get the current location
      this.contextMenuData.axis = {
        x, y
      }
    },
    home() {
      alert("主页")
    },
    deletedata() {
      console.log('delete!')
    }
  }
}
</script>
