package keyboards

import telebot "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func StartQuizKeyboard(n int) any {
	var keyboard = telebot.NewInlineKeyboardMarkup()

	switch n {
	case 1:
		keyboard := telebot.NewInlineKeyboardMarkup(
			telebot.NewInlineKeyboardRow(
				telebot.NewInlineKeyboardButtonData("Слава Богам!", "1_start_test_fail"),
			),
			telebot.NewInlineKeyboardRow(
				telebot.NewInlineKeyboardButtonData("Героям слава!", "1_start_test_ok"),
			),
			telebot.NewInlineKeyboardRow(
				telebot.NewInlineKeyboardButtonData("Хай живе!", "1_start_test_fail"),
			),
			telebot.NewInlineKeyboardRow(
				telebot.NewInlineKeyboardButtonData("Бендері слава!", "1_start_test_fail"),
			),
		)
		return keyboard
	case 2:
		keyboard := telebot.NewInlineKeyboardMarkup(
			telebot.NewInlineKeyboardRow(
				telebot.NewInlineKeyboardButtonData("Слава Україні!", "2_start_test_fail"),
			),
			telebot.NewInlineKeyboardRow(
				telebot.NewInlineKeyboardButtonData("Смерть ворогам!", "2_start_test_ok"),
			),
			telebot.NewInlineKeyboardRow(
				telebot.NewInlineKeyboardButtonData("Смерть путіну!", "2_start_test_fail"),
			),
			telebot.NewInlineKeyboardRow(
				telebot.NewInlineKeyboardButtonData("Смерть москалям!", "2_start_test_fail"),
			),
		)
		return keyboard
	default:
		return keyboard
	}
}

func MainMenuKeyboard() any {

	menu := telebot.NewInlineKeyboardMarkup(
		telebot.NewInlineKeyboardRow(
			telebot.NewInlineKeyboardButtonData("Новини", "news_source_list"),
		),
		telebot.NewInlineKeyboardRow(
			telebot.NewInlineKeyboardButtonData("Мудрість", "menu_wisdom"),
		),
		// telebot.NewInlineKeyboardRow(
		// 	telebot.NewInlineKeyboardButtonData("Вакансії", "menu_vacancies"),
		// ),
		telebot.NewInlineKeyboardRow(
			telebot.NewInlineKeyboardButtonData("Передбачення", "menu_oracul"),
		),
	)
	return menu
}

func NewsKeyboard() any {

	menu := telebot.NewInlineKeyboardMarkup(
		telebot.NewInlineKeyboardRow(
			telebot.NewInlineKeyboardButtonData("<<<", "news_prev"),
			telebot.NewInlineKeyboardButtonData(">>>", "news_next"),
		),
		telebot.NewInlineKeyboardRow(
			telebot.NewInlineKeyboardButtonData("Назад до переліку", "news_source_list"),
		))
	return menu
}

func NewsSourceKeyboard() any {
	menu := telebot.NewInlineKeyboardMarkup(
		telebot.NewInlineKeyboardRow(
			telebot.NewInlineKeyboardButtonData("Головне в Україні і світі (LIGA)", "news_list news_collection_liga"),
		),
		telebot.NewInlineKeyboardRow(
			telebot.NewInlineKeyboardButtonData("Новини зі світу IT (DOU)", "news_list news_collection_dou"),
		),
	)
	return menu
}
