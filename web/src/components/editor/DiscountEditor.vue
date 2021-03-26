<!-- eslint-disable -->
<template>
  <!--  todo: 这个组件的命名可以改下，因为它只是包括了 item,而 item 只能在 form 中显示-->
  <div>
    <el-form-item label="优惠">
      <el-radio v-model="curFavorIndex" v-for="(favor,index) in favors" :label="index"
                :key="index">
        {{ favor.name }}
      </el-radio>
    </el-form-item>
    <component :is="favors[curFavorIndex].component" @confirmFavor="confirmFavor"></component>
    <el-form-item label="已选折扣">
      <el-tag v-for="(favor,index) in selectedFavors" :key="index" closable style="margin-right: 10px">
        {{ getFavorTagName(favor) }}
      </el-tag>
    </el-form-item>
  </div>
</template>

<script>
/* eslint-disable */
import DiscountComponent from "../baby/DiscountComponent";
import FullReductionComponent from "../baby/FullReductionComponent";
import FreeComponent from "../baby/FreeComponent";
import cst from "../../common/cst";
import utils from "../../common/utils";

export default {
  name: "DiscountEditor",
  data() {
    return {
      selectedFavors: [],
      curFavorIndex: 0,
      favors: [
        {
          name: cst.FAVOR_TYPE.NONE.NAME,
        },
        {
          name: cst.FAVOR_TYPE.REBATE.NAME,
          component: DiscountComponent
        },
        {
          name: cst.FAVOR_TYPE.FULL_REDUCTION.NAME,
          component: FullReductionComponent
        },
        {
          name: cst.FAVOR_TYPE.FREE.NAME,
          component: FreeComponent
        }
      ],
    }
  },
  methods: {
    getFavorTagName(favor) {
      return utils.GetFavorTagName(favor)
    },
    confirmFavor(favor) {
      this.selectedFavors.push(favor)
    }
  }
}
</script>

<style scoped>

</style>
