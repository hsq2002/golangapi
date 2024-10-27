package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

var consumer sarama.Consumer

func InitializeConsumer(brokersUrl []string) error {
	var err error
	consumer, err = sarama.NewConsumer(brokersUrl, nil)
	if err != nil {
		return err
	}
	return nil
}

func ConsumeMessages(topic string) {
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		log.Printf("Error getting partitions for topic %s: %v", topic, err)
		return
	}

	for _, partition := range partitionList {
		pc, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			log.Printf("Error consuming partition %d: %v", partition, err)
			return
		}
		defer pc.AsyncClose()

		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				log.Printf("Received message: %s\n", string(message.Value))
			}
		}(pc)
	}
}
