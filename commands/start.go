package commands

type StartQuestions struct {
	name          string
	next_question int
}

func Start(number int) StartQuestions {

	switch number {
	case 0:
		n := "Вітаю Вас! Для початку прошу пройти невеличкий тест."
		stq := StartQuestions{n, 1}
		return stq
	case 1:
		n := "Слава Україні!"
		stq := StartQuestions{n, 2}
		return stq
	case 2:
		n := "Слава Нації!!!"
		stq := StartQuestions{n, 3}
		return stq
	case 3:
		n := "Що таке Паляниця?"
		stq := StartQuestions{n, 4}
		return stq
	default:
		n := "Що таке Паляниця?"
		stq := StartQuestions{n, 5}
		return stq
	}

}
