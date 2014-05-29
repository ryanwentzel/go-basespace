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

type ApiListResponse struct {
	Items          []interface{} `json:"items,omitempty"`
	DisplayedCount int           `json:"displayedcount"`
	TotalCount     int           `json:"totalcount"`
	Offset         int           `json:"offset"`
	Limit          int           `json:"limit"`
	SortDir        string        `json:"sortdir"`
	SortBy         string        `json:"sortby"`
}

type ApiResponse struct {
	Response map[string]interface{} `json:"response,omitempty"`
	Status   *ResponseStatus        `json:"responsestatus,omitempty"`
}
