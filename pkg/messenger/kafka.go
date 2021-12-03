package messenger

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"strings"
	"worker/core/domain/messenger"
	"worker/core/port"
)

type kafkaMessenger struct {
	Url string
	Topic string
	Group string
}

func NewKafkaMessenger(Url string, Topic string, Group string) port.Messenger {
	return kafkaMessenger{
		Url: Url,
		Topic: Topic,
		Group: Group,
	}
}

func (msgr kafkaMessenger) Write(Message messenger.Message) error {
	kafkaURL := msgr.Url
	topic := msgr.Topic

	kafkaWriter := getKafkaWriter(kafkaURL, topic)

	defer kafkaWriter.Close()

	msg := kafka.Message{
		Key: Message.Key,
		Value: Message.Value,
	}
	return kafkaWriter.WriteMessages(context.TODO(), msg)
}

func (msgr kafkaMessenger) Read(Callback func (message messenger.Message), ErrCallback func (err error)) {
	kafkaURL := msgr.Url
	topic := msgr.Topic
	groupID := msgr.Group

	reader := getKafkaReader(kafkaURL, topic, groupID)

	defer reader.Close()

	fmt.Println("start reading ... !!")
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			ErrCallback(err)
			return
		}
		//fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		msg := messenger.Message{
			Topic: m.Topic,
			Key: m.Key,
			Value: m.Value,
		}
		Callback(msg)
	}
}

func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}

func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}
