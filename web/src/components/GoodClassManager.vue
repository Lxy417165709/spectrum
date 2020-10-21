<!-- eslint-disable -->
<template>
  <div>
    <el-table
      ref="multipleTable"
      :data="goodClasses"
      style="width: 100%"
      tooltip-effect="dark"
      @selection-change="handleElTableSelectionChange"
    >
      <el-table-column
        type="selection"
        width="55">
      </el-table-column>
      <el-table-column label="商品类名"  show-overflow-tooltip>
        <template slot-scope="props">
          <el-input v-if="props.row.editing === true" v-model="props.row.name" size="mini"></el-input>
          <template v-else>{{ props.row.name }}</template>
        </template>
      </el-table-column>
      <el-table-column label="商品名">
        <template slot-scope="props">
          <el-tag v-for="(good, goodIndex) in props.row.goods" :key="goodIndex" closable
                  style="margin-right:2px" type="success"
                  @close="delOption(props.row.editing,props.$index,goodIndex)">
            <template v-if="props.row.editing === true">
              <el-input v-model="good.name" placeholder="请输入选项" size="mini">
              </el-input>
            </template>
            <template v-else>{{ good.name }}</template>
          </el-tag>
          <template v-if="props.row.editing === true">
            <el-button circle icon="el-icon-plus" type="primary" @click="createGood(props.$index)"></el-button>
            <el-button circle icon="el-icon-check" type="success" @click="addGood(props.$index)"></el-button>
          </template>
        </template>
      </el-table-column>
    </el-table>
    <div style="padding-top: 20px">
      <el-button v-if="selectOptionClasses.length === 0" plain type="primary" @click="createGoodClass">添加商品类</el-button>
      <el-button v-else plain type="danger" @click="delGoodClasses">删除商品类</el-button>
    </div>
  </div>
</template>

<script>
/* eslint-disable */
import utils from "../common/utils"
import global from "../common/global_object/global"
import init from "../common/global_object/init"

export default {
  name: "OptionClassManager",
  data() {
    return {
      goodClasses: [],
      selectOptionClasses: [],
    }
  },
  async mounted() {
    await init.globalGoodClasses()
    this.goodClasses = global.goodClasses
    for (let i = 0; i < this.goodClasses.length; i++) {
      this.goodClasses[i].edit = false
    }
  },
  methods: {
    handleElTableSelectionChange(selectOptionClasses) {
      this.selectOptionClasses = selectOptionClasses;
    },
    createGood(index) {
      this.goodClasses[index].goods.push({
        name: ""
      })
    },
    handleSelectionChange(val) {
      this.selections = val
    },
    createGoodClass() {
      this.goodClasses.push({
        name: "",
        goods: [],
        editing: true,
      })
    },
    addGood(index) {
      console.log(
        "addGood_parameters",
        "index:", index,
        "goodClass:", this.goodClasses[index]
      )
      let model = utils.getRequestModel("mvp", "AddGoodClass", {
        goodClass: this.goodClasses[index]
      })
      utils.sendRequestModel(model).then(res => {
        if (!utils.hasRequestSuccess(res)) {
          this.$message.error(res.data.err)
          return
        }
        this.$message.success(res.data.msg)
        this.goodClasses[index].editing = false
      })
    },
    delOptionClasses() {
      let model = utils.getRequestModel("mvp", "DelOptionClass", {
        "optionClasses": this.selectOptionClasses
      })
      utils.sendRequestModel(model).then(async res => {
        if (!utils.hasRequestSuccess(res)) {
          this.$message.error(res.data.err)
          return
        }
        await init.globalOptionClasses()
        this.$message.success(res.data.msg)
        this.optionClasses = global.optionClasses
      })
    },
    delOption(editing, optionClassIndex, optionIndex) {
      console.log(
        "delOption_parameters",
        "editing:", editing,
        "optionClassIndex:", optionClassIndex,
        "optionIndex:", optionIndex
      )
      if (editing === true) {
        this.optionClasses[optionClassIndex].options = utils.removeIndex(
          this.optionClasses[optionClassIndex].options,
          optionIndex
        )
        return
      }
      let model = utils.getRequestModel("mvp", "DelOption", {
        "className": this.optionClasses[optionClassIndex].name,
        "optionName": this.optionClasses[optionClassIndex].options[optionIndex].name
      })
      utils.sendRequestModel(model).then(async res => {
        if (!utils.hasRequestSuccess(res)) {
          this.$message.error(res.data.err)
          return
        }
        this.$message.success(res.data.msg)
        await init.globalOptionClasses()
        this.optionClasses[optionClassIndex].options = utils.removeIndex(
          this.optionClasses[optionClassIndex].options,
          optionIndex
        )
      })
    },
  }
}
</script>
