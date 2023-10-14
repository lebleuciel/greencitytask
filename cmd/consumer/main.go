package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"sync"

	"github.com/IBM/sarama"
)

func main() {
	// Create a new Sarama configuration for Kafka consumer
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	// Initialize the Kafka consumer to connect to Kafka broker at "kafka:9092"
	consumer, err := sarama.NewConsumer([]string{"kafka:9092"}, config)
	if err != nil {
		log.Fatal("Error creating Kafka consumer:", err)
	}
	defer consumer.Close()
	// Define the Kafka topic to subscribe to and partitions to listen on
	topic := "greencity-topic"
	partitions := []int32{0, 1, 2, 3}

	var mu sync.Mutex
	receivedMessages := make(map[int32][]string)

	fmt.Println("Kafka Consumer started. Listening for messages...")
	// Iterate over specified partitions, create partition consumers, and start Goroutines
	for _, partition := range partitions {
		partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			log.Fatal("Error creating partition consumer:", err)
		}

		go func(partition int32, partitionConsumer sarama.PartitionConsumer) {
			// Continuously listen for and process Kafka messages in each partition
			for msg := range partitionConsumer.Messages() {
				mu.Lock()
				receivedMessages[partition] = append(receivedMessages[partition], string(msg.Value))
				mu.Unlock()
				verifyMessages(&mu, receivedMessages, partitions)
			}
		}(partition, partitionConsumer)
	}
	// The consumer runs indefinitely to keep listening for Kafka messages
	select {}
}

// isVerified checks if a message contains only allowed words
func isVerified(msg string) bool {
	allowedWords := []string{
		"Green",
		"City",
		"Tehran",
		"User",
		"Product",
		"Golang",
		"Mutex",
		"Channel",
	}

	pattern := fmt.Sprintf("^(%s)*$", strings.Join(allowedWords, "|"))
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(msg)
}

// verifyMessages checks and logs verification status for messages received from different partitions
func verifyMessages(mu *sync.Mutex, receivedMessages map[int32][]string, partitions []int32) {
	mu.Lock()
	defer mu.Unlock()
	var msg string
	for _, partition := range partitions {
		if len(receivedMessages[partition]) == 0 {
			return
		}
		msg += receivedMessages[partition][0]
		receivedMessages[partition] = receivedMessages[partition][1:]
	}

	log.Printf("Verify message in partitions: %v", isVerified(msg))
}
