package main

import (
	"encoding/json"
	"fmt"
)

type RequestError struct {
	StatusCode int
	Err        error
}

func (r *RequestError) ErrorCode() int {
	return r.StatusCode
}

func (r *RequestError) Error() string {
	var errMsg []byte
	var err error
	if errMsg, err = json.Marshal(r); err != nil {
		return fmt.Sprintf("json.Marshal: %v", err)
	}
	return string(errMsg)
}
