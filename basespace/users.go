package basespace

import (
	"encoding/json"
	"errors"
)

type User struct {
	Id             *string  `json:"id,omitempty"`
	Href           *string  `json:"href,omitempty"`
	Name           *string  `json:"name,omitempty"`
	Email          *string  `json:"email,omitempty"`
	GravatarUrl    *string  `json:"gravatarurl,omitempty"`
	DateLastActive *string  `json:"datelastactive,omitempty"`
	DateCreated    *string  `json:"datecreated,omitempty"`
	HrefRuns       *string  `json:"hrefruns,omitempty"`
	HrefProjects   *string  `json:"hrefprojects,omitempty"`
	Roles          []string `json:"roles,omitempty"`
}

// JSON returns a JSON representation of the user.
func (u *User) JSON() string {
	userBytes, err := json.Marshal(u)
	if err != nil {
		return ""
	}

	return string(userBytes)
}

// UsersService handles communication with the BaseSpace Users API
type UsersService struct {
	client *Client
	mapper *Mapper
}

// GetCurrent fetches the current user.
// GET /users/current
func (u *UsersService) GetCurrent() (*User, error) {
	req, err := u.client.NewRequest("GET", "users/current", nil)
	if err != nil {
		return nil, err
	}

	resp, err := u.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.Status != nil && resp.Status.ErrorCode != "" {
		return nil, errors.New(resp.Status.Error())
	}

	user, err := u.mapper.User(resp.Response)
	if err != nil {
		return nil, err
	}

	return user, nil
}
