<!-- eslint-disable -->
<template>
  <el-card :body-style="{ padding: '0px' }" shadow="hover" style="position:relative;">
    <img style="height:200px;width:200px; border: none"
         :src="'api/file/' + good.mainElement.sizeInfos[good.mainElement.selectedIndex].pictureStorePath">
    <div style="width: 200px;height: 30px;position:relative;padding: 0;margin: 0; border: none">
      <span style="position:absolute; left:5px">{{ good.mainElement.name }}</span>
      <span style="position:absolute; right:10px;color: red;font-size: 1.2em">{{good.mainElement.sizeInfos[good.mainElement.selectedIndex].price}}</span>
    </div>
  </el-card>

</template>

<script>
/* eslint-disable */
export default {
  props: {
    good: Object
  },
  data() {
    return {};
  },
  methods: {}
}
</script>
