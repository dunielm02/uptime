package notifications

import (
	"fmt"
	"lifeChecker/config"
)

type NotificationChannel interface {
	AliveNotification(string) error
	DeadNotification(string) error
}

func GetNotificationChannelFromConfig(cfg config.NotificationChannelsConfig) (NotificationChannel, error) {
	switch cfg.Type {
	case "telegram":
		return getTelegramBot(cfg)
	case "slack":
		return getSlackBot(cfg)
	case "discord":
		return getDiscordBot(cfg)
	case "microsoft-teams":
		return getTeamsBot(cfg)
	case "google-workspace":
		return getGwsBot(cfg)
	}

	return nil, fmt.Errorf("type \"%s\" not found", cfg.Type)
}
