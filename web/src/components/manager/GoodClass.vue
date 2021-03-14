<!-- eslint-disable -->
<template>
  <div>

    <!--    1. 顶栏-->
    <el-row style="height: 40px;margin-left:10px;margin-bottom: 10px">
      <el-button v-show="cpt_canBackButtonShow" style="height: 40px"
                 @click="turnToParentComponentMode"
                 type="primary">
        退回
      </el-button>
    </el-row>

    <!--    2. 商品类展示-->
    <div v-show="cpt_canClassListShow">
      <!--      普通商品类-->
      <el-divider content-position="left">普通商品类</el-divider>
      <good-class-list :goodClasses="db_goodClasses" :props_isAdminView="props_isAdminView"
                       @turnToGoodListMode="turnToGoodListMode"
                       @openGoodClassEditor="openGoodClassEditor"></good-class-list>

      <!--      附属类-->
      <el-divider content-position="left" v-if="props_isAdminView">附属类</el-divider>
      <good-option-class-list :goodOptionClasses="db_goodOptionClasses" :props_isAdminView="props_isAdminView"
                              @turnToGoodOptionListMode="turnToGoodOptionListMode"
                              v-if="props_isAdminView"></good-option-class-list>
    </div>

    <!--    3. 商品类内的商品展示-->
    <div v-show="cpt_canGoodListShow">
      <el-divider content-position="left">元素</el-divider>
      <good-list v-if="cpt_isGoodListExist"
                 ref="GoodList"

                 :props_isAdminView="props_isAdminView"
                 :className="db_goodClasses[curGoodClassIndex].name"
                 @turnToClassListMode="turnToClassListMode"
                 @openGoodEditorOfAdmin="openGoodEditorOfAdmin"
                 @openGoodEditorOfUser="openGoodEditorOfUser"></good-list>
    </div>

    <!--    4. 附属选项类内的附属选项展示-->
    <div v-show="cpt_canGoodOptionListShow">
      <el-divider content-position="left">元素</el-divider>
      <good-option-list v-if="cpt_isGoodOptionListExist"
                        ref="GoodOptionList"
                        :props_isAdminView="props_isAdminView"
                        @openGoodOptionEditorOfAdmin="openGoodOptionEditorOfAdmin"
      ></good-option-list>
    </div>


    <!--    4. 商品添加、编辑框-->
    <el-dialog
      title="商品添加/编辑"
      :visible.sync="GoodEditorOfAdminVisible"
      width="30%">
      <good-editor-of-admin ref="GoodEditorOfAdmin"></good-editor-of-admin>
    </el-dialog>

    <!--    4. 商品选项添加、编辑框-->
    <el-dialog
      title="商品选项添加/编辑"
      :visible.sync="GoodOptionEditorOfAdminVisible"
      width="30%">
      <good-option-editor-of-admin ref="GoodOptionEditorOfAdmin"></good-option-editor-of-admin>
    </el-dialog>

    <!--    5. 商品下单框-->
    <el-dialog
      title="商品点单"
      :visible.sync="GoodEditorOfUserVisible"
      width="30%">
      <good-editor-of-user ref="GoodEditorOfUser"></good-editor-of-user>
    </el-dialog>

    <!--    6. 商品类添加、编辑框-->
    <el-dialog
      title="商品类添加/编辑"
      :visible.sync="GoodClassEditorVisible"
      width="30%">
      <good-class-editor ref="GoodClassEditor"></good-class-editor>
    </el-dialog>
  </div>
</template>

<script>
/* eslint-disable */
import GoodClassList from "../list/GoodClassList";
import GoodList from "../list/GoodList";
import GoodEditorOfAdmin from "../editor/GoodEditorOfAdmin";
import GoodEditorOfUser from "../editor/GoodEditorOfUser";
import GoodClassEditor from "../editor/GoodClassEditor";
import test from "../../common/test/test";
import utils from "../../common/utils";
import GoodOptionClassList from "../list/GoodOptionClassList";
import GoodOptionList from "../list/GoodOptionList";
import cst from "../../common/cst";
import GoodOptionEditorOfAdmin from "../editor/GoodOptionEditorOfAdmin";

