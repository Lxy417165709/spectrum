<!-- eslint-disable -->
<!-- todo: 这里要实现下，默认选中第一个 tab-->
<template>
  <el-tabs v-model="curSizeIndexName" type="border-card" @tab-click="" editable @edit="handleTabsEdit"
           @tab-add="handleClick">
    <el-tab-pane v-for="(sizeInfo,index) in sizeInfos" :label="sizeInfo.name" :name="sizeInfo.name" :key="index">
      <good-size-info-editor :name="sizeInfo.name" :price="sizeInfo.price"></good-size-info-editor>
    </el-tab-pane>
  </el-tabs>
</template>

<script>
/* eslint-disable */
import utils from "../../common/utils";
import GoodSizeInfoEditor from "./GoodSizeInfoEditor";

export default {
  name: "GoodSizeEditor",
  components: {GoodSizeInfoEditor},
  props: {
    curSizeIndex: Number,
    sizeInfos: Array
  },
  mounted() {
    // this.$nextTick(()=>{
    //   this.curSizeIndexName = this.sizeInfos[this.curSizeIndex].name
    // })

  },
  data() {
    return {
      addTabCount: 0,
      curSizeIndexName: ""
    }
  },
  methods: {
    handleClick(tab, event) {
      this.addTabCount++
      let name = "未设定规格" + this.addTabCount
      this.sizeInfos.push({
        name: name,
        price: 30
      })
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
