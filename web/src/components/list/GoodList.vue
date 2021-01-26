<!-- eslint-disable -->
<template>
    <el-row>
      <el-col v-for="(good,goodIndex) in goods" :key="goodIndex"
              style="height: 300px; width: 202px; margin-left: 10px; border: none">
        <span>
          <good-card :name="good.name" :pictureStorePath="good.pictureStorePath"></good-card>
        </span>
      </el-col>
      <el-col v-if="isEditMode" style="height: 300px; width: 202px; margin-left: 10px; border: none">
        <good-spacial-card></good-spacial-card>
      </el-col>
    </el-row>

</template>

<script>
/* eslint-disable */
import GoodCard from "../card/GoodCard";
import GoodSpacialCard from "../card/GoodSpacialCard";

export default {
  components: {GoodSpacialCard, GoodCard},
  props: {
    goods: Array,
    isEditMode: Boolean
  },
  data() {
    return {};
  },
  methods: {
    handleButtonClick(){
      this.$emit("turnToGoodClassListMode")
    }
  }
}
</script>
