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
