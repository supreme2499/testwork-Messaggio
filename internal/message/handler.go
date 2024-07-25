package message

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"testingwork-kafka/internal/config"
	"testingwork-kafka/internal/message/database"

	"testingwork-kafka/internal/handlers"
	"testingwork-kafka/pkg/clients/postresql"
	"testingwork-kafka/pkg/logging"

	"github.com/julienschmidt/httprouter"
)

type handler struct {
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	logger.Info("возвращение структуры обработчика")
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	logger := logging.GetLogger()
	logger.Info("Запуск обработчика")
	router.POST("/message", h.Postmessage)
	router.GET("/statistics", h.GetStatistics)
}

// пока что в запросах просто затычки
func (h *handler) GetStatistics(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	w.Write([]byte("Статистика обработанных сообщений"))
}

func (h *handler) Postmessage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	logger := logging.GetLogger()
	cfg := config.GetConfig()
	clientdb, err := postresql.NewClient(context.TODO(), 3, cfg.Storage)
	if err != nil {
		log.Fatal("ошибка клиента постгрес: ", err)
	}
	repository := database.NewRepository(clientdb)

	logger.Info("сервер начал получение сообщений")
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	for {
		content := database.Contents{}
		if err := decoder.Decode(&content); err == io.EOF {
			break
		} else if err != nil {
			logger.Error("Ошибка json")
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}
		if content.Status == "" {
			content.Status = "received"
		}
		err = repository.Message(context.TODO(), content.Content, content.Status)
		if err != nil {
			log.Fatal("ошибка записи сообщения в бд handler:", err)
		}

		logger.Infof("Успешно получено сообщение: %s", content.Content)
	}
	logger.Info("сервер закончил получать сообщения")
}
