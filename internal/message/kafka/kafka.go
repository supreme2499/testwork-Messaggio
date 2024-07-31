package kafka

import (
	"log"
	"testingwork-kafka/pkg/logging"

	samara "github.com/IBM/sarama"
)

func connectProduser(brokers []string) (samara.SyncProducer, error) {
	config := samara.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = samara.WaitForAll
	config.Producer.Retry.Max = 5

	return samara.NewSyncProducer(brokers, config)
}

func PushMessageToQueue(topic string, message []byte) error {
	logger := logging.GetLogger()
	brokers := []string{"kafka:9092"}

	producer, err := connectProduser(brokers)
	if err != nil {
		log.Fatal("ошибка подключения к кафке", err)
	}
	logger.Info("Успешное подключение к кафка")
	defer producer.Close()

	msg := &samara.ProducerMessage{
		Topic: topic,
		Value: samara.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatal("ошибка отправки сообщения в кафку: ", err)
	}
	logger.Infof("сообщение в топик: %s , partition: %d, offset: %d", topic, partition, offset)

	return nil
}
