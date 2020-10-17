<!--eslint-disable-->
<template>
  <div>
    <el-table
      ref="multipleTable"
      :data="optionClasses"
      style="width: 100%"
      tooltip-effect="dark"
      type="index">
      <el-table-column
        type="selection"
        width="55">
      </el-table-column>
      <el-table-column
        label="选项类名"
        prop="className"
        show-overflow-tooltip>
      </el-table-column>
      <el-table-column
        label="选项名">
        <template slot-scope="props">
          <el-tag v-for="soc in props.row.optionNames" closable type="success" @close="delOption(props.$index,soc)">
            {{ soc }}
          </el-tag>
        </template>
      </el-table-column>
    </el-table>
    <div style="margin-top: 20px">
      <el-button>切换第二、第三行的选中状态</el-button>
      <el-button>取消选择</el-button>
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
  watch: {
    optionClasses(n, o) {
      global.optionClasses = n
    }
  },
  data() {
    return {
      optionClasses: [],
    }
  },
  async mounted() {
    await init.globalOptionClasses()
    console.log(global.optionClasses)
    this.optionClasses = global.optionClasses
  },
  methods: {
    delOption(optionClassIndex, optionName) {
      let model = utils.getRequestModel("mvp", "DelOption", {
        "className": global.optionClasses[optionClassIndex].className,
        "optionName": optionName
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
