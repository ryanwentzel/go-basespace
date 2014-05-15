package basespace

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestLoadSettings(t *testing.T) {
	var json = `{
    "apiUrl": "http://api.mybasespace.com/",
    "accessToken": "myAccessToken"
  }
  `

	file, err := ioutil.TempFile("", "bssettings")
	if err != nil {
		t.Errorf("LoadSettings could not create temp file: %v", err)
	}

	defer file.Close()

	_, e := file.WriteString(json)
	if e != nil {
		t.Errorf("LoadSettings could not write json: %v", e)
	}

	file.Sync()

	name := file.Name()
	fmt.Println("Attempting to read from", name)
	settings, err := LoadSettings(name)
	if err != nil {
		t.Errorf("LoadSetttings: %v", err)
	}

	if settings.ApiURL != "http://api.mybasespace.com/" {
		t.Errorf("LoadSettings ApiURL = %v, expected %v", settings.ApiURL, "http://api.mybasespace.com/")
	}
}
