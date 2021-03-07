<!-- eslint-disable -->
<template>
  <el-row>
    <el-col v-for="(goodOptionClass,goodOptionClassIndex) in goodOptionClasses" :key="goodOptionClassIndex"
            style="height: 300px; width: 202px; margin-left: 10px; border: none">
      <good-option-class-card
                              :goodOptionClass="goodOptionClass"
                              @click.native="handleGoodOptionClassCardClick(goodOptionClassIndex)"
      ></good-option-class-card>

    </el-col>
    <!--    <el-col v-if="isAdminView" style="height: 300px; width: 202px; margin-left: 10px; border: none">-->
    <!--      <good-option-special-card></good-option-special-card>-->
    <!--    </el-col>-->
  </el-row>
</template>

<script>
/* eslint-disable */

import GoodOptionCard from "../card/GoodOptionCard";
import GoodOptionSpecialCard from "../card/GoodOptionSpecialCard";
import GoodOptionClassCard from "../card/GoodOptionClassCard";

export default {
  components: {GoodOptionClassCard, GoodOptionCard, GoodOptionSpecialCard},
  props: {
    goodOptionClasses: Array,
    isAdminView: Boolean,
  },
  data() {
    return {};
  },
  methods: {
    handleGoodOptionClassCardClick(goodOptionClassIndex) {
      this.$emit("turnToGoodOptionListMode", goodOptionClassIndex)
    }
  }
}
</script>
