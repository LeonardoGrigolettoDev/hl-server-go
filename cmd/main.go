package main

import (
	"log"

	"github.com/LeonardoGrigolettoDev/hl-server-go/cmd/controller"
	postgres "github.com/LeonardoGrigolettoDev/hl-server-go/cmd/database/postgre"
	"github.com/LeonardoGrigolettoDev/hl-server-go/cmd/mqtt"
	"github.com/LeonardoGrigolettoDev/hl-server-go/cmd/redis"
	"github.com/LeonardoGrigolettoDev/hl-server-go/cmd/repository"
	"github.com/LeonardoGrigolettoDev/hl-server-go/cmd/usecase"
	websockets "github.com/LeonardoGrigolettoDev/hl-server-go/cmd/websocket"
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
	mqtt.StartMQTTListen()
	redis.ConnectRedis()

	//Repository layer
	DeviceRepository := repository.NewDeviceRepository(db)
	//Usecase layer
	DeviceUseCase := usecase.NewDeviceUsecase(DeviceRepository)
	//Controller layer
	DeviceControler := controller.NewDeviceController(DeviceUseCase)

	//Routes
	//	DEVICE
	server.GET("/api/devices", DeviceControler.GetDevices)
	server.POST("/api/device", DeviceControler.CreateDevice)
	server.GET("/api/device/:id", DeviceControler.GetDeviceById)
	server.PUT("/api/device/:id", DeviceControler.UpdateDeviceById)
	server.GET("/ws", websockets.StreamVideoCapture)
	server.POST("/api/ws/device/publish", websockets.PublishMessage)
	server.Run(":8080")
	log.Println("Server running on port 8080")
}
