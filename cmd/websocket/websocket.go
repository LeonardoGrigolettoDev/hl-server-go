package websockets

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type StreamMessage struct {
	Device string `json:"device"`
	Stream int    `json:"stream"` // 1 ou 0
}

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
var connections = make(map[*websocket.Conn]bool) // Armazena as conexões ativas
var mu sync.Mutex                                // Mutex para proteger as conexões

func StreamVideoCapture(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		c.JSON(500, gin.H{"error": "Failed to upgrade connection"})
		return
	}
	defer conn.Close()

	mu.Lock()
	connections[conn] = true // Adiciona a nova conexão
	mu.Unlock()

	log.Println("WebSocket connection established")

	// Recebe mensagens do cliente
	for {
		// Aguardar a próxima mensagem do cliente
		messageType, r, err := conn.NextReader() // Lê a próxima mensagem
		if err != nil {
			log.Println("WebSocket connection closed:", err)
			break
		}

		// Para mensagens de texto ou binárias
		var buffer bytes.Buffer
		if _, err := io.Copy(&buffer, r); err != nil {
			log.Println("Error reading message:", err)
			continue
		}

		// Agora você pode imprimir o conteúdo do buffer
		log.Printf("Received message of type %d: %s\n", messageType, buffer.String())

		// Enviar a mensagem para todas as conexões
		mu.Lock()
		for conn := range connections {
			err := conn.WriteMessage(messageType, buffer.Bytes())
			if err != nil {
				log.Println("Error sending message:", err)
				conn.Close()
				delete(connections, conn) // Remove a conexão se houver erro
			}
		}
		mu.Unlock()
	}

	mu.Lock()
	delete(connections, conn) // Remove a conexão quando ela é fechada
	mu.Unlock()
	log.Println("WebSocket connection closed")
}

func PublishMessage(c *gin.Context) {
	var msg StreamMessage

	// Tenta decodificar a mensagem JSON recebida
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Converte a estrutura de mensagem para JSON
	messageJSON, err := json.Marshal(msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error marshaling message"})
		return
	}

	log.Println("Publishing message:", string(messageJSON))
	c.JSON(http.StatusOK, gin.H{"status": "Message sent", "message": string(messageJSON)})

	mu.Lock()
	// Enviar a mensagem para todas as conexões
	for conn := range connections {
		err := conn.WriteMessage(websocket.TextMessage, messageJSON)
		if err != nil {
			log.Println("Error sending message:", err)
			conn.Close()
			delete(connections, conn) // Remove a conexão se houver erro
		}
	}
	mu.Unlock()
}
