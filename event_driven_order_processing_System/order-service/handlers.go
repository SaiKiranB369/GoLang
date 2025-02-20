package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Order struct {
	ID         uint `gorm:"primaryKey"`
	UserID     int
	ProductID  int
	Quantity   int
	TotalPrice float64
	Status     string
}

type OrderRequest struct {
	UserID     int     `json:"user_id"`
	ProductID  int     `json:"product_id"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}

type OrderCreatedEvent struct {
	OrderID    string  `json:"order_id"`
	UserID     int     `json:"user_id"`
	TotalPrice float64 `json:"total_price"`
}

func createOrderHandler(db *gorm.DB, producer sarama.SyncProducer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req OrderRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			fmt.Println("JSON Bind Error:", err) // Debugging
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("Received Order Request: %+v\n", req) // Debugging

		// Save to database
		order := Order{
			UserID:     req.UserID,
			ProductID:  req.ProductID,
			Quantity:   req.Quantity,
			TotalPrice: req.TotalPrice,
			Status:     "PENDING",
		}

		if err := db.Create(&order).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
			return
		}

		// Publish Kafka event
		event := OrderCreatedEvent{
			OrderID:    strconv.Itoa(int(order.ID)),
			UserID:     order.UserID,
			TotalPrice: order.TotalPrice,
		}

		if err := publishOrderCreatedEvent(producer, event); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to publish event"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"order_id": order.ID,
			"status":   order.Status,
		})
	}
}

func publishOrderCreatedEvent(producer sarama.SyncProducer, event OrderCreatedEvent) error {
	message, _ := json.Marshal(event)
	kafkaMsg := &sarama.ProducerMessage{
		Topic: "order_created",
		Value: sarama.StringEncoder(message),
	}

	_, _, err := producer.SendMessage(kafkaMsg)
	return err
}

func getOrdersHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var orders []Order
		if err := db.Find(&orders).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
			return
		}
		c.JSON(http.StatusOK, orders)
	}
}
