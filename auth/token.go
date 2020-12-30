package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"slack-topic/utils"
)

type Token struct {
	Value string `json:"token"`
}

func GetJSONToken() Token {
	filename := utils.GetFileName()
	jsonFile, err := os.Open(filename)

	if err != nil {
		fmt.Println("error loading file")
		os.Exit(1)
	}
	defer jsonFile.Close()
	tokenValue, _ := ioutil.ReadAll(jsonFile)

	var token Token

	json.Unmarshal(tokenValue, &token)

	return token
}
