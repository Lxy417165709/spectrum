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

    <el-form-item label="类别">
      <el-select autocomplete="on" v-model="form.goodClassName" :filter-method="test" filterable placeholder="请选择">
        <el-option
          v-for="(goodClass,index) in this.goodClasses"
          :key="index"
          :label="goodClass.name"
          :value="goodClass.name">
        </el-option>
      </el-select>
    </el-form-item>

<!--    <el-form-item label="单卖品">-->
<!--      <el-switch v-model="form.type"></el-switch>-->
<!--    </el-form-item>-->
    <el-form-item v-show="form.type && optionClasses.length > 0" label="附属选项">
      <el-select v-model="selectOptionClassName" placeholder="请选择附属选项">
          <el-option :label="optionClass.name" :value="optionClass.name" v-for="(optionClass,index) in optionClasses"></el-option>
      </el-select>
      <el-button type="primary" @click="addOptionClass">添加</el-button>
    </el-form-item>
    <el-form-item v-show="form.type && selectOptionClasses.length > 0" label="已选选项">
      <el-tag v-for="(selectOptionClass,index) in selectOptionClasses" closable type="success" @close="delSelectOption(index)" :key="index">
        {{ selectOptionClass.name }}
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
    await init.globalGoodClasses()
    // for (let i = 0; i < global.optionClasses.length; i++) {
    //   this.optionClassNames.push(global.optionClasses[i].className)
    // }
    // this.setDefaultSelectOption()
    this.optionClasses = global.optionClasses
    this.goodClasses = global.goodClasses
  },
  data() {
    return {
      form: {
        price: '',
        name: '',
        type: true,
        goodClassName: ""
      },
      dialogImageUrl: '',
      dialogVisible: false,

      pictureStorePath: '',




      optionClasses: [],
      goodClasses: [],
      selectOptionClasses: [],
      selectOptionClassName: "",
    }
  },
  methods: {
    test(val) {
      console.log(val,this.goodClasses[this.goodClasses.length-1].name)
    },
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
      this.form.pictureStorePath = res.data.fileStorePath;
      console.log(this.form.pictureStorePath)
    },
    handlePictureCardPreview(file) {
      this.dialogImageUrl = file.url;
      this.dialogVisible = true;
    },
    addGood() {
      let model = utils.getRequestModel("mvp", "AddGood", {
        name: this.form.name,
        price: parseFloat(this.form.price),
        type: this.form.type === true ? 0 : 1,
        pictureStorePath: this.form.pictureStorePath,
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
    addOptionClass() {
      this.selectOptionClasses.push({
        name:this.selectOptionClassName,
      })
      // this.optionClasses = utils.removeElement(this.optionClasses,this.selectOptionClass)
      // this.setDefaultSelectOption()
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
