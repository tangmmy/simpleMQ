package handlers

import (
	"net/http"
	"simpleMQ/internal/models"
	"simpleMQ/internal/service"

	"github.com/gin-gonic/gin"
)

//handler layer, will be responsible for transforming request to flowcontext

func HandleMessage(c *gin.Context) {

}

func HandleProducerMessage(c *gin.Context) {

	var produceRequest models.ProducerRequest
	err := c.ShouldBindBodyWithJSON(&produceRequest)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "error in parsing",
		})
		return
	}

	var context models.FlowContext = make(models.FlowContext)
	context[models.TOPIC] = produceRequest.Topic
	context[models.PAYLOAD] = produceRequest.Payload
	err_2 := service.AddMessage(context)

	if err_2 != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "error in handling",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": produceRequest,
	})

}
