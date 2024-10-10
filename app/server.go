package app

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fakh1m/LavanyaStore/database/seeders"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
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

	server.initializeRoutes()
}

func (server *Server) initializeDB(dbConfig DBConfig) {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbConfig.DBHost, dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBPort)
	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}
}

func (server *Server) dbMigrate() {
	for _, model := range RegisterModels() {
		err := server.DB.AutoMigrate(model.Model)

		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Database migrate successfully...")
}

func (server *Server) initCommands(config AppConfig, dbConfig DBConfig) {
	server.initializeDB(dbConfig)

	cmdApp := cli.NewApp()
	cmdApp.Commands = []cli.Command{
		{
			Name: "db:migrate",
			Usage: "menjalankan dbMigrate",
			Action: func (c *cli.Context) error {
				server.dbMigrate()
				return nil
			},
		},
		{
			Name: "db:seed",
			Usage: "menjalankan seeder",
			Action: func (c *cli.Context) error {
				err := seeders.DBSeed(server.DB)
				if err != nil{
					log.Fatal(err)
				}
				return nil
			},
		},
	}
	err := cmdApp.Run(os.Args)
	if err!= nil{
		log.Fatal(err)
	}
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

	flag.Parse()
	arg := flag.Arg(0)
	if arg != "" {
		server.initCommands(appConfig, dbConfig)
	} else {
		server.Initialize(appConfig, dbConfig)
		server.Run(":" + appConfig.AppPort)
	}
}