export default {
  name: "GoodClass",
  components: {
    GoodOptionEditorOfAdmin,
    GoodOptionClassList,
    GoodClassEditor,
    GoodEditorOfAdmin,
    GoodList,
    GoodClassList,
    GoodEditorOfUser,
    GoodOptionList
  },
  props: {
    props_isAdminView: Boolean,
    props_haveParentComponent: Boolean,
  },
  async mounted() {
    this.db_goodOptionClasses = test.goodOptionClasses
    await this.getAllGoodClasses()
  },
  data() {
    return {
      viewMode: cst.VIEW_MODE.CLASS_LIST_MODE,

      GoodEditorOfAdminVisible: false,
      GoodClassEditorVisible: false,
      GoodEditorOfUserVisible: false,
      GoodOptionEditorOfAdminVisible: false,

      curGoodClassIndex: cst.INDEX.INVALID_INDEX,
      curGoodOptionClassIndex: cst.INDEX.INVALID_INDEX,
      curDeskIndex: cst.INDEX.INVALID_INDEX,

      db_goodClasses: [],
      db_goodOptionClasses: [],
    }
  },
  methods: {
    async getAllGoodClasses() {
      let model = utils.getRequestModel("mvp", "GetAllGoodClasses", {})
      await utils.sendRequestModel(model).then(res => {
        if (!utils.hasRequestSuccess(res)) {
          console.log("getAllGoodClasses.res", res)
          this.$message.error(res.data.err)
          return
        }
        this.db_goodClasses = res.data.data.goodClasses
        this.$message.success(res.data.msg)
      })
    },
    turnToClassListMode(deskIndex) {
      this.viewMode = cst.VIEW_MODE.CLASS_LIST_MODE
      this.curDeskIndex = deskIndex
    },
    async turnToGoodOptionListMode(goodOptionClassIndex) {
      let model = utils.getRequestModel("mvp", "GetAllGoodOptions", {
        className: this.db_goodOptionClasses[goodOptionClassIndex].name,
      })
      await utils.sendRequestModel(model).then(res => {
        console.log("GetAllGoodOptions.res", res)
        if (!utils.hasRequestSuccess(res)) {
          this.$message.error(res.data.err)
          return
        }
        this.$message.success(res.data.msg)

        this.viewMode = cst.VIEW_MODE.GOOD_OPTION_LIST_MODE
        this.curGoodOptionClassIndex = goodOptionClassIndex

        this.$nextTick(() => {
          this.$refs.GoodOptionList.goodOptions = res.data.data.elements
          this.$refs.GoodOptionList.className = this.db_goodOptionClasses[this.curGoodOptionClassIndex].name
        })
      })
    },
    async turnToGoodListMode(goodClassIndex) {
      let model = utils.getRequestModel("mvp", "GetAllGoods", {
        className: this.db_goodClasses[goodClassIndex].name,
      })
      await utils.sendRequestModel(model).then(res => {
        console.log("GetAllGoods.res", res)
        if (!utils.hasRequestSuccess(res)) {
          this.$message.error(res.data.err)
          return
        }
        this.$message.success(res.data.msg)

        this.curGoodClassIndex = goodClassIndex
        this.viewMode = cst.VIEW_MODE.GOOD_LIST_MODE
        this.$nextTick(() => {
          this.$refs.GoodList.goods = res.data.data.goods
        })
      })
    },
    turnToParentComponentMode() {
      if (this.viewMode === cst.VIEW_MODE.CLASS_LIST_MODE) {
        this.$emit("turnToParentComponentMode")
      }
      if (this.viewMode === cst.VIEW_MODE.GOOD_OPTION_LIST_MODE || this.viewMode === cst.VIEW_MODE.GOOD_LIST_MODE) {
        this.viewMode = cst.VIEW_MODE.CLASS_LIST_MODE
      }
    },
    openGoodEditorOfAdmin(good, className) {
      this.GoodEditorOfAdminVisible = true
      this.$nextTick(() => {
        console.log("openGoodEditorOfAdmin", good)
        this.$refs.GoodEditorOfAdmin.good = good
        this.$refs.GoodEditorOfAdmin.className = className
        console.log("openGoodEditorOfAdmin,this.$refs.GoodEditorOfAdmin", this.$refs.GoodEditorOfAdmin.good)
      })
    },
    openGoodOptionEditorOfAdmin(option, className) {
      this.GoodOptionEditorOfAdminVisible = true
      this.$nextTick(() => {
        this.$refs.GoodOptionEditorOfAdmin.option = option
        this.$refs.GoodOptionEditorOfAdmin.className = className
      })
    },
    openGoodEditorOfUser(good) {
      this.GoodEditorOfUserVisible = true
      this.$nextTick(() => {
        this.$refs.GoodEditorOfUser.good = good
      })
    },
    openGoodClassEditor(goodClass) {
      this.GoodClassEditorVisible = true
      this.$nextTick(() => {
        this.$refs.GoodClassEditor.goodClass = goodClass
      })
    },
  },
  computed: {
    cpt_canBackButtonShow() {
      return this.viewMode === cst.VIEW_MODE.CLASS_LIST_MODE && this.props_haveParentComponent ||
        this.viewMode === cst.VIEW_MODE.GOOD_LIST_MODE ||
        this.viewMode === cst.VIEW_MODE.GOOD_OPTION_LIST_MODE
    },
    cpt_canGoodListShow() {
      return this.viewMode === cst.VIEW_MODE.GOOD_LIST_MODE
    },
    cpt_canGoodOptionListShow() {
      return this.viewMode === cst.VIEW_MODE.GOOD_OPTION_LIST_MODE
    },
    cpt_canClassListShow() {
      return this.viewMode === cst.VIEW_MODE.CLASS_LIST_MODE
    },
    cpt_isGoodListExist() {
      return this.curGoodClassIndex !== cst.INDEX.INVALID_INDEX
    },
    cpt_isGoodOptionListExist() {
      return this.curGoodOptionClassIndex !== cst.INDEX.INVALID_INDEX
    }
  },

}
</script>

<style scoped>

</style>
