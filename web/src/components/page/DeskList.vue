<!-- eslint-disable -->
<template>
  <el-row>
    <!--    1. 桌子列表部分-->
    <div v-if="isDeskListVisible">
      <el-row style="height: 16px;"></el-row>
      <el-divider content-position="left">
        {{ className }}
      </el-divider>
      <!--    1.1 展示 桌位 部分-->
      <el-col v-for="(desk,deskIndex) in desks" :key="deskIndex"
              style="height: 300px; width: 202px; margin-left: 10px; border: none">
        <desk-card :desk="desk" @dblclick.native="handleDbClick(desk)"
                   @click.native="handleClick(deskIndex)"></desk-card>
      </el-col>

      <!--    1.2 添加 桌位 部分-->
      <el-col style="height: 300px; width: 202px; margin-left: 10px; border: none">
        <desk-spacial-card @click.native="tryToAddDesk"></desk-spacial-card>
      </el-col>
    </div>


    <!--     2. 商品类列表部分-->
    <div v-if="!isDeskListVisible">
      <good-class :orderID="curDesk.orderID" :deskSpaceName="curDesk.space.name" :props_isAdminView="false"
                  :props_haveParentComponent="true"
                  ref="ref_goodClass" @turnToParentComponentMode="isDeskListVisible=true"></good-class>
    </div>

    <!--    3. 桌子添加、编辑-->
    <el-dialog
      title="桌子添加/编辑"
      :visible.sync="DeskEditorOfAdminVisible"
      width="30%">
      <desk-editor-of-admin ref="DeskEditorOfAdmin" :className="className"
                            @successToAddDesk="successToAddDesk"></desk-editor-of-admin>
    </el-dialog>


  </el-row>
</template>

<script>
/* eslint-disable */
import DeskCard from "../card/DeskCard";
import DeskSpacialCard from "../card/DeskSpacialCard";
import test from "../../common/test";
import utils from "../../common/utils"
import DeskEditorOfAdmin from "../editor/DeskEditorOfAdmin";
import GoodClass from "./ManageGoodPage";
import cst from "../../common/cst";

let time = null
export default {
  components: {DeskSpacialCard, DeskCard, DeskEditorOfAdmin, GoodClass},
  props: {
    className: String,
  },
  watch: {
    async className(n, o) {
      console.log(n, o)
      await this.getAllDesks()
    }
  },
  data() {
    return {
      desks: [],
      DeskEditorOfAdminVisible: false,
      isDeskListVisible: true,

      curDesk: {},
    };
  },
  async created() {
    await this.getAllDesks()
  },
  methods: {
    async getAllDesks() {
      await utils.GetAllDesks(this, {
        className: this.className
      }, (res) => {
        this.desks = res.data.data.desks
      })
    },
    successToAddDesk(desk) {
      this.desks.push(desk)
      this.DeskEditorOfAdminVisible = false
    },


    async turnToClassListMode(deskIndex) {
      // todo: this.desks[deskIndex].id === undefined 时，此时 this.desks[deskIndex].id !== 0 也成立..
      if (this.desks[deskIndex].id !== 0) {
        this.curDeskIndex = deskIndex
        this.curDesk = this.desks[deskIndex]
        this.isDeskListVisible = false
        console.log("curDesk", this.desks[deskIndex])
        return
      }
      // 进行点单
      await utils.OrderDesk(this, {
        desk: this.desks[deskIndex],
      }, (res) => {
        this.desks[deskIndex].id = res.data.data.order.desk.id;
        this.desks[deskIndex].orderID = res.data.data.order.id;
      })
    },
    tryToAddDesk() {
      this.openDeskEditorOfAdmin(test.blankDesk)
    },
    openDeskEditorOfAdmin(desk) {
      this.DeskEditorOfAdminVisible = true
      this.$nextTick(() => {
        this.$refs.DeskEditorOfAdmin.desk = desk
      })
    },
    // 单击事件函数
    handleClick(deskIndex) {
      clearTimeout(time);
      time = setTimeout(() => {
        this.turnToClassListMode(deskIndex)
      }, 500);
    },
    // 双击事件函数
    handleDbClick(desk) {
      clearTimeout(time);
      this.openDeskEditorOfAdmin(desk)
    }
  }
}
</script>
