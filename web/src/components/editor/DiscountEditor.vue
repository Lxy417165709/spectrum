<!-- eslint-disable -->
<template>
	<div>
		<el-form-item label="优惠">
			<el-radio v-model="curFavorIndex"  v-for="(favor,index) in selectableFavors" :label="index"
					  :key="index" >
				{{ favor.name }}
			</el-radio>
		</el-form-item>
		<component :is="selectableFavors[curFavorIndex].component" @confirmFavor="addFavor"></component>
		<el-form-item label="已选折扣">
			<el-tag v-for="(favor,index) in favors" @close="delFavor(index,favor)"  :key="index" closable style="margin-right: 10px">
				{{ getFavorTagName(favor) }}
			</el-tag>
		</el-form-item>
	</div>
</template>

<script>
/* eslint-disable */
import DiscountComponent from "../baby/DiscountComponent";
import FullReductionComponent from "../baby/FullReductionComponent";
import FreeComponent from "../baby/FreeComponent";
import cst from "../../common/cst";
import utils from "../../common/utils";

export default {
	name: "DiscountEditor",
	props: {
		index: Number,
		chargeableObjName: String,
		favors: Array
	},
	data() {
		return {
			curFavorIndex: 0,
			selectableFavors: [
				{
					name: cst.FAVOR_TYPE.NONE.NAME,
				},
				{
					name: cst.FAVOR_TYPE.REBATE.NAME,
					component: DiscountComponent
				},
				{
					name: cst.FAVOR_TYPE.FULL_REDUCTION.NAME,
					component: FullReductionComponent
				},
				{
					name: cst.FAVOR_TYPE.FREE.NAME,
					component: FreeComponent
				}
			],
		}
	},
	methods: {
		getFavorTagName(favor) {
			return utils.GetFavorTagName(favor)
		},
		addFavor(favor) {
			if (this.chargeableObjName === "order") {
				this.$emit("addFavorForOrder", favor, this.index)
			}
			if (this.chargeableObjName === "good") {
				this.$emit("addFavorForGood", favor, this.index)
			}
		},
		delFavor(favorIndex,favor) {
			if (this.chargeableObjName === "order") {
				this.$emit("delFavorForOrder", favor, this.index)
			}
			if (this.chargeableObjName === "good") {
				this.$emit("delFavorForGood", favorIndex)
			}
		}
	}
}
</script>

<style scoped>

</style>
