package basespace

import (
	"fmt"
)

type ResponseStatus struct {
	ErrorCode string `json:"errorcode"`
	Message   string `json:"message"`
}

func (r *ResponseStatus) Error() string {
	return fmt.Sprintf("%v: %v", r.ErrorCode, r.Message)
}

type ApiResponse struct {
	Response map[string]interface{} `json:"response,omitempty"`
	Status   *ResponseStatus        `json:"responsestatus,omitempty"`
}
