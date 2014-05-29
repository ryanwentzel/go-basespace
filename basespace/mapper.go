package basespace

import (
	"github.com/mitchellh/mapstructure"
)

type Mapper struct{}

func (m *Mapper) User(data map[string]interface{}) (*User, error) {
	var user User
	err := mapstructure.Decode(data, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *Mapper) ListResponse(data map[string]interface{}) (*ApiListResponse, error) {
	var resp ApiListResponse
	err := mapstructure.Decode(data, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil;
}

func (m *Mapper) AppSessions(response ApiListResponse) (*[]AppSession, error) {
	var appSessions []AppSession
	e := mapstructure.Decode(response.Items, &appSessions)
	if e != nil {
		return nil, e
	}

	return &appSessions, nil
}
