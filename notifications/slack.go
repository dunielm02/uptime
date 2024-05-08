package notifications

import (
	"fmt"
	"lifeChecker/config"

	"github.com/mitchellh/mapstructure"
)

type SlackBot struct {
	WebHookUrl string `mapstructure:"webhook"`
}

func getSlackBot(cfg config.NotificationChannelsConfig) (*SlackBot, error) {
	var ret SlackBot
	err := mapstructure.Decode(cfg.Spec, &ret)
	return &ret, err
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

	return sendToWebHook(bot.WebHookUrl, body, slackResponseHandler)
}
