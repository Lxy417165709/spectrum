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
  goodClassToPbGoodClass(goodClass) {
    console.log("goodClassToPbGoodClass", "goodClass:", goodClass)
    let pbGoodClass = {
      name: goodClass.name,
      pictureStorePath: goodClass.pictureStorePath,
      // goods: this.goodsToPbGoods(goodClass.goods)
    }
    console.log("goodClassToPbGoodClass", "pbGoodClass:", pbGoodClass)
    return pbGoodClass
  },

  PbGoodToGood(pbGood) {
    console.log("utils.pbGoodToGood", "pbGood:", pbGood)
    let good = {
      id: pbGood.id,
      name: pbGood.mainElement.name,
    }

  },

  goodToPbGood(good) {
    let pbGood = {
      id: good.id,
      mainElement: this.getGoodPbMainElement(good),
      attachElements: this.getGoodPbAttachElements(good),
      favors: this.getGoodPbFavors(good),
      expenseInfo: {}
    }
    console.log("utils.goodToPbGood", "pbGood:", pbGood)
    return pbGood
  },
  getGoodPbMainElement(good) {
    let pbSizeInfos = []
    for (let i = 0; i < good.sizeInfos.length; i++) {
      pbSizeInfos.push(this.sizeInfoToPbSizeInfo(good.curSizeIndex, i, good.sizeInfos[i]))
    }
    return {
      name: good.name,
      sizeInfos: pbSizeInfos,
      type: 0,
    }
  },
  getGoodPbAttachElements(good) {
    let pbAttachElements = []
    for (let i = 0; i < good.attachElements.length; i++) {
      pbAttachElements.push(this.attachElementToPbAttachElement(good.attachElements[i]))
    }
    return pbAttachElements
  },
  getGoodPbFavors(good) {
    let pbFavors = []
    for (let i = 0; i < good.favors.length; i++) {
      pbFavors.push(this.favorToPbFavor(good.favors[i]))
    }
    return pbFavors
  },
  attachElementToPbAttachElement(attachElement) {
    return {
      name: attachElement.name,
      type: 1,
      sizeInfos: this.sizeInfosToPbSizeInfos(attachElement.curSizeIndex, attachElement.sizeInfos),
    }
  },
  favorToPbFavor(favor) {
    return {
      name: favor.name,
      // type: 2,
      sizeInfos: this.sizeInfosToPbSizeInfos(favor.curSizeIndex, favor.sizeInfos),
    }
  },
  sizeInfosToPbSizeInfos(defaultSelectIndex, sizeInfos) {
    let pbSizeInfos = []
    if (sizeInfos === undefined) {
      return pbSizeInfos
    }
    for (let i = 0; i < sizeInfos.length; i++) {
      pbSizeInfos.push(this.sizeInfoToPbSizeInfo(defaultSelectIndex, i, sizeInfos[i]))
    }
    return pbSizeInfos
  },
  sizeInfoToPbSizeInfo(defaultSelectIndex, curIndex, sizeInfo) {
    return {
      size: sizeInfo.name,
      price: sizeInfo.price,
      pictureStorePath: sizeInfo.pictureStorePath,
      isSelected: defaultSelectIndex === curIndex,
    }
  },

  NewBlankSizeInfo(name) {
    return {
      size: name,
      price: "0"
    }
  }

}
