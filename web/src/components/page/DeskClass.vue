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
        <el-divider content-position="left" v-if="cpt_canDeskListShow">{{
            db_deskClasses[curDeskClassIndex].name
          }}
        </el-divider>
        <desk-list v-if="cpt_canDeskListShow"
                   ref="DeskList"
                   :className="db_deskClasses[curDeskClassIndex].name"
                   @openDeskEditorOfAdmin="openDeskEditorOfAdmin"
                   @turnToClassListMode="turnToClassListMode"
                   @attachOrderID="attachOrderID"></desk-list>
      </div>
      <div v-show="cpt_canClassListShow">
        <good-class :props_isAdminView="false" :props_haveParentComponent="true" ref="ref_goodClass"
                    @turnToParentComponentMode="turnToDeskListMode"></good-class>
      </div>
    </el-main>

    <!--    3. 桌子添加、编辑-->
    <el-dialog
      title="桌类添加/编辑"
      :visible.sync="deskEditorOfAdminVisible"
      width="30%">
      <desk-editor-of-admin ref="DeskEditorOfAdmin"></desk-editor-of-admin>
    </el-dialog>

    <!--    4. 桌类添加、编辑-->
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
import utils from "../../common/utils";
import DeskEditorOfAdmin from "../editor/DeskEditorOfAdmin";

export default {
  name: 'DeskClass',
  components: {DeskEditorOfAdmin, DeskClassEditor, GoodClass, DeskList},
  async created() {
    await this.getAllDeskClasses()
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
    async getAllDeskClasses() {
      let model = utils.getRequestModel("mvp", "GetAllDeskClasses", {})
      await utils.sendRequestModel(model).then(res => {
        console.log("GetAllDeskClasses.res", res)
        if (!utils.hasRequestSuccess(res)) {
          this.$message.error(res.data.err)
          return
        }
        this.$message.success(res.data.msg)
        this.db_deskClasses = res.data.data.deskClasses
      })
    },
    async getAllDesks(className) {
      let model = utils.getRequestModel("mvp", "GetAllDesks", {
        className: className
      })
      await utils.sendRequestModel(model).then(res => {
        console.log("GetAllDesks.res", res)
        if (!utils.hasRequestSuccess(res)) {
          this.$message.error(res.data.err)
          return
        }
        this.$message.success(res.data.msg)

        this.$nextTick(() => {
          this.$refs.DeskList.desks = res.data.data.desks
        })
      })
    },

    handleDeskClassClick(deskClassIndex) {
      this.curDeskClassIndex = deskClassIndex

      this.curDeskIndex = cst.INDEX.INVALID_INDEX

      this.viewMode = cst.VIEW_MODE.DESK_LIST_MODE

      this.$nextTick(() => {
        this.$refs.ref_goodClass.viewMode = cst.VIEW_MODE.CLASS_LIST_MODE;
      })

      this.getAllDesks(this.db_deskClasses[this.curDeskClassIndex].name)
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
    attachOrderID(orderID) {
      this.viewMode = cst.VIEW_MODE.CLASS_LIST_MODE
      this.$nextTick(() => {
        this.$refs.ref_goodClass.orderID = orderID
        console.log("attachOrderID", this.$refs.ref_goodClass.orderID)
      })
    },


    turnToDeskListMode() {
      this.viewMode = cst.VIEW_MODE.DESK_LIST_MODE
      this.getAllDesks(this.db_deskClasses[this.curDeskClassIndex].name)
    },

    openDeskEditorOfAdmin(desk, className) {
      this.deskEditorOfAdminVisible = true
      this.$nextTick(() => {
        this.$refs.DeskEditorOfAdmin.desk = desk
        this.$refs.DeskEditorOfAdmin.className = className
      })
    }
  },
  computed: {
    cpt_canClassListShow() {
      return this.viewMode === cst.VIEW_MODE.CLASS_LIST_MODE
    },
    cpt_canDeskListShow() {
      return this.viewMode === cst.VIEW_MODE.DESK_LIST_MODE && this.curDeskClassIndex !== cst.INDEX.INVALID_INDEX
    }
  }
}
</script>

<style scoped>

</style>
