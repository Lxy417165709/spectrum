<!-- eslint-disable -->
<template>
  <el-container style="height: 800px; border: 1px solid #eee">
    <!--    1. 桌类选择-->
    <el-aside width="200px">
      <el-menu>
        <el-menu-item v-for="(deskSet,deskSetIndex) in deskSets" :key="deskSetIndex"
                      @click="handleDeskSetClick(deskSetIndex)">
          <template slot="title"><i class="el-icon-message"></i><span>{{ deskSet.name }}</span></template>
        </el-menu-item>
      </el-menu>
      <el-button style="margin-left: 20px;margin-top: 10px" type="primary" @click="handleDeskButtonClick">添加桌位
      </el-button>
    </el-aside>

    <!--    2. 桌类详情-->
    <el-main>
      <div v-show="viewMode === 0">
        <el-row style="height: 40px;margin-left:10px;margin-bottom: 10px">
        </el-row>
        <el-divider content-position="left">{{ deskSets[curDeskSetIndex].name }}</el-divider>
        <desk-list :desks="deskSets[curDeskSetIndex].desks"
                   @turnToGoodClassListMode="turnToGoodClassListMode"></desk-list>
      </div>
      <div v-show="viewMode !== 0">
        <good-class :isEditMode="false" :hasFather="true" ref="goodClass"
                    @turnToFatherMode="turnToDeskListMode"></good-class>
      </div>
    </el-main>
    <el-dialog
      title="桌类添加/编辑"
      :visible.sync="deskSetEditorVisible"
      width="30%">
      <desk-set-editor></desk-set-editor>
    </el-dialog>
  </el-container>
</template>
<script>
/* eslint-disable */
import DeskList from "../list/DeskList";
import GoodClass from "../manager/GoodClass";
import DeskSetEditor from "../editor/DeskSetEditor";
import test from "../../common/test/test";

export default {
  name: 'DeskSet',
  components: {DeskSetEditor, GoodClass, DeskList},
  created() {
    this.deskSets = test.deskSets
  },
  data() {
    return {
      viewMode: 0, // todo: 这里可以弄个枚举
      deskSetEditorVisible: false,
      curDeskSetIndex: 0,
      curDeskIndex: -1,

      // 数据库读取属性
      deskSets: [],
    }
  },
  methods: {
    handleDeskSetClick(index) {
      this.curDeskSetIndex = index
      this.viewMode = 0
      this.curDeskIndex = -1
      this.$refs.goodClass.viewMode = 1; // 父传子
    },
    turnToGoodClassListMode(deskIndex) {
      this.viewMode = 1
      this.curDeskIndex = deskIndex
    },
    handleDeskButtonClick() {
      this.deskSetEditorVisible = true
    },
    turnToDeskListMode() {
      this.viewMode = 0
    }
  }
}
</script>

<style scoped>

</style>
