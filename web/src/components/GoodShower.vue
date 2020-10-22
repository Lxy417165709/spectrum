<!-- eslint-disable -->
<template>
  <el-row>
    <el-col v-for="(good,goodIndex) in goods" :key="goodIndex" style="height: 300px; width: 200px; margin-left: 10px">
      <span @dblclick="order(good)">
        <el-card :body-style="{ padding: '0px' }" shadow="hover" style="position:relative;">
        <img
          :src="'api/file/' + good.pictureStorePath"
          style="height:200px;width:200px">
        <div style="width: 200px;height: 30px;position:relative;padding: 0;margin: 0">
          <span style="position:absolute; left:5px">{{ good.name }}</span>
          <time style="color: #dd6161; position:absolute; right:5px">{{ good.price }}</time>
        </div>
        <div style="
          display:flex;
          align-items: center;
          justify-content: center;
          color: white;
          width:20px;
          height:20px;
          position: absolute;
          right:10px;
          top:10px;
          background-color: #0db0f8;
          border-radius: 50%">
          {{ good.count === undefined ? 0 : good.count }}
        </div>
      </el-card>

      </span>

    </el-col>
  </el-row>
</template>

<script>
/* eslint-disable */
export default {
  name: "GoodShower",
  props: {
    goods: Array,
  },
  data() {
    return {};
  },
  methods: {
    order(good) {
      this.$emit("passGoodToParent", good)
      this.$message.success("点单成功")
      good.count++
    },
  }
}
</script>
