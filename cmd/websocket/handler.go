package websockets

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/LeonardoGrigolettoDev/hl-server-go/cmd/redis"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type StreamMessage struct {
	Device string `json:"device"`
	Stream int    `json:"stream"` // 1 ou 0
}

var (
	upgrader = websocket.Upgrader{ReadBufferSize: 1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			// Permite conexões de qualquer origem (ajuste conforme necessário)
			return true
		}} // Use this to upgrade HTTP connection to WebSocket
	clients     = make(map[string][]*websocket.Conn) // Map to store clients for each iddevice
	clientsMux  = sync.Mutex{}
	devices     = make(map[string][]*websocket.Conn) // Map to store clients for each iddevice
	devicesMux  = sync.Mutex{}                       // Mutex for safe access to clients map
	redisClient = redis.ConnectRedis()
	ctx         = context.Background()
)

func VideoCaptureHandler(c *gin.Context) {
	deviceId := c.Param("id")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		c.JSON(500, gin.H{"error": "Failed to upgrade connection"})
		return
	}
	defer conn.Close()

	devicesMux.Lock()
	devices[deviceId] = append(devices[deviceId], conn)
	devicesMux.Unlock()

	defer func() {
		devicesMux.Lock()
		for i, c := range devices[deviceId] {
			if c == conn {
				devices[deviceId] = append(devices[deviceId][:i], devices[deviceId][i+1:]...)
				break
			}
		}
		devicesMux.Unlock()
	}()
	log.Println("WebSocket connection established")
	for {
		_, r, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket connection closed:", err)
			break
		}
		fmt.Println("Publishing on: " + "device:" + deviceId)
		redisClient.Publish(ctx, "device:"+deviceId, r)
	}
	log.Println("WebSocket connection closed")
}

func WatchVideoHandler(c *gin.Context) {
	deviceId := c.Param("id")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	clientsMux.Lock()
	clients[deviceId] = append(clients[deviceId], conn)
	clientsMux.Unlock()

	defer func() {
		clientsMux.Lock()
		for i, c := range clients[deviceId] {
			if c == conn {
				clients[deviceId] = append(clients[deviceId][:i], clients[deviceId][i+1:]...)
				break
			}
		}
		clientsMux.Unlock()
	}()
	pubsub := redisClient.Subscribe(ctx, "device:"+deviceId)
	defer pubsub.Close()

	for msg := range pubsub.Channel() {
		// Redistribui a mensagem para os clientes conectados
		clientsMux.Lock()
		for _, clientConn := range clients[deviceId] {
			err := clientConn.WriteMessage(websocket.BinaryMessage, []byte(msg.Payload))
			if err != nil {
				clientConn.Close()
			}
		}
		clientsMux.Unlock()
	}
}

func PublishDeviceMessage(c *gin.Context) {
	var msg StreamMessage

	// Tenta decodificar a mensagem JSON recebida
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	deviceId := msg.Device
	// Converte a estrutura de mensagem para JSON
	messageJSON, err := json.Marshal(msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error marshaling message"})
		return
	}

	log.Println("Publishing message:", string(messageJSON))
	broadcastMessageDevices(deviceId, messageJSON)
	c.JSON(http.StatusOK, gin.H{"status": "Message sent", "message": string(messageJSON)})
}

// func broadcastMessageClients(idDevice string, message []byte) {
// 	clientsMux.Lock()
// 	for _, conn := range clients[idDevice] {
// 		conn.WriteMessage(websocket.TextMessage, message)
// 		fmt.Println("sending stream" + idDevice)
// 	}
// 	clientsMux.Unlock()
// }

func broadcastMessageDevices(idDevice string, message []byte) {
	devicesMux.Lock()
	for _, conn := range devices[idDevice] {
		conn.WriteMessage(websocket.TextMessage, message)
	}
	devicesMux.Unlock()
}
