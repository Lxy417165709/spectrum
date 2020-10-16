<!--eslint-disable-->
<template>
  <el-form ref="form" :model="form" label-width="80px">
    <el-form-item label="照片">
      <el-upload
        action="https://jsonplaceholder.typicode.com/posts/"
        list-type="picture-card"
        :on-preview="handlePictureCardPreview"
        :on-remove="handleRemove">
        <i class="el-icon-plus"></i>
      </el-upload>
    </el-form-item>
    <el-form-item label="商品名">
      <el-input v-model="form.name"></el-input>
    </el-form-item>
    <el-form-item label="价格">
      <el-input v-model="form.price"></el-input>
    </el-form-item>
    <el-form-item v-show="form.goodOptionClasses.length > 0" label="附属选项">
      <el-select v-model="form.selectOptionClass" placeholder="请选择附属选项">
        <template v-for="gp in form.goodOptionClasses">
          <el-option :label="gp" :value="gp"></el-option>
        </template>
      </el-select>
      <el-button type="primary" @click="addOption">添加</el-button>
    </el-form-item>
    <el-form-item v-show="form.selectOptionClasses.length > 0" label="已选选项">
      <el-row :gutter="10" type="flex">
        <template v-for="gp in form.selectOptionClasses">
          <el-col :span="3"  size="mini"><div class="grid-content bg-purple-dark">{{gp}}</div></el-col>
        </template>
      </el-row>

    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">立即创建</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
/* eslint-disable */
import axios from 'axios';

export default {
  name: 'GoodAdder',
  mounted() {
    this.initAllOptionClasses()
  },
  data() {
    return {
      form: {
        selectOptionClass: '',
        goodOptionClasses: [],
        selectOptionClasses: [],
        price: '',
        name: '',
      },
      dialogImageUrl: '',
      dialogVisible: false
    }
  },
  methods: {
    handleRemove(file, fileList) {
      console.log(file, fileList);
    },
    handlePictureCardPreview(file) {
      this.dialogImageUrl = file.url;
      this.dialogVisible = true;
    },
    onSubmit() {
      let data = {
        object:"mvp",
        function : "AddGood",
        parameters : this.form
      }
      axios.post("/api/distributor",data).then(res => {
        console.log(res)
      })
    },
    initAllOptionClasses() {
      let data = {
        object:"mvp",
        function : "GetAllOptionClasses",
        parameters : {}
      }
      axios.post("/api/distributor", data).then(res => {
        let returnResult = res.data
        if (returnResult.err === "" || returnResult.err === undefined || returnResult.err === null) {
          this.form.goodOptionClasses = returnResult.data.optionClassNames
          this.$message.success('选项类获取成功!')
        }else{
          this.$message.error(returnResult.err)
        }
      })
    },
    addOption() {
      this.form.selectOptionClasses.push(this.form.selectOptionClass)
      this.form.goodOptionClasses = this.removeElement(this.form.goodOptionClasses,this.form.selectOptionClass)
      if (this.form.goodOptionClasses.length !== 0) {
        this.form.selectOptionClass = this.form.goodOptionClasses[0]
      } else {
        this.form.selectOptionClass = ""
      }
    },
    removeElement(array,element) {
      let result = [];
      for (let i = 0;i<array.length;i++){
        if (array[i]!==element){
          result.push(array[i])
        }
      }
      return result
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
  text-align: center;/*让文字水平居中*/
}

</style>
