package main

import (
	"bufio"
	"context"
	"log"
	"os"
	"os/signal"
	"ss/internal/app"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	var wg *sync.WaitGroup = new(sync.WaitGroup)
	wg.Add(2)

	app := app.NewApp(ctx, wg)
	app.RunApp()

	defer func(c context.CancelFunc) {
		if err := recover(); err != nil {
			log.Println("PANIC:", err)
			cancel()
			wg.Wait()
		}
	}(cancel)

	log.Println("Core started")

	go func() {
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		cancel()
	}()

	go func() {
		signalCtrlC := make(chan os.Signal, 1)
		signal.Notify(signalCtrlC, os.Interrupt)

		<-signalCtrlC
		cancel()
	}()

	go func() {
		time.Sleep(24 * time.Hour)
		cancel()
	}()

	wg.Wait()
}
