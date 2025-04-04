package app

import (
	"context"
	"fmt"
	"ss/internal/config"
	"sync"
)

type App struct {
	bot  *Bot
	web  *Web
	conf *config.AppConf

	ctx context.Context
	wg  *sync.WaitGroup
}

func NewApp(ctx context.Context, w *sync.WaitGroup, conf *config.AppConf) *App {
	bot, err := NewBot()
	if err != nil {
		fmt.Println(err)
	}

	server := NewServer()

	app := &App{
		bot:  bot,
		web:  server,
		ctx:  ctx,
		wg:   w,
		conf: conf,
	}

	return app
}

func (app *App) RunApp() {
	app.bot.Run()
	app.web.Run()
}
