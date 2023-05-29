package wkafka

import (
	"github.com/Shopify/sarama"
	"log"
	"strings"
)

var (
	producer sarama.AsyncProducer
	topic    = "default_message"
)

// InitProducer 生产者
func InitProducer(topicInput, hosts string) {
	topic = topicInput
	config := sarama.NewConfig()
	config.Producer.Compression = sarama.CompressionGZIP
	client, err := sarama.NewClient(strings.Split(hosts, ","), config)
	if err != nil {
		log.Printf("init kafka client err: %v\n", err)
	}
	producer, err = sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		log.Printf("init kafka async client err: %v\n", err)
	}
}

func Send(data []byte) {
	b := sarama.ByteEncoder(data)
	producer.Input() <- &sarama.ProducerMessage{
		Topic: topic,
		Key:   nil,
		Value: b,
	}
}
