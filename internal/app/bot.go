package app

import (
	"time"

	tg "gopkg.in/telebot.v4"
)

type Bot struct {
	telegramBot *tg.Bot
}

// Инициализация бота
func NewBot() (*Bot, error) {
	pref := tg.Settings{
		Poller:    &tg.LongPoller{Timeout: 10 * time.Second},
		ParseMode: tg.ModeMarkdown,
	}

	pref.Token = "7997570231:AAEPahDlgyoYcwdrdDZ_YvVRS7mOweCjGlA"

	// Создание бота
	bot, err := tg.NewBot(pref)
	if err != nil {
		return nil, err
	}

	return &Bot{
		telegramBot: bot,
	}, nil
}

func (b *Bot) Run() {
	b.telegramBot.Start()
}

func (b *Bot) Stop() {
	b.telegramBot.Stop()
}
