package domain

import (
	"encoding/json"
	"github.com/djumanoff/amqp"
)

type AmqpScala struct {
	JsonClass     string          `json:"jsonClass"`
	Body          json.RawMessage `json:"body"`
	Headers       amqp.Table      `json:"headers"`
	RoutingKey    string          `json:"routingKey"`
	ReplyTo       string          `json:"replyTo,omitempty"`
}

type ErrorMessage struct {
	AppID   string          `json:"app_id,omitempty"`
	Message json.RawMessage `json:"message"`
	Body    json.RawMessage `json:"body"`
}
