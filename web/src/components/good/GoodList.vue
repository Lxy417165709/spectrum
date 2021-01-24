<!-- eslint-disable -->
<template>
  <el-row>
    <el-col v-for="(good,goodIndex) in goods" :key="goodIndex"
            style="height: 300px; width: 202px; margin-left: 10px; border: none">
      <span>
        <component :is="good.component" :name="good.name"
                   :pictureStorePath="good.pictureStorePath"></component>
      </span>
    </el-col>
  </el-row>
</template>

<script>
/* eslint-disable */
export default {
  props: {
    goods: Array
  },
  data() {
    return {};
  },
  methods: {}
}
</script>
