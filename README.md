# golangapi
Real-time streaming API with Golang and Kafka




# Creating a Topic via Docker
Create Topic:
kafka-topics --create --topic test-topic --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1

List of Topics:
kafka-topics --list --bootstrap-server localhost:9092

producing a message:
kafka-console-producer --topic test-topic --bootstrap-server localhost:9092

consuming messages:
docker exec -it golang-api-kafka-1 bash (or whatever the name of your created docker container is)
    
