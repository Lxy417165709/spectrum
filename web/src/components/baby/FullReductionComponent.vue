<!-- eslint-disable -->
<template>
  <el-form label-width="80px">
    <el-form-item label="满">
      <el-input v-model="curFull" style="width: 30%"></el-input>
      减
      <el-input v-model="curReduction" style="width: 30%"></el-input>
      <el-button @click="handleClick"> 确定</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
/* eslint-disable */
import cst from "../../common/cst";

export default {
  name: "DiscountComponent",
  data() {
    return {
      curFull: 100,
      curReduction: 6,
    }
  },
  methods: {
    handleClick() {
      this.$emit("confirmFavor", {
        favorType: cst.FAVOR_TYPE.FULL_REDUCTION.VALUE,
        parameters: [this.curFull+ '', this.curReduction+ '']
      })
    }
  }
}
</script>

<style scoped>

</style>
