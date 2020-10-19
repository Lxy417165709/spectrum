<template>
  <el-container style="height: 800px; border: 1px solid #eee">
    <el-aside width="200px">
      <el-menu>
        已点商品
        <el-submenu v-for="good in goods" :index="good.name">
          <template slot="title"><i class="el-icon-message" @click="addTab(good)"></i>{{good.name}}</template>
        </el-submenu>
      </el-menu>
    </el-aside>
    <el-main>
        <el-tabs>
            <el-tab-pane  v-for="unit in units" :label="unit.good.name">
              <component :is="unit.component" :good="unit.good"></component>
            </el-tab-pane>
        </el-tabs>
    </el-main>
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
    console.log(this.goods)
  },
  data() {
    return {
      goods : [],
      goodMap : {},
      units : [],
      activeName:""
    }
  },
  methods: {
    addTab(good){
      let unit = {
        component: GoodOrder,
        good: good
      }
      this.units.push(unit)
      // this.activeName = tabName
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
