<template>
  <el-container style="height: 800px; border: 1px solid #eee">
    <el-aside width="200px">
      <el-menu>
        已点商品
        <el-submenu v-for="good in goods" :index="good.name" :key="good.name">
          <template slot="title"><i class="el-icon-message" @click="addTab(good)"></i>{{good.name}}</template>
        </el-submenu>
      </el-menu>
    </el-aside>
    <el-main>
        <el-tabs>
            <el-tab-pane  v-for="unit in units" :label="unit.good.name">
              <component :is="unit.component" :good="unit.good" @passGoodToParent="receiveChildGood"></component>
            </el-tab-pane>
        </el-tabs>
    </el-main>
    <el-drawer
      title="订单"
      :visible.sync="drawer"
      :direction="direction">
<!--      <el-collapse v-model="activeNames" @change="handleChange" accordion>-->
      <el-collapse accordion>
        <el-collapse-item v-for="(good,index) in orderGoods" :title=" (index+1) + '. ' + good.name" style="padding-left: 10px">
          <el-collapse accordion>
            <el-collapse-item title="商品详情" name="1" style="padding-left: 20px;">
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
                <el-form-item>
                  <el-button type="primary">修改</el-button>
                </el-form-item>
              </el-form>
            </el-collapse-item>
            <el-collapse-item title="折扣处理" name="2" style="padding-left: 20px;">
              <div>控制反馈：通过界面样式和交互动效让用户可以清晰的感知自己的操作；</div>
              <div>页面反馈：操作后，通过页面元素的变化清晰地展现当前状态。</div>
            </el-collapse-item>
          </el-collapse>





        </el-collapse-item>
      </el-collapse>
    </el-drawer>
    <el-button style="position:absolute;bottom:100px;left:50px" type="primary" @click="drawer=true">打开订单</el-button>
  </el-container>

</template>



<script>

import init from "../common/global_object/init"
import global from "../common/global_object/global"
import GoodOrder from "./GoodOrder";


export default {
  name: "GoodSelector",
  async created() {
    await init.globalGoods()
    this.goods = global.goods;
  },
  data() {
    return {
      goods : [],
      goodMap : {},
      units : [],
      activeName:"",
      drawer: false,
      direction: 'rtl',
      orderGoods: []
    }
  },
  methods: {
    receiveChildGood(good){
      this.orderGoods.push(JSON.parse(JSON.stringify(good)))  // 深拷贝
    },
    addTab(good){
      let unit = {
        component: GoodOrder,
        good: JSON.parse(JSON.stringify(good))  // 深拷贝
      }
      this.units.push(unit)
    },
    handleOpen(key, keyPath) {
      console.log(key, keyPath);
    },
    handleClose(key, keyPath) {
      console.log(key, keyPath);
    },
  }
}
</script>
