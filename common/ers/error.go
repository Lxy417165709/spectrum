package ers

import (
	"encoding/json"
	"fmt"
)

type Error struct {
	//Code   int32  `json:"code"`
	Detail string `json:"detail"`
}

func (e *Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func New(detailFormat string, parameters ...interface{}) error {
	return &Error{
		//Code:   code,
		Detail: fmt.Sprintf(detailFormat, parameters...),
	}
}
