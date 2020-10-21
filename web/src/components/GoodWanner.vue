<template>
  <el-container style="height: 800px; border: 1px solid #eee">
    <el-aside width="200px">
      <el-menu @click="handleClick">
        <el-submenu v-for="(goodClass,goodClassIndex) in goodClasses" :key="goodClassIndex" :index="goodClass.name">
          <template slot="title"><i class="el-icon-message"></i><span
            @click="addTab(goodClass)">{{ goodClass.name }}</span></template>
        </el-submenu>
      </el-menu>
    </el-aside>
    <el-main>
      <el-tabs>
        <el-tab-pane v-for="unit in units" :label="unit.goodClassName">
          <component :is="unit.component" :goods="unit.goods"></component>
        </el-tab-pane>
      </el-tabs>
    </el-main>
  </el-container>
</template>
<script>

import GoodShower from "./GoodShower";
import utils from "../common/utils";
import init from "../common/global_object/init";
import global from "../common/global_object/global";

export default {
  async mounted() {
    await init.globalGoodClasses()
    this.goodClasses = global.goodClasses
  },
  data() {
    return {
      isCollapse: true,
      goodClasses: [],
      units: []
    };
  },
  methods: {
    addTab(goodClass) {
      this.units.push(
        {
          goodClassName: goodClass.name,
          component: GoodShower,
          goods: utils.deepCopy(goodClass.goods)
        }
      )
    },
    handleOpen(key, keyPath) {
      console.log(key, keyPath);
    },
    handleClose(key, keyPath) {
      console.log(key, keyPath);
    },
    handleClick() {
      console.log(111)
    }
  }
}
</script>


