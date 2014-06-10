package basespace

import (
	"encoding/json"
	"errors"
)

type AppSession struct {
	Id            *string `json:"id,omitempty"`
	Name          *string `json:"name,omitempty"`
	Href          *string `json:"href,omitempty"`
	UserCreatedBy *User   `json:"usercreatedby,omitempty"`
	Status        *string `json:"status,omitempty"`
	StatusSummary *string `json:statussummary,omitempty"`
	DateCreated   *string `json:datecreated,omitempty"`
	ModifiedOn    *string `json:modifiedon,omitempty"`
	TotalSize     int64   `json:"totalsize,omitempty"`
}

func (a *AppSession) JSON() string {
	data, err := json.Marshal(a)
	if err != nil {
		return ""
	}

	return string(data)
}

type AppSessionsService struct {
	client *Client
	mapper *Mapper
}

// List returns the AppSessions associated with the current user
func (s *AppSessionsService) List() ([]AppSession, *ApiListResponse, error) {
	req, err := s.client.NewRequest("GET", "users/current/appsessions", nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	if resp.Status != nil && resp.Status.ErrorCode != "" {
		return nil, nil, errors.New(resp.Status.Error())
	}

	listResponse, err := s.mapper.ListResponse(resp.Response)
	if err != nil {
		return nil, nil, err
	}

	list, err := s.mapper.AppSessions(*listResponse)
	if err != nil {
		return nil, listResponse, err
	}

	return *list, listResponse, nil
}
