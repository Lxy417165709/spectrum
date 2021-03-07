<!-- eslint-disable -->
<template>
  <div>

    <!--    1. 顶栏-->
    <el-row style="height: 40px;margin-left:10px;margin-bottom: 10px">
      <el-button v-show="cpt_canBackButtonShow" style="height: 40px"
                 @click="turnToFatherMode"
                 type="primary">
        退回
      </el-button>
    </el-row>

    <!--    2. 商品类展示-->
    <div v-show="cpt_canClassListShow">
      <el-divider content-position="left">普通商品类</el-divider>
      <good-class-list :goodClasses="goodClasses" :isAdminView="isAdminView"
                       @turnToGoodListMode="turnToGoodListMode"
                       @openGoodClassEditor="openGoodClassEditor"></good-class-list>

      <el-divider content-position="left" v-if="isAdminView">附属类</el-divider>
      <good-option-class-list :goodOptionClasses="goodOptionClasses" :isAdminView="isAdminView"
                              @turnToGoodOptionListMode="turnToGoodOptionListMode"
                              v-if="isAdminView"></good-option-class-list>
    </div>

    <!--    3. 商品类内的商品展示-->
    <div v-show="cpt_canGoodListShow">
      <el-divider content-position="left">元素</el-divider>
      <good-list v-if="curGoodClassIndex!==-1"
                 :isAdminView="isAdminView"
                 :goods="goodClasses[curGoodClassIndex].goods"
                 :className="goodClasses[curGoodClassIndex].name"
                 @turnToClassListMode="turnToClassListMode"
                 @openGoodEditorOfAdmin="openGoodEditorOfAdmin"
                 @openGoodEditorOfUser="openGoodEditorOfUser"></good-list>
    </div>

    <!--    4. 附属选项类内的附属选项展示-->
    <div v-show="cpt_canGoodOptionListShow">
      <el-divider content-position="left">元素</el-divider>
      <good-option-list v-if="curGoodOptionClassIndex!==-1"
                        :isAdminView="isAdminView"
                        :goodOptions="goodOptionClasses[curGoodOptionClassIndex].goodOptions"
                        :className="goodOptionClasses[curGoodOptionClassIndex].name"></good-option-list>
    </div>

    <!--    4. 商品添加、编辑框-->
    <el-dialog
      title="商品添加/编辑"
      :visible.sync="GoodEditorOfAdminVisible"
      width="30%">
      <good-editor-of-admin ref="GoodEditorOfAdmin"></good-editor-of-admin>
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

export default {
  name: "GoodClass",
  components: {
    GoodOptionClassList,
    GoodClassEditor,
    GoodEditorOfAdmin,
    GoodList,
    GoodClassList,
    GoodEditorOfUser,
    GoodOptionList
  },
  props: {
    isAdminView: Boolean,
    hasFather: Boolean,
  },
  created() {
    console.log("created", cst.VIEW_MODE)
  },
  async mounted() {
    this.goodOptionClasses = test.goodOptionClasses
    console.log("mounted", this.goodOptionClasses)

    await this.getAllGoodClasses()
  },
  data() {
    return {
      viewMode: cst.VIEW_MODE.CLASS_LIST_MODE,
      GoodEditorOfAdminVisible: false,
      GoodClassEditorVisible: false,
      GoodEditorOfUserVisible: false,
      curGoodClassIndex: -1,
      curGoodOptionClassIndex: -1,

      // 数据库读取属性
      goodClasses: [],
      goodOptionClasses: [],
    }
  },
  methods: {
    async getAllGoodClasses() {
      let model = utils.getRequestModel("mvp", "GetAllGoodClasses", {})
      await utils.sendRequestModel(model).then(res => {
        if (!utils.hasRequestSuccess(res)) {
          this.$message.error(res.data.err)
          return
        }
        this.goodClasses = res.data.data.goodClasses
        this.$message.success(res.data.msg)
        console.log("getAllGoodClasses-this.goodClasses", this.goodClasses)
      })
    },
    turnToClassListMode() {
      this.viewMode = cst.VIEW_MODE.CLASS_LIST_MODE
    },
    turnToGoodOptionListMode(goodOptionClassIndex) {
      this.viewMode = cst.VIEW_MODE.GOOD_OPTION_LIST_MODE
      this.curGoodOptionClassIndex = goodOptionClassIndex
    },
    turnToGoodListMode(goodClassIndex) {
      this.viewMode = cst.VIEW_MODE.GOOD_LIST_MODE
      this.curGoodClassIndex = goodClassIndex
    },
    turnToFatherMode() {
      if (this.viewMode === cst.VIEW_MODE.CLASS_LIST_MODE) {
        this.$emit("turnToFatherMode")
      }
      if (this.viewMode === cst.VIEW_MODE.GOOD_OPTION_LIST_MODE || this.viewMode === cst.VIEW_MODE.GOOD_LIST_MODE) {
        this.viewMode = cst.VIEW_MODE.CLASS_LIST_MODE
      }
    },
    openGoodEditorOfAdmin(good, className) {
      this.GoodEditorOfAdminVisible = true
      this.$nextTick(() => {
        this.$refs.GoodEditorOfAdmin.good = good
        this.$refs.GoodEditorOfAdmin.className = className
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
      return this.viewMode === cst.VIEW_MODE.CLASS_LIST_MODE && this.hasFather ||
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
  },

}
</script>

<style scoped>

</style>
