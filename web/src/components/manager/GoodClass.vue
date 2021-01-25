<!-- eslint-disable -->
<template>
  <div>
    <div v-show="viewMode === 1">
      <good-class-list :goodClasses="goodClasses" @passGoodsViewMode="changeToGoodsMode"></good-class-list>
    </div>
    <div v-show="viewMode === 2">
      <good-list v-if="curGoodClassIndex!==-1" :goods="goodClasses[curGoodClassIndex].goods"></good-list>
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
    changeToGoodsMode(mode, goodClassIndex) {
      this.viewMode = mode
      this.curGoodClassIndex = goodClassIndex
    }
  }
}
</script>

<style scoped>

</style>
