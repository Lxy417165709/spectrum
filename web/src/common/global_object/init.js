/* eslint-disable */

import utils from "../utils";
import global from "./global";

export default{
  async globalOptionClasses() {
    let model = utils.getRequestModel("mvp", "GetAllOptionClasses", {})
    await utils.sendRequestModel(model).then(res => {
      if (!utils.hasRequestSuccess(res)) {
        console.log(res.data.err)
        return
      }
      global.optionClasses = res.data.data.optionClasses
    })
  },
  async globalGoods() {
    let model = utils.getRequestModel("mvp", "GetAllGoods", {})
    await utils.sendRequestModel(model).then(res => {
      if (!utils.hasRequestSuccess(res)) {
        console.log(res.data.err)
        return
      }
      global.goods = res.data.data.goods
      console.log("global_goods",global.goods)
    })
  },
  async globalGoodClasses() {
    let model = utils.getRequestModel("mvp", "GetAllGoodClasses", {})
    await utils.sendRequestModel(model).then(res => {
      if (!utils.hasRequestSuccess(res)) {
        console.log(res.data.err)
        return
      }

      global.goodClasses = res.data.data.goodClasses
      console.log("global_good_classes",global.goodClasses)
    })
  }

}
