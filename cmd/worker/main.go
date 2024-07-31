package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"testingwork-kafka/internal/config"
	"testingwork-kafka/internal/message/database"
	"testingwork-kafka/internal/worker"
	"testingwork-kafka/pkg/clients/postresql"
	"testingwork-kafka/pkg/logging"

	"github.com/IBM/sarama"
)

func main() {
	logger := logging.GetLogger()
	topic := "message"
	msgCnt := 0
	worker, err := worker.ConnectConsumer([]string{"kafka:9092"})
	if err != nil {
		log.Fatal("ошибка коннекта консумера", err)
	}

	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatal("ошибка получения задачи из очереди: ", err)
	}
	logger.Info("Старт consumer")

	signalchan := make(chan os.Signal, 1)
	signal.Notify(signalchan, syscall.SIGINT, syscall.SIGTERM)

	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				logger.Error(err)
			case msg := <-consumer.Messages():
				msgCnt++
				logger.Infof("Topic: %s | Message: %s", string(msg.Topic), string(msg.Value))
				content := string(msg.Value)
				logger.Infof("(Сообщение обработано: %s", content)
				cfg := config.GetConfig()
				clientdb, err := postresql.NewClient(context.TODO(), 3, cfg.Storage)
				if err != nil {
					log.Fatal("ошибка клиента постгрес: ", err)
				}
				defer clientdb.Close()

				repository := database.NewRepository(clientdb)

				err = repository.MessageWork(context.TODO(), content)
				if err != nil {
					log.Fatal("ошибка записи сообщения в бд handler:", err)
				}
			case <-signalchan:
				doneCh <- struct{}{}
			}

		}
	}()

	<-doneCh
	if err := worker.Close(); err != nil {
		log.Fatal("ошибка закрытия")
	}
}
