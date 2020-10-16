<!--eslint-disable-->
<template>
  <el-form ref="form" :model="form" label-width="80px">
    <el-form-item label="选项类">
      <el-input v-model="form.optionClassName"></el-input>
    </el-form-item>
    <el-form-item label="选项">

      <template v-for="i in form.options.length">
        <el-row type="flex">
          <el-input v-model="form.options[i-1]"></el-input>
          <el-button type="danger" plain icon="el-icon-circle-close" @click="delOption(i-1)">删除选项</el-button>
        </el-row>
      </template>
      <el-button icon="el-icon-circle-plus" plain size="mini" type="primary" @click='form.options.push("")'>添加选项</el-button>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="addOptionClass">添加选项类</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
/* eslint-disable */
import axios from 'axios';

export default {
  name: 'GoodAdder',
  data() {
    return {
      form: {
        optionClassName: "",
        options: [""]
      }
    }
  },
  methods: {
    delOption(i) {
      this.form.options = this.form.options.slice(0,i).concat(this.form.options.slice(i+1))
    },
    addOptionClass() {
      let data = {
        object: "option",
        function: "add",
        parameters: this.form
      }
      console.log(data)
      axios.post("/api/distributor", data).then(res => {
        console.log(res)
      })
    },
  }
}
</script>
