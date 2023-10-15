# Greencity task

This project involves the implementation of two microservices in Golang that communicate via the Kafka message broker. The first microservice acts as a producer, while the second microservice is a consumer.




## Consumer and Producer Microservices Explanation
This Go program sets up a Kafka consumer. It listens for messages from a Kafka topic named "greencity-topic" on partitions 0, 1, 2, and 3. When a message arrives, it checks if the message consists of specific allowed words like "Green," "City," "Tehran," and others. If the message contains only these allowed words, it's considered verified. The consumer collects and verifies messages from different partitions, and if a complete message is obtained across partitions, it prints a verification log. This program uses the IBM/sarama package for Kafka communication and works to ensure that messages are correctly received and verified.
# Installation
## Setting Up Kafka and Zookeeper

To run Kafka and Zookeeper in Docker, execute the following command:

```bash
make run-compose name="zookeeper kafka"
```

To create a Kafka topic with 4 partitions, use the following command:

```bash
make create-kafka-topic
```
## Running the Producer and Consumer

To run the Kafka Producer and Consumer, use the following command:
```bash
make run-compose 
```

You can now view logs for the Producer and Consumer using the following commands:
```bash
docker logs producer
docker logs consumer
```


## License

[MIT](https://choosealicense.com/licenses/mit/)