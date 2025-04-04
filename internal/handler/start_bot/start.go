package handler

import (
	tg "gopkg.in/telebot.v4"
)

// Обработчик команды /start
func HandleStart(c tg.Context) error {
	// Проверка на необрабатываемые сообщения
	// Проверка на спам

	// Создание меню и веб-приложения
	menu := &tg.ReplyMarkup{ResizeKeyboard: true}
	webApp := &tg.WebApp{URL: "https://localhost:3000/"}

	var appButton, channelButton tg.Btn
	var text string

	appButton = menu.WebApp("Запустить", webApp)
	channelButton = menu.URL("Сообщество", "https://t.me/ton_argon_ru")

	text = "Заупустить приложение"

	// Добавление кнопок в меню
	menu.Inline(
		menu.Row(appButton),
		menu.Row(channelButton),
	)

	return c.Send(text, menu)
}
