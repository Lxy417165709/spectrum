<!-- eslint-disable -->
<template>
  <el-form label-width="80px">
    <el-form-item label="桌类名">{{ deskSet.name }}</el-form-item>
    <!--    <el-form-item label="照片">-->
    <!--      <el-upload-->
    <!--        action="/api/upload"-->
    <!--        list-type="picture-card">-->
    <!--        <i class="el-icon-plus"></i>-->
    <!--      </el-upload>-->
    <!--    </el-form-item>-->
    <desk-info-editor style="margin-bottom: 20px"></desk-info-editor>
    <!--    todo: 这里显示的样式有问题...-->
    <el-form-item>
      <el-button type="primary">确定</el-button>
    </el-form-item>

  </el-form>
</template>

<script>
/* eslint-disable */
import DeskInfoEditor from "./DeskSizeEditor";
export default {
  name: "DeskSetEditor",
  components: {DeskInfoEditor},
  data() {
    return {
      deskSet: {
        name: "桌球"
      }
    }
  }
}
</script>

<style scoped>

</style>
