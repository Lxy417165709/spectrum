<!-- eslint-disable -->
<template>
  <el-container style="height: 800px; border: 1px solid #eee">
    <!--    1. 桌类选择-->
    <el-aside width="200px">
      <el-menu>
        <el-menu-item v-for="(deskClass,deskClassIndex) in db_deskClasses" :key="deskClassIndex"
                      @click="handleDeskClassClick(deskClassIndex)">
          <template slot="title"><i class="el-icon-message"></i><span>{{ deskClass.name }}</span></template>
        </el-menu-item>
      </el-menu>
      <el-button style="margin-left: 20px;margin-top: 10px" type="primary" @click="handleDeskButtonClick">添加桌位
      </el-button>
    </el-aside>

    <!--    2. 桌类详情-->
    <el-main>
      <div v-show="cpt_canDeskListShow">
        <el-row style="height: 40px;margin-left:10px;margin-bottom: 10px">
        </el-row>
        <el-divider content-position="left">{{ db_deskClasses[curDeskClassIndex].name }}</el-divider>
        <desk-list :desks="db_deskClasses[curDeskClassIndex].desks"
                   @turnToClassListMode="turnToClassListMode"></desk-list>
      </div>
      <div v-show="cpt_canClassListShow">
        <good-class :props_isAdminView="false" :props_haveParentComponent="true" ref="ref_goodClass"
                    @turnToParentComponentMode="turnToDeskListMode"></good-class>
      </div>
    </el-main>

    <!--    3. 桌类添加、编辑-->
    <el-dialog
      title="桌类添加/编辑"
      :visible.sync="deskClassEditorVisible"
      width="30%">
      <desk-class-editor ref="DeskClassEditor"></desk-class-editor>
    </el-dialog>
  </el-container>
</template>
<script>
/* eslint-disable */
import DeskList from "../list/DeskList";
import GoodClass from "../manager/GoodClass";
import DeskClassEditor from "../editor/DeskClassEditor";
import test from "../../common/test/test";
import cst from "../../common/cst";

export default {
  name: 'DeskClass',
  components: {DeskClassEditor, GoodClass, DeskList},
  created() {
    this.db_deskClasses = test.deskClasses
  },
  data() {
    return {
      viewMode: cst.VIEW_MODE.DESK_LIST_MODE,

      deskClassEditorVisible: false,

      curDeskClassIndex: cst.INDEX.FIRST_INDEX,
      curDeskIndex: cst.INDEX.INVALID_INDEX,

      db_deskClasses: [],
    }
  },
  methods: {
    handleDeskClassClick(deskClassIndex) {
      this.curDeskClassIndex = deskClassIndex

      this.viewMode = cst.VIEW_MODE.DESK_LIST_MODE
      this.curDeskIndex = cst.INDEX.INVALID_INDEX
      this.$refs.ref_goodClass.viewMode = cst.VIEW_MODE.CLASS_LIST_MODE;
    },
    turnToClassListMode(deskIndex) {
      this.viewMode = cst.VIEW_MODE.CLASS_LIST_MODE
      this.curDeskIndex = deskIndex
    },
    handleDeskButtonClick() {
      this.deskClassEditorVisible = true
      this.$nextTick(() => {
        this.$refs.DeskClassEditor.deskClass = test.blankDeskClass
      })
    },
    turnToDeskListMode() {
      this.viewMode = cst.VIEW_MODE.DESK_LIST_MODE
    }
  },
  computed: {
    cpt_canClassListShow() {
      return this.viewMode === cst.VIEW_MODE.CLASS_LIST_MODE
    },
    cpt_canDeskListShow() {
      return this.viewMode === cst.VIEW_MODE.DESK_LIST_MODE
    }
  }
}
</script>

<style scoped>

</style>
