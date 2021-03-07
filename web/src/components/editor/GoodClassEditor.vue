<!-- eslint-disable -->
<template>
  <el-form label-width="80px">
    <!--    1. 商品类的图片-->
    <el-form-item label="照片">
      <el-upload
        action="/api/upload"
        list-type="picture-card">
        <i class="el-icon-plus"></i>
      </el-upload>
    </el-form-item>

    <!--    2. 商品类的名字-->
    <el-form-item label="类名">
      <el-input v-model="goodClass.name" style="width: 70%"></el-input>
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
        goodClass: utils.goodClassToPbGoodClass(goodClass),
      })
      await utils.sendRequestModel(model).then(res => {
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
