/* eslint-disable */

import axios from "axios";

export default {
  getRequestModel(obj, fun, par) {
    return {
      object: obj,
      function: fun,
      parameters: par,
    }
  },
  sendRequestModel(model) {
    return axios.post("/api/distributor", model)
  },
  hasRequestSuccess(res) {
    let returnResult = res.data
    return returnResult.err === "" || returnResult.err === undefined || returnResult.err === null
  },
  removeElement(array, element) {
    let result = [];
    for (let i = 0; i < array.length; i++) {
      if (array[i] !== element) {
        result.push(array[i])
      }
    }
    return result
  },
  removeElementByField(array, filed, value) {
    let result = [];
    for (let i = 0; i < array.length; i++) {
      if (array[i][filed] !== value) {
        result.push(array[i])
      }
    }
    return result
  },
  removeIndex(array, index) {
    return array.slice(0, index).concat(array.slice(index + 1))
  },
  deepCopy(obj) {
    return JSON.parse(JSON.stringify(obj))
  },
  hasData(obj) {
    return obj !== undefined && obj !== null
  },
  NewBlankSizeInfo(name) {
    return {
      size: name,
      price: "0"
    }
  },
  async GetAllGoodOptions(obj,par, then) {
    let model = this.getRequestModel("mvp", "GetAllGoodOptions", par)
    await this.sendRequestModel(model).then((res) => {
      console.log("GetAllGoodOptions.res", res.data)
      if (!this.hasRequestSuccess(res)) {
        obj.$message.error(res.data.err)
        return
      }
      obj.$message.success(res.data.msg)
      then(res)
    })

    // await this.sendRequestModel(model).then(res => {
    //   console.log("GetAllGoodOptions.res", res.data)
    //   if (!this.hasRequestSuccess(res)) {
    //     // this.$message.error(res.data.err)
    //     return
    //   }
    //   // this.$message.success(res.data.msg)
    //   return res
    // })
  }
}
