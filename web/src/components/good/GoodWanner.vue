<!-- eslint-disable -->
<template>
  <el-container style="height: 800px; border: 1px solid #eee">
    <el-aside width="200px">
      <el-menu>
        <el-submenu v-for="(goodClass,goodClassIndex) in goodClasses" :key="goodClassIndex" :index="goodClass.name">
          <template slot="title"><i class="el-icon-message"></i><span
            @click="addTab(goodClass)">{{ goodClass.name }}</span></template>
        </el-submenu>
      </el-menu>
    </el-aside>
    <el-main>
      <el-tabs>
        <el-tab-pane v-for="(unit,index) in units" :key="index" :label="unit.goodClassName">
          <component :is="unit.component" :goods="unit.goods" @passGoodToParent="receiveChildGood"></component>
        </el-tab-pane>
      </el-tabs>
    </el-main>
    <el-drawer
      :visible.sync="drawer"
      direction="rtl"
      title="当前订单">
      <div style="width: 100%; border:none;max-height: 400px;overflow:scroll; padding: 0">
        <el-collapse accordion>
          <el-collapse-item v-for="(orderGoodUnit,index) in orderGoodUnits"
                            :key="index" style="padding-left: 10px;position: relative">
            <template slot="title">
              <span>{{ (index + 1) + '. ' + orderGoodUnit.good.name }}</span>
              <el-button circle icon="el-icon-close" size="mini" style="position: absolute;right: 35px" type="danger"
                         @click.stop="delGood(index)"></el-button>
            </template>

            <el-collapse accordion>
              <el-collapse-item name="1" style="padding-left: 20px;" title="商品详情">
                <component :is="orderGoodUnit.detailComponent" :good="orderGoodUnit.good"></component>
              </el-collapse-item>
              <el-collapse-item name="2" style="padding-left: 20px;" title="折扣处理">
                <Discounter></Discounter>
              </el-collapse-item>
            </el-collapse>
          </el-collapse-item>
        </el-collapse>
      </div>
      <el-form label-width="80px" style="width: 100%; margin-top:20px">
        <Discounter></Discounter>
        <el-form-item>
          <el-button @click="sendOrder()">确定</el-button>
        </el-form-item>
      </el-form>

    </el-drawer>
    <el-button style="position:absolute;bottom:100px;left:50px" type="primary" @click="drawer=true">打开当前订单</el-button>
  </el-container>
</template>
<script>
/* eslint-disable */
import GoodShower from "./GoodShower";
import utils from "../../common/utils";
import init from "../../common/global_object/init";
import global from "../../common/global_object/global";
import GoodEditor from "./GoodEditor";
import Discounter from "../discount/Discounter";

export default {
  components: {Discounter},
  async mounted() {
    await init.globalGoodClasses()
    this.goodClasses = global.goodClasses
  },
  data() {
    return {
      isCollapse: true,
      goodClasses: [],
      units: [],
      goodEditorUnits: [],
      orderGoodUnits: [],
      drawer: false,
    };
  },
  methods: {
    sendOrder() {
      let goods = []
      for (let i = 0; i < this.orderGoodUnits.length; i++) {
        goods.push(this.orderGoodUnits[i].good)
      }

      let order = {
        goods: goods,
        createdTime: new Date().getTime(),
        // discountsMap: {
        //   0: []   // 先不使用折扣
        // }
      }
      let model = utils.getRequestModel("mvp", "Order", {
        order: order,
      })
      console.log("order", order)
      utils.sendRequestModel(model).then(async res => {
        console.log(res)
      })
    },
    receiveChildGood(good) {
      let newGood = utils.deepCopy(good) // 深拷贝
      for (let i = 0; newGood.attachGoodClasses !== undefined && i < newGood.attachGoodClasses.length; i++) {
        newGood.attachGoodClasses[i].selectGoodNames = [] // 添加 selectGoodIndexes 字段，这样子组件才能根据 selectGoodIndexes 进行响应
      }
      this.orderGoodUnits.push({
        detailComponent: GoodEditor,
        good: newGood
      })
    },
    delGood(index) {
      for (let i = 0; i < this.units.length; i++) {
        for (let t = 0; t < this.units[i].goods.length; t++) {
          if (this.units[i].goods[t].name === this.orderGoodUnits[index].good.name) {
            this.units[i].goods[t].count--
            break
          }
        }
      }
      this.orderGoodUnits = utils.removeIndex(this.orderGoodUnits, index)
    },
    addTab(goodClass) {
      for (let i = 0; i < goodClass.goods.length; i++) {
        goodClass.goods[i].count = 0 // 添加 count 字段，这样子组件才能根据 count 进行响应
      }
      this.units.push({
        goodClassName: goodClass.name,
        component: GoodShower,
        goods: utils.deepCopy(goodClass.goods)
      })
    },
  }
}
</script>


