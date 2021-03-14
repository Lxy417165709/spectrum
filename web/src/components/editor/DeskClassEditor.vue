<!-- eslint-disable -->
<template>
  <el-form label-width="80px">

    <el-form-item label="桌类名">
      <el-input v-model="deskClass.name" style="width: 70%"></el-input>
    </el-form-item>

    <!--    @edit="handleTabsEdit"-->
    <el-tabs type="border-card" @tab-click="" editable style="margin-bottom: 10px"
             @tab-add="handleClick">
      <el-tab-pane v-for="(desk,index) in deskClass.desks" :label="desk.space.name" :name="desk.space.name"
                   :key="index">
        <el-form label-width="80px">

          <el-form-item label="照片">
            <el-upload
              action="/api/upload"
              list-type="picture-card">
              <i class="el-icon-plus"></i>
            </el-upload>
          </el-form-item>

          <el-form-item label="价格">
            <el-input v-model="desk.space.price" style="width: 70%"></el-input>
          </el-form-item>
        </el-form>
      </el-tab-pane>
    </el-tabs>

    <el-form-item>
      <el-button type="primary">确定</el-button>
    </el-form-item>

  </el-form>
</template>

<script>
/* eslint-disable */
import DeskInfoEditor from "./_DeskSizeEditor";

export default {
  name: "DeskClassEditor",
  components: {DeskInfoEditor},
  data() {
    return {
      deskClass: {},
      addTabCount: 0
    }
  },
  methods: {
    handleClick(tab, event) {
      this.addTabCount++
      let name = "未设定规格" + this.addTabCount
      this.deskClass.desks.push({
        space: {
          name: name,
          price: "10"
        }
      })
    },
    // handleTabsEdit(name, event) {
    //   if (name === this.curSizeInfoName) {
    //     return
    //   }
    //   this.sizeInfos = utils.removeElementByField(this.sizeInfos, "name", name)
    // }
  }
}
</script>

<style scoped>

</style>
