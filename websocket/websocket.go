package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Aceitar conexões de qualquer origem
	},
}

func StreamVideoCapture(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error during connection upgrade:", err)
		return
	}
	defer conn.Close()

	// Cria uma janela para exibir o vídeo
	// window := gocv.NewWindow("Video Stream")
	// defer window.Close()

	for {
		// Lê a mensagem do WebSocket
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		fmt.Println(msg)

		// Converte os bytes recebidos para uma matriz (frame)
		// img, err := gocv.IMDecode(msg, gocv.IMReadColor)
		// if err != nil {
		// 	log.Println("Error decoding image:", err)
		// 	continue
		// }

		// // Mostra o frame na janela
		// window.IMShow(img)
		// if window.WaitKey(1) >= 0 {
		// 	break
		// }
	}
}
