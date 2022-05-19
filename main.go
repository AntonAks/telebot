package main

import (
	"log"
	"os"
	"telebot/callbacks"
	"telebot/commands"

	telebot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	bot, err := telebot.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := telebot.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		if update.CallbackQuery != nil {
			chat_id := update.CallbackQuery.Message.Chat.ID
			msg_id := update.CallbackQuery.Message.MessageID

			answer := callbacks.HandleCallback(*update.CallbackQuery)
			msg := telebot.NewMessage(chat_id, "")

			if answer.DeletePrev() {
				del := telebot.NewDeleteMessage(chat_id, msg_id)
				bot.Request(del)
			}

			if update.CallbackQuery.Data == "news_next" || update.CallbackQuery.Data == "news_prev" {
				del := telebot.NewDeleteMessage(chat_id, msg_id)
				bot.Request(del)
			}

			msg.Text = answer.Name()
			if answer.Keyboard() != "" {
				msg.ReplyMarkup = answer.Keyboard()
			}

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}

		}

		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		// COMMANDS
		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		} else if update.Message.IsCommand() {
			// Create a new MessageConfig. We don't have text yet,
			// so we leave it empty.
			msg := telebot.NewMessage(update.Message.Chat.ID, "")

			// Extract the command from the Message.
			command_name := update.Message.Command()
			answer := commands.GetAnswer(command_name, update.Message.Chat.ID)
			msg.Text = answer.Text
			if answer.Blocked {
				msg.ReplyMarkup = answer.Keyboard
			}

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
	}
}
