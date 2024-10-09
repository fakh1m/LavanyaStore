package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fakh1m/LavanyaStore/database/seeders"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

type DBConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

func (server *Server) Initialize(appConfig AppConfig, dbConfig DBConfig) {
	fmt.Println("Welcome to " + appConfig.AppName)

	server.initializeDB(dbConfig)
	server.initializeRoutes()
	seeders.DBSeed(server.DB)
}

func (server *Server) initializeDB(dbConfig DBConfig) {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbConfig.DBHost, dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBPort)
	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	for _, model := range RegisterModels() {
		err = server.DB.AutoMigrate(model.Model)

		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Database migrate successfully...")
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func (server *Server) Run(addr string) {
	fmt.Printf("Listenig to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func Run() {
	var server = Server{}
	var appConfig = AppConfig{}
	var dbConfig = DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loadding .env file!")
	}

	appConfig.AppName = getEnv("APP_NAME", "Lavanya Store")
	appConfig.AppEnv = getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT", "9000")

	dbConfig.DBHost = getEnv("DB_HOST", "localhost")
	dbConfig.DBUser = getEnv("DB_USER", "kimbozy")
	dbConfig.DBPassword = getEnv("DB_PASSWORD", "password")
	dbConfig.DBName = getEnv("DB_NAME", "lavanya_store")
	dbConfig.DBPort = getEnv("DB_PORT", "5432")

	server.Initialize(appConfig, dbConfig)
	server.Run(":" + appConfig.AppPort)
}
