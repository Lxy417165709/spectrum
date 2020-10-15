<template>
  <el-form ref="form" :model="form" label-width="80px">
    <el-form-item label="商品名">
      <el-input v-model="form.name"></el-input>
    </el-form-item>
    <el-form-item label="价格">
      <el-input v-model="form.price"></el-input>
    </el-form-item>
        <el-form-item label="附属选项">
          <el-select v-model="form.goodOption" placeholder="请选择附属选项">
            <el-option label="温度" value="wei_du"></el-option>
            <el-option label="冰量" value="bing_liang"></el-option>
          </el-select>
          <el-button type="primary" @click="addOption">添加</el-button>
        </el-form-item>
    <el-form-item label="活动时间">
      <el-col :span="11">
        <el-date-picker v-model="form.date1" placeholder="选择日期" style="width: 100%;" type="date"></el-date-picker>
      </el-col>
      <el-col :span="2" class="line">-</el-col>
      <el-col :span="11">
        <el-time-picker v-model="form.date2" placeholder="选择时间" style="width: 100%;"></el-time-picker>
      </el-col>
    </el-form-item>
    <el-form-item label="即时配送">
      <el-switch v-model="form.delivery"></el-switch>
    </el-form-item>
    <el-form-item label="活动性质">
      <el-checkbox-group v-model="form.type">
        <el-checkbox v-for="gp in form.goodOptions">{{gp}}</el-checkbox>
<!--        <el-checkbox label="美食/餐厅线上活动" name="type"></el-checkbox>-->
<!--        <el-checkbox label="地推活动" name="type"></el-checkbox>-->
<!--        <el-checkbox label="线下主题活动" name="type"></el-checkbox>-->
<!--        <el-checkbox label="单纯品牌曝光" name="type"></el-checkbox>-->
      </el-checkbox-group>
    </el-form-item>
<!--    <el-form-item label="特殊资源">-->
<!--      <el-radio-group v-model="form.resource">-->
<!--        <el-radio label="线上品牌商赞助"></el-radio>-->
<!--        <el-radio label="线下场地免费"></el-radio>-->
<!--      </el-radio-group>-->
<!--    </el-form-item>-->
    <el-form-item label="活动形式">
      <el-input v-model="form.desc" type="textarea"></el-input>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">立即创建</el-button>
      <el-button>取消</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
export default {
  name: 'App',
  data() {
    return {
      form: {
        goodOption: '',
        price: 0,
        name: '',
        region: '',
        date1: '',
        date2: '',
        delivery: false,
        type: [],
        resource: '',
        desc: '',
        goodOptions: []
      }
    }
  },
  methods: {
    onSubmit() {
      console.log('submit!');
    },
    addOption(){
      this.form.goodOptions.push(this.form.goodOption);
      console.log(this.form.goodOption,this.form.goodOptions);
    }
  }
}
</script>

<style>
</style>
