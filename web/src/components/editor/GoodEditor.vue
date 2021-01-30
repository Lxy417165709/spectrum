<!-- eslint-disable -->
<template>
  <el-form ref="form" label-width="80px">
    <el-form-item label="商品名">{{ good.name }}</el-form-item>
    <!--    todo: 这里会报错，因为子组件会修改 curSizeIndex..-->
    <good-size-editor style="margin-bottom: 20px" ref="goodSizeEditor" :curSizeIndex="good.curSizeIndex"
                      :sizeInfos="good.sizeInfos"></good-size-editor>
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
    <el-form-item>
      <el-button type="primary">确定</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
/* eslint-disable */
import GoodSizeEditor from "./GoodSizeEditor";
import test from "../../common/test/test";

export default {
  name: "GoodEditor",
  components: {GoodSizeEditor},
  mounted() {
    this.selectableElement = test.selectableElement
  },
  data() {
    return {
      good: {},
      selectableElement: {}
    }
  }
}
</script>

<style scoped>

</style>
