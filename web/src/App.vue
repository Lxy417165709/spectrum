<!--eslint-disable-->

<template>
  <el-form ref="form" :model="form" label-width="80px">
    <el-form-item label="商品名">
      <el-input v-model="form.name"></el-input>
    </el-form-item>
    <el-form-item label="价格">
      <el-input v-model="form.price"></el-input>
    </el-form-item>
    <el-form-item v-show="form.goodOptions.length > 0" label="附属选项">
      <el-select v-model="form.selectOption" placeholder="请选择附属选项">
        <template v-for="gp in form.goodOptions">
          <el-option :label="gp" :value="gp"></el-option>
        </template>
      </el-select>
      <el-button type="primary" @click="addOption">添加</el-button>
    </el-form-item>
    <!--    <el-form-item label="活动时间">-->
    <!--      <el-col :span="11">-->
    <!--        <el-date-picker v-model="form.date1" placeholder="选择日期" style="width: 100%;" type="date"></el-date-picker>-->
    <!--      </el-col>-->
    <!--      <el-col :span="2" class="line">-</el-col>-->
    <!--      <el-col :span="11">-->
    <!--        <el-time-picker v-model="form.date2" placeholder="选择时间" style="width: 100%;"></el-time-picker>-->
    <!--      </el-col>-->
    <!--    </el-form-item>-->
    <!--    <el-form-item label="即时配送">-->
    <!--      <el-switch v-model="form.delivery"></el-switch>-->
    <!--    </el-form-item>-->
    <el-form-item v-show="form.selectOptions.length > 0" label="已选选项">
      <template v-for="gp in form.selectOptions">
        <el-checkbox-group v-model="form.type">
          <el-checkbox :label="gp"></el-checkbox>
          <!--        <el-checkbox label="地推活动" name="type"></el-checkbox>-->
          <!--        <el-checkbox label="线下主题活动" name="type"></el-checkbox>-->
          <!--        <el-checkbox label="单纯品牌曝光" name="type"></el-checkbox>-->
        </el-checkbox-group>
      </template>
    </el-form-item>
    <!--    <el-form-item label="特殊资源">-->
    <!--      <el-radio-group v-model="form.resource">-->
    <!--        <el-radio label="线上品牌商赞助"></el-radio>-->
    <!--        <el-radio label="线下场地免费"></el-radio>-->
    <!--      </el-radio-group>-->
    <!--    </el-form-item>-->
    <!--    <el-form-item label="活动形式">-->
    <!--      <el-input v-model="form.desc" type="textarea"></el-input>-->
    <!--    </el-form-item>-->
    <el-form-item>
      <el-button type="primary" @click="onSubmit">立即创建</el-button>
      <el-button>取消</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
/* eslint-disable */
import axios from 'axios';

export default {
  name: 'App',
  data() {
    return {
      form: {
        selectOption: '',
        price: 0,
        name: '',
        region: '',
        date1: '',
        date2: '',
        delivery: false,
        type: [],
        resource: '',
        desc: '',
        goodOptions: ["温度", "冰量"],
        selectOptions: [],
      }
    }
  },
  methods: {
    onSubmit() {
      let data = {
        object:"yes",
        function : "test",
        parameters : {
          "msg":"吃饭没～"
        }
      }
      axios.post("/api/distributor",data).then(res => {
        console.log(res)
      })
    },

    addOption() {
      this.form.selectOptions.push(this.form.selectOption)
      this.form.goodOptions = this.removeElement(this.form.selectOption)
      if (this.form.goodOptions.length !== 0) {
        this.form.selectOption = this.form.goodOptions[0]
      } else {
        this.form.selectOption = ""
      }
    },
    removeElement(elementName) {
      let index = this.form.goodOptions.indexOf(elementName)
      let arr1 = this.form.goodOptions.splice(0, index)
      let arr2 = this.form.goodOptions.splice(index + 1)
      return arr1.concat(arr2)
    }
  }
}
</script>

<style>
</style>
