<!-- eslint-disable -->
<template>
  <el-form label-width="80px">
    <el-form-item label="照片">
      <el-upload
        action="/api/upload"
        list-type="picture-card">
        <i class="el-icon-plus"></i>
      </el-upload>
    </el-form-item>
    <el-form-item label="类名">
      <el-input v-model="className" style="width: 70%"></el-input>
    </el-form-item>
    <!--    todo: 这里可以优化为 for 循环-->
<!--    <el-form-item label="类别">-->
<!--      <el-radio v-model="classType" label="1">普通商品</el-radio>-->
<!--      <el-radio v-model="classType" label="2">附属属性</el-radio>-->
<!--      <el-radio v-model="classType" label="3">附属商品</el-radio>-->
<!--    </el-form-item>-->

    <el-form-item label=" ">
      <el-button type="primary">确定</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
/* eslint-disable */
export default {
  name: "GoodClassEditor",
  data() {
    return {
      className: "奶茶系列",
      classType: "1"
    }
  }
}
</script>

<style scoped>

</style>
