<!-- eslint-disable -->
<template>
  <el-form label-width="80px">
    <el-form-item label="折扣">
      <el-input v-model="curDiscount" style="width: 20%"></el-input>
      <el-button @click="handleClick"> 确定</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
/* eslint-disable */
export default {
  name: "DiscountComponent",
  data() {
    return {
      curDiscount: 10,
    }
  },
  methods: {
    handleClick() {
      this.$emit("confirmFavor", {
        name: "打 " + this.curDiscount + " 折"
      })
    }
  }
}
</script>

<style scoped>

</style>
