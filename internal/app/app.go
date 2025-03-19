package app

import (
	"context"
	"fmt"
	"sync"
)

type App struct {
	bot *Bot
	web *Web

	ctx context.Context
	wg  *sync.WaitGroup
}

func NewApp(ctx context.Context, w *sync.WaitGroup) *App {
	bot, err := NewBot()
	if err != nil {
		fmt.Println(err)
	}

	server := NewServer()

	app := &App{
		bot: bot,
		web: server,
		ctx: ctx,
		wg:  w,
	}

	return app
}

func (app *App) RunApp() {
	app.bot.Run()
	app.web.Run()
}
