# Greencity task

This project implements a system of two microservices using Golang that communicate via the Kafka message broker.

# Installation
## Setting Up Kafka and Zookeeper

To run Kafka and Zookeeper in Docker, execute the following command:

```bash
make run-compose name="zookeeper kafka"
```

To create a Kafka topic with 4 partitions, use the following command:

```bash
make run-compose name="zookeeper kafka"
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

# Conclusion

## License

[MIT](https://choosealicense.com/licenses/mit/)