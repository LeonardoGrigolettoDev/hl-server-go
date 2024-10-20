package redis

import (
	"context"
	"fmt"
	"log"

	redis "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// Inicializar a conexão com o Redis
func ConnectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // endereço do Redis
		Password: "",               // sem senha, por padrão
		DB:       0,                // usar o banco de dados padrão
	})

	// Testando a conexão
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Erro ao conectar ao Redis: %v", err)
	}
	fmt.Printf("Conexão com Redis: %s\n", pong)
	return rdb
}
