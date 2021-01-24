<!-- eslint-disable -->
<template>
  <el-row>
    <el-col v-for="(num,numIndex) in desk.nums" :key="numIndex" style="height: 300px; width: 202px; margin-left: 10px; border: none">
      <span >
        <el-card :body-style="{ padding: '0px' }" shadow="hover" style="position:relative;">
        <img style="height:200px;width:200px; border: none">
        <div style="width: 200px;height: 30px;position:relative;padding: 0;margin: 0; border: none">
          <span style="position:absolute; left:5px">{{ num }}</span>
        </div>
        </el-card>
      </span>
    </el-col>
  </el-row>
</template>

<script>
/* eslint-disable */
export default {
  props: {
    desk: Object,
  },
  data() {
    return {};
  },
  methods: {
  }
}
</script>
