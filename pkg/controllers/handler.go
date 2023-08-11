package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"mosb_go/pkg/models"
	"net/http"
	"os"
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

	// If the text contains marco, call the `sayPolo` function, which
	// is defined below
	if err := sendBotResponse(body.Message.Chat.ID, strings.ToLower(body.Message.Text)); err != nil {
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

func sendBotResponse(chatID int64, text string) error {
	response := ""
	if fn, ok := models.InputMap[text]; ok {
		response = fn()
	} else {
		fmt.Println("Key not found:", text)
		response = models.InputMap["/notFound"]()
	}

	// Create the request body struct
	reqBody := &sendMessageReqBody{
		ChatID: chatID,
		Text:   response,
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
	bot_token := os.Getenv("BOT_TOKEN")
	if bot_token == "" {
		return bot_token, errors.New("env: Bot token not found in environment variable")
	}

	return bot_token, nil
}

func ConnectBot() {
	http.ListenAndServe(":3000", http.HandlerFunc(Handler))
}
