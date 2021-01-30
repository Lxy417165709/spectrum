<!-- eslint-disable -->
<template>
  <el-container>
    <el-header>
      <el-row>
        <el-col :span="4">
          是否结账
          <el-select value="全部">
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
            <el-button type="primary">确定</el-button>
          </div>
        </el-col>
      </el-row>
    </el-header>
    <el-main>
      <el-collapse v-model="activeName" accordion>
        <el-collapse-item name="1">
          <template slot="title">
            <div><span>序号: 1</span></div>
            <!--            <div><span>桌类: 台球</span></div>-->
            <!--            <div><span>桌号: 2</span></div>-->
          </template>

          <el-divider content-position="left">桌费</el-divider>
          <!--todo: 表项内容有些不是纯文本，可能是按钮、标签-->
          <el-table
            :data="order.deskInfos"
            border
            style="width: 100%">
            <el-table-column
              type="index"
              width="50">
            </el-table-column>
            <el-table-column
              prop="name"
              label="桌位名"
              width="180">
            </el-table-column>
            <el-table-column
              prop="num"
              label="桌位号"
              width="180">
            </el-table-column>
            <el-table-column
              prop="beginTime"
              label="开桌时间"
              width="180px">
            </el-table-column>
            <el-table-column
              prop="endTime"
              label="关桌时间"
              width="180px">
            </el-table-column>
            <el-table-column
              prop="duration"
              label="占用时长">
            </el-table-column>
            <el-table-column
              prop="countWay"
              label="计费方式">
            </el-table-column>
            <el-table-column
              prop="price"
              label="计费价格">
            </el-table-column>
            <el-table-column
              prop="nonFavorExpense"
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
              prop="expense"
              label="折后花费">
            </el-table-column>

            <el-table-column

              label="结账">
            </el-table-column>
          </el-table>

          <el-divider content-position="left">商品费</el-divider>
          <!--todo: 表项内容有些不是纯文本，可能是按钮、标签-->
          <el-table
            :data="order.goodInfos"
            border
            style="width: 100%">
            <el-table-column
              type="index"
              width="50">
            </el-table-column>
            <el-table-column
              prop="name"
              label="商品名"
              width="180">
            </el-table-column>
            <el-table-column
              prop="nonFavorExpense"
              label="原花费">
            </el-table-column>
            <el-table-column
              label="优惠"
              width="180px">
              <template slot-scope="scope">
                <!--            todo: 这个可以提出为一个组件 -->
                <el-tag v-for="(favor,index) in scope.row.favors" :key="index" style="margin-right: 10px">
                  {{
                    favor.name
                  }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column
              prop="expense"
              label="折后花费">
            </el-table-column>

            <el-table-column

              label="结账">
            </el-table-column>
          </el-table>


          <el-divider>优惠设置</el-divider>
          <!--          todo: 这里要抽出一个组件-->
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

export default {
  name: "Order",
  components: {DiscountEditor},
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
      order: {
        deskInfos: [
          {
            name: "桌球",
            num: 1,
            beginTime: "2021年01月29日16:08:07",
            endTime: "2021年01月30日16:08:07",
            duration: "2小时",
            countWay: "计时",
            price: 18,
            nonFavorExpense: 36,
            favors: [
              {
                name: "8 折"
              },
              {
                name: "满 100 减 6"
              }
            ],
            expense: 32.4

          }
        ],
        goodInfos: [
          {
            name: "波霸奶茶",
            nonFavorExpense: 18.0,
            favors: [
              {
                name: "8 折"
              },
              {
                name: "满 100 减 6"
              }
            ],
            expense: 14.4,
          },
          {
            name: "红豆奶茶",
            nonFavorExpense: 15.0,
            favors: [
              {
                name: "9 折"
              },
              {
                name: "满 1090 减 6"
              }
            ],
            expense: 15,
          }
        ]
      }
    }
  }
}
</script>

<style scoped>

</style>
