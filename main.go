package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	utils "mosb_go/utils"
	"net/http"
	"strings"
)

type webhookReqBody struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

func Handler(res http.ResponseWriter, req *http.Request) {
	body := &webhookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		fmt.Println("could not decode request body", err)
		return
	}

	if !strings.Contains(strings.ToLower(body.Message.Text), "marco") {
		return
	}

	// If the text contains marco, call the `sayPolo` function, which
	// is defined below
	if err := sayPolo(body.Message.Chat.ID); err != nil {
		fmt.Println("error in sending reply:", err)
		return
	}

	// log a confirmation message if the message is sent successfully
	fmt.Println("reply sent")
}

type sendMessageReqBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func sayPolo(chatID int64) error {
	// Create the request body struct
	reqBody := &sendMessageReqBody{
		ChatID: chatID,
		Text:   "Polo!!",
	}
	// Create the JSON body from the struct
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	// Get token
	token, err := getBotToken()
	if err != nil {
		return err
	}
	link := "https://api.telegram.org/bot" + token + "/sendMessage"

	// Send a post request with your token
	res, err := http.Post(link, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	return nil
}

func getBotToken() (string, error) {
	config, err := utils.LoadConfig("config.json")
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return "", err
	}

	return config.BotToken, nil
}

// FInally, the main funtion starts our server on port 3000
func main() {

	http.ListenAndServe(":3000", http.HandlerFunc(Handler))
}
