package main

import (
	"encoding/json"
	"log"
	"os"
)

var (
	config Configuration
)

type Configuration struct {
	TelegramBotToken 	string 		`json:"TelegramBotToken"`
	LaunchURL 			string 		`json:"LaunchURL"`
}

func init() {
	config = ReadConfigJson()
}

func ReadConfigJson() Configuration{
	file, _ := os.Open("config.json")
	defer func() {
		if r := recover(); r != nil {
            log.Println("defer recovered from panic:", r)
        }
		file.Close()
	}()

	decoder := json.NewDecoder(file)
	err := decoder.Decode(&config)
	if err != nil {
		log.Println("error:", err)
		panic(err)
	}

	return config
}