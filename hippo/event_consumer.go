package hippo

import (
	"github.com/Shopify/sarama"

)

func InitConsumer() error {

	conf := sarama.NewConfig()
	hippoClient, err:= sarama.NewConsumer([]string{"localhost:9092"}, conf)
	if err != nil {
		return err
	}

	err = hippoClient.Close()
	if err != nil {
		return err
	}

	return nil
}