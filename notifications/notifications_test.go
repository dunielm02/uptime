package notifications

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotifications(t *testing.T) {
	t.Run("Sending a Message to Telegram", testTelegram)
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
