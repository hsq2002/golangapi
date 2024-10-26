package main

import (
	"fmt"
	appKafka "./kafka"
)

func main() {

	fmt.Println("okay");
	appKafka()
	fmt.Println("Kafka has been started...")

	time.Sleep(10 * time.Minute)
}
