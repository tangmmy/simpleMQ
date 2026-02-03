package models

type ProducerRequest struct {
	Payload string `json:"payload"`
	Topic   string `json:"topic"`
}

type ConsumerRequest struct {
	Topic string `json:"topic"`
}

type FlowContext map[FlowContextKey]any

type FlowContextKey string

const (
	TOPIC   FlowContextKey = "TOPIC"
	PAYLOAD FlowContextKey = "PAYLOAD"
)
