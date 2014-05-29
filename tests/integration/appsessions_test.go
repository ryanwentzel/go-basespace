package tests

import (
  "log"
  "testing"
)

func TestAppSessions_List(t *testing.T) {
  client := createClient()
  _, resp, err := client.AppSessions.List()
  if err != nil {
    t.Fatalf("AppSessions.List returned error: %v", err)
  }

  log.Printf("AppSessions.List returned %v items", resp.DisplayedCount)
}
