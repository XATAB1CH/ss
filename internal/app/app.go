package app

import (
	"context"
	"fmt"
	"sync"
)

type App struct {
	bot *Bot

	ctx context.Context
	wg  *sync.WaitGroup
}

func NewApp(ctx context.Context, w *sync.WaitGroup) *App {
	bot, err := NewBot()
	if err != nil {
		fmt.Println(err)
	}

	app := &App{
		bot: bot,
		ctx: ctx,
		wg:  w,
	}

	return app
}

func (app *App) RunApp() {
	app.bot.Run()
}
