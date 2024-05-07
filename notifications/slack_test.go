package notifications

import "testing"

func TestSlackNotification(t *testing.T) {
	bot := SlackBot{
		webHookUrl: "https://hooks.slack.com/services/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",
	}

	bot.AliveNotification("Hello, World!")
	bot.DeadNotification("Hello, World!")
}
