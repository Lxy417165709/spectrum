package controller

var curOrderID int // 这里不考虑高并发

func getOrderID() int {
	return curOrderID
}

func nextOrderID() {
	curOrderID++
}

var curThingID int // 这里不考虑高并发

func getThingID() int {
	return curThingID
}

func nextThingID() {
	curThingID++
}
