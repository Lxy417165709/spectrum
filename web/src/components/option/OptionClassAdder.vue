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
          <el-button icon="el-icon-circle-close" plain type="danger" @click="delOption(i-1)">删除选项</el-button>
        </el-row>
      </template>
      <el-button icon="el-icon-circle-plus" plain size="mini" type="primary" @click='form.options.push("")'>添加选项
      </el-button>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="addOptionClass">添加选项类</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
/* eslint-disable */
import utils from "../../common/utils";

export default {
  name: 'GoodAdder',
  data() {
    return {
      form: {
        optionClassName: "",
        options: []
      }
    }
  },
  methods: {
    delOption(i) {
      this.form.options = this.form.options.slice(0, i).concat(this.form.options.slice(i + 1))
    },
    addOptionClass() {
      let model = utils.getRequestModel("mvp", "AddOptionClass", this.form)
      utils.sendRequestModel(model).then(res => {
        if (!utils.hasRequestSuccess(res)) {
          this.$message.error(res.data.err)
          return
        }
        this.$message.success(res.data.msg)
      })
    },
  }
}
</script>
