<!-- eslint-disable -->
<template>
  <el-row>
    <el-col v-for="(goodClass,goodClassIndex) in goodClasses" :key="goodClassIndex"
            style="height: 300px; width: 202px; margin-left: 10px; border: none">
      <span @click="handleGoodClassCardClick(goodClassIndex)">
        <good-class-card  :name="goodClass.name"
                   :pictureStorePath="goodClass.pictureStorePath"></good-class-card>
      </span>
    </el-col>
  </el-row>
</template>

<script>
/* eslint-disable */
import GoodClassCard from "./GoodClassCard";
export default {
  components: {GoodClassCard},
  props: {
    goodClasses: Array
  },
  data() {
    return {};
  },
  methods: {
    handleGoodClassCardClick(goodClassIndex){
      console.log("handleDeskCardClick")
      this.$emit("passGoodsViewMode", 2, goodClassIndex)
    }
  }
}
</script>
