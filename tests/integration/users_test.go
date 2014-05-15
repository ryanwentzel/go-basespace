package tests

import (
	"log"
	"testing"
)

func TestUsers_GetCurrent(t *testing.T) {
	client := createClient()
	u, err := client.Users.GetCurrent()
	if err != nil {
		t.Errorf("GetCurrentUser - %v", err)
	}

	log.Printf("GetCurrentUser - %v", u.JSON())
}
