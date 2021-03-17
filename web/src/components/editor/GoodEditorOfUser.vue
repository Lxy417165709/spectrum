<!-- eslint-disable -->
<template>
  <el-form label-width="80px">
    <el-form-item label="照片">
      <img style="height:200px;width:200px; border: none;margin-top: 12px">

    </el-form-item>
    <el-form-item label="名字">
      <span v-if="good.mainElement!==undefined">{{ good.mainElement.name }}</span>
    </el-form-item>
    <el-form-item label="规格" v-if="good.mainElement!==undefined">
      <el-radio v-model="good.mainElement.selectedIndex" v-for="(sizeInfo,index) in good.mainElement.sizeInfos"
                :label="index"
                :key="index">
        {{ sizeInfo.size }}
      </el-radio>
    </el-form-item>
    <el-form-item v-for="(attachElement,index) in good.attachElements" :key="index" :label="attachElement.name"
                  v-if="attachElement!==undefined  && attachElement.type===1">
      <el-radio v-model="attachElement.selectedIndex" v-for="(sizeInfo,index) in attachElement.sizeInfos" :label="index"
                :key="index">
        {{ sizeInfo.size }}
      </el-radio>
    </el-form-item>
    <el-form-item v-if="!needAttachGood">
      <el-button @click="needAddAttachGood">需要加料</el-button>
    </el-form-item>
    <el-form-item v-for="(attachElement,index) in good.attachElements" :key="index" :label="attachElement.name"
                  v-if="attachElement!==undefined && attachElement.type===2 && needAttachGood">
      <el-radio v-model="attachElement.selectedIndex" v-for="(sizeInfo,index) in attachElement.sizeInfos" :label="index"
                :key="index">
        {{ sizeInfo.size }}
      </el-radio>
    </el-form-item>
    <discount-editor></discount-editor>

    <el-form-item label="价格">
      <span>{{ countPrice }}</span>
    </el-form-item>
    <el-form-item label="备注">
      <el-input style="width: 70%"></el-input>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="orderGood">确定</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
/* eslint-disable */
import DiscountEditor from "./DiscountEditor";
import cst from "../../common/cst";
import utils from "../../common/utils";

export default {
  name: "GoodEditorOfUser",
  components: {DiscountEditor},
  props: {
    orderID: Number,
  },
  data() {
    return {
      needAttachGood: false,
      good: {},
      countPrice: 0,
    }
  },
  methods: {
    needAddAttachGood() {
      this.needAttachGood = true
    },
    orderGood() {
      utils.OrderGood(this, {
        goods: [
          this.good,
        ],
        orderID: this.orderID
      }, (res) => {
      })
    }
  }
}
</script>

<style scoped>

</style>
