<!-- eslint-disable -->
<template>
  <el-form label-width="80px">
    <el-form-item label="照片">
      <img style="height:200px;width:200px; border: none;margin-top: 12px">

    </el-form-item>
    <el-form-item label="名字">
      <span>{{ good.name }}</span>
    </el-form-item>
    <el-form-item label="规格">
      <el-radio v-model="good.curSizeIndex" v-for="(sizeInfo,index) in good.sizeInfos" :label="index" :key="index">
        {{ sizeInfo.name }}
      </el-radio>
    </el-form-item>
    <el-form-item v-for="(attachElement,index) in good.attachElements" :key="index" :label="attachElement.name"
                  v-if="attachElement.elementType===1">
      <el-radio v-model="attachElement.curSizeIndex" v-for="(sizeInfo,index) in attachElement.sizeInfos" :label="index"
                :key="index">
        {{ sizeInfo.name }}
      </el-radio>
    </el-form-item>
    <el-form-item v-if="!needAttachGood">
      <el-button @click="needAddAttachGood">需要加料</el-button>
    </el-form-item>
    <el-form-item v-for="(attachElement,index) in good.attachElements" :key="index" :label="attachElement.name"
                  v-if="attachElement.elementType===2 && needAttachGood">
      <el-radio v-model="attachElement.curSizeIndex" v-for="(sizeInfo,index) in attachElement.sizeInfos" :label="index"
                :key="index">
        {{ sizeInfo.name }}
      </el-radio>
    </el-form-item>
    <discount-editor></discount-editor>

    <el-form-item label="价格">
      <span>{{ good.price }}</span>
    </el-form-item>
    <el-form-item label="备注">
      <el-input style="width: 70%"></el-input>
    </el-form-item>
    <el-form-item>
      <el-button type="primary">确定</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
/* eslint-disable */
import DiscountEditor from "./DiscountEditor";

export default {
  name: "GoodSellEditor",
  components: {DiscountEditor},
  props: {
  },
  data() {
    return {
      needAttachGood: false,
      good: {}
    }
  },
  methods: {
    needAddAttachGood() {
      this.needAttachGood = true
    },
  }
}
</script>

<style scoped>

</style>
