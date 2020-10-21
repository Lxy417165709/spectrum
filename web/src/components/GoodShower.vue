<template>
  <el-row>
    <el-col v-for="(good,goodIndex) in goods" :key="goodIndex" style="height:300px;width:200px;margin-left: 10px">
      <el-card :body-style="{ padding: '0px' }" shadow="hover">
        <img
          :src="'api/file/' + good.pictureStorePath"
             style="height:200px;width:200px">
        <div style="width: 200px;height: 30px;position:relative;padding: 0;margin: 0">
          <span style="position:absolute; left:5px">{{good.name}}</span>
          <time style="color: #dd6161; position:absolute; right:5px">{{good.price}}</time>
        </div>
      </el-card>
    </el-col>
  </el-row>
</template>

<script>
import init from "../common/global_object/init"
import global from "../common/global_object/global"
export default {
  name: "GoodShower",
  props: {
    goods:Array,
  },
  async mounted() {
    // await this.getAllGoods()
  },
  data() {
    return {
      // goods : []
    };
  },
  methods: {
    async getAllGoods() {
      // await init.globalGoods()
      // this.goods = global.goods;
    }
  }
}
</script>
