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
        {{ favor.name }}
      </el-tag>
    </el-form-item>
  </div>
</template>

<script>
/* eslint-disable */
import DiscountComponent from "../baby/DiscountComponent";

export default {
  name: "DiscountEditor",
  data() {
    return {
      selectedFavors: [
        {
          name: "打 8 折"
        },
        {
          name: "满 100 减 6"
        },
      ],
      curFavorIndex: 0,
      // todo: favor 被选中后，要展示 favor 对应的组件
      favors: [
        {
          name: "无",
        },
        {
          name: "打折",
          component: DiscountComponent
        },
        {
          name: "满减",
        },
        {
          name: "免单",
        }
      ],
    }
  },
  methods: {
    confirmFavor(favor) {
      this.selectedFavors.push(favor)
    }
  }
}
</script>

<style scoped>

</style>
