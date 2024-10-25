package mqtt

import (
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func StartMQTTListen() {
	broker := "tcp://localhost:1883" // Substitua pelo endereço do broker
	clientID := "go-mqtt-client"
	topic := "device/#" // Tópico para publicar e subscrever

	// Configura as opções do cliente MQTT
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)

	// Define o callback para quando uma mensagem é recebida
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Recebida a mensagem: %s do tópico: %s\n", msg.Payload(), msg.Topic())
	})

	// Define o callback para reconexão automática
	opts.OnConnect = func(client mqtt.Client) {
		fmt.Println("Conectado ao broker MQTT")
		if token := client.Subscribe(topic, 1, nil); token.Wait() && token.Error() != nil {
			log.Fatal(token.Error())
		}
		fmt.Printf("Inscrito no tópico %s\n", topic)
	}

	// Inicializa o cliente MQTT
	client := mqtt.NewClient(opts)

	// Conecta ao broker
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
}
