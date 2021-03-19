<!-- eslint-disable -->
<template>
  <el-form label-width="80px">
    <!--    1. 商品类的名字-->
    <el-form-item label="类名">
      <el-input v-model="goodClass.name" style="width: 70%"></el-input>
    </el-form-item>

    <!--    2. 商品类的图片-->
    <el-form-item label="照片">
      <el-upload
        v-if="!canImageShow(goodClass)"
        action="/api/upload"
        :on-success="imageUploadSuccess"
        list-type="picture-card">
        <i class="el-icon-plus"></i>
      </el-upload>
      <el-image
        v-if="canImageShow(goodClass)"
        @dblclick.native="cleanSizeInfoPictureStorePath"
        :src="'api/file/' + goodClass.pictureStorePath"
        style="width: 148px; height: 148px;"></el-image>
    </el-form-item>


    <!--    3. 确认按钮-->
    <el-form-item label="">
      <el-button type="primary" @click="addGoodClass(goodClass)">确定</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
/* eslint-disable */
import utils from "../../common/utils";

export default {
  name: "GoodClassEditor",
  data() {
    return {
      goodClass: {}
    }
  },
  methods: {
    async addGoodClass(goodClass) {
      let model = utils.getRequestModel("mvp", "AddGoodClass", {
        goodClass: goodClass,
      })
      await utils.sendRequestModel(model).then(res => {
        if (!utils.hasRequestSuccess(res)) {
          this.$message.error(res.data.err)
          return
        }
        this.$message.success(res.data.msg)
        this.goodClass = res.data.data.goodClass
      })
    },
    cleanSizeInfoPictureStorePath() {
      this.goodClass.pictureStorePath = ""
    },
    imageUploadSuccess(res, file, fileList) {
      this.goodClass.pictureStorePath = res.data.fileStorePath;
    },
    canImageShow(obj) {
      return obj.pictureStorePath !== undefined && obj.pictureStorePath !== ''
    }
  },
  computed: {}
}
</script>

<style scoped>

</style>
