package basespace

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type ClientSettings struct {
	ApiURL      string `json:"apiUrl"`
	AccessToken string `json:"accessToken"`
}

func LoadSettings(filePath string) (*ClientSettings, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Could not read file %v", filePath)
		log.Println("")
	}

	var settings ClientSettings
	json.Unmarshal(file, &settings)

	return &settings, nil
}
