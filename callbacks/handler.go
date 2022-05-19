package callbacks

import (
	"log"
	"strconv"
	"strings"
	"telebot/aws_helper"
	"telebot/db_helper"
	"telebot/keyboards"

	telebot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type answer struct {
	text        string
	keyboard    any
	delete_prev bool
}

func HandleCallback(callback telebot.CallbackQuery) answer {

	limit := 21
	text := callback.Data
	chat_id := callback.Message.Chat.ID
	text_2 := callback.Message.Text
	var news_source string

	if strings.Contains(callback.Data, "news_list") {
		text = strings.Fields(callback.Data)[0]
		news_source = strings.Fields(callback.Data)[1]
	}

	switch text {
	case "1_start_test_fail":
		db_helper.AddWarning(chat_id)
		return answer{"Відповідь не вірна золотце. Будеш помилятися, твій ID буде заблокований...", "", true}
	case "2_start_test_fail":
		db_helper.AddWarning(chat_id)
		return answer{"Відповідь не вірна золотце. Будеш помилятися, твій ID буде заблокований...", "", true}
	case "1_start_test_ok":
		return answer{"Слава Нації!", keyboards.StartQuizKeyboard(2), true}
	case "2_start_test_ok":
		db_helper.AddChecked(chat_id)
		return answer{"Непогано! Чим я можу тобі допомогти?", keyboards.MainMenuKeyboard(), true}

	case "news_source_list":
		return answer{"Які саме новини вас цікавлять?", keyboards.NewsSourceKeyboard(), true}

	case "news_list":
		news := db_helper.News(news_source)
		return answer{"1 / 20 " + SourceHeader(news_source) + "\n" + news[0], keyboards.NewsKeyboard(), true}
	case "news_next":

		news_source = HeaderToSource(strings.Fields(text_2)[3])
		news := db_helper.News(news_source)

		page_str := strings.Fields(text_2)

		if len(page_str) == 0 {
			page_str = append(page_str, "1")
		}
		page_num, _ := strconv.Atoi(page_str[0])
		page_num = page_num + 1

		if page_num == limit {
			page_num = 1
		}

		return answer{strconv.Itoa(page_num) + " / 20 " + SourceHeader(news_source) + "\n" + news[page_num-1], keyboards.NewsKeyboard(), true}

	case "news_prev":

		news_source = HeaderToSource(strings.Fields(text_2)[3])
		news := db_helper.News(news_source)

		page_str := strings.Fields(text_2)
		if len(page_str) == 0 {
			page_str = append(page_str, "1")
		}

		page_num, _ := strconv.Atoi(page_str[0])
		page_num = page_num - 1
		log.Println("1 >>>>", page_num)
		if page_num <= 0 {
			page_num = 1
		}
		log.Println("2 >>>>", page_num)
		return answer{strconv.Itoa(page_num) + " / 20 " + SourceHeader(news_source) + "\n" + news[page_num-1], keyboards.NewsKeyboard(), true}

	case "menu_wisdom":
		return answer{aws_helper.Wisdom(), "", false}
	case "menu_vacancies":
		return answer{"VACANCIES", "", false}
	case "menu_oracul":
		return answer{aws_helper.Prediction(), "", false}

	default:
		return answer{"", "", true}
	}
}

func (a *answer) Name() string {
	return a.text
}

func (a *answer) Keyboard() any {
	return a.keyboard
}

func (a *answer) DeletePrev() bool {
	return a.delete_prev
}

func SourceHeader(source_name string) string {

	switch source_name {
	case "news_collection_liga":
		return "Liga"
	case "news_collection_dou":
		return "DOU"
	default:
		return ""
	}
}

func HeaderToSource(header_name string) string {

	switch header_name {
	case "Liga":
		return "news_collection_liga"
	case "DOU":
		return "news_collection_dou"
	default:
		return ""
	}
}
