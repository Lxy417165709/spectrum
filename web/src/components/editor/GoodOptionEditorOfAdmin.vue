<!-- eslint-disable -->
<template>
  <el-form ref="form" label-width="80px">

    <!-- 1. 选项名编辑器 -->
    <el-form-item label="选项名">
      <el-input style="width: 70%" v-if="option!==undefined"
                v-model="option.name"></el-input>
<!--      <el-input style="width: 70%" v-if="option!==undefined && canModifyName"-->
<!--                v-model="option.name"></el-input>-->
<!--      <span v-if="option!==undefined && !canModifyName">{{ option.name }}</span>-->
    </el-form-item>

    <!-- 2. 规格编辑器 -->
    <el-tabs type="border-card" @tab-click="tabClick" addable v-model="curSizeInfoIndexString"
             @tab-add="handleClick" style="margin-bottom: 10px">
      <el-tab-pane v-for="(sizeInfo,index) in option.sizeInfos" :label="sizeInfo.size"
                   :name="index.toString()"
                   :key="index">

        <el-form-item label="规格名">
                    <el-input style="width: 70%" v-if="option.selectedIndex !== curSizeInfoIndexInt"
                              v-model="sizeInfo.size"></el-input>
                    <span v-if="option.selectedIndex === curSizeInfoIndexInt">{{ sizeInfo.size }}</span>
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
          <el-form-item label="默认选中" v-if="option.selectedIndex !== index && sizeInfo.id !== 0">
            <el-button @click="handleChangeDefaultSizeInfo(index)" type="primary">选中</el-button>
          </el-form-item>
          <el-form-item label="是否删除"
                        v-if="option.selectedIndex !== index && option.sizeInfos.length>1">
            <el-button @click="deleteElementSizeInfo(index)" type="danger">删除</el-button>
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
      curSizeInfoIndexString: '0',

      canModifyName: false,
    }
  },
  methods: {
    handleClick(tab, event) {
      this.option.sizeInfos.push(utils.deepCopy(test.blankSizeInfo))
    },
    handleTabsEdit(name, event) {
      this.option.sizeInfos = utils.removeElementByField(this.option.sizeInfos, "size", name)
    },
    handleChangeDefaultSizeInfo(index) {
      this.option.selectedIndex = index
    },
    deleteElementSizeInfo(index) {
      if (this.option.sizeInfos[this.curSizeInfoIndexInt].id === 0) {
        this.option.sizeInfos = utils.removeIndex(this.option.sizeInfos, index)
        if (this.curSizeInfoIndexInt >= this.option.sizeInfos.length) {
          this.curSizeInfoIndexString = (this.option.sizeInfos.length - 1).toString()
        }
        return
      }

      utils.DeleteElementSizeInfo(this, {
        elementName: this.option.name,
        sizeInfoSize: this.option.sizeInfos[index].size,
      }, (res) => {
        this.option.sizeInfos = utils.removeIndex(this.option.sizeInfos, index)
        // if (this.option.selectedIndex > this.curSizeInfoIndexInt) {
        //   this.option.selectedIndex--
        // }  // todo: 这个 默认选中应该要调整下，防止删除后默认选中也跟着改变
        if (this.curSizeInfoIndexInt >= this.option.sizeInfos.length) {
          this.curSizeInfoIndexString = (this.option.sizeInfos.length - 1).toString()
        }
      })
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
      // this.curSizeInfoIndex = tab.index
      // console.log(this.curSizeInfoIndex)
    },
    cleanSizeInfoPictureStorePath() {
      this.option.sizeInfos[this.curSizeInfoIndexInt].pictureStorePath = ""
    },
    imageUploadSuccess(res, file, fileList) {
      this.option.sizeInfos[this.curSizeInfoIndexInt].pictureStorePath = res.data.fileStorePath;
    },
  },
  computed: {
    curSizeInfoIndexInt() {
      return this.curSizeInfoIndexString - 0
    }
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
