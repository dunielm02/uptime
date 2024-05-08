package notifications

import (
	"fmt"
	"lifeChecker/config"

	"github.com/mitchellh/mapstructure"
)

type TeamsBot struct {
	WebHookUrl string `mapstructure:"webhook"`
}

func getTeamsBot(cfg config.NotificationChannelsConfig) (*TeamsBot, error) {
	var ret TeamsBot
	err := mapstructure.Decode(cfg.Spec, &ret)
	return &ret, err
}

func (bot *TeamsBot) AliveNotification(name string) error {
	return bot.sendNotification(aliveMessage(name))
}

func (bot *TeamsBot) DeadNotification(name string) error {
	return bot.sendNotification(deadMessage(name))
}

func TeamsResponseHandler(resp []byte) error {
	fmt.Println(string(resp))

	return nil
}

func (bot *TeamsBot) sendNotification(message string) error {
	body := map[string]string{
		"text": message,
	}

	return sendToWebHook(bot.WebHookUrl, body, TeamsResponseHandler)
}
