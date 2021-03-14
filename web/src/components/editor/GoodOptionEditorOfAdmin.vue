<!-- eslint-disable -->
<template>
  <el-form ref="form" label-width="80px">

    <!-- 1. 选项名编辑器 -->
    <el-form-item label="选项名">
      <el-input style="width: 70%" v-model="option.name"></el-input>
    </el-form-item>

    <!-- 2. 规格编辑器 -->
    <el-tabs type="border-card" @tab-click="" editable @edit="handleTabsEdit"
             @tab-add="handleClick" style="margin-bottom: 10px">
      <el-tab-pane v-for="(sizeInfo,index) in option.sizeInfos" :label="sizeInfo.size"
                   :name="index.toString()"
                   :key="index">
        <el-form label-width="80px">
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
          <el-form-item label="默认选中" v-if="sizeInfo.isSelected===false">
            <el-button @click="handleChangeDefaultSizeInfo(index)">确定</el-button>
          </el-form-item>
        </el-form>
      </el-tab-pane>
    </el-tabs>

    <!--    5. 提交按钮-->
    <el-form-item>
      <el-button type="primary" @click="addGoodOption(option)">确定</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
/* eslint-disable */
import test from "../../common/test/test";
import utils from "../../common/utils";
import global from "../../common/global_object/global";

export default {
  name: "GoodOptionEditorOfAdmin",
  components: {},
  mounted() {
    this.selectableElement = test.selectableElement
  },
  data() {
    return {
      option: {},
      className: "",

      addTabCount: 0,
    }
  },
  methods: {
    handleClick(tab, event) {
      this.addTabCount++
      let name = "未设定规格" + this.addTabCount
      this.option.sizeInfos.push(utils.NewBlankSizeInfo(name))
    },
    handleTabsEdit(name, event) {
      this.option.sizeInfos = utils.removeElementByField(this.option.sizeInfos, "size", name)
    },
    handleChangeDefaultSizeInfo(index) {
      for (let i = 0; i < this.option.sizeInfos.length; i++) {
        this.option.sizeInfos[i].isSelected = false
      }
      this.option.sizeInfos[index].isSelected = true
    },
    async addGoodOption(option) {
      let model = utils.getRequestModel("mvp", "AddElement", {
        element: option,
        className: this.className,
      })
      await utils.sendRequestModel(model).then(res => {
        console.log("addGoodOption.res", res)
        if (!utils.hasRequestSuccess(res)) {
          this.$message.error(res.data.err)
          return
        }
        this.$message.success(res.data.msg)
      })
    }
  }
}
</script>

<style scoped>

</style>
