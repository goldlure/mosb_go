package utils

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	BotToken string `json:"bot_token"`
}

func LoadConfig(filename string) (Configuration, error) {
	var config Configuration
	file, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	return config, err
}
