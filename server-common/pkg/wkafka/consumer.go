package wkafka

import (
	"github.com/Shopify/sarama"
	"log"
	"strings"
)

var consumer sarama.Consumer

// InitConsumer 消费者
func InitConsumer(hosts string) {
	config := sarama.NewConfig()
	client, err := sarama.NewClient(strings.Split(hosts, ","), config)
	if err != nil {
		log.Printf("init kafka consumer client error: %v\n", err)
	}
	consumer, err = sarama.NewConsumerFromClient(client)
	if err != nil {
		log.Printf("init kafka consumer client error: %v\n", err)
	}
}

type ConsumerCallback func(data []byte)

// ConsumerMsg 消费消息 通过回调函数进行
func ConsumerMsg(callBack ConsumerCallback) {
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if nil != err {
		log.Printf("iConsumePartition error: %v\n", err)
		return
	}

	defer partitionConsumer.Close()
	for {
		msg := <-partitionConsumer.Messages()
		if nil != callBack {
			callBack(msg.Value)
		}
	}
}
