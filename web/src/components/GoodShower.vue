
<template>
  <div>
    <el-carousel :interval="4000" type="card" height="500px">
      <el-carousel-item v-for="good in goods">
        <h3 class="medium">{{ good.name }}</h3>
        <el-image
          :src="'api/file/' + good.pictureStorePath"
          fit="scale-down"></el-image>
        <template v-for="optionClass in good.optionClasses">
          <p>{{optionClass.className}}</p>
        </template>
      </el-carousel-item>
    </el-carousel>
    <el-button @click="getAllGoods">获得所有商品</el-button>
  </div>

</template>



<script>
import utils from "../common/utils";

export default {
  name:"GoodShower",
  data() {
    return {
      urls: [
        'api/file/static/upload/bg.jpg',
        'api/file/static/upload/bg.jpg',
        'api/file/static/upload/bg.jpg'],
      goods: []
    }
  },
  methods :{
    getAllGoods() {
      let model = utils.getRequestModel("mvp", "GetAllGoods", {})
      utils.sendRequestModel(model).then(res => {
        if (!utils.hasRequestSuccess(res)) {
          this.$message.error(res.data.err)
          return
        }
        this.goods = res.data.data.goods
        console.log(this.goods)
        this.$message.success(res.data.msg)
      })
    }
  }
}
</script>

<style>
.el-carousel__item{
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
