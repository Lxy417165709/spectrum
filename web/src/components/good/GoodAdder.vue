<!--eslint-disable-->


<template>
  <el-form ref="form" :model="form" label-width="120px">
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



    <el-form-item label="附属品">
      <el-switch v-model="isAttachGood"></el-switch>
    </el-form-item>

    <el-form-item label="类别">
      <el-select autocomplete="on" v-model="form.goodClassName" filterable placeholder="请选择">
        <template v-if="isAttachGood === false">
          <el-option
            v-for="(goodClass,index) in this.goodClasses"
            :key="index"
            :label="goodClass.name"
            :value="goodClass.name">
          </el-option>
        </template>
        <template v-else>
          <el-option
            v-for="(goodClass,index) in this.attachGoodClasses"
            :key="index"
            :label="goodClass.name"
            :value="goodClass.name">
          </el-option>
        </template>
      </el-select>
    </el-form-item>
    <template v-if="isAttachGood === false">
      <el-form-item v-show="optionClasses.length > 0" label="附属选项类">
        <el-select v-model="selectOptionClassName" placeholder="请选择附属选项类">
          <el-option :label="optionClass.name" :value="optionClass.name" v-for="(optionClass,index) in optionClasses" :key="index"></el-option>
        </el-select>
        <el-button type="primary" @click="addOptionClass">添加</el-button>
      </el-form-item>
      <el-form-item v-show="selectOptionClasses.length > 0" label="已选选项类">
        <el-tag v-for="(selectOptionClass,index) in selectOptionClasses" closable type="success" @close="delSelectOption(index)" :key="index">
          {{ selectOptionClass.name }}
        </el-tag>
      </el-form-item>

      <el-form-item v-show="attachGoodClasses.length > 0" label="附属产品类">
        <el-select v-model="selectAttachGoodClassName" placeholder="附属产品类选项">
          <el-option :label="attachGoodClass.name" :value="attachGoodClass.name" v-for="(attachGoodClass,index) in attachGoodClasses" :key="index"></el-option>
        </el-select>
        <el-button type="primary" @click="addAttachGoodClass">添加</el-button>
      </el-form-item>
      <el-form-item v-show="selectAttachGoodClasses.length > 0" label="已选附属产品类">
        <el-tag v-for="(selectAttachGoodClass,index) in selectAttachGoodClasses" closable type="success" @close="delSelectGoodClass(index)" :key="index">
          {{ selectAttachGoodClass.name }}
        </el-tag>
      </el-form-item>
    </template>

    <el-form-item>
      <el-button type="primary" @click="addGood">立即创建</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
/* eslint-disable */
import axios from 'axios';
import utils from '../../common/utils'
import global from "../../common/global_object/global"
import init from "../../common/global_object/init";

export default {
  name: 'GoodAdder',
  async mounted() {
    await init.globalOptionClasses()
    await init.globalGoodClasses()
    this.optionClasses = global.optionClasses
    for (let i=0;i<global.goodClasses.length;i++){
      if (global.goodClasses[i].classType===undefined || global.goodClasses[i].classType === 0) {
        this.goodClasses.push(global.goodClasses[i])
      }else{
        this.attachGoodClasses.push(global.goodClasses[i])
      }
    }
  },
  data() {
    return {
      form: {
        price: '',
        name: '',
        goodClassName: ""
      },
      dialogImageUrl: '',
      dialogVisible: false,
      isAttachGood:false,
      pictureStorePath: '',
      attachGoodClasses: [],



      optionClasses: [],
      goodClasses: [],
      selectOptionClasses: [],
      selectOptionClassName: "",
      selectAttachGoodClassName:"",
      selectAttachGoodClasses: [],
    }
  },
  methods: {
    setDefaultSelectOption() {
      if (this.optionClasses.length !== 0) {
        this.selectOptionClassName = this.optionClasses[0].name
      } else {
        this.selectOptionClassName = {name: ""}
      }
    },
    delSelectGoodClass(index) {
      this.attachGoodClasses.push(this.selectAttachGoodClasses[index])
      this.selectAttachGoodClasses = utils.removeIndex(this.selectAttachGoodClasses,index)
      // this.setDefaultSelectOption()
    },
    delSelectOption(index) {
      this.optionClasses.push(this.selectOptionClasses[index])
      this.selectOptionClasses = utils.removeIndex(this.selectOptionClasses,index)
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
        good:{
          name: this.form.name,
          price: parseFloat(this.form.price),
          pictureStorePath: this.form.pictureStorePath,
          optionClasses: this.selectOptionClasses,
          attachGoodClasses:this.selectAttachGoodClasses,
        },
        goodClassName: this.form.goodClassName,
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
        name: this.selectOptionClassName,
      })
      this.optionClasses = utils.removeElementByField(this.optionClasses, "name", this.selectOptionClassName)
      this.setDefaultSelectOption()
    },
    addAttachGoodClass() {
      this.selectAttachGoodClasses.push({
        name: this.selectAttachGoodClassName,
      })
      this.attachGoodClasses = utils.removeElementByField(this.attachGoodClasses, "name", this.selectAttachGoodClassName)
    }
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
