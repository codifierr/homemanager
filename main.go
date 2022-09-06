package main

import (
	"conf"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("Initializing home manager!")

	//get config
	config, err := conf.GetConf()
	if err != nil {
		log.Fatal().AnErr("Error", err).Msg("Fatal Error")
	}

	log.Info().Interface("Config", config).Msg("Config")

	bot, err := tgbotapi.NewBotAPI(config.Telegram.Token)
	if err != nil {
		log.Fatal().AnErr("Error", err)
	}

	log.Info().Str("Authorized on account", bot.Self.UserName).Msg("UserName")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Info().Str("UserName", update.Message.From.UserName).Str("Message", update.Message.Text).Msg("Received message")

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ack")
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
	}
}
