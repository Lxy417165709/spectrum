<!-- eslint-disable -->
<template>
  <el-form ref="form" label-width="80px">
    <el-form-item label="商品名">{{ good.name }}</el-form-item>


    <!--    todo: 这里要提出一个小组件 -->
    <el-form-item label="照片">
      <el-image :src="'api/file/' + good.pictureStorePath"
                style="width: 100px; height: 100px;padding-top: 15px"></el-image>
    </el-form-item>


    <el-form-item label="价格">{{ good.price }}</el-form-item>

    <el-form-item label="附属选项">
      <el-select v-model="selectableElement.curAttachOptionName" placeholder="附属选项">
        <el-option v-for="(element,index) in selectableElement.attachElements" :key="index"
                   v-show="element.elementType=== 1"
                   :label="element.name" :value="element.name"></el-option>
      </el-select>
      <el-button type="primary">添加</el-button>
    </el-form-item>
    <el-form-item label="已选">
      <el-tag v-for="(element,index) in good.attachElements" :key="index" closable style="margin-right: 5px"
              type="success" v-show="element.elementType=== 1">
        {{ element.name }}
        <!--        @close="delSelectGoodClass(index)"-->
      </el-tag>
    </el-form-item>


    <el-form-item label="附属配料">
      <el-select v-model="selectableElement.curAttachGoodName" placeholder="附属配料">
        <el-option v-for="(element,index) in selectableElement.attachElements" :key="index"
                   v-show="element.elementType=== 2"
                   :label="element.name" :value="element.name"></el-option>
      </el-select>
      <el-button type="primary">添加</el-button>
    </el-form-item>

    <el-form-item label="已选">
      <el-tag v-for="(element,index) in good.attachElements" :key="index" closable style="margin-right: 5px"
              type="success" v-show="element.elementType=== 2">
        {{ element.name }}
      </el-tag>
    </el-form-item>

  </el-form>
</template>

<script>
/* eslint-disable */
export default {
  name: "GoodEditor",
  data() {
    return {
      good: {
        name: "波霸奶茶",
        price: 18.0,
        attachElements: [
          {
            name: "冰量",
            elementType: 1,
          },
          {
            name: "温度",
            elementType: 1,
          },
          {
            name: "珍珠",
            elementType: 2,
          }
        ]
      },
      selectableElement: {
        curAttachOptionName: "",
        curAttachGoodName: "",

        attachElements: [
          {
            name: "冰量",
            elementType: 1,
          },
          {
            name: "温度",
            elementType: 1,
          },
          {
            name: "珍珠",
            elementType: 2,
          }
        ]
      }


    }
  }
}
</script>

<style scoped>

</style>
