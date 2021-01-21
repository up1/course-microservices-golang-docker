package service1

import (
	"log"
	"service1/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Start() {
	r := gin.New()
	r.GET("/api/v1/users/:name", createNewUser)
	r.GET("/api/v2/users/:name", createNewUserWithAsync)
	r.Run(":8080")
}

func createNewUser(c *gin.Context) {
	name := c.Param("name")
	log.Printf("Receive data with name=%s", name)
	c.JSON(200, gin.H{
		"message": "Hello " + name,
	})
}

// keep waiting channels for reply messages from rabbit
var allChans = make(map[string](chan model.HelloMessage))

// channel to publish rabbit messages
var pchan = make(chan model.HelloMessage, 10)

func createNewUserWithAsync(c *gin.Context) {
	name := c.Param("name")

	// Create channel
	responseChan := make(chan model.HelloMessage)
	uuid := uuid.New().String()
	message := model.HelloMessage{uuid, name}
	allChans[uuid] = responseChan
	log.Printf("Receive data with message=%+v", message)
	pchan <- message

	// Wait for response
	waitForResponse(uuid, responseChan, c)
}

func waitForResponse(uuid string, responseChan chan model.HelloMessage, c *gin.Context) {
	for {
		select {
		case response := <-responseChan:
			// responses received
			log.Printf("INFO: received response: %s with uuid: %s", response.Message, response.Id)

			// send response back to client
			c.JSON(200, gin.H{
				"message": response.Message,
			})

			// remove channel from allChans
			delete(allChans, uuid)
			return

		case <-time.After(10 * time.Second):
			// timeout
			log.Printf("ERROR: request timeout uid: %+v", responseChan)

			// send response back to client
			c.JSON(408, gin.H{
				"message": "Timeout with uuid=" + uuid,
			})

			// remove channel from allChans
			delete(allChans, uuid)
			return
		}
	}
}
