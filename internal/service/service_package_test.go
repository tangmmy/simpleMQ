package service

import (
	"simpleMQ/internal/models"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAddMessage(t *testing.T) {
	gin.SetMode(gin.TestMode)

	var context = make(models.FlowContext)
	context[models.TOPIC] = "1234"
	context[models.PAYLOAD] = "12345678"

	err := AddMessage(&context)

	if err != nil {
		t.Fatalf("expected success but got %s", err)
	}
}

func TestRemoveMessage(t *testing.T) {
	gin.SetMode(gin.TestMode)

	var context = make(models.FlowContext)
	context[models.TOPIC] = "1234"
	context[models.PAYLOAD] = "12345678"

	//add first
	err := AddMessage(&context)

	if err != nil {
		t.Fatalf("expected success but got %s", err)
	}

	//then try remove
	message, err := RemoveMessage(&context)

	if err != nil {
		t.Fatalf("expected success but got %s", err)
	}

	if message.Payload != "12345678" {
		t.Fatalf("expected same payload but got %s", message.Payload)
	}
}
