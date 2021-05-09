<!-- eslint-disable -->
<template>
  <el-container style="height: 800px; border: 1px solid #eee">
    <!--    1. 桌类选择、添加-->
    <el-aside width="200px">
      <el-menu>
        <el-menu-item v-for="(deskClass,deskClassIndex) in db_deskClasses" :key="deskClassIndex"
                      @click.native="handleDeskClassItemClick(deskClassIndex)"
                      @dblclick.native="handleDeskClassItemDbClick(deskClass)">
          <template slot="title"><i class="el-icon-message"></i><span>{{ deskClass.name }}</span></template>
        </el-menu-item>
      </el-menu>
      <el-button style="margin-left: 20px;margin-top: 10px" type="primary" @click="tryToAddDeskClass">添加桌位
      </el-button>
    </el-aside>

    <!--    2. 桌类详情、商品点单-->
    <el-main>
      <desk-list
        v-if="cpt_curDeskClass!==undefined"
        :className="cpt_curDeskClass.name"></desk-list>
    </el-main>

    <!--    4. 桌类添加、编辑-->
    <el-dialog
      title="桌类添加/编辑"
      :visible.sync="deskClassEditorVisible"
      width="30%">
      <desk-class-editor ref="DeskClassEditor" @successToAddDeskClass="successToAddDeskClass"></desk-class-editor>
    </el-dialog>
  </el-container>
</template>
<script>
/* eslint-disable */
import DeskList from "./DeskList";
import GoodClass from "./ManageGoodPage";
import DeskClassEditor from "../editor/DeskClassEditor";
import test from "../../common/test";
import cst from "../../common/cst";
import utils from "../../common/utils";
import DeskEditorOfAdmin from "../editor/DeskEditorOfAdmin";

let time = null;
export default {
  name: 'ManageDeskPage',
  components: {DeskEditorOfAdmin, DeskClassEditor, GoodClass, DeskList},
  async created() {
    await this.flashAllDeskClasses()
  },
  data() {
    return {
      viewMode: cst.VIEW_MODE.DESK_LIST_MODE,

      deskClassEditorVisible: false,
      deskEditorOfAdminVisible: false,

      curDeskClassIndex: cst.INDEX.INVALID_INDEX,
      curDeskIndex: cst.INDEX.INVALID_INDEX,

      db_deskClasses: [],
    }
  },
  methods: {
    async flashAllDeskClasses() {
      await utils.GetAllDeskClasses(this, {}, (res) => {
        this.db_deskClasses = res.data.data.deskClasses
      })
    },

    async showDeskClassDetail(deskClassIndex) {
      this.curDeskClassIndex = deskClassIndex
      this.curDeskIndex = cst.INDEX.INVALID_INDEX
      this.viewMode = cst.VIEW_MODE.DESK_LIST_MODE
    },
    turnToClassListMode(deskIndex, deskID, orderID) {
      this.viewMode = cst.VIEW_MODE.CLASS_LIST_MODE
      this.curDeskIndex = deskIndex
      this.$refs.ref_goodClass.orderID = orderID
    },
    openDeskClassEditor(deskClass) {
      this.deskClassEditorVisible = true
      this.$nextTick(() => {
        this.$refs.DeskClassEditor.deskClass = deskClass
      })
    },
    // 单击事件函数
    handleDeskClassItemClick(deskClassIndex) {
      clearTimeout(time);
      time = setTimeout(() => {
        this.showDeskClassDetail(deskClassIndex)
      }, 500);
    },
    // 双击事件函数
    handleDeskClassItemDbClick(deskClass) {
      clearTimeout(time);
      this.openDeskClassEditor(deskClass)
    },
    tryToAddDeskClass() {
      this.openDeskClassEditor(test.blankDeskClass)
    },
    successToAddDeskClass() {
      this.deskClassEditorVisible = false
      this.flashAllDeskClasses()
    }
  },
  computed: {
    cpt_canClassListShow() {
      return this.viewMode === cst.VIEW_MODE.CLASS_LIST_MODE
    },
    cpt_curDeskClass() {
      return this.db_deskClasses[this.curDeskClassIndex]
    }
  }
}
</script>

<style scoped>

</style>
