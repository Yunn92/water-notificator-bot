package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/yunn92/waterNotificationBot/eternal"
)

var Bot eternal.Bot

func init() {
	Bot = eternal.BotInitilisation()
	Bot.BotAPI = initiateBotAPI()
}

func main() {
	log.Printf("Authorized on account %s", Bot.BotAPI.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := Bot.BotAPI.GetUpdatesChan(u)

	for update := range updates {
		userID := update.Message.From.ID

		if update.Message.IsCommand() {
			command := update.Message.Command()

			switch command {
			case "start":
				Bot.SayHello(userID)
			case "settings":
				Bot.ShowSettings(userID)
			case "joke":
				Bot.TellAJoke(userID)
			}

		}
		if update.Message != nil {
			log.Printf("%s[%d]: %s", update.Message.From.UserName, update.Message.From.ID, update.Message.Text)

			data := update.Message.Text
			//userID := update.Message.From.ID
			switch Bot.UserList[userID].Step {
			case "initiation":
				switch data {
				case "Включить напоминание":
					Bot.UserList[userID].ChangeMode()
				}

			}
		}
	}
}

func initiateBotAPI() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI("YourAwesomeToken")
	if err != nil {
		log.Panic(err)
	}

	//bot.Debug = true

	return bot
}
