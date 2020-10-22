<!-- eslint-disable -->
<template>
  <div>
    <el-table
        ref="multipleTable"
        :data="optionClasses"
        style="width: 100%"
        tooltip-effect="dark"
        @selection-change="handleElTableSelectionChange"
    >
      <el-table-column
          type="selection"
          width="55">
      </el-table-column>
      <el-table-column label="选项类名" prop="className" show-overflow-tooltip>
        <template slot-scope="props">
          <el-input v-if="props.row.editing === true" v-model="props.row.name" size="mini"></el-input>
          <template v-else>{{ props.row.name }}</template>
        </template>
      </el-table-column>
      <el-table-column label="选项名">
        <template slot-scope="props">
          <el-tag v-for="(option, optionIndex) in props.row.options" :key="optionIndex" closable
                  style="margin-right:2px" type="success"
                  @close="delOption(props.row.editing,props.$index,optionIndex)">
            <template v-if="props.row.editing === true">
              <el-input v-model="option.name" placeholder="请输入选项" size="mini">
              </el-input>
            </template>
            <template v-else>{{ option.name }}</template>
          </el-tag>
          <template v-if="props.row.editing === true">
            <el-button circle icon="el-icon-plus" type="primary" @click="createOption(props.$index)"></el-button>
            <el-button circle icon="el-icon-check" type="success" @click="addOption(props.$index)"></el-button>
          </template>
        </template>
      </el-table-column>
    </el-table>
    <div style="padding-top: 20px">
      <el-button v-if="selectOptionClasses.length === 0" plain type="primary" @click="createOptionClass">添加选项类</el-button>
      <el-button v-else plain type="danger" @click="delOptionClasses">删除选项类</el-button>
    </div>
  </div>
</template>

<script>
/* eslint-disable */
import utils from "../../common/utils";
import global from "../../common/global_object/global"
import init from "../../common/global_object/init"

export default {
  name: "OptionClassManager",
  data() {
    return {
      optionClasses: [],
      selectOptionClasses: [],
    }
  },
  async mounted() {
    await init.globalOptionClasses()
    this.optionClasses = global.optionClasses
    for (let i = 0; i < this.optionClasses.length; i++) {
      this.optionClasses[i].edit = false
    }
  },
  methods: {
    handleElTableSelectionChange(selectOptionClasses) {
      this.selectOptionClasses = selectOptionClasses;
    },
    createOption(index) {
      this.optionClasses[index].options.push({
        name: ""
      })
    },
    handleSelectionChange(val) {
      this.selections = val
    },
    createOptionClass() {
      this.optionClasses.push({
        name: "",
        options: [],
        editing: true,
      })
    },
    addOption(index) {
      console.log(
          "addOption_parameters",
          "index:", index,
          "optionClass:", this.optionClasses[index]
      )
      let model = utils.getRequestModel("mvp", "AddOptionClass", {
        optionClass: this.optionClasses[index]
      })
      utils.sendRequestModel(model).then(res => {
        if (!utils.hasRequestSuccess(res)) {
          this.$message.error(res.data.err)
          return
        }
        this.$message.success(res.data.msg)
        this.optionClasses[index].editing = false
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
