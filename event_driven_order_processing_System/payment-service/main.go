package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/IBM/sarama"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Database connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Kafka consumer
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	consumer, err := sarama.NewConsumerGroup(
		[]string{"kafka:9092"},
		"payment-service-group",
		config,
	)
	if err != nil {
		log.Fatal("Failed to create consumer group:", err)
	}

	// Kafka consumer setup
	processor := &PaymentProcessor{db: db}
	go startKafkaConsumer(consumer, processor)

	// Start HTTP server
	http.HandleFunc("/payments", func(w http.ResponseWriter, r *http.Request) {
		// Logic for payments endpoint (e.g., query DB or other actions)
		// Placeholder response for testing:
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("Payments service is running"))
		if err != nil {
			log.Printf("Failed to write response: %v", err)
		}
	})

	log.Println("Payment service started...")
	log.Fatal(http.ListenAndServe(":8080", nil)) // Start the HTTP server
}

// startKafkaConsumer handles the Kafka consumer logic
func startKafkaConsumer(consumer sarama.ConsumerGroup, processor *PaymentProcessor) {
	for {
		err := consumer.Consume(context.Background(), []string{"order_created"}, processor)
		if err != nil {
			log.Println("Consume error:", err)
		}
	}
}
