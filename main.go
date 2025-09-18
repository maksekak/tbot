package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	firstMenu       = "<b>Menu 1</b>\n\nменю меню меню меню ура."
	SecondMenu      = "<b>Menu 2</b>\n\nвап."
	nextButton      = "вперед"
	backButton      = "назад"
	text            string
	msg             tgbotapi.MessageConfig
	firstMenuMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow( //кнопка
			tgbotapi.NewInlineKeyboardButtonData(nextButton, nextButton),
		),
	)
)

/*
	func knoki(query *tgbotapi.CallbackQuery){
		var text string
		if query.Data == nextButton {
					text = SecondMenu



				}else if query.Data == backButton {
			text = firstMenu

		}

}
*/
func main() {

	bot, err := tgbotapi.NewBotAPI("7952741128:AAFiVKZsXqqxS9427LpKGGd9QmjzxnRWVoE")

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	up := bot.GetUpdatesChan(u)
	tgbotapi.NewInlineKeyboardButtonData("fdgfdsg", "")
	for update := range up {

		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			//chatId=tgbotapi.NewMessage(update.Message.Chat.ID,)
			text = firstMenu

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text) //вывод меню
			msg.ParseMode = tgbotapi.ModeHTML                        //как в хтмл если теги добавить текст поеняется
			msg.ReplyMarkup = firstMenuMarkup                        //вывод кнопки
			fmt.Println(update.CallbackQuery)

			bot.Send(msg)

		}
		if update.CallbackQuery != nil {
			text = SecondMenu
		} else if update.CallbackQuery == nil {
			text = firstMenu
		}
		fmt.Println(text)
		fmt.Println(update.CallbackQuery)
	}
}
