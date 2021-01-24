<!-- eslint-disable -->
<template>
  <el-row>
    <el-col v-for="(goodClass,goodClassIndex) in goodClasses" :key="goodClassIndex"
            style="height: 300px; width: 202px; margin-left: 10px; border: none">
      <span @click="handleGoodClassCardClick(goodClassIndex)">
        <component :is="goodClass.component" :name="goodClass.name"
                   :pictureStorePath="goodClass.pictureStorePath"></component>
      </span>
    </el-col>
  </el-row>
</template>

<script>
/* eslint-disable */
export default {
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
