<!-- eslint-disable -->
<template>
  <el-card :body-style="{ padding: '0px' }" shadow="hover" style="position:relative;">
    <img style="height:200px;width:200px; border: none" :src="goodOption.pictureStorePath">
    <div style="width: 200px;height: 30px;position:relative;padding: 0;margin: 0; border: none">
      <span style="position:absolute; left:5px">{{ goodOption.name }}</span>
    </div>
  </el-card>
</template>

<script>
/* eslint-disable */
export default {
  props: {
    goodOption:Object,
  },
  data() {
    return {};
  },
  methods: {}
}
</script>
