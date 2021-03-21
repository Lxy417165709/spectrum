<!-- eslint-disable -->
<template>
  <el-container>
    <!--    1. 搜索条件-->
    <el-header>
      <el-row>
        <el-col :span="4">
          是否结账
          <el-select v-model="check_out_value">
            <el-option label="全部" value="全部"></el-option>
            <el-option label="未结账" value="未结账"></el-option>
            <el-option label="已结账" value="已结账"></el-option>
          </el-select>
        </el-col>
        <el-col :span="7">
          时间范围
          <el-date-picker
            v-model="value2"
            type="datetimerange"
            :picker-options="pickerOptions"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            align="right">
          </el-date-picker>
        </el-col>
        <el-col :span="8">
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
            <div><span>订单号: {{ order.id }} </span></div>
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
                {{ scope.row.endAt === 0 ? "-" : timestampToTime(scope.row.endAt) }}
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
                    favor.name
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
                      favor.name
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
              </el-table-column>
            </el-table>
          </div>


          <el-divider>优惠设置</el-divider>
          <el-form label-width="80px">
            <discount-editor></discount-editor>
            <el-form-item>
              <el-button type="primary">结账</el-button>
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
import test from "../../common/test/test";
import utils from "../../common/utils";

export default {
  name: "ManageOrderPage",
  components: {DiscountEditor},
  mounted() {
    this.order = test.order
  },
  data() {
    return {
      activeName: "1",
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
      value1: '',
      value2: '',
      db_orders: {},
      check_out_value: "全部",
    }
  },
  methods: {
    startToGetOrder() {
      utils.GetOrder(this, {
        // orderID: 2,
        checkOutState: 0,
      }, (res) => {
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
      if (endAt === 0) {
        endAt = Date.parse(new Date()) / 1000;
      }
      let timestamp = (endAt - startAt);
      let h = Math.floor(timestamp / 3600) + "小时 "
      let m = Math.floor((timestamp % 3600) / 60) + "分 "
      let s = timestamp % 60 + '秒'
      return h + m + s;
    }
  },

}
</script>

<style scoped>

</style>
