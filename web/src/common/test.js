/* eslint-disable */

import cst from "./cst";

export default {
  blankGood: {
    mainElement: {
      name: "请输入",
      selectedIndex: 0,
      sizeInfos: [{
        id: 0,
        size: "小规格",
        price: "10.0",
        pictureStorePath: ""
      }],
    },
    attachElements: [],
    expenseInfo: {
      nonFavorExpense: 0.0,
      Expense: 0,
      CheckOutAt: cst.TIMESTAMP.NIL,
    },
    favors:[]
  },
  blankGoodClass: {
    name: "请输入",
    pictureStorePath: "",
  },
  blankGoodOption: {
    name: "请输入",
    type: cst.ELEMENT_TYPE.OPTION,
    selectedIndex: 0,
    sizeInfos: [{
      id: 0,
      size: "",
      price: "0.0",
      pictureStorePath: ""
    }],
  },
  blankGoodIngredient: {
    name: "请输入",
    type: cst.ELEMENT_TYPE.INGREDIENT,
    selectedIndex: 0,
    sizeInfos: [{
      id: 0,
      size: "",
      price: "0.0",
      pictureStorePath: ""
    }],
  },
  goodOptionClasses: [
    {
      name: cst.ATTACH_CLASS_NAME.GOOD_OPTION_CLASS_NAME,
      // pictureStorePath: "https://ss1.bdstatic.com/70cFuXSh_Q1YnxGkpoWK1HF6hhy/it/u=2186952895,3925242332&fm=26&gp=0.jpg",
      pictureStorePath: "api/file/" + "static/upload/温热1.jpeg",
    },
    {
      name: cst.ATTACH_CLASS_NAME.GOOD_INGREDIENT_CLASS_NAME,
      pictureStorePath: "api/file/" + "static/upload/小红1.jpeg",
      // pictureStorePath: "https://ss1.bdstatic.com/70cFuXSh_Q1YnxGkpoWK1HF6hhy/it/u=2186952895,3925242332&fm=26&gp=0.jpg",
    }
  ],
  blankDesk: {
    space: {
      className: "",
      name: "请输入",
      price: "10",
      pictureStorePath: ""
    }
  },
  blankDeskClass: {
    name: "请输入",
    pictureStorePath: "",
  },
  blankSizeInfo: {
    id: 0,
    size: "规格名",
    price: "0",
    pictureStorePath: ""
  }
}
