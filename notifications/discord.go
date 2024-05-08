package notifications

import "fmt"

type DiscordBot struct {
	webHookUrl string
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

	return sendToWebHook(bot.webHookUrl, body, discordResponseHandler)
}
