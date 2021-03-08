package app

import (
	"fmt"
	"log"

	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/jonathanwamsley/banking_auth/config"
	"github.com/jonathanwamsley/banking_auth/domain"
	"github.com/jonathanwamsley/banking_auth/service"

	// "github.com/jonathanwamsley/banking_auth/domain"
	"github.com/jonathanwamsley/banking_auth/logger"
	// "github.com/jonathanwamsley/banking_auth/service"
)

// getDbClient loads and returns db connection. The db makes connection is confirmed via Ping
func getDbClient(connectionInfo string) *sqlx.DB {

	// dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)

	client, err := sqlx.Open("mysql", connectionInfo)
	if err != nil {
		panic(err)
	}

	if err = client.Ping(); err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}

// Start helps decouples from running the whole entire application
// it connects the handlers, starts the server, and any other configuration setup
func Start() {
	if err := godotenv.Load(); err != nil {
		logger.Fatal("no .env file found")
		panic(err)
	}
	config := config.NewConfig()
	dbClient := getDbClient(config.GetMySQLInfo())
	serverInfo := config.GetServerInfo()

	ah := AuthHandler{service.NewLoginService(domain.NewAuthRepository(dbClient), domain.GetRolePermissions())}
	router := mux.NewRouter()
	router.HandleFunc("/auth/login", ah.Login).Methods(http.MethodPost)
	router.HandleFunc("/auth/verify", ah.Verify).Methods(http.MethodGet)

	logger.Info(fmt.Sprintf("Starting server on %s ...", serverInfo))
	log.Fatal(http.ListenAndServe(serverInfo, router))
}
