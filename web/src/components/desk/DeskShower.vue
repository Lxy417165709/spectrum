<!-- eslint-disable -->
<template>
  <el-container style="height: 800px; border: 1px solid #eee">
    <el-aside width="200px">
      <el-menu>
        <el-menu-item v-for="(unit,unitIndex) in units" :key="unitIndex" :index="unit.name"
                      @click="handleUnitClick(unit)">
          <template slot="title"><i class="el-icon-message"></i><span>{{ unit.desk.name }}</span></template>
        </el-menu-item>
      </el-menu>
    </el-aside>
    <el-main>
      <span>{{curUnit.desk.name}}</span>
      <el-tabs>
        <component :is="curUnit.component" :desk="curUnit.desk"></component>
      </el-tabs>
    </el-main>
  </el-container>
</template>
<script>
/* eslint-disable */
import DeskCard from "./DeskCard";

export default {
  name: 'DeskShower',
  components: {DeskCard},

  mounted() {
    this.curUnit = this.units[0]
  },
  data() {
    return {
      units: [
        {
          component: DeskCard,
          desk: {
            name: '台球桌',
            nums: [1, 2, 3]
          },
        },
        {
          component: DeskCard,
          desk: {
            name: '麻将桌',
            nums: [1, 2]
          },
        },


      ],
      curUnit: {}
    }
  },
  methods: {
    handleUnitClick(unit) {
      this.curUnit = unit
    }
  }
}
</script>

<style scoped>

</style>
