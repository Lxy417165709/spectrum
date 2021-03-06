<!-- eslint-disable -->
<template>
  <el-form ref="form" label-width="80px">

    <el-form-item label="商品名">
      <el-input style="width: 70%" v-model="good.name"></el-input>
    </el-form-item>
    <good-size-editor style="margin-bottom: 20px" ref="goodSizeEditor"
                      :originCurSizeIndex="good.curSizeIndex"
                      :originSizeInfos="good.sizeInfos"></good-size-editor>
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
      <el-button type="primary" @click="addGood(good)">确定</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
/* eslint-disable */
import GoodSizeEditor from "./GoodSizeEditor";
import test from "../../common/test/test";
import utils from "../../common/utils";
import global from "../../common/global_object/global";

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
  },
  methods: {
    async addGood(good) {
      // todo: 将前端的 good，转换为协议的 good
      let model = utils.getRequestModel("mvp", "AddGood", {
        good: {
          mainElement: {
            name: good.name,
          }
        },
        className: "test_class",
      })
      await utils.sendRequestModel(model).then(res => {
        if (!utils.hasRequestSuccess(res)) {
          console.log(res.data.err)
          return
        }
        if (utils.hasData(res.data.data.goods)) {
          global.goods = res.data.data.goods
        }
        console.log("global_goods", global.goods)
      })
    }
  }
}
</script>

<style scoped>

</style>
