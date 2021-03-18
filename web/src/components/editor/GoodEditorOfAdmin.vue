<!-- eslint-disable -->
<template>
  <el-form ref="form" label-width="80px">

    <!-- 1. 商品名编辑器 -->
    <el-form-item label="商品名">
      <el-input style="width: 70%" v-if="good.mainElement!==undefined" v-model="good.mainElement.name"></el-input>
    </el-form-item>

    <!-- 2. 规格编辑器 -->
    <el-tabs type="border-card" @tab-click="" editable @edit="handleTabsEdit"
             @tab-add="handleClick" style="margin-bottom: 10px" v-if="good.mainElement!==undefined">
      <el-tab-pane v-for="(sizeInfo,index) in good.mainElement.sizeInfos" :label="sizeInfo.size"
                   :name="index.toString()"
                   :key="index">
        <el-form label-width="80px">
          <el-form-item label="规格名">
            <el-input v-model="sizeInfo.size" style="width: 70%"></el-input>
          </el-form-item>
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
          <el-form-item label="默认选中" v-if="index!==good.mainElement.selectedIndex">
            <el-button @click="handleChangeDefaultSizeInfo(index)">确定</el-button>
          </el-form-item>
        </el-form>
      </el-tab-pane>
    </el-tabs>

    <!-- 3. 附属选项编辑器 -->
    <el-form-item label="附属选项">
      <el-select v-model="curGoodOptionName" placeholder="附属选项">
        <el-option v-for="(element,index) in selectableElements" :key="index" v-if="element.type===1"
                   :label="element.name" :value="element.name"></el-option>
      </el-select>
      <el-button type="primary" @click="addGoodOption()">添加</el-button>
    </el-form-item>

    <el-form-item label="已选">
      <el-tag v-for="(element,index) in good.attachElements" :key="index" closable style="margin-right: 5px"
              v-if="element.type===1"
              type="success">
        {{ element.name }}
        <!--        @close="delSelectGoodClass(index)"-->
      </el-tag>
    </el-form-item>

    <!--     4. 附属配料编辑器 -->
    <el-form-item label="附属配料">
      <el-select v-model="curGoodIngredientName" placeholder="附属配料">
        <el-option v-for="(element,index) in selectableElements" :key="index" v-if="element.type===2"
                   :label="element.name" :value="element.name"></el-option>
      </el-select>
      <el-button type="primary" @click="addGoodIngredient()">添加</el-button>
    </el-form-item>

    <el-form-item label="已选">
      <el-tag v-for="(element,index) in good.attachElements" :key="index" closable style="margin-right: 5px"
              v-if="element.type===2"
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
import utils from "../../common/utils";

export default {
  name: "GoodEditorOfAdmin",
  components: {},
  async mounted() {
    await utils.GetAllGoodOptions(this, {}, (res) => {
      this.selectableElements = res.data.data.elements
    })
  },
  props: {
    className: String,
  },
  data() {
    return {
      good: {},
      curGoodOptionName: "",
      curGoodIngredientName: "",

      selectableElements: [],
      addTabCount: 0,
    }
  },
  methods: {
    addGoodOption() {
      for (let i = 0; i < this.selectableElements.length; i++) {
        if (this.selectableElements[i].name === this.curGoodOptionName) {
          this.good.attachElements.push(this.selectableElements[i])
          break
        }
      }
    },
    addGoodIngredient() {
      for (let i = 0; i < this.selectableElements.length; i++) {
        if (this.selectableElements[i].name === this.curGoodIngredientName) {
          this.good.attachElements.push(this.selectableElements[i])
        }
      }
    },
    handleClick(tab, event) {
      this.addTabCount++
      // todo: name 应该是可以编辑的
      let name = "未设定规格" + this.addTabCount
      this.good.mainElement.sizeInfos.push(utils.NewBlankSizeInfo(name))
    },
    handleTabsEdit(name, event) {
      // this.good.mainElement.sizeInfos = utils.removeElementByField(this.good.mainElement.sizeInfos, "name", name)
    },
    handleChangeDefaultSizeInfo(index) {
      this.good.mainElement.selectedIndex = index
    },
    async addGood(good) {
      let model = utils.getRequestModel("mvp", "AddGood", {
        good: good,
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
