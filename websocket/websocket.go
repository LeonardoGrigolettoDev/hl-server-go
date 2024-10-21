package websockets

import (
	"fmt"
	"log"
	"net/http"

	_ "image/png"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func StreamVideoCapture(w http.ResponseWriter, r *http.Request) {
	// Atualizando para WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	for {
		// Recebendo mensagem via WebSocket
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		fmt.Println(message)

		// Verificando se é uma imagem válida
		// if isValidImage(message) {
		// 	// Salvando a imagem no disco
		// 	err = ioutil.WriteFile("image.jpg", message, 0644)
		// 	if err != nil {
		// 		log.Println("Error saving image:", err)
		// 	}
		// } else {
		// 	log.Println("Received invalid image")
		// }
	}
}

// func isValidImage(imageBytes []byte) bool {
// 	img, _, err := image.Decode(bytes.NewReader(imageBytes))
// 	if err != nil {
// 		log.Println("Invalid image:", err)
// 		return false
// 	}

// 	// (Opcional) Redimensionar a imagem se necessário
// 	resizedImg := resize.Resize(800, 0, img, resize.Lanczos3)

// 	// Salvar a imagem redimensionada (caso necessário)
// 	out, err := ioutil.TempFile(".", "resized_*.jpg")
// 	if err != nil {
// 		log.Println("Error creating temp file:", err)
// 		return false
// 	}
// 	defer out.Close()

// 	err = jpeg.Encode(out, resizedImg, nil)
// 	if err != nil {
// 		log.Println("Error encoding image:", err)
// 		return false
// 	}

// 	return true
// }
