package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/IBM/sarama"
	"github.com/eapache/queue"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewManualPartitioner

	producer, err := sarama.NewSyncProducer([]string{"kafka:9092"}, config)
	if err != nil {
		log.Fatal("Error initializing Kafka producer:", err)
	}
	defer producer.Close()

	q := queue.New()

	topic := "greencity-topic"
	partitions := 4

	go func() {
		for i := 0; ; i++ {
			data := generateRandomData()
			q.Add(data)
			if q.Length()%1024 == 0 {
				go sendMessagesToKafka(producer, topic, partitions, q)
			}
			time.Sleep(10 * time.Millisecond)
		}
	}()
	select {}
}

func sendMessagesToKafka(producer sarama.SyncProducer, topic string, partitions int, q *queue.Queue) {
	for q.Length() > 0 {
		msg := q.Peek()
		for i := 0; i < 4; i++ {
			if data, ok := msg.([]byte); ok {
				producerMessages := make([]*sarama.ProducerMessage, 4)
				for partition := 0; partition < partitions; partition++ {
					producerMessages[partition] = &sarama.ProducerMessage{
						Topic:     topic,
						Partition: int32(partition),
						Value:     sarama.ByteEncoder(data[partition*len(data)/partitions : (partition+1)*len(data)/partitions]),
					}
				}

				for partition := 0; partition < partitions; partition++ {
					_, _, err := producer.SendMessage(producerMessages[partition])
					if err != nil {
						log.Printf("Error producing message: %v", err)
					} else {
						log.Printf("Produced message: %s", data)
					}
				}
			} else {
				log.Println("Error: Unable to assert data as []byte")
			}
		}
	}
}

func generateRandomData() []byte {
	sentences := []string{
		"Green",
		"City",
		"Tehran",
		"User",
		"Product",
		"Golang",
		"Mutex",
		"Channel",
	}

	var data string

	for len(data) <= 320 {
		randomIndex := rand.Intn(len(sentences))
		data += sentences[randomIndex]
	}

	return []byte(data)
}
