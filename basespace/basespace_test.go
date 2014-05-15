package basespace

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient(nil)

	if client.ApiURL.String() != defaultApiURL {
		t.Errorf("NewClient ApiURL = %v, expected %v", client.ApiURL.String(), defaultApiURL)
	}
}
