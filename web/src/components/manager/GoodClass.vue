<!-- eslint-disable -->
<template>
  <div>
    <el-row style="height: 40px;background-color: #99a9bf;margin-left:10px;margin-bottom: 10px">
      <el-button style="height: 40px" @click="handleButtonClick">退回</el-button>
    </el-row>
    <div v-show="viewMode === 1">
      <good-class-list :goodClasses="goodClasses" :isEditMode="isEditMode"
                       @turnToGoodListMode="turnToGoodListMode"></good-class-list>
    </div>
    <div v-show="viewMode === 2">
      <good-list v-if="curGoodClassIndex!==-1" :isEditMode="isEditMode" :goods="goodClasses[curGoodClassIndex].goods"
                 @turnToGoodClassListMode="turnToGoodClassListMode"></good-list>
    </div>
  </div>
</template>

<script>
/* eslint-disable */
import GoodClassList from "../list/GoodClassList";
import GoodList from "../list/GoodList";

export default {
  name: "GoodClass",
  components: {GoodList, GoodClassList},
  props: {
    goodClasses: Array,
    isEditMode: Boolean,
  },
  mounted() {
  },
  data() {
    return {
      viewMode: 1,
      curGoodClassIndex: -1,
    }
  },
  methods: {
    turnToGoodListMode(mode, goodClassIndex) {
      this.viewMode = mode
      this.curGoodClassIndex = goodClassIndex
    },
    turnToGoodClassListMode() {
      this.viewMode = 1
    },
    handleButtonClick() {
      if (this.viewMode === 1) {
        this.$emit("turnToDeskListMode")
      }
      if (this.viewMode === 2) {
        this.viewMode = 1
      }
    }
  }
}
</script>

<style scoped>

</style>
