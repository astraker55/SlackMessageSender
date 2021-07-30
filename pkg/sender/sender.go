package sender

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type MessageSender struct {
	channels map[string]string
}

func newMessageSender() *MessageSender {
	return &MessageSender{channels: map[string]string{}}
}

func InitMessageSender() (*MessageSender, error) {
	conf, err := os.Open("hooks.json")
	if err != nil {
		return nil, errors.New("File 'hooks.json' doesn't exist")
	}
	ms := newMessageSender()
	decoder := json.NewDecoder(conf)
	decoder.Decode(&ms.channels)
	return *&ms, nil
}

func (m *MessageSender) SendMessage(channel string, data string) (int, error) {
	message := fmt.Sprintf(`{"text": "%s"}`, data)
	fmt.Println("Sending message", message)
	ch_url, ok := m.channels[channel]
	if !ok {
		return 400, errors.New("Unknown channel")
	}
	resp, _ := http.Post(ch_url, "application/json", strings.NewReader(message))
	if resp.StatusCode != 200 {
		return resp.StatusCode, errors.New("Can't reach channel")
	}
	return resp.StatusCode, nil
}
