package app

import (
	"net/http"
	"os"
	"ss/internal/pkg/tools"

	"github.com/gin-gonic/gin"
)

// Структура web-сервера
type Web struct {
	router *gin.Engine
}

// Инициализация сервера
func NewServer() *Web {
	router := gin.Default()

	return &Web{
		router: router,
	}
}

// Запуск web-сервера
func (w *Web) Run() {
	ln, cfg := tools.GetAnchorLnCnf()

	// Запуск сервера с использованием Gin и TLS
	server := &http.Server{
		Addr:      ":" + os.Getenv("HTTPS_PORT"),
		Handler:   w.router,
		TLSConfig: cfg,
	}

	if err := server.Serve(ln); err != nil {
		panic(err)
	}
}
