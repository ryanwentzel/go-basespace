package tests

import (
	"log"
	"testing"
)

func TestUsers_GetCurrent(t *testing.T) {
	client := createClient()
	u, err := client.Users.GetCurrent()
	if err != nil {
		t.Errorf("Users.GetCurrent returned error: %v", err)
	}

	log.Printf("Users.GetCurrent returned: %v", *u.Email)
}
