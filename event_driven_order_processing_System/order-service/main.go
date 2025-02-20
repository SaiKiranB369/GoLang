package main

import (
	"fmt"
	"log"
	"os"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
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

	// Ensure the orders table exists, auto-migrate if necessary
	if err := db.AutoMigrate(&Order{}); err != nil {
		log.Fatal("Failed to auto-migrate database:", err)
	}

	// Kafka producer
	producer, err := sarama.NewSyncProducer([]string{os.Getenv("KAFKA_BROKERS")}, nil)
	if err != nil {
		log.Fatal("Failed to create Kafka producer:", err)
	}
	defer producer.Close()

	// Setup routes
	router := gin.Default()
	router.POST("/orders", createOrderHandler(db, producer))
	router.GET("/orders", getOrdersHandler(db))

	log.Fatal(router.Run(":8080"))
}
