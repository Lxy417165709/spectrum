<!-- eslint-disable -->
<template>


  <el-card :body-style="{ padding: '0px' }" shadow="hover" style="position:relative;">
    <el-badge class="item" :hidden="desk===undefined || desk.id===0" value="已占用" style="position: absolute;right: 10px;top:5px">
    </el-badge>
    <img style="height:200px;width:200px; border: none" :src="'api/file/' + desk.space.pictureStorePath"
         v-if="desk.space!==undefined">
    <div
      style="width: 200px;height: 30px;position:relative;padding: 0;margin: 0; border: none;">
      <span style="position:absolute; left:5px" v-if="desk.space!==undefined">{{ desk.space.name }}</span>
    </div>
  </el-card>

</template>

<script>
/* eslint-disable */

// todo: desk.id === undefined 时，此时 desk.id !== 0 也成立..
export default {
  props: {
    desk: {}
  },
  data() {
    return {};
  },
  methods: {}
}
</script>

<style>
.isOpening {
  background-color: #ff0000;
}
</style>
