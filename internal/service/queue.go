package service

import (
	"errors"
	"simpleMQ/internal/models"
)

type Message struct {
	Topic         string
	IsPersistence bool
	Payload       any
}

const maxSize = 10

type Queue struct {
	Queue      [maxSize]Message
	StartIndex int // meaning that element from at this position will be removed first
	EndIndex   int // meaning that a new element will be insert at this position
}

var topicMap map[string]*Queue = make(map[string]*Queue)

func AddMessage(context models.FlowContext) error {

	message := Message{
		Topic:         context[models.TOPIC].(string),
		IsPersistence: false,
		Payload:       context[models.TOPIC].(string),
	}
	targetQueue := topicMap[message.Topic]

	//if cannot find, initialize a new one
	if targetQueue == nil {
		topicMap[message.Topic] = &Queue{}
		targetQueue = topicMap[message.Topic]
	}

	//the queue is full, throw error
	if (targetQueue.EndIndex+1)%maxSize == targetQueue.StartIndex {
		return errors.New("queue is full")
	}

	//the queue is not full
	targetQueue.Queue[targetQueue.EndIndex] = message
	targetQueue.EndIndex = (targetQueue.EndIndex + 1) % maxSize

	return nil
}
