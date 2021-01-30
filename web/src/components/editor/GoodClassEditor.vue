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
    <el-form-item label="">
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
