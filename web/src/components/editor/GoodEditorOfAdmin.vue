<!-- eslint-disable -->
<template>
  <el-form ref="form" label-width="80px">

    <!-- 1. 商品名编辑器 -->
    <el-form-item label="商品名">
      <el-input style="width: 70%" v-model="good.name"></el-input>
    </el-form-item>

    <!-- 2. 规格编辑器 -->
    <el-tabs type="border-card" @tab-click="" editable @edit="handleTabsEdit"
             @tab-add="handleClick" style="margin-bottom: 10px">
      <el-tab-pane v-for="(sizeInfo,index) in good.sizeInfos" :label="sizeInfo.name" :name="index.toString()"
                   :key="index">
        <el-form label-width="80px">
          <el-form-item label="照片">
            <el-upload
              action="/api/upload"
              list-type="picture-card">
              <i class="el-icon-plus"></i>
            </el-upload>
          </el-form-item>
          <el-form-item label="价格">
            <el-input v-model="sizeInfo.price" style="width: 70%"></el-input>
          </el-form-item>
          <el-form-item label="默认选中" v-if="good.curSizeIndex!==index">
            <el-button @click="handleChangeDefaultSizeInfo(index)">确定</el-button>
          </el-form-item>
        </el-form>
      </el-tab-pane>
    </el-tabs>

    <!-- 3. 附属选项编辑器 -->
    <el-form-item label="附属选项">
      <el-select v-model="selectableElement.curAttachElementName" placeholder="附属选项">
        <el-option v-for="(element,index) in selectableElement.attachElements" :key="index"
                   :label="element.name" :value="element.name"></el-option>
      </el-select>
      <el-button type="primary" @click="addAttachElement()">添加</el-button>
    </el-form-item>
    <el-form-item label="已选">
      <el-tag v-for="(element,index) in good.attachElements" :key="index" closable style="margin-right: 5px"
              type="success">
        {{ element.name }}
        <!--        @close="delSelectGoodClass(index)"-->
      </el-tag>
    </el-form-item>

    <!-- 4. 附属配料编辑器 -->
    <el-form-item label="附属配料">
      <el-select v-model="selectableElement.curFavorName" placeholder="附属配料">
        <el-option v-for="(element,index) in selectableElement.favors" :key="index"
                   :label="element.name" :value="element.name"></el-option>
      </el-select>
      <el-button type="primary" @click="addFavor()">添加</el-button>
    </el-form-item>

    <el-form-item label="已选">
      <el-tag v-for="(element,index) in good.favors" :key="index" closable style="margin-right: 5px"
              type="success">
        {{ element.name }}
      </el-tag>
    </el-form-item>

    <!--    5. 提交按钮-->
    <el-form-item>
      <el-button type="primary" @click="addGood(good)">确定</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
/* eslint-disable */
import GoodSizeEditor from "./_GoodSizeEditor";
import test from "../../common/test/test";
import utils from "../../common/utils";
import global from "../../common/global_object/global";

export default {
  name: "GoodEditorOfAdmin",
  components: {GoodSizeEditor},
  mounted() {
    this.selectableElement = test.selectableElement
  },
  data() {
    return {
      good: {},
      className: "",

      selectableElement: {},
      addTabCount: 0,
    }
  },
  methods: {
    addAttachElement() {
      // todo: 添加后，应该把元素从可选列表中删除
      for (let i = 0; i < this.selectableElement.attachElements.length; i++) {
        if (this.selectableElement.attachElements[i].name === this.selectableElement.curAttachElementName) {
          this.good.attachElements.push(this.selectableElement.attachElements[i])
        }
      }
    },
    addFavor() {
      // todo: 添加后，应该把元素从可选列表中删除
      for (let i = 0; i < this.selectableElement.favors.length; i++) {
        if (this.selectableElement.favors[i].name === this.selectableElement.curFavorName) {
          this.good.favors.push(this.selectableElement.favors[i])
        }
      }
    },
    handleClick(tab, event) {
      this.addTabCount++
      // todo: name 应该是可以编辑的
      let name = "未设定规格" + this.addTabCount
      // todo: sizeInfo 应该是一个对象，要有构造函数
      this.good.sizeInfos.push({
        name: name,
        price: 30
      })
    },
    handleTabsEdit(name, event) {
      this.good.sizeInfos = utils.removeElementByField(this.good.sizeInfos, "name", name)
    },
    handleChangeDefaultSizeInfo(index) {
      this.good.curSizeIndex = index
    },
    async addGood(good) {
      let model = utils.getRequestModel("mvp", "AddGood", {
        good: utils.goodToPbGood(good),
        className: this.className,
      })
      await utils.sendRequestModel(model).then(res => {
        if (!utils.hasRequestSuccess(res)) {
          this.$message.error(res.data.err)
          return
        }
        this.$message.success(res.data.msg)
      })
    }
  }
}
</script>

<style scoped>

</style>
