package commands

import (
	"telebot/db_helper"
	"telebot/keyboards"
)

type ComandAnswer struct {
	Text     string
	Blocked  bool
	Keyboard any
}

func GetAnswer(commandText string, userId int64) ComandAnswer {

	user := db_helper.NewUser(userId)

	switch commandText {
	case "start":

		if user.Checked {
			return ComandAnswer{"Ви вже пройшли тест. Чим я можу тобі допомогти?", true, keyboards.MainMenuKeyboard()}
		}

		if user.Warnings <= 2 {
			return ComandAnswer{"Вітаю Вас! Для початку прошу пройти невеличкий тест. Слава Україні!", true, keyboards.StartQuizKeyboard(1)}
		} else {
			return ComandAnswer{"Ваш ID заблоковано. Ви не пройшли тест :(", false, ""}
		}

	case "menu":
		if !user.Checked {
			return ComandAnswer{"Для початку прошу пройти невеличкий тест. Слава Україні!", true, keyboards.StartQuizKeyboard(1)}
		}

		if user.Warnings <= 2 {
			return ComandAnswer{"Чим я можу тобі допомогти?", true, keyboards.MainMenuKeyboard()}
		} else {
			return ComandAnswer{"Ваш ID заблоковано. Ви не пройшли тест :(", false, ""}
		}
	case "help":
		return ComandAnswer{"Я звичайний бот, який знає новини, афоризми та може робити передбачення.\n Доступні команди: /start /menu /help", false, ""}
	default:
		if !user.Checked {
			return ComandAnswer{"Для початку прошу пройти невеличкий тест. Слава Україні!", true, keyboards.StartQuizKeyboard(1)}
		}

		if user.Warnings <= 2 {
			return ComandAnswer{"Не знаю такої команди, може меню допоможе?", true, keyboards.MainMenuKeyboard()}
		} else {
			return ComandAnswer{"Ваш ID заблоковано. Ви не пройшли тест :(", false, ""}
		}
	}
}
