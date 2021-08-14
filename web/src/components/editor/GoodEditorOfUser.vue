<!-- eslint-disable -->
<template>
  <el-form label-width="80px">
    <el-form-item label="照片">
      <el-image
        v-if="good.mainElement!==undefined"
        :src="'api/file/' + good.mainElement.sizeInfos[good.mainElement.selectedIndex].pictureStorePath"
        style="height:200px;width:200px; border: none;margin-top: 12px"></el-image>

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
    <discount-editor ref="DiscountEditor" :index="-1" chargeableObjName="good"
                    @addFavorForGood="addFavorForGood"
                    @delFavorForGood="delFavorForGood"
                     :favors="good.favors"
    ></discount-editor>

    <el-form-item label="原价格" v-if="good.mainElement!==undefined">
      <span style="font-size: 1.4em;color: red;">{{ cpt_price.toFixed(2) }} 元</span>
    </el-form-item>
    <el-form-item label="折后价格" v-if="good.mainElement!==undefined && this.good.favors!== null && this.good.favors.length>0">
      <span style="font-size: 1.4em;color: red;">{{ expense.toFixed(2) }} 元</span>
    </el-form-item>
<!--    <el-form-item label="备注">-->
<!--      <el-input style="width: 70%"></el-input>-->
<!--    </el-form-item>-->
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
      expense : 0,
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
    },
    addFavorForGood(favor, index) {
      if (utils.IsNil(this.good.favors)){
        this.good.favors = []
      }
      this.good.favors.push(favor)
      this.flashExpense()
    },
    delFavorForGood(favorIndex) {
      this.good.favors = utils.removeIndex(this.good.favors,favorIndex)
      this.flashExpense()
    },
    async flashExpense() {
      let price = this.cpt_price;
      await utils.GetExpense(this,{
        notFavorExpense:price,
        favors:this.good.favors
      },(res)=>{
       this.expense =  res.data.data.expense
      })
    },
  },
  computed: {
    cpt_price() {
      let price = 0
      price += this.good.mainElement.sizeInfos[this.good.mainElement.selectedIndex].price - 0
      for (let i = 0; i < this.good.attachElements.length; i++) {
        price += this.good.attachElements[i].sizeInfos[this.good.attachElements[i].selectedIndex].price - 0
      }
      this.flashExpense()
      return price
    }
  }
}
</script>

<style scoped>

</style>
