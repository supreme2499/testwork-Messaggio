package main

import (
	"fmt"
	"net"
	"net/http"
	"testingwork-kafka/internal/config"
	mess "testingwork-kafka/internal/message"
	"testingwork-kafka/pkg/logging"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	time.Sleep(30 * time.Second)
	cfg := config.GetConfig()

	//передаём написанный наш логер в функцию меин
	logger := logging.GetLogger()

	//мультиплекстор или роутер. его мы используем для выполнения наших http запросов
	logger.Info("Создание роутера")
	router := httprouter.New()

	//вызываем функцию который возвращает ссылку на структуру handler которая нам нужна
	//для работы наших методов обработки событий. Тоесть ты передаём нашу структуру в
	//наш мэин что бы мы могли им пользоваться здесь
	handler := mess.NewHandler(logger)

	//вызываем метод Register используя нашу структуру, которую мы передали ранее
	//а в метод мы передаём роутер. Метод регистер это наш обработчик событий
	//окоторый отвечает за ответы на запросы к серверу
	handler.Register(router)

	//вызов функции которая запускает наш сервер
	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	//передаём логер в функцию
	logger := logging.GetLogger()

	logger.Info("создание листинга tcp")
	listener, ListenErr := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
	logger.Infof("сервер запущен на %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)

	//так же мы отдельно вынесли обработчик событий т. к. нам проще просто в конце один раз проверить на ошибку
	if ListenErr != nil {
		logger.Fatal(ListenErr)
	}
	server := &http.Server{
		//хендлер(обработчик событий)
		Handler: router,
		//время ожидания на запись
		WriteTimeout: 15 * time.Second,
		//время ожидания на чтение
		ReadTimeout: 15 * time.Second,
	}
	//запуск сервера + если появится ошибка код выйдет в лог фатал
	logger.Fatal(server.Serve(listener))
}
