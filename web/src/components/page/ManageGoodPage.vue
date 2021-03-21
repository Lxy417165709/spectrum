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
      <el-divider content-position="left" v-if="desk !== undefined && desk.space!==undefined">{{ desk.space.name }} - 普通商品类</el-divider>
      <good-class-list :goodClasses="db_goodClasses" :props_isAdminView="props_isAdminView"
                       @turnToGoodListMode="turnToGoodListMode"></good-class-list>

      <!--      附属类-->
      <el-divider content-position="left" v-if="props_isAdminView">附属类</el-divider>
      <good-option-class-list :goodOptionClasses="db_goodOptionClasses" :props_isAdminView="props_isAdminView"
                              @turnToGoodOptionListMode="turnToGoodOptionListMode"
                              v-if="props_isAdminView"></good-option-class-list>
    </div>

    <!--    3. 商品类内的商品展示-->
    <div v-show="cpt_canGoodListShow">
      <good-list v-if="cpt_isGoodListExist" :desk="desk"
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
import test from "../../common/test/test";
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

    desk: Object,
  },
  async created() {
    this.db_goodOptionClasses = test.goodOptionClasses
    await utils.GetAllGoodClasses(this, {}, (res) => {
      this.db_goodClasses = res.data.data.goodClasses
    })
  },
  data() {
    return {
      viewMode: cst.VIEW_MODE.CLASS_LIST_MODE,

      GoodClassEditorVisible: false,
      GoodOptionEditorOfAdminVisible: false,

      curGoodClassIndex: cst.INDEX.INVALID_INDEX,
      curGoodOptionClassIndex: cst.INDEX.INVALID_INDEX,

      db_goodClasses: [],
      db_goodOptionClasses: [],
    }
  },
  methods: {
    turnToClassListMode(deskIndex) {
      this.curDeskIndex = deskIndex
      this.viewMode = cst.VIEW_MODE.CLASS_LIST_MODE
    },
    async turnToGoodOptionListMode(goodOptionClassIndex) {
      await utils.GetAllGoodOptions(this, {
        className: this.db_goodOptionClasses[goodOptionClassIndex].name,
      }, (res) => {
        this.viewMode = cst.VIEW_MODE.GOOD_OPTION_LIST_MODE
        this.curGoodOptionClassIndex = goodOptionClassIndex
        this.$nextTick(() => {
          this.$refs.GoodOptionList.goodOptions = res.data.data.elements
        })
      })
    },
    async turnToGoodListMode(goodClassIndex) {
      await utils.GetAllGoods(this, {
        className: this.db_goodClasses[goodClassIndex].name,
      }, (res) => {
        this.viewMode = cst.VIEW_MODE.GOOD_LIST_MODE
        this.curGoodClassIndex = goodClassIndex
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
      return this.db_goodClasses[this.curGoodClassIndex]
    },
    cpt_curGoodOptionClass() {
      return this.db_goodOptionClasses[this.curGoodOptionClassIndex]
    }
  },

}
</script>

<style scoped>

</style>
