package kafka

import (
	"github.com/segmentio/kafka-go"
	"fmt"
)

func StartKafka() {
	conf := Kafka.ReaderConfig{
		Brokers: []string {"localhost:9892"},
		Topic: "mytopic",
		GroupId: "gl",
		MaxBytes: 10,
	}

	reader := kafka.NewReader(conf)

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Some error occurred", err)
			continue
		}
		fmt.Println("Message is: ", string(m.Value))
	}
}
