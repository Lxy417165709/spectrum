<!-- eslint-disable -->
<template>
  <el-tabs v-model="curSizeInfoName" type="border-card" @tab-click="" editable @edit="handleTabsEdit"
           @tab-add="handleClick">
    <el-tab-pane v-for="(sizeInfo,index) in sizeInfos" :label="sizeInfo.size" :name="sizeInfo.size" :key="index">
      <desk-size-info-editor :price="sizeInfo.price"></desk-size-info-editor>
    </el-tab-pane>
  </el-tabs>
</template>

<script>
/* eslint-disable */
import DeskSizeInfoEditor from "./DeskSizeInfoEditor";
import utils from "../../common/utils";
// todo: 这里变量名、变量内容需要改变下
export default {
  name: "DeskInfoEditor",
  components: {DeskSizeInfoEditor},
  data() {
    return {
      curSizeInfoName: "小规格",
      addTabCount: 0,
      sizeInfos: [
        {
          size: "小规格",
          price: "18.0",
        },
        {
          size: "中规格",
          price: "20.0",
        },
        {
          size: "大规格",
          price: "25.0",
        }
      ]
    }
  },
  methods: {
    handleClick(tab, event) {
      this.addTabCount++
      let name = "未设定规格" + this.addTabCount
      this.sizeInfos.push(utils.NewBlankSizeInfo(name))
      this.curSizeInfoName = name
    },
    handleTabsEdit(name, event) {
      // todo: 这里要处理，被关闭的tab是当前展示的tab的情况
      if (name === this.curSizeInfoName) {
        return
      }
      this.sizeInfos = utils.removeElementByField(this.sizeInfos, "name", name)
    }
  }
}
</script>

<style scoped>

</style>
