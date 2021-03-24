<!-- eslint-disable -->
<template>
  <el-container>
    <!--    1. 搜索条件-->
    <el-header>
      <el-row>
        <el-col :span="3">
          订单号:
          <el-input v-model="searchOrderID" style="width: 70px"></el-input>
        </el-col>
        <el-col :span="6">
          是否结账:
          <el-select v-model="checkOutValue">
            <el-option label="全部" value="全部"></el-option>
            <el-option label="未结账" value="未结账"></el-option>
            <el-option label="已结账" value="已结账"></el-option>
          </el-select>
        </el-col>
        <el-col :span="9">
          时间范围:
          <el-date-picker
            v-model="timeInterval"
            type="datetimerange"
            :picker-options="pickerOptions"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期">
          </el-date-picker>
        </el-col>
        <el-col :span="3">
          <div>
            <el-button type="primary" @click="startToGetOrder">确定</el-button>
          </div>
        </el-col>
      </el-row>
    </el-header>
    <el-main>
      <el-collapse v-model="activeName" accordion>
        <el-collapse-item :name="index.toString()" :key="index" v-for="(order,index) in db_orders">
          <template slot="title">
            <div>
              <span>订单号: {{ order.id }} </span>
              <el-divider direction="vertical"></el-divider>
              <span>{{ order.desk.space.className }}</span>
              <el-divider direction="vertical"></el-divider>
              <span> {{ order.desk.space.name }}</span>
            </div>
          </template>

          <!--          2.1 桌费-->
          <el-divider content-position="left">桌费</el-divider>
          <el-table
            :data="[order.desk]"
            border
            style="width: 100%">
            <el-table-column
              type="index"
              width="50">
            </el-table-column>
            <el-table-column
              prop="space.className"
              label="桌类名"
              width="180">
            </el-table-column>
            <el-table-column
              prop="space.name"
              label="桌名"
              width="180">
            </el-table-column>
            <el-table-column
              label="开桌时间"
              width="180px">
              <template slot-scope="scope">
                {{ timestampToTime(scope.row.startAt) }}
              </template>
            </el-table-column>
            <el-table-column
              label="关桌时间"
              width="180px">
              <template slot-scope="scope">
                {{ getTimeShow(scope.row.endAt) }}
              </template>
            </el-table-column>
            <el-table-column
              prop="duration"
              label="占用时长">
              <template slot-scope="scope">
                {{ getDuration(scope.row.endAt, scope.row.startAt) }}
              </template>
            </el-table-column>
            <el-table-column
              prop="countWay"
              label="计费方式">
            </el-table-column>
            <el-table-column
              prop="space.price"
              label="计费价格">
            </el-table-column>
            <el-table-column
              prop="expenseInfo.nonFavorExpense"
              label="原花费">
            </el-table-column>
            <el-table-column
              label="优惠"
              width="180px">
              <template slot-scope="scope">
                <el-tag v-for="(favor,index) in scope.row.favors" :key="index" style="margin-right: 10px">
                  {{
                    getFavorString(favor)
                  }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column
              prop="expenseInfo.expense"
              label="折后花费">
            </el-table-column>

            <el-table-column
              label="结账">
              <template slot-scope="scope">
                <el-button type="primary" @click="checkOutDesk(scope.row)">结账</el-button>
              </template>
            </el-table-column>
          </el-table>

          <!--          2.1 商品费-->
          <div v-if="order.goods !== null && order.goods.length!==0">
            <el-divider content-position="left">商品费</el-divider>
            <el-table
              :data="order.goods"
              border
              style="width: 100%">
              <el-table-column
                type="index"
                width="50">
              </el-table-column>
              <el-table-column
                prop="mainElement.name"
                label="商品名"
                width="180">
              </el-table-column>
              <el-table-column
                label="规格"
                width="180">
                <template slot-scope="scope">
                  {{ scope.row.mainElement.sizeInfos[scope.row.mainElement.selectedIndex].size }}
                </template>
              </el-table-column>
              <el-table-column
                label="附属品"
                width="180">
                <template slot-scope="scope">
                  <el-tag v-for="(attachElement,index) in scope.row.attachElements" :key="index"
                          style="margin-right: 10px">
                    {{ attachElement.name }} ({{ attachElement.sizeInfos[attachElement.selectedIndex].size }})
                  </el-tag>
                </template>
              </el-table-column>
              <el-table-column
                prop="expenseInfo.nonFavorExpense"
                label="原花费">
              </el-table-column>
              <el-table-column
                label="优惠"
                width="180px">
                <template slot-scope="scope">
                  <el-tag v-for="(favor,index) in scope.row.favors" :key="index" style="margin-right: 10px">
                    {{
                      getFavorString(favor)
                    }}
                  </el-tag>
                </template>
              </el-table-column>
              <el-table-column
                prop="expenseInfo.expense"
                label="折后花费">
              </el-table-column>
              <el-table-column
                label="结账">
                <template slot-scope="scope">
                  <el-button type="primary" @click="checkOutGood(scope.row)">结账</el-button>
                </template>
              </el-table-column>
            </el-table>
          </div>

          <el-divider>优惠设置</el-divider>
          <el-form label-width="80px">
            <discount-editor></discount-editor>
            <el-form-item>
              <el-button type="primary" @click="checkOutOrder(order)">结账</el-button>
            </el-form-item>
          </el-form>

        </el-collapse-item>
      </el-collapse>
    </el-main>
  </el-container>
</template>

<script>
/* eslint-disable */
import DiscountEditor from "../editor/DiscountEditor";
import test from "../../common/test";
import utils from "../../common/utils";
import cst from "../../common/cst";

export default {
  name: "ManageOrderPage",
  components: {DiscountEditor},
  mounted() {
    this.order = test.order
  },
  data() {
    return {
      activeName: "0",
      pickerOptions: {
        shortcuts: [
          {
            text: '最近一天',
            onClick(picker) {
              const end = new Date();
              const start = new Date();
              start.setTime(start.getTime() - 3600 * 1000 * 24);
              picker.$emit('pick', [start, end]);
            }
          },
          {
            text: '最近一周',
            onClick(picker) {
              const end = new Date();
              const start = new Date();
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 7);
              picker.$emit('pick', [start, end]);
            }
          }, {
            text: '最近一个月',
            onClick(picker) {
              const end = new Date();
              const start = new Date();
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 30);
              picker.$emit('pick', [start, end]);
            }
          }, {
            text: '最近三个月',
            onClick(picker) {
              const end = new Date();
              const start = new Date();
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 90);
              picker.$emit('pick', [start, end]);
            }
          }]
      },
      timeInterval: [],
      db_orders: {},
      checkOutValue: "全部",
      searchOrderID: ""
    }
  },
  created() {
    const end = new Date();
    const start = new Date();
    start.setTime(start.getTime() - 3600 * 1000 * 24);
    this.timeInterval = [start, end]
  },
  methods: {
    startToGetOrder() {
      let par = {
        orderID: this.searchOrderID - 0,
        checkOutState: this.cpt_checkOutState,
        startAt: Date.parse(this.timeInterval[0]) / 1000,
        endAt: Date.parse(this.timeInterval[1]) / 1000,
      }
      console.log("par", par)
      utils.GetOrder(this, par, (res) => {
        this.db_orders = res.data.data.orders
      })
    },
    timestampToTime(timestamp) {
      let date = new Date(timestamp * 1000);//时间戳为10位需*1000，时间戳为13位的话不需乘1000
      let Y = date.getFullYear() + '-';
      let M = (date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1) + '-';
      let D = (date.getDate() < 10 ? '0' + date.getDate() : date.getDate()) + ' ';
      let h = (date.getHours() < 10 ? '0' + date.getHours() : date.getHours()) + ':';
      let m = (date.getMinutes() < 10 ? '0' + date.getMinutes() : date.getMinutes()) + ':';
      let s = date.getSeconds() < 10 ? '0' + date.getSeconds() : date.getSeconds();
      return Y + M + D + h + m + s;
    },
    getDuration(endAt, startAt) {
      if (endAt === cst.TIMESTAMP.NIL) {
        endAt = Date.parse(new Date()) / 1000;
      }
      let timestamp = (endAt - startAt);
      let h = Math.floor(timestamp / 3600) + "小时 "
      let m = Math.floor((timestamp % 3600) / 60) + "分 "
      let s = timestamp % 60 + '秒'
      return h + m + s;
    },
    getTimeShow(timestamp) {
      if (timestamp === cst.TIMESTAMP.NIL) {
        return "-"
      }
      return this.timestampToTime(scope.row.endAt)
    },
    getFavorString(favor) {
      if (favor.favorType === cst.FAVOR_TYPE.NONE.VALUE) {
        return cst.FAVOR_TYPE.NONE.NAME
      }
      if (favor.favorType === cst.FAVOR_TYPE.REBATE.VALUE) {
        return cst.FAVOR_TYPE.REBATE.NAME
      }
      if (favor.favorType === cst.FAVOR_TYPE.FULL_REDUCTION.VALUE) {
        return cst.FAVOR_TYPE.FULL_REDUCTION.NAME
      }
      if (favor.favorType === cst.FAVOR_TYPE.FREE.VALUE) {
        return cst.FAVOR_TYPE.FREE.NAME
      }
      return ""
    },
    checkOutGood(good){
      console.log("checkOutGood",good)
    },
    checkOutDesk(desk){
      console.log("checkOutDesk",desk)
    },
    checkOutOrder(order){
      console.log("checkOutOrder",order)
    }
  },
  computed: {
    cpt_checkOutState() {
      if (this.checkOutValue === "全部") {
        return 0
      }
      if (this.checkOutValue === "未结账") {
        return 1
      }
      if (this.checkOutValue === "已结账") {
        return 2
      }
      return 0
    },
  }


}
</script>

<style scoped>

</style>
