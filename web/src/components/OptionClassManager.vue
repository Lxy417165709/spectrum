<!--eslint-disable-->
<template>
  <div>
    <el-table
      ref="multipleTable"
      :data="optionClasses"
      style="width: 100%"
      tooltip-effect="dark"
      @selection-change="handleSelectionChange"
    >
      <el-table-column
        type="selection"
        width="55">
      </el-table-column>
      <el-table-column
        label="选项类名"
        prop="className"
        show-overflow-tooltip>
        <template slot-scope="props">
          <el-input size="mini" v-model="props.row.className" v-if="props.row.edit === true"></el-input>
          <template v-else>{{props.row.className}}</template>
        </template>
      </el-table-column>
      <el-table-column
        label="选项名">
        <template slot-scope="props">
          <el-tag  v-for="(soc,idx) in props.row.optionNames" closable type="success" @close="delOption(props.row.edit,props.$index,idx)" :key="idx" style="margin-right:2px">
            <template v-if="props.row.edit === true" >
              <el-input size="mini" placeholder="请输入选项" v-model="props.row.optionNames[idx]">
              </el-input>
            </template>
            <template v-else>{{ soc }}</template>
          </el-tag>
          <el-button type="primary" icon="el-icon-plus" circle v-if="props.row.edit === true" @click="addOption(props.$index)"></el-button>
          <el-button type="success" icon="el-icon-check" circle v-if="props.row.edit === true" @click="saveOption(props.$index)"></el-button>
        </template>
      </el-table-column>
    </el-table>
    <div style="padding-top: 20px">
      <el-button v-if="selections.length === 0" type="primary" plain @click="addOptionClass">添加选项类</el-button>
      <el-button v-if="selections.length !== 0" type="danger" plain @click="delOptionClasses">删除选项类</el-button>
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
      optionClasses: [],
      selections : [],
    }
  },
  async mounted() {
    await init.globalOptionClasses()
    console.log(global.optionClasses)
    this.optionClasses = global.optionClasses
    for (let i=0;i<this.optionClasses.length;i++){
      this.optionClasses[i].edit = false
    }
  },
  methods: {
    handleClick(soc) {
      console.log(soc)
      soc.editing = true
    },
    addOption(index) {
      console.log(index)
      this.optionClasses[index].optionNames.push("")
    },
    saveOption(index) {
      console.log(this.optionClasses[index])
      let model = utils.getRequestModel("mvp", "AddOptionClass", this.optionClasses[index])
      utils.sendRequestModel(model).then(res => {
        if (!utils.hasRequestSuccess(res)) {
          this.$message.error(res.data.err)
          return
        }
        this.$message.success(res.data.msg)
      })
    },
    handleSelectionChange(val){
      this.selections = val
    },
    addOptionClass(){
      this.optionClasses.push({
        className:"",
        optionNames:[],
        edit:true,
      })
    },
    delOptionClasses() {
      let optionClassNames = []
      for (let i = 0;i<this.selections.length;i++) {
        optionClassNames.push(this.selections[i].className)
      }
      global.optionClasses[0].className
      let model = utils.getRequestModel("mvp","DelOptionClass",{
        "optionClassNames":optionClassNames
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
    delOption(editing,optionClassIndex, optionIndex) {
      console.log("par",editing,optionClassIndex, optionIndex)
      if (editing === true) {
        this.optionClasses[optionClassIndex].optionNames = this.optionClasses[optionClassIndex].optionNames
          .slice(0, optionIndex)
          .concat(this.optionClasses[optionClassIndex].optionNames.slice(optionIndex + 1))
        return
      }
      let model = utils.getRequestModel("mvp", "DelOption", {
        "className": global.optionClasses[optionClassIndex].className,
        "optionName": this.optionClasses[optionClassIndex].optionNames[optionIndex]
      })
      utils.sendRequestModel(model).then(async res => {
        if (!utils.hasRequestSuccess(res)) {
          this.$message.error(res.data.err)
          return
        }
        this.$message.success(res.data.msg)
        await init.globalOptionClasses()
        this.optionClasses = global.optionClasses
      })


    },
    handleClose(optionClass) {
      this.allOptionClasses = utils.removeElement(this.allOptionClasses, optionClass)
    },
  }
}
</script>

<style scoped>

</style>
