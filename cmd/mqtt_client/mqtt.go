package mqtt_client

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type ClientConfig struct {
	Broker                string
	ClientID              string
	Username              string
	Password              string
	DefaultPublishHandler mqtt.MessageHandler
	OnConnectHandler      mqtt.OnConnectHandler
}

func ConnectToBroker(clientConfig *ClientConfig) mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(clientConfig.Broker)
	opts.SetClientID(clientConfig.ClientID)
	opts.SetUsername(clientConfig.Username)
	opts.SetDefaultPublishHandler(clientConfig.DefaultPublishHandler)
	opts.OnConnect = clientConfig.OnConnectHandler
	clientConn := mqtt.NewClient(opts)
	return clientConn
}

func SubscribeOnTopic(client mqtt.Client) {
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
}
