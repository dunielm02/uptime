package notifications

import (
	"encoding/json"
	"fmt"
	"lifeChecker/config"

	"github.com/mitchellh/mapstructure"
)

type TelegramBot struct {
	Token  string `mapstructure:"token"`
	ChatId string `mapstructure:"chat-id"`
}

func getTelegramBot(cfg config.NotificationChannelsConfig) (*TelegramBot, error) {
	var ret TelegramBot
	err := mapstructure.Decode(cfg.Spec, &ret)
	return &ret, err
}

func (bot *TelegramBot) AliveNotification(name string) error {
	return bot.sendNotification(aliveMessage(name))
}

func (bot *TelegramBot) DeadNotification(name string) error {
	return bot.sendNotification(deadMessage(name))
}

func (bot *TelegramBot) url() string {
	return fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates", bot.Token)
}

func telegramResponseHandler(respData []byte) error {
	var res map[string]any
	err := json.Unmarshal(respData, &res)
	if err != nil {
		return err
	}

	if res["ok"] != true {
		return fmt.Errorf("telegram api did not accept the message\n\t eror_code: %v \n\t description: %v", res["error_code"], res["description"])
	}
	return nil
}

func (bot *TelegramBot) sendNotification(message string) error {
	body := map[string]string{
		"chat_id": bot.ChatId,
		"text":    message,
	}

	return sendToWebHook(bot.url(), body, telegramResponseHandler)
}
