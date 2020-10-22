<!-- eslint-disable -->
<template>
  <div>
    <el-form ref="form" label-width="80px">
      <el-form-item label="折扣类型">
        <el-select v-model="discountTypeName" placeholder="请选择附属选项">
          <el-option v-for="(discountOption,index) in discountOptions" :key="index"
                     :label="discountOption.name" :value="discountOption.name">
          </el-option>
        </el-select>
      </el-form-item>
      <template v-if="discountTypeName === '打折'">
        <el-form-item label="折扣">
          <el-col :span="5">
            <el-input v-model="rebate" type="primary"></el-input>
          </el-col>
        </el-form-item>
      </template>
      <template v-if="discountTypeName === '满减'">
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
import global from "../../common/global_object/global";

export default {
  name: "Discounter",
  data() {
    return {
      discountTypeName: global.discountType.NONE,
      discountOptions: [
        {
          name: global.discountType.NONE
        },
        {
          name: global.discountType.REBATE
        },
        {
          name: global.discountType.FULL_REDUCTION
        },
        {
          name: global.discountType.FREE
        }
      ],
      rebate: 10,
      full: 10000,
      reduction: 10,
    }
  }
}
</script>
