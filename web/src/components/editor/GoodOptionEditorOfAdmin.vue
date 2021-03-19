<!-- eslint-disable -->
<template>
  <el-form ref="form" label-width="80px">

    <!-- 1. 选项名编辑器 -->
    <el-form-item label="选项名">
      <el-input style="width: 70%" v-if="option!==undefined && canModifyName"
                v-model="option.name"></el-input>
      <span v-if="option!==undefined && !canModifyName">{{ option.name }}</span>
    </el-form-item>

    <!-- 2. 规格编辑器 -->
    <el-tabs type="border-card" @tab-click="tabClick" editable @edit="handleTabsEdit"
             @tab-add="handleClick" style="margin-bottom: 10px">
      <el-tab-pane v-for="(sizeInfo,index) in option.sizeInfos" :label="sizeInfo.size"
                   :name="index.toString()"
                   :key="index">

        <el-form-item label="规格名">
          <el-input style="width: 70%" v-if="sizeInfo.size!==undefined && canModifyName"
                    v-model="sizeInfo.size"></el-input>
          <span v-if="sizeInfo.size!==undefined && !canModifyName">{{ sizeInfo.size }}</span>
        </el-form-item>


        <el-form label-width="80px">
          <el-form-item label="照片">
            <el-upload
              v-if="sizeInfo.pictureStorePath===''"
              action="/api/upload"
              :on-success="imageUploadSuccess"
              list-type="picture-card">

              <i class="el-icon-plus"></i>
            </el-upload>
            <el-image
              v-if="sizeInfo.pictureStorePath!==''"
              @dblclick.native="cleanSizeInfoPictureStorePath"
              :src="'api/file/' + sizeInfo.pictureStorePath"
              style="width: 148px; height: 148px;"></el-image>
          </el-form-item>


          <el-form-item label="价格">
            <el-input v-model="sizeInfo.price" style="width: 70%"></el-input>
          </el-form-item>
          <el-form-item label="默认选中" v-if="option.selectedIndex !== index">
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

export default {
  name: "GoodOptionEditorOfAdmin",
  components: {},
  props: {
    className: String,
  },
  mounted() {
    this.selectableElement = test.selectableElement
  },
  data() {
    return {
      option: {},

      addTabCount: 0,
      curSizeInfoIndex: 0,

      canModifyName: false,
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
      this.option.selectedIndex = index
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
    },
    tabClick(tab) {
      this.curSizeInfoIndex = tab.index
    },
    cleanSizeInfoPictureStorePath() {
      this.option.sizeInfos[this.curSizeInfoIndex].pictureStorePath = ""
    },
    imageUploadSuccess(res, file, fileList) {
      this.option.sizeInfos[this.curSizeInfoIndex].pictureStorePath = res.data.fileStorePath;
    },
  }
}
</script>

<style>
.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}

.avatar-uploader .el-upload:hover {
  border-color: #409EFF;
}

.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  line-height: 178px;
  text-align: center;
}

.avatar {
  width: 178px;
  height: 178px;
  display: block;
}
</style>
