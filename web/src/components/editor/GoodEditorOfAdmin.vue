<!-- eslint-disable -->
<template>
  <el-form ref="form" label-width="80px">

    <!-- 1. 商品名编辑器 -->
    <el-form-item label="商品名">
      <el-input style="width: 70%" v-if="good.mainElement!==undefined" v-model="good.mainElement.name"></el-input>
    </el-form-item>

    <!-- 2. 规格编辑器 -->
    <el-tabs type="border-card" @tab-click="tabClick" addable @tab-add="handleClick" style="margin-bottom: 10px"
             v-if="good.mainElement!==undefined">
      <el-tab-pane v-for="(sizeInfo,index) in good.mainElement.sizeInfos" :label="sizeInfo.size"
                   :name="index.toString()"
                   :key="index">
        <el-form label-width="80px">
          <el-form-item label="规格名">
            <el-input v-model="sizeInfo.size" style="width: 70%"></el-input>
          </el-form-item>

          <el-form-item label="照片">
            <el-upload
              v-if="sizeInfo.pictureStorePath===''"
              action="/api/upload"
              :on-success="imageUploadSuccess"
              list-type="picture-card">

              <i class="el-icon-plus"></i>
            </el-upload>
            <el-image
              v-if="sizeInfo.pictureStorePath!==''"
              @dblclick.native="cleanSizeInfoPictureStorePath"
              :src="'api/file/' + sizeInfo.pictureStorePath"
              style="width: 148px; height: 148px;"></el-image>
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
        <el-option v-for="(element,index) in selectableElements" :key="index" v-if="isGoodOption(element)"
                   :label="element.name" :value="element.name"></el-option>
      </el-select>
      <el-button type="primary" @click="addGoodOption()">添加</el-button>
    </el-form-item>

    <el-form-item label="已选" v-if="cpt_canSelectedGoodOptionShow">
      <el-tag v-for="(element,index) in good.attachElements" :key="index" closable style="margin-right: 5px"
              v-if="isGoodOption(element)"
              type="success">
        {{ element.name }}
      </el-tag>
    </el-form-item>

    <!--     4. 附属配料编辑器 -->
    <el-form-item label="附属配料">
      <el-select v-model="curGoodIngredientName" placeholder="附属配料">
        <el-option v-for="(element,index) in selectableElements" :key="index" v-if="isGoodIngredient(element)"
                   :label="element.name" :value="element.name"></el-option>
      </el-select>
      <el-button type="primary" @click="addGoodIngredient()">添加</el-button>
    </el-form-item>

    <el-form-item label="已选" v-if="cpt_canSelectedGoodIngredientShow">
      <el-tag v-for="(element,index) in good.attachElements" :key="index" closable style="margin-right: 5px"
              v-if="isGoodIngredient(element)"
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
import cst from "../../common/cst";

export default {
  name: "GoodEditorOfAdmin",
  components: {},
  async mounted() {
    await utils.GetAllGoodOptions(this, {}, (res) => {
      this.selectableElements = res.data.data.elements
      for (let i = 0; i < this.selectableElements.length; i++) {
        if (this.curGoodOptionName === "" && this.isGoodOption(this.selectableElements[i])) {
          this.curGoodOptionName = this.selectableElements[i].name;
        }
        if (this.curGoodIngredientName === "" && this.isGoodIngredient(this.selectableElements[i])) {
          this.curGoodIngredientName = this.selectableElements[i].name;
        }
      }
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
      curSizeInfoIndex: 0,

      selectableElements: [],
      addTabCount: 0,

      canModifyGoodName: false,
    }
  },
  methods: {
    addGoodOption() {
      for (let i = 0; i < this.selectableElements.length; i++) {
        if (this.selectableElements[i].name !== this.curGoodOptionName) {
          continue
        }
        if (this.good.attachElements === null) {
          this.good.attachElements = []
        }
        if (!utils.isExist(this.good.attachElements, "name", this.selectableElements[i].name)) {
          this.good.attachElements.push(this.selectableElements[i])
        }
      }
    },
    addGoodIngredient() {
      for (let i = 0; i < this.selectableElements.length; i++) {
        if (this.selectableElements[i].name !== this.curGoodIngredientName) {
          continue
        }
        if (this.good.attachElements === null) {
          this.good.attachElements = []
        }
        if (!utils.isExist(this.good.attachElements, "name", this.selectableElements[i].name)) {
          this.good.attachElements.push(this.selectableElements[i])
        }
        break
      }
    },
    handleClick(tab, event) {
      this.addTabCount++
      let name = "未设定规格" + this.addTabCount
      this.good.mainElement.sizeInfos.push(utils.NewBlankSizeInfo(name))
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
    },
    tabClick(tab) {
      this.curSizeInfoIndex = tab.index
    },
    cleanSizeInfoPictureStorePath() {
      this.good.mainElement.sizeInfos[this.curSizeInfoIndex].pictureStorePath = ""
    },
    imageUploadSuccess(res, file, fileList) {
      this.good.mainElement.sizeInfos[this.curSizeInfoIndex].pictureStorePath = res.data.fileStorePath;
    },
    isGoodOption(element) {
      return element.type === cst.ELEMENT_TYPE.OPTION
    },
    isGoodIngredient(element) {
      return element.type === cst.ELEMENT_TYPE.INGREDIENT
    }
  },
  computed: {
    cpt_canSelectedGoodOptionShow() {
      if (this.good === undefined || this.good.attachElements === null || this.good.attachElements === undefined || this.good.attachElements.length === 0) {
        return false
      }
      return utils.isExist(this.good.attachElements, "type", cst.ELEMENT_TYPE.OPTION)
    },
    cpt_canSelectedGoodIngredientShow() {
      if (this.good === undefined || this.good.attachElements === null || this.good.attachElements === undefined || this.good.attachElements.length === 0) {
        return false
      }
      return utils.isExist(this.good.attachElements, "type", cst.ELEMENT_TYPE.INGREDIENT)
    }
  }
}
</script>

<style scoped>

</style>
