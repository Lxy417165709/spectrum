<!-- eslint-disable -->
<template>
  <el-form ref="form" label-width="80px">
    <el-form-item label="照片">
      <el-image
        :src="'api/file/' + good.pictureStorePath"
        style="width: 100px; height: 100px;padding-top: 15px"></el-image>
    </el-form-item>
    <el-form-item label="商品名">
      {{ good.name }}
    </el-form-item>
    <el-form-item label="价格">
      {{ good.price }}
    </el-form-item>
    <template v-for="optionClass in good.optionClasses">
      <el-form-item :label="optionClass.name">
        <el-radio-group v-for="(option,index) in optionClass.options" :key="index"
                        v-model="optionClass.selectOptionIndex">
          <el-radio :label="index+1" style="padding-right:10px;">{{ option.name }}</el-radio>
        </el-radio-group>
      </el-form-item>
    </template>
  </el-form>
</template>
<script>
/* eslint-disable */
export default {
  mounted() {
    console.log(this.good)
  },
  props: {
    good: Object
  }
}
</script>

<style scoped>

</style>
