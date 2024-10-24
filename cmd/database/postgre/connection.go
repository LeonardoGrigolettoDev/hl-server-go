package postgre

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

// Função para conectar ao banco de dados PostgreSQL
func PostgreSQLConnectDB() (*sql.DB, error) {
	// Pegando variáveis de ambiente
	dbHost := os.Getenv("DB_HOST_PG")
	dbPort := os.Getenv("DB_PORT_PG")
	dbUser := os.Getenv("DB_USER_PG")
	dbPassword := os.Getenv("DB_PASSWORD_PG")
	dbName := os.Getenv("DB_NAME_PG")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Conectado ao banco de dados com sucesso!")
	return db, nil
}
