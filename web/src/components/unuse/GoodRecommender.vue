<!--eslint-disable-->
<template>
  <div>
    <el-carousel :interval="4000" height="560px" type="card">
      <el-carousel-item v-for="(good,goodIndex) in goods" :key="goodIndex">
        <h3 class="medium">{{ good.name }}</h3>
        <el-image
          :src="'api/file/' + good.pictureStorePath"
          fit="scale-down"></el-image>
        <template v-for="optionClass in good.optionClasses">
          <p>{{ optionClass.className }}</p>
          <el-radio-group v-for="(optionName,index) in optionClass.optionNames"
                          v-model="selectOptionIndex[optionClass]">
            <el-radio :label="index">{{ optionName }}</el-radio>
          </el-radio-group>
        </template>
      </el-carousel-item>
    </el-carousel>
    <el-button @click="getAllGoods">获得所有商品</el-button>
  </div>

</template>

<script>
/* eslint-disable */

export default {
  name: "GoodRecommender",
  data() {
    return {
      urls: [
        'api/file/static/upload/bg.jpg',
        'api/file/static/upload/bg.jpg',
        'api/file/static/upload/bg.jpg'],
      goods: [],
      selectOptionIndex: {}
    }
  },
  methods: {
    async getAllGoods() {
      // await init.globalGoods()
      // this.goods = global.goods;
      // for (let good in this.goods) {
      //   console.log("www",good.optionClasses)
      //   for(let optionClass in good.optionClasses){
      //     this.selectOptionIndex[optionClass] = 0
      //   }
      // }
    }
  }
}
</script>

<!--eslint-disable-->
<style>
.el-carousel__item {
  text-align: center;
}

.el-carousel__item h3 {
  color: #475669;
  font-size: 14px;
  opacity: 0.75;
  line-height: 20px;
  margin: 0;
}


.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n+1) {
  background-color: #d3dce6;
}
</style>
