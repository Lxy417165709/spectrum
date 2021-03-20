<!-- eslint-disable -->
<template>
  <el-form ref="form" label-width="80px">

    <!-- 1. 商品名编辑器 -->
    <el-form-item label="桌名">
      <el-input style="width: 70%" v-if="desk.space!==undefined" v-model="desk.space.name"></el-input>
    </el-form-item>

    <el-form-item label="照片">
      <el-upload
        v-if="desk.space!==undefined && desk.space.pictureStorePath===''"
        action="/api/upload"
        :on-success="imageUploadSuccess"
        list-type="picture-card">

        <i class="el-icon-plus"></i>
      </el-upload>
      <el-image
        v-if="desk.space!==undefined && desk.space.pictureStorePath!==''"
        @dblclick.native="cleanSizeInfoPictureStorePath"
        :src="'api/file/' + desk.space.pictureStorePath"
        style="width: 148px; height: 148px;"></el-image>
    </el-form-item>

    <el-form-item label="价格">
      <el-input v-if="desk.space!==undefined" v-model="desk.space.price" style="width: 70%"></el-input>
    </el-form-item>
    <!--    5. 提交按钮-->
    <el-form-item>
      <el-button type="primary" @click="addDesk(desk)">确定</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
/* eslint-disable */
import utils from "../../common/utils";

export default {
  name: "DeskEditorOfAdmin",
  components: {},
  props: {
    className: String,
  },
  data() {
    return {
      desk: {},
    }
  },
  methods: {
    async addDesk() {
      await utils.AddDesk(this, {
        desk: this.desk,
        className: this.className,
      }, (res) => {

      })
    },
    cleanSizeInfoPictureStorePath() {
      this.desk.space.pictureStorePath = ""
    },
    imageUploadSuccess(res, file, fileList) {
      this.desk.space.pictureStorePath = res.data.fileStorePath;
    },
  }
}
</script>

<style scoped>

</style>
