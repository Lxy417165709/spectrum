<!-- eslint-disable -->
<template>
  <el-form label-width="80px">
    <el-form-item label="是否免单">
      <el-button @click="handleClick"> 确定</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
/* eslint-disable */
export default {
  name: "FreeComponent",
  methods: {
    handleClick() {
      this.$emit("confirmFavor", {
        name: "免单"
      })
    }
  }
}
</script>

<style scoped>

</style>
