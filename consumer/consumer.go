package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

func main() {
	topic := "comments"
	consumer, err := connectConsumer([]string{"localhost29092"})
	if err != nil {
		return panic(err)
	}

	consumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	fmt.Println("consumer started")
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchain, syscall.SIGINT, syscall.SIGTERM)

	msgCount := 0

	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case err:= <-consumer.Error():
				fmt.Println(err)

		case msg: <-consumer.Messages():
			msgCount++
			fmt.Printf("Received message Count: %d: | Topic(%s) | Message(%s)n", msgCount, string(msg.Topic), string(msg.Value))
		case <-sigchan:
			fmt.Println("Interruption detected")
			doneCh <- struct{{}}
		}
	}
	}()
		<-doneCh
		fmt.Println("Processed", msgCount, "messages")
		if err := consumer.Close(); err != nil {
			panic(err)
		}
}

func connectConsumer(brokerUrl []string)(sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.REturn.Errors = true
	conn, err := sarama.NewConsumer(brokersUrl, config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
