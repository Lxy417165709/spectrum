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
  getGoodClassName(classType) {
    switch (classType) {
      case undefined:
        return "单品"
      case 0:
        return "单品"
      case 1:
        return "附属品"
    }
    return "未定义"
  },
  hasData(obj) {
    return obj !== undefined && obj !== null
  },
  goodToPbGood(good) {
    console.log("utils.goodToPbGood", good)
    let pbSizeInfos = []
    for (let i = 0; i < good.sizeInfos.length; i++) {
      pbSizeInfos.push(this.sizeInfoToPbSizeInfo(good.curSizeIndex, i, good.sizeInfos[i]))
    }
    return {
      mainElement: {
        name: good.name,
        sizeInfos: pbSizeInfos,
      },
      // todo: 还有一些字段
    }
  },
  sizeInfoToPbSizeInfo(defaultSelectIndex, curIndex, sizeInfo) {
    return {
      size: sizeInfo.name,
      price: sizeInfo.price,
      pictureStorePath: sizeInfo.pictureStorePath,
      isSelected: defaultSelectIndex === curIndex,
    }
  }
}
