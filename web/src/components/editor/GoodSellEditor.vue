<!-- eslint-disable -->
<template>
  <el-form label-width="80px">
    <el-form-item label="照片">
      <img src="" style="width: 200px;height: 200px;margin-top: 10px">
    </el-form-item>
    <el-form-item label="名字">
      <span>{{ good.name }}</span>
    </el-form-item>
    <el-form-item label="规格">
      <el-radio v-model="good.curSizeIndex" v-for="(sizeInfo,index) in good.sizeInfos" :label="index" :key="index">
        {{ sizeInfo.name }}
      </el-radio>
    </el-form-item>
    <el-form-item v-for="(attachElement,index) in good.attachElements" :key="index" :label="attachElement.name"
                  v-if="attachElement.elementType===1">
      <el-radio v-model="attachElement.curSizeIndex" v-for="(sizeInfo,index) in attachElement.sizeInfos" :label="index"
                :key="index">
        {{ sizeInfo.name }}
      </el-radio>
    </el-form-item>
    <el-form-item v-if="!needAttachGood">
      <el-button @click="needAddAttachGood">需要加料</el-button>
    </el-form-item>
    <el-form-item v-for="(attachElement,index) in good.attachElements" :key="index" :label="attachElement.name"
                  v-if="attachElement.elementType===2 && needAttachGood">
      <el-radio v-model="attachElement.curSizeIndex" v-for="(sizeInfo,index) in attachElement.sizeInfos" :label="index"
                :key="index">
        {{ sizeInfo.name }}
      </el-radio>
    </el-form-item>
    <el-form-item label="优惠">
      <el-radio v-model="curFavorIndex" v-for="(favor,index) in favors" :label="index"
                :key="index">
        {{ favor.name }}
      </el-radio>
    </el-form-item>
    <component :is="favors[curFavorIndex].component" @confirmFavor="confirmFavor"></component>
    <el-form-item label="已选折扣">
      <el-tag v-for="(favor,index) in selectedFavors" :key="index" closable style="margin-right: 10px">
        {{ favor.name }}
      </el-tag>
    </el-form-item>

    <el-form-item label="价格">
      <span>{{ good.price }}</span>
    </el-form-item>
    <el-form-item label="备注">
      <el-input style="width: 70%"></el-input>
    </el-form-item>
    <el-form-item>
      <el-button type="primary">确定</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
/* eslint-disable */
import DiscountComponent from "../baby/DiscountComponent";

export default {
  name: "GoodSellEditor",
  data() {
    return {
      selectedFavors: [
        {
          name: "打 8 折"
        },
        {
          name: "满 100 减 6"
        },
      ],
      needAttachGood: false,
      curFavorIndex: 0,
      // todo: favor 被选中后，要展示 favor 对应的组件
      favors: [
        {
          name: "无",
        },
        {
          name: "打折",
          component: DiscountComponent
        },
        {
          name: "满减",
        },
        {
          name: "免单",
        }
      ],
      good: {
        price: 19,
        curSizeIndex: 0,
        name: "波霸奶茶",
        attachElements: [
          {
            name: "冰量",
            elementType: 1,
            curSizeIndex: 0,
            sizeInfos: [
              {
                name: "少冰",
                price: 0,
              },
              {
                name: "中冰",
                price: 0,
              },
              {
                name: "多冰",
                price: 0,
              },
            ]
          },
          {
            name: "温度",
            curSizeIndex: 0,
            elementType: 1,
            sizeInfos: [
              {
                name: "冷饮",
                price: 0,
              },
              {
                name: "常温",
                price: 0,
              },
              {
                name: "热饮",
                price: 0,
              },
            ]
          },
          {
            name: "珍珠",
            elementType: 2,
            curSizeIndex: 0,
            sizeInfos: [
              {
                name: "无",
                price: 0,
              },
              {
                name: "少量",
                price: 1,
              },
              {
                name: "中量",
                price: 1,
              },
              {
                name: "大量",
                price: 1,
              },
            ]
          }
        ],
        sizeInfos: [
          {
            name: "小规格",
            price: 18.0,
          },
          {
            name: "中规格",
            price: 20.0,
          },
          {
            name: "大规格",
            price: 25.0,
          }
        ]
      }
    }
  },
  methods: {
    needAddAttachGood() {
      this.needAttachGood = true
    },
    confirmFavor(favor) {
      this.selectedFavors.push(favor)
    }
  }
}
</script>

<style scoped>

</style>
