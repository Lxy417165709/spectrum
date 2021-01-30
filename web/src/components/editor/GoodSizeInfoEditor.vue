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
    <el-form-item label="价格">
      <el-input v-model="price" style="width: 70%"></el-input>
    </el-form-item>
  </el-form>
</template>

<script>
/* eslint-disable */
export default {
  props: {
    price: Number,
  },
  name: "GoodSizeInfoEditor"
}
</script>

<style scoped>

</style>
