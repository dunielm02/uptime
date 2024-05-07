package notifications

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SlackBot struct {
	webHookUrl string
}

func (bot *SlackBot) AliveNotification(name string) error {
	return bot.sendNotification(aliveMessage(name))
}

func (bot *SlackBot) DeadNotification(name string) error {
	return bot.sendNotification(deadMessage(name))
}

func (bot *SlackBot) sendNotification(message string) error {
	body := map[string]string{
		"text": message,
	}

	jsonFormatted, _ := json.Marshal(body)

	res, err := http.Post(bot.webHookUrl, "application/json", bytes.NewBuffer(jsonFormatted))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	read, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if string(read) != "ok" {
		return fmt.Errorf("error sending Slack message: %s", string(read))
	}

	return nil
}
