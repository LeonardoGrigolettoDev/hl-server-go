package redis

import (
	"context"
	"fmt"
	"log"
	"os"

	redis "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// Inicializar a conexão com o Redis
func ConnectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"), // endereço do Redis
		Password: os.Getenv("REDIS_PASS"), // sem senha, por padrão
		DB:       0,                       // usar o banco de dados padrão
	})

	// Testando a conexão
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Erro ao conectar ao Redis: %v", err)
	}
	fmt.Printf("Conexão com Redis: %s\n", pong)
	return rdb
}
