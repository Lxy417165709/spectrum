<!--eslint-disable-->
<template>
  <el-form ref="form" label-width="80px">
    <el-form-item label="照片">
      <el-image
        :src="'api/file/' + good.pictureStorePath"
        style="width: 100px; height: 100px;padding-top: 15px"></el-image>
    </el-form-item>
    <el-form-item label="商品名">
      {{ good.name }}
    </el-form-item>
    <el-form-item label="价格">
      {{ good.price }}
    </el-form-item>
    <template v-for="optionClass in good.optionClasses">
      <el-form-item :label="optionClass.className">
        {{optionClass.defaultSelectOptionIndex}}
          <el-radio-group v-model="optionClass.defaultSelectOptionIndex" v-for="(optionName,index) in optionClass.optionNames">
            <el-radio :label="index" style="padding-right:10px;">{{ optionName}}</el-radio>
          </el-radio-group>
      </el-form-item>
    </template>

  </el-form>
</template>

<script>
/* eslint-disable */

export default {
  name: "GoodOrder",
  async mounted() {
    if (this.good.optionClasses!== undefined){
      for(let i=0;i<this.good.optionClasses.length;i++){
        this.$set(this.selectOptionIndex,this.good.optionClasses[i].className,0)  // 涉及 vue 响应原理..
      }
    }
  },
  props: {
    good: {
      pictureStorePath:String
    }
  },
  data() {
    return {
      // good:{
      //   type: true,
      //   name:"波霸奶茶",
      //   price:50,
      //   pictureStorePath:"static/upload/bg.jpg",
      //   optionClasses:[
      //     {
      //       className:"冰量",
      //       optionNames:["正常冰","少冰","多冰"]
      //     },
      //     {
      //       className:"温度",
      //       optionNames:["常温","冷饮","热饮"]
      //     }
      //   ]
      // },
      selectOptionIndex: {},
    }
  },
  methods: {}
}
</script>
