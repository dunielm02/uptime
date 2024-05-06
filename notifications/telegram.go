package notifications

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func aliveMessage(name string) string {
	return fmt.Sprintf("\u2705 The Service: \"%s\" is Up\u2705", name)
}
func deadMessage(name string) string {
	return fmt.Sprintf("\u203c\ufe0f The Service \"%s\" is Down \u203c\ufe0f", name)
}

type TelegramBot struct {
	token  string
	chatId string
}

func (bot *TelegramBot) AliveNotification(name string) error {
	return bot.sendNotification(aliveMessage(name))
}

func (bot *TelegramBot) DeadNotification(name string) error {
	return bot.sendNotification(deadMessage(name))
}

func (bot *TelegramBot) url() string {
	return fmt.Sprintf("https://api.telegram.org/bot%s", bot.token)
}

func (bot *TelegramBot) sendNotification(message string) error {
	body := map[string]string{
		"chat_id": bot.chatId,
		"text":    message,
	}

	jsonFormatted, _ := json.Marshal(body)

	res, err := http.Post(bot.url()+"/sendMessage", "application/json", bytes.NewBuffer(jsonFormatted))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	read, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var resData map[string]any
	json.Unmarshal(read, &resData)

	if resData["ok"] != true {
		return fmt.Errorf("telegram api did not accept the message\n\t eror_code: %v \n\t description: %v", resData["error_code"], resData["description"])
	}

	return nil
}
