package notifications

import (
	"fmt"
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

func slackResponseHandler(resp []byte) error {
	if string(resp) != "ok" {
		return fmt.Errorf("error sending Slack message: %s", string(resp))
	}

	return nil
}

func (bot *SlackBot) sendNotification(message string) error {
	body := map[string]string{
		"text": message,
	}

	return sendToWebHook(bot.webHookUrl, body, slackResponseHandler)
}
