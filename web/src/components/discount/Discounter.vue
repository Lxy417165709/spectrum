<!-- eslint-disable -->
<template>
  <div>
    <el-form ref="form" label-width="80px">
      <el-form-item label="折扣类型">
        <el-select v-model="FavorTypeName" placeholder="请选择附属选项">
          <el-option v-for="(FavorOption,index) in FavorOptions" :key="index"
                     :label="FavorOption.name" :value="FavorOption.name">
          </el-option>
        </el-select>
      </el-form-item>
      <template v-if="FavorTypeName === '打折'">
        <el-form-item label="折扣">
          <el-col :span="5">
            <el-input v-model="rebate" type="primary"></el-input>
          </el-col>
        </el-form-item>
      </template>
      <template v-if="FavorTypeName === '满减'">
        <el-form-item label="满">
          <el-col :span="5">
            <el-input v-model="full" type="primary"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item label="减">
          <el-col :span="5">
            <el-input v-model="reduction" type="primary"></el-input>
          </el-col>
        </el-form-item>
      </template>

    </el-form>
  </div>
</template>

<script>
/* eslint-disable */
import cst from "../../common/cst";

export default {
  name: "Discounter",
  data() {
    return {
      FavorTypeName: cst.FAVOR_TYPE.NONE.NAME,
      FavorOptions: [
        cst.FAVOR_TYPE.NONE.NAME,
        cst.FAVOR_TYPE.REBATE.NAME,
        cst.FAVOR_TYPE.FULL_REDUCTION.NAME,
        cst.FAVOR_TYPE.FREE.NAME,
      ],
      rebate: 10,
      full: 10000,
      reduction: 10,
    }
  }
}
</script>
