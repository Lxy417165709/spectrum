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
      <span>{{
          deskSets[curDeskSetIndex].name
        }} {{ curDeskIndex === -1 ? '' : " - " + deskSets[curDeskSetIndex].desks[curDeskIndex].space.num }}
      {{ curGoodClassIndex === -1 ? '' : " - " + goodClasses[curGoodClassIndex].name }}
      </span>

      <div v-show="viewMode === 0">
        <el-tabs>
          <component :is="deskSets[curDeskSetIndex].component" :desks="deskSets[curDeskSetIndex].desks"
                     @passViewMode="changeMode"></component>
        </el-tabs>
      </div>
      <div v-show="viewMode === 1">
        <el-tabs>
          <component :is="goodClassList" :goodClasses="goodClasses" @passGoodsViewMode="changeToGoodsMode"></component>
        </el-tabs>
      </div>
      <div v-show="viewMode === 2">
        <el-tabs>
          <component v-if="curGoodClassIndex!==-1" :is="goodList" :goods="goodClasses[curGoodClassIndex].goods" ></component>
        </el-tabs>
      </div>
    </el-main>
  </el-container>
</template>
<script>
/* eslint-disable */
import DeskList from "./DeskList";
import DeskCard from "./DeskCard";
import GoodClassList from "../good/GoodClassList";
import GoodClassCard from "../good/GoodClassCard";
import GoodCard from "../good/GoodCard";
import GoodList from "../good/GoodList";
// todo: 这个文件的 el-main 应该可以提出来
export default {
  name: 'DeskShower',
  components: {DeskList},
  data() {
    return {
      viewMode: 0, // todo: 这里可以弄个枚举
      goodClassList: GoodClassList,
      goodList: GoodList,
      deskSets: [
        {
          component: DeskList,
          name: "台球桌",
          desks: [
            {
              component: DeskCard,
              space: {
                name: "台球桌",
                num: 1,
              },
            },
            {
              component: DeskCard,
              space: {
                name: "台球桌",
                num: 2,
              },
            },
            {
              component: DeskCard,
              space: {
                name: "台球桌",
                num: 3,
              },
            }
          ],
        },
        {
          component: DeskList,
          name: "麻将桌",
          desks: [
            {
              component: DeskCard,
              space: {
                name: "麻将桌",
                num: 1,
              },
            },
            {
              component: DeskCard,
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
          component: GoodClassCard,
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
          component: GoodClassCard,
          name: "水果茶系列",
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
          component: GoodClassCard,
          name: "小吃系列",
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
        }
      ],
      curDeskSetIndex: 0,
      curDeskIndex: -1,
      curGoodClassIndex: -1,
    }
  },
  methods: {
    handleDeskSetClick(index) {
      this.curDeskSetIndex = index
      this.viewMode = 0
      this.curDeskIndex = -1
      this.curGoodClassIndex = -1
    },
    changeMode(mode, deskIndex) {
      this.viewMode = mode
      this.curDeskIndex = deskIndex
    },
    changeToGoodsMode(mode, goodClassIndex) {
      this.viewMode = mode
      this.curGoodClassIndex = goodClassIndex
    }
  }
}
</script>

<style scoped>

</style>
