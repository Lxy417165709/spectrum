<!-- eslint-disable -->
<template>
  <el-tabs v-model="curSizeInfoName" type="border-card" @tab-click="" editable @edit="handleTabsEdit"
           @tab-add="handleClick">
    <el-tab-pane v-for="(sizeInfo,index) in sizeInfos" :label="sizeInfo.name" :name="sizeInfo.name" :key="index">
      <el-form>
        <el-form-item label="照片">
          <el-upload
            action="/api/upload"
            list-type="picture-card">
            <i class="el-icon-plus"></i>
          </el-upload>
        </el-form-item>
        <el-form-item label="价格">
          <el-input v-model="sizeInfo.price" style="width: 70%"></el-input>
        </el-form-item>
      </el-form>
    </el-tab-pane>

    <!--    <el-tab-pane label="+" name="add" closable></el-tab-pane>-->
  </el-tabs>
</template>

<script>
/* eslint-disable */
import utils from "../../common/utils";

export default {
  name: "GoodSizeEditor",
  data() {
    return {

      curSizeInfoName: "小规格",
      addTabCount: 0,
      sizeInfos: [
        {
          name: "小规格",
          price: 18.0,
        },
        {
          name: "中规格",
          price: 20.0,
        },
        {
          name: "大规格",
          price: 25.0,
        }
      ]
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
