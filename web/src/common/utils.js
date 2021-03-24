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
  isExist(array, field, value) {
    for (let i = 0; i < array.length; i++) {
      if (array[i][field] === value) {
        return true
      }
    }
    return false
  },
  removeElementByField(array, field, value) {
    let result = [];
    for (let i = 0; i < array.length; i++) {
      if (array[i][field] !== value) {
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
      id: 0,
      size: name,
      price: "0",
      pictureStorePath: ""
    }
  },
  async GetAllGoodOptions(obj, par, then) {
    let model = this.getRequestModel("mvp", "GetAllGoodOptions", par)
    await this.sendRequestModel(model).then((res) => {
      console.log("GetAllGoodOptions.res", res)
      if (!this.hasRequestSuccess(res)) {
        obj.$message.error(res.data.err)
        return
      }
      obj.$message.success(res.data.msg)
      then(res)
    })
  },
  async OrderDesk(obj, par, then) {
    let model = this.getRequestModel("mvp", "OrderDesk", par)
    await this.sendRequestModel(model).then((res) => {
      console.log("OrderDesk.res", res)
      if (!this.hasRequestSuccess(res)) {
        obj.$message.error(res.data.err)
        return
      }
      obj.$message.success(res.data.msg)
      then(res)
    })
  },
  async OrderGood(obj, par, then) {
    let model = this.getRequestModel("mvp", "OrderGood", par)
    await this.sendRequestModel(model).then((res) => {
      console.log("OrderGood.res", res)
      if (!this.hasRequestSuccess(res)) {
        obj.$message.error(res.data.err)
        return
      }
      obj.$message.success(res.data.msg)
      then(res)
    })
  },
  async GetAllDesks(obj, par, then) {
    let model = this.getRequestModel("mvp", "GetAllDesks", par)
    await this.sendRequestModel(model).then((res) => {
      console.log("GetAllDesks.res", res)
      if (!this.hasRequestSuccess(res)) {
        obj.$message.error(res.data.err)
        return
      }
      obj.$message.success(res.data.msg)
      then(res)
    })
  },
  async GetAllDeskClasses(obj, par, then) {
    let model = this.getRequestModel("mvp", "GetAllDeskClasses", par)
    await this.sendRequestModel(model).then((res) => {
      console.log("GetAllDeskClasses.res", res)
      if (!this.hasRequestSuccess(res)) {
        obj.$message.error(res.data.err)
        return
      }
      obj.$message.success(res.data.msg)
      then(res)
    })
  },
  async GetAllGoodClasses(obj, par, then) {
    let model = this.getRequestModel("mvp", "GetAllGoodClasses", par)
    await this.sendRequestModel(model).then((res) => {
      console.log("GetAllGoodClasses.res", res)
      if (!this.hasRequestSuccess(res)) {
        obj.$message.error(res.data.err)
        return
      }
      obj.$message.success(res.data.msg)
      then(res)
    })
  },
  async GetAllGoods(obj, par, then) {
    let model = this.getRequestModel("mvp", "GetAllGoods", par)
    await this.sendRequestModel(model).then((res) => {
      console.log("GetAllGoods.res", res)
      if (!this.hasRequestSuccess(res)) {
        obj.$message.error(res.data.err)
        return
      }
      obj.$message.success(res.data.msg)
      then(res)
    })
  },
  async AddDesk(obj, par, then) {
    let model = this.getRequestModel("mvp", "AddDesk", par)
    await this.sendRequestModel(model).then((res) => {
      console.log("AddDesk.res", res)
      if (!this.hasRequestSuccess(res)) {
        obj.$message.error(res.data.err)
        return
      }
      obj.$message.success(res.data.msg)
      then(res)
    })
  },
  async AddGoodClass(obj, par, then) {
    let model = this.getRequestModel("mvp", "AddGoodClass", par)
    await this.sendRequestModel(model).then((res) => {
      console.log("AddGoodClass.res", res)
      if (!this.hasRequestSuccess(res)) {
        obj.$message.error(res.data.err)
        return
      }
      obj.$message.success(res.data.msg)
      then(res)
    })
  },
  async DeleteElementSizeInfo(obj, par, then) {
    let model = this.getRequestModel("mvp", "DeleteElementSizeInfo", par)
    await this.sendRequestModel(model).then((res) => {
      console.log("DeleteElementSizeInfo.res", res)
      if (!this.hasRequestSuccess(res)) {
        obj.$message.error(res.data.err)
        return
      }
      obj.$message.success(res.data.msg)
      then(res)
    })
  },
  async GetOrder(obj, par, then) {
    let model = this.getRequestModel("mvp", "GetOrder", par)
    await this.sendRequestModel(model).then((res) => {
      console.log("GetOrder.res", res)
      if (!this.hasRequestSuccess(res)) {
        obj.$message.error(res.data.err)
        return
      }
      obj.$message.success(res.data.msg)
      then(res)
    })
  },
  IsNil(obj) {
    return obj === undefined || obj === null || isNaN(obj)
  }
}
