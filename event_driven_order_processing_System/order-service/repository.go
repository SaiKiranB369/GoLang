package main

import (
	"encoding/json"
	"fmt"

	"event_driven_order_processing_System/common"

	"github.com/IBM/sarama"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db       *gorm.DB
	producer sarama.SyncProducer
}

func NewOrderRepository(db *gorm.DB, producer sarama.SyncProducer) *OrderRepository {
	return &OrderRepository{db: db, producer: producer}
}

func (r *OrderRepository) CreateOrder(order *common.Order) error {
	// Save the order to the database
	if err := r.db.Create(order).Error; err != nil {
		return fmt.Errorf("database create failed: %w", err)
	}

	// Convert order.ID to string safely
	orderID := fmt.Sprintf("%v", order.ID)
	if orderID == "" {
		return fmt.Errorf("invalid order ID: %v", order.ID)
	}

	// Prepare and publish the order created event
	event := common.OrderCreatedEvent{
		OrderID:    orderID,
		UserID:     order.UserID,
		TotalPrice: order.TotalPrice,
	}

	if err := r.publishOrderCreatedEvent(event); err != nil {
		return fmt.Errorf("failed to publish event: %w", err)
	}

	return nil
}

func (r *OrderRepository) publishOrderCreatedEvent(event common.OrderCreatedEvent) error {
	// Marshal event into JSON
	jsonData, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	// Create Kafka message
	msg := &sarama.ProducerMessage{
		Topic: "order_created",
		Value: sarama.ByteEncoder(jsonData),
	}

	// Send the message
	if _, _, err := r.producer.SendMessage(msg); err != nil {
		return fmt.Errorf("failed to publish event: %w", err)
	}

	return nil
}
