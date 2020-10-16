package ers

import "fmt"

type MeaningError struct{
	Code int `json:"code"`
	Desc string `json:"desc"`
}

func (e *MeaningError)Error() string{
	return fmt.Sprintf("desc:%s, code: %d",e.Desc,e.Code)
}



var InvalidName = &MeaningError{
	Code: 10000,
	Desc: "非法命名",
}
