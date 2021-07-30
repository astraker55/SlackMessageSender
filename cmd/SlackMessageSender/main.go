package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/astraker55/SlackMessageSender/pkg/sender"
)

func main() {
	chanList := struct {
		Channels []struct {
			Channel string `json:"channel"`
			Text    string `json:"text"`
		} `json:"channels"`
	}{}
	sender, err := sender.InitMessageSender()
	if err != nil {
		fmt.Print(err.Error())
	}
	f, _ := os.Open("message.json")
	byte_msgs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Print(err)
	}
	err = json.Unmarshal(byte_msgs, &chanList)
	if err != nil {
		fmt.Print(err)
	}
	for _, obj := range chanList.Channels {
		sender.SendMessage(obj.Channel, obj.Text)
	}
}
