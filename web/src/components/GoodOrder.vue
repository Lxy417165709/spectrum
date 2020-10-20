<!--eslint-disable-->
<template>
  <div>
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
        <el-form-item :label="optionClass.className">
          <el-radio-group v-model="optionClass.defaultSelectOptionIndex" v-for="(optionName,index) in optionClass.optionNames">
            <el-radio :label="index+1" style="padding-right:10px;">{{ optionName}}</el-radio>
          </el-radio-group>
        </el-form-item>
      </template>
      <el-form-item >
        <el-button type="primary" @click="order">点单</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
/* eslint-disable */

export default {
  name: "GoodOrder",
  async mounted() {
  },
  props: {
    good: Object
  },
  data() {
    return {
    }
  },
  methods: {
    order() {
      this.$emit("passGoodToParent",this.good)
      this.$message.success("点单成功")
      // console.log(this.good.optionClasses[0].defaultSelectOptionIndex)
    }
  }
}
</script>
