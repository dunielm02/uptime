package notifications

import (
	"lifeChecker/tests/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotifications(t *testing.T) {
	mocks.StartWebhooksMock()

	t.Run("Sending a Message to Telegram", testTelegram)

	t.Run("Sending a Message to Slack", testSlack)

	t.Run("Sending a Message to Discord", testDiscord)

	t.Run("Sending a Message to Microsoft Teams", testMicrosoftTeams)

	t.Run("Sending a Message to Google Workspace", testGoogleWorkspace)
}

func testTelegram(t *testing.T) {
	bot := TelegramBot{
		Token:  os.Getenv("TELEGRAM_BOT"),
		ChatId: "827317315",
	}

	err := bot.AliveNotification("Testing")
	assert.Nil(t, err)
	err = bot.DeadNotification("Testing")
	assert.Nil(t, err)
}

func testSlack(t *testing.T) {
	bot := SlackBot{
		WebHookUrl: "http://localhost:3202/slack",
	}

	err := bot.AliveNotification("Testing")
	assert.Nil(t, err)
	err = bot.DeadNotification("Testing")
	assert.Nil(t, err)
}

func testDiscord(t *testing.T) {
	bot := DiscordBot{
		WebHookUrl: "http://localhost:3202/discord",
	}

	err := bot.AliveNotification("Testing")
	assert.Nil(t, err)
	err = bot.DeadNotification("Testing")
	assert.Nil(t, err)
}

func testMicrosoftTeams(t *testing.T) {
	bot := TeamsBot{
		WebHookUrl: "http://localhost:3202/teams",
	}

	err := bot.AliveNotification("Testing")
	assert.Nil(t, err)
	err = bot.DeadNotification("Testing")
	assert.Nil(t, err)
}

func testGoogleWorkspace(t *testing.T) {
	bot := GwsBot{
		WebHookUrl: "http://localhost:3202/gws",
	}

	err := bot.AliveNotification("Testing")
	assert.Nil(t, err)
	err = bot.DeadNotification("Testing")
	assert.Nil(t, err)
}
