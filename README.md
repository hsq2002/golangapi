# golangapi
Real-time streaming API with Golang and Kafka



# start Zookeeper and Kafka in golangapi
- docker-compose up -d

# verify appropriate containers are running
- docker ps

# Creating a Topic via Docker
Create Topic:
- first - bash into container - docker exec -it golangapi-kafka-1 bash
- kafka-topics --create --topic test-topic --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1

List of Topics:
- kafka-topics --list --bootstrap-server localhost:9092

producing a message:
- kafka-console-producer --topic test-topic --bootstrap-server localhost:9092

consuming messages:
- docker exec -it golangapi-kafka-1 kafka-console-consumer --topic test-topic --from-beginning --bootstrap-server localhost:9092


# Testing the API endpoints
- build docker images - docker-compose build
- run docker - docker-compose up
- Test functionality via Postman, Insomnia, etc
