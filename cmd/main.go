package main

import (
	"log"

	"github.com/LeonardoGrigolettoDev/hl-server-go/cmd/controller"
	postgres "github.com/LeonardoGrigolettoDev/hl-server-go/cmd/database/postgre"
	"github.com/LeonardoGrigolettoDev/hl-server-go/cmd/redis"
	"github.com/LeonardoGrigolettoDev/hl-server-go/cmd/repository"
	"github.com/LeonardoGrigolettoDev/hl-server-go/cmd/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv" // Importando o godotenv
	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load()
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

	redis.ConnectRedis()

	//Repository layer
	DeviceRepository := repository.NewDeviceRepository(db)
	//Usecase layer
	DeviceUseCase := usecase.NewDeviceUsecase(DeviceRepository)
	//Controller layer
	DeviceControler := controller.NewDeviceController(DeviceUseCase)

	//Routes
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.GET("/api/device", DeviceControler.GetDevices)

	// r.HandleFunc("/ws", websockets.StreamVideoCapture) // Endpoint para WebSocket
	// http.HandleFunc("/image", func(w http.ResponseWriter, r *http.Request) {
	// 	// Verifica se a imagem existe
	// 	imgPath := "./websocket/image.jpg"
	// 	if _, err := os.Stat(imgPath); err == nil {
	// 		http.ServeFile(w, r, imgPath) // Serve a imagem atual
	// 	} else {
	// 		http.Error(w, "Image not found", http.StatusNotFound)
	// 	}
	// })
	server.Run(":8080")
	log.Println("Server running on port 8080")
}
