package errorshelper

import (
	"encoding/json"
	"log"
)

type HttpError struct {
	Code        uint16 `json:"code"`
	Err         string `json:"err"`
	Context     string `json:"-"`
	ErrInternal string `json:"-"`
}

func (item *HttpError) Error() string {
	b, err := json.Marshal(item)
	if err != nil {
		log.Println(err, item.Err)
	}
	return string(b)
}

func NewHttpError(err string, code uint16, errInternal string) *HttpError {
	return &HttpError{Code: code, Err: err, ErrInternal: errInternal}
}
