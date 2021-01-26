<!-- eslint-disable -->
<template>
  <el-row>
    <el-col v-for="(goodClass,goodClassIndex) in goodClasses" :key="goodClassIndex"
            style="height: 300px; width: 202px; margin-left: 10px; border: none">
      <span @click="handleGoodClassCardClick(goodClassIndex)">
        <good-class-card :name="goodClass.name"
                         :pictureStorePath="goodClass.pictureStorePath"></good-class-card>
      </span>
    </el-col>
    <el-col v-if="isEditMode" style="height: 300px; width: 202px; margin-left: 10px; border: none">
      <good-class-spacial-card></good-class-spacial-card>
    </el-col>
  </el-row>
</template>

<script>
/* eslint-disable */
import GoodClassCard from "../card/GoodClassCard";
import GoodClassSpacialCard from "../card/GoodClassSpacialCard";

export default {
  components: {GoodClassSpacialCard, GoodClassCard},
  props: {
    goodClasses: Array,
    isEditMode: Boolean
  },
  data() {
    return {};
  },
  methods: {
    handleGoodClassCardClick(goodClassIndex) {
      console.log("handleDeskCardClick")
      this.$emit("turnToGoodListMode", 2, goodClassIndex)
    },
  }
}
</script>
