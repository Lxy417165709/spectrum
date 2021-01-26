<!-- eslint-disable -->
<template>
  <el-container style="height: 800px; border: 1px solid #eee">
    <el-aside width="200px">
      <el-menu>
        <el-menu-item v-for="(deskSet,deskSetIndex) in deskSets" :key="deskSetIndex"
                      @click="handleDeskSetClick(deskSetIndex)">
          <template slot="title"><i class="el-icon-message"></i><span>{{ deskSet.name }}</span></template>
        </el-menu-item>
      </el-menu>
    </el-aside>
    <el-main>
      <!--            <span>{{-->
      <!--                deskSets[curDeskSetIndex].name-->
      <!--              }} {{ curDeskIndex === -1 ? '' : " - " + deskSets[curDeskSetIndex].desks[curDeskIndex].space.num }}-->
      <!--            {{ curGoodClassIndex === -1 ? '' : " - " + goodClasses[curGoodClassIndex].name }}-->
      <!--            </span>-->
      <div v-show="viewMode === 0">
        <desk-list :desks="deskSets[curDeskSetIndex].desks"
                   @passViewMode="changeToGoodClassListMode"></desk-list>
      </div>
      <div v-show="viewMode !== 0">
        <good-class :goodClasses="goodClasses" ref="goodClassSon" :isEditMode="false" @turnToDeskListMode="turnToDeskListMode"></good-class>
      </div>

    </el-main>
  </el-container>
</template>
<script>
/* eslint-disable */
import DeskList from "../list/DeskList";
import GoodCard from "../card/GoodCard";
import GoodClass from "../manager/GoodClass";
// todo: 这个文件的 el-main 应该可以提出来
export default {
  name: 'DeskShower',
  components: {GoodClass, DeskList},
  data() {
    return {
      viewMode: 0, // todo: 这里可以弄个枚举
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
    changeToGoodClassListMode(deskIndex) {
      this.viewMode = 1
      this.curDeskIndex = deskIndex
      console.log("test",this.curDeskIndex,this.viewMode,this.goodClasses)
    },
    turnToDeskListMode(){
      this.viewMode = 0
    }
  }
}
</script>

<style scoped>

</style>
