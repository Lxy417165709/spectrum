<!-- eslint-disable -->
<template>
  <div>

    <!--    1. 顶栏-->
    <el-row style="height: 8px" v-if="props_isAdminView"></el-row>
    <el-row style="height: 16px;">
      <el-button v-show="cpt_canBackButtonShow" style="position: absolute;top:-8px;left:10px;" size="mini"
                 @click="turnToParentComponentMode"
                 type="primary" round>
        退回
      </el-button>
    </el-row>

    <!--    2. 商品类展示-->
    <div v-show="cpt_canClassListShow">
      <!--      普通商品类-->
      <el-divider content-position="left">{{ deskSpaceName }} {{ deskSpaceName !== undefined ? "-" : "" }} 普通商品类
      </el-divider>
      <good-class-list :goodClasses="goodClasses" :props_isAdminView="props_isAdminView"
                       @turnToGoodListMode="turnToGoodListMode"
                       @successToAddGoodClass="successToAddGoodClass"></good-class-list>

      <!--      附属类-->
      <el-divider content-position="left" v-if="props_isAdminView">附属类</el-divider>
      <good-option-class-list :goodOptionClasses="goodOptionClasses" :props_isAdminView="props_isAdminView"
                              @turnToGoodOptionListMode="turnToGoodOptionListMode"
                              v-if="props_isAdminView"></good-option-class-list>
    </div>

    <!--    3. 商品类内的商品展示-->
    <div v-show="cpt_canGoodListShow">
      <good-list v-if="cpt_isGoodListExist" :orderID="orderID"
                 ref="GoodList"
                 :props_isAdminView="props_isAdminView"
                 :className="cpt_curGoodClass.name"
                 @turnToClassListMode="turnToClassListMode"
      ></good-list>
    </div>

    <!--    4. 附属选项类内的附属选项展示-->
    <div v-show="cpt_canGoodOptionListShow">
      <good-option-list v-if="cpt_isGoodOptionListExist"
                        ref="GoodOptionList"
                        @turnToGoodOptionListMode="turnToGoodOptionListMode"
                        :props_isAdminView="props_isAdminView"
                        :className="cpt_curGoodOptionClass.name">
      </good-option-list>
    </div>

  </div>
</template>

<script>
/* eslint-disable */
import GoodClassList from "../list/GoodClassList";
import GoodList from "../list/GoodList";
import GoodEditorOfAdmin from "../editor/GoodEditorOfAdmin";
import GoodEditorOfUser from "../editor/GoodEditorOfUser";
import GoodClassEditor from "../editor/GoodClassEditor";
import test from "../../common/test";
import utils from "../../common/utils";
import GoodOptionClassList from "../list/GoodOptionClassList";
import GoodOptionList from "../list/GoodOptionList";
import cst from "../../common/cst";
import GoodOptionEditorOfAdmin from "../editor/GoodOptionEditorOfAdmin";

export default {
  name: "ManageGoodPage",
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

    deskSpaceName: String,
    orderID: Number,
  },
  async created() {
    this.goodOptionClasses = test.goodOptionClasses
    await this.flashAllGoodClasses()
  },
  data() {
    return {
      viewMode: cst.VIEW_MODE.CLASS_LIST_MODE,

      GoodClassEditorVisible: false,
      GoodOptionEditorOfAdminVisible: false,

      curGoodClassIndex: cst.INDEX.INVALID_INDEX,
      curGoodOptionClassIndex: cst.INDEX.INVALID_INDEX,

      goodClasses: [],
      goodOptionClasses: [],
    }
  },
  methods: {
    flashAllGoodClasses(){
      utils.GetAllGoodClasses(this, {}, (res) => {
        this.goodClasses = res.data.data.goodClasses
      })
    },
    turnToClassListMode(deskIndex) {
      this.curDeskIndex = deskIndex
      this.viewMode = cst.VIEW_MODE.CLASS_LIST_MODE
    },
    async turnToGoodOptionListMode(goodOptionClassIndex) {
      this.viewMode = cst.VIEW_MODE.GOOD_OPTION_LIST_MODE
      this.curGoodOptionClassIndex = goodOptionClassIndex
      await this.$nextTick(() => {
        this.$refs.GoodOptionList.flashAllGoodOptions()
      })
    },
    async turnToGoodListMode(goodClassIndex) {
      this.viewMode = cst.VIEW_MODE.GOOD_LIST_MODE
      this.curGoodClassIndex = goodClassIndex
      await this.$nextTick(() => {
        this.$refs.GoodList.flashAllGoods()
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
    successToAddGoodClass() {
      this.flashAllGoodClasses()
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
    },
    cpt_curGoodClass() {
      return this.goodClasses[this.curGoodClassIndex]
    },
    cpt_curGoodOptionClass() {
      return this.goodOptionClasses[this.curGoodOptionClassIndex]
    }
  },

}
</script>

<style scoped>

</style>
