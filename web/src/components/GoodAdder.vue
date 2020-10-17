<!--eslint-disable-->


<template>
  <el-form ref="form" :model="form" label-width="80px">
    <el-form-item label="照片">
      <el-upload
        :on-preview="handlePictureCardPreview"
        :on-remove="handleRemove"
        :on-success="handleSuccess"
        action="/api/upload"
        list-type="picture-card">
        <i class="el-icon-plus"></i>
      </el-upload>
    </el-form-item>
    <el-form-item label="商品名">
      <el-input v-model="form.name"></el-input>
    </el-form-item>
    <el-form-item label="价格">
      <el-input v-model="form.price"></el-input>
    </el-form-item>
    <el-form-item label="单卖品">
      <el-switch v-model="form.type"></el-switch>
    </el-form-item>
    <el-form-item v-show="form.type && optionClassNames.length > 0" label="附属选项">
      <el-select v-model="form.selectOptionClass" placeholder="请选择附属选项">
        <template v-for="gp in optionClassNames">
          <el-option :label="gp" :value="gp"></el-option>
        </template>
      </el-select>
      <el-button type="primary" @click="addOption">添加</el-button>
    </el-form-item>
    <el-form-item v-show="form.type && form.selectOptionClasses.length > 0" label="已选选项">
      <el-tag v-for="soc in form.selectOptionClasses" closable type="success" @close="delSelectOption(soc)">
        {{ soc }}
      </el-tag>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="addGood">立即创建</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
/* eslint-disable */
import axios from 'axios';
import utils from '../common/utils'
import global from "../common/global_object/global"
import init from "../common/global_object/init";

export default {
  name: 'GoodAdder',
  async mounted() {
    await init.globalOptionClasses()
    for (let i = 0; i < global.optionClasses.length; i++) {
      this.optionClassNames.push(global.optionClasses[i].className)
    }
    this.setDefaultSelectOption()
  },
  data() {
    return {
      form: {
        selectOptionClass: '',
        goodOptionClasses: [],
        selectOptionClasses: [],
        price: '',
        name: '',
        type: true,
      },
      dialogImageUrl: '',
      dialogVisible: false,
      optionClassNames: []
    }
  },
  methods: {
    setDefaultSelectOption() {
      if (this.optionClassNames.length !== 0) {
        this.form.selectOptionClass = this.optionClassNames[0]
      } else {
        this.form.selectOptionClass = ""
      }
    },
    delSelectOption(optionName) {
      this.form.selectOptionClasses = utils.removeElement(this.form.selectOptionClasses, optionName)
      this.optionClassNames.push(optionName)
      this.setDefaultSelectOption()
    },
    handleRemove(file, fileList) {
      console.log(file, fileList);
    },
    handleSuccess(res) {
      if (!utils.hasRequestSuccess(res)) {
        this.$message.error(res.err)
        return
      }
      this.$message.success(res.msg)
      console.log(res.data.fileStorePath)
    },
    handlePictureCardPreview(file) {
      this.dialogImageUrl = file.url;
      this.dialogVisible = true;
    },
    addGood() {
      let model = utils.getRequestModel("mvp", "AddGood", {
        name: this.form.name,
        price: parseFloat(this.form.price),
        type: this.form.type === true? 0 : 1,
        pictureStorePath:"test",
        optionClassNames: this.form.selectOptionClasses,
      })
      utils.sendRequestModel(model).then(res => {
        if (!utils.hasRequestSuccess(res)) {
          this.$message.error(res.data.err)
          return
        }
        this.$message.success(res.data.msg)
      })
    },
    addOption() {
      this.form.selectOptionClasses.push(this.form.selectOptionClass)
      this.optionClassNames = utils.removeElement(this.optionClassNames, this.form.selectOptionClass)
      this.setDefaultSelectOption()
    },
  }
}
</script>

<style>
.bg-purple-dark {
  background: #ee625e;
}

.grid-content {
  border-radius: 4px;
  color: white;
  text-align: center; /*让文字水平居中*/
}

</style>
