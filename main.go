package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	postgres "github.com/LeonardoGrigolettoDev/fly-esp-server-go/database/postgre"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv" // Importando o godotenv
	_ "github.com/lib/pq"
)

var _V1_DB_PATH = "./database/postgre/migration/v1/sql"
var _ACTION_PATH = "./database/postgre/actions"
var dbHost string
var dbPort string
var dbUser string
var dbPassword string
var dbName string
var dbString string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error on load file .env: %v", err)
	}
	dbHost = os.Getenv("DB_HOST_PG")
	dbPort = os.Getenv("DB_PORT_PG")
	dbUser = os.Getenv("DB_USER_PG")
	dbPassword = os.Getenv("DB_PASSWORD_PG")
	dbName = os.Getenv("DB_NAME_PG")
	dbString = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
}

func main() {
	// Carregando variáveis de ambiente do arquivo .env

	if len(os.Args) >= 2 {
		var command = os.Args[1]
		log.Println(command)
		switch command {
		case "migrate":
			executeMigrationType()
			return
		case "action":
			executeDBActionType()
			return
		default:
			log.Println("Optional action command not found, initializing normal server.")
		}
	}
	db, err := postgres.PostgreSQLConnectDB()
	if err != nil {
		log.Fatalf("Error on connect DB (PostgreSQL): %v", err)
	}
	defer db.Close()

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("API running!"))
	}).Methods("GET")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func executeDBActionType() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar .env: %v", err)
	}

	if len(os.Args) < 3 {
		log.Println("Action is not specified.")
		return
	}
	var actionType = os.Args[2]
	removeAllDB := func() {
		log.Println("Removing all DB")

		cmd := exec.Command("goose", "-dir", _ACTION_PATH+"/turn_virgin_db_sql", "postgres", dbString, "up")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			log.Fatalf("Erro ao rodar o goose: %v", err)
		}
		fmt.Printf("PostgreSQL is clear.")
	}
	// if actionType == "remove-db" {
	// 	removeAllDB()
	// 	return
	// }
	if actionType == "recreate-db" {
		runUpMigration()
		removeAllDB()
	}

}

func executeMigrationType() {
	if len(os.Args) < 3 {
		runUpMigration()
		return
	}
	var migrationType = os.Args[2]
	if migrationType == "up" {
		runUpMigration()
		return
	}
	if migrationType == "down" {
		log.Println("Migrating DB with the previous migrations instrutions.")
		cmd := exec.Command("goose", "-dir", _V1_DB_PATH, "postgres", dbString, "down")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Fatalf("Erro ao rodar o goose: %v", err)
		}
		fmt.Printf("PostgreSQL DB is built with the previous app DB version.\n")
		return
	}
}

func runUpMigration() {
	log.Println("Migrating DB with the lastest migrations instrutions.")
	cmd := exec.Command("goose", "-dir", _V1_DB_PATH, "postgres", dbString, "up")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Erro ao rodar o goose: %v", err)
	}
	fmt.Printf("PostgreSQL DB is built with the lastest app DB version.\n")
}
