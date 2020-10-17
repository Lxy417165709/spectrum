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
}
