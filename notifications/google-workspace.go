package notifications

import (
	"fmt"
	"lifeChecker/config"

	"github.com/mitchellh/mapstructure"
)

// GwsBot contain the functions to send a notification to
// Google Workspace.
type GwsBot struct {
	WebHookUrl string `mapstructure:"webhook"`
}

func getGwsBot(cfg config.NotificationChannelsConfig) (*GwsBot, error) {
	var ret GwsBot
	err := mapstructure.Decode(cfg.Spec, &ret)
	return &ret, err
}

func (bot *GwsBot) AliveNotification(name string) error {
	return bot.sendNotification(aliveMessage(name))
}

func (bot *GwsBot) DeadNotification(name string) error {
	return bot.sendNotification(deadMessage(name))
}

func GwsResponseHandler(resp []byte) error {
	fmt.Println(string(resp))

	return nil
}

func (bot *GwsBot) sendNotification(message string) error {
	body := map[string]string{
		"text": message,
	}

	return sendToWebHook(bot.WebHookUrl, body, GwsResponseHandler)
}
