package notifications

import (
	"fmt"
	"lifeChecker/config"

	"github.com/mitchellh/mapstructure"
)

type DiscordBot struct {
	WebHookUrl string `mapstructure:"webhook"`
}

func getDiscordBot(cfg config.NotificationChannelsConfig) (*DiscordBot, error) {
	var ret DiscordBot
	err := mapstructure.Decode(cfg.Spec, &ret)
	return &ret, err
}

func (bot *DiscordBot) AliveNotification(name string) error {
	return bot.sendNotification(aliveMessage(name))
}

func (bot *DiscordBot) DeadNotification(name string) error {
	return bot.sendNotification(deadMessage(name))
}

func discordResponseHandler(resp []byte) error {
	fmt.Println(string(resp))
	// if string(resp) != "ok" {
	// 	return fmt.Errorf("error sending Slack message: %s", string(resp))
	// }

	return nil
}

func (bot *DiscordBot) sendNotification(message string) error {
	body := map[string]string{
		"content": message,
	}

	return sendToWebHook(bot.WebHookUrl, body, discordResponseHandler)
}
