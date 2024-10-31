package main

import (
	"fmt"
	"log"
	"os"

	"github.com/LeonardoGrigolettoDev/hl-server-go/cmd/controller"
	postgres "github.com/LeonardoGrigolettoDev/hl-server-go/cmd/database/postgre"
	"github.com/LeonardoGrigolettoDev/hl-server-go/cmd/mqtt_client"
	"github.com/LeonardoGrigolettoDev/hl-server-go/cmd/redis"
	"github.com/LeonardoGrigolettoDev/hl-server-go/cmd/repository"
	"github.com/LeonardoGrigolettoDev/hl-server-go/cmd/usecase"
	websockets "github.com/LeonardoGrigolettoDev/hl-server-go/cmd/websocket"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv" // Importando o godotenv
	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error on load file .env: %v", err)
	}
}

func main() {
	// Conectando ao PostgreSQL
	db, err := postgres.PostgreSQLConnectDB()
	if err != nil {
		log.Fatalf("Error on connect DB (PostgreSQL): %v", err)
	}
	defer db.Close()

	server := gin.Default()
	mqttConfig := mqtt_client.ClientConfig{
		Broker:   os.Getenv("MQTT_BROKER_ADDR_TCP"),
		ClientID: os.Getenv("MQTT_CLIENT_ID"),
		Username: os.Getenv("MQTT_SERVER_USER"),
		Password: os.Getenv("MQTT_SERVER_USER_PASS"),
		DefaultPublishHandler: func(client mqtt.Client, msg mqtt.Message) {
			fmt.Printf("Recebida a mensagem: %s do tópico: %s\n", msg.Payload(), msg.Topic())
		},
		OnConnectHandler: func(client mqtt.Client) {
			topic := "/device/#"
			fmt.Println("Conectado ao broker MQTT")
			if token := client.Subscribe(topic, 1, nil); token.Wait() && token.Error() != nil {
				log.Fatal(token.Error())
			}
			fmt.Printf("Inscrito no tópico %s\n", topic)
		},
	}
	mqttClient := mqtt_client.ConnectToBroker(&mqttConfig)
	mqtt_client.SubscribeOnTopic(mqttClient)
	redis.ConnectRedis()

	//Repository layer
	DeviceRepository := repository.NewDeviceRepository(db)
	//Usecase layer
	DeviceUseCase := usecase.NewDeviceUsecase(DeviceRepository)
	//Controller layer
	DeviceControler := controller.NewDeviceController(DeviceUseCase)

	//Routes
	//	DEVICE
	//		CRUD
	server.GET("/devices", DeviceControler.GetDevices)
	server.POST("/device", DeviceControler.CreateDevice)
	server.GET("/device/:id", DeviceControler.GetDeviceById)
	server.PUT("/device/:id", DeviceControler.UpdateDeviceById)
	//		COMMUNICATION (WebSocket)
	server.GET("/device/capture/:id", websockets.VideoCaptureHandler)
	server.GET("/device/stream/:id", websockets.WatchVideoHandler)
	server.POST("/device/publish", websockets.PublishDeviceMessage)

	server.Run(":8080")
	log.Println("Server running on port 8080")
}
