<!-- eslint-disable -->
<template>
  <el-row>
    <el-col v-for="(goodClass,goodClassIndex) in goodClasses" :key="goodClassIndex"
            v-if="classType === 1 && goodClass.classType===1 || classType === -1
            && (goodClass.classType=== 2 || goodClass.classType=== 3)"

            style="height: 300px; width: 202px; margin-left: 10px; border: none">
      <span @click="handleGoodClassCardClick(goodClassIndex)">
        <good-class-card :name="goodClass.name"
                         :pictureStorePath="goodClass.pictureStorePath"
        ></good-class-card>
      </span>
    </el-col>
    <el-col v-if="isEditMode && classType===1" style="height: 300px; width: 202px; margin-left: 10px; border: none">
      <good-class-spacial-card @click.native="openGoodClassEditor"></good-class-spacial-card>
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
    classType: Number,
    goodClasses: Array,
    isEditMode: Boolean
  },
  mounted() {
    console.log(this.classType)
    for (let i = 0; i < this.goodClasses.length; i++) {
      console.log(i, this.goodClasses[i])
    }
    // console.log(this.goodClasses)
  },
  data() {
    return {};
  },
  methods: {
    handleGoodClassCardClick(goodClassIndex) {
      this.$emit("turnToGoodListMode", 2, goodClassIndex)
    },
    openGoodClassEditor() {
      this.$emit("openGoodClassEditor")
    }
  }
}
</script>
