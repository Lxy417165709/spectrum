<!-- eslint-disable -->
<template>
  <el-container style="height: 800px; border: 1px solid #eee">
    <!--    1. 桌类选择-->
    <el-aside width="200px">
      <el-menu>
        <el-menu-item v-for="(deskSet,deskSetIndex) in deskSets" :key="deskSetIndex"
                      @click="handleDeskSetClick(deskSetIndex)">
          <template slot="title"><i class="el-icon-message"></i><span>{{ deskSet.name }}</span></template>
        </el-menu-item>
      </el-menu>
      <el-button style="margin-left: 20px;margin-top: 10px" type="primary" @click="handleDeskButtonClick">添加桌位
      </el-button>
    </el-aside>

    <!--    2. 桌类详情-->
    <el-main>
      <div v-show="viewMode === 0">
        <el-row style="height: 40px;margin-left:10px;margin-bottom: 10px">
        </el-row>
        <el-divider content-position="left">{{ deskSets[curDeskSetIndex].name }}</el-divider>
        <desk-list :desks="deskSets[curDeskSetIndex].desks"
                   @turnToGoodClassListMode="turnToGoodClassListMode"></desk-list>
      </div>
      <div v-show="viewMode !== 0">
        <good-class :goodClasses="goodClasses" ref="goodClassSon" :isEditMode="false"
                    @turnToFatherMode="turnToDeskListMode" :hasFather="true"></good-class>
      </div>
    </el-main>
    <el-dialog
      title="桌类添加/编辑"
      :visible.sync="deskSetEditorVisible"
      width="30%">
      <!--      <span>这是一段信息</span>-->

      <!--      <span slot="footer" class="dialog-footer">-->
      <!--        <el-button @click="dialogVisible = false">取 消</el-button>-->
      <!--        <el-button type="primary" @click="dialogVisible = false">确 定</el-button>-->
      <!--      </span>-->
      <desk-set-editor></desk-set-editor>
    </el-dialog>
  </el-container>
</template>
<script>
/* eslint-disable */
import DeskList from "../list/DeskList";
import GoodCard from "../card/GoodCard";
import GoodClass from "../manager/GoodClass";
import DeskSetEditor from "../editor/DeskSetEditor";

export default {
  name: 'DeskShower',
  components: {DeskSetEditor, GoodClass, DeskList},
  data() {
    return {
      viewMode: 0, // todo: 这里可以弄个枚举
      deskSetEditorVisible: false,
      deskSets: [
        {
          name: "台球桌",
          desks: [
            {
              space: {
                name: "台球桌",
                num: 1,
              },
            },
            {
              space: {
                name: "台球桌",
                num: 2,
              },
            },
            {
              space: {
                name: "台球桌",
                num: 3,
              },
            }
          ],
        },
        {
          name: "麻将桌",
          desks: [
            {
              space: {
                name: "麻将桌",
                num: 1,
              },
            },
            {
              space: {
                name: "麻将桌",
                num: 2,
              },
            }
          ],
        },
      ],
      goodClasses: [
        {
          name: "奶茶系列",
          classType: 1,
          goods: [
            {
              component: GoodCard,
              name: "波霸奶茶",
            },
            {
              component: GoodCard,
              name: "红豆奶茶",
            }
          ]
        },
        {
          name: "水果茶系列",
          classType: 1,
          goods: [
            {
              component: GoodCard,
              name: "超神水果茶",
            },
            {
              component: GoodCard,
              name: "我爱水果茶",
            }
          ]
        },
        {
          name: "小吃系列",
          classType: 1,
          goods: [
            {
              component: GoodCard,
              name: "炸鸡",
            },
            {
              component: GoodCard,
              name: "薯条",
            }
          ]
        },
        {
          name: "附属选项",
          classType: 2,
          goods: [
            {
              component: GoodCard,
              name: "冰量",
            },
            {
              component: GoodCard,
              name: "温度",
            }
          ]
        },
        {
          name: "附属商品",
          classType: 3,
          goods: [
            {
              component: GoodCard,
              name: "珍珠",
            },
            {
              component: GoodCard,
              name: "红豆",
            }
          ]
        }
      ],
      curDeskSetIndex: 0,
      curDeskIndex: -1,
    }
  },
  methods: {
    handleDeskSetClick(index) {
      this.curDeskSetIndex = index
      this.viewMode = 0
      this.curDeskIndex = -1
      this.$refs.goodClassSon.viewMode = 1; // 父传子
    },
    turnToGoodClassListMode(deskIndex) {
      this.viewMode = 1
      this.curDeskIndex = deskIndex
      console.log("test", this.curDeskIndex, this.viewMode, this.goodClasses)
    },
    handleDeskButtonClick() {
      this.deskSetEditorVisible = true
    },
    turnToDeskListMode() {
      this.viewMode = 0
    }
  }
}
</script>

<style scoped>

</style>
