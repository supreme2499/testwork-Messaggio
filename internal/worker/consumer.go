package worker

import samara "github.com/IBM/sarama"

func ConnectConsumer(brokers []string) (samara.Consumer, error) {
	config := samara.NewConfig()
	config.Producer.Return.Errors = true

	return samara.NewConsumer(brokers, config)
}
