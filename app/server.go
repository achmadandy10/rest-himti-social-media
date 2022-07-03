package app

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"service_social_media/app/databases/seeder"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

type AppConfig struct {
	AppName     string
	AppEnv      string
	AppPort     string
	AppTimeZone string
}

type DBConfig struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUsername string
	DBPassword string
	DBDriver   string
}

type GINConfig struct {
	GINMode string
}

func (server *Server) Initialize(appConfig AppConfig, dbConfig DBConfig) {
	fmt.Println("Service " + appConfig.AppName)

	server.InitializeRoutes()
	server.InitializeDB(dbConfig, appConfig)
}

func (server *Server) InitializeDB(dbConfig DBConfig, appConfig AppConfig) {
	var err error

	if dbConfig.DBDriver == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.DBUsername, dbConfig.DBPassword, dbConfig.DBHost, dbConfig.DBPort, dbConfig.DBName)
		server.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", dbConfig.DBHost, dbConfig.DBUsername, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBPort, appConfig.AppTimeZone)
		server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		panic("Failed on connecting to the database server")
	}
}

func (server *Server) dbMigrate() {
	for _, model := range RegisterModels() {
		err := server.DB.Debug().AutoMigrate(model.Model)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Database migrated successfully.")
}

func (server *Server) InitializeCommands(dbConfig DBConfig, appConfig AppConfig) {
	server.InitializeDB(dbConfig, appConfig)

	cmdApp := cli.NewApp()
	cmdApp.Commands = []cli.Command{
		{
			Name: "db:migrate",
			Action: func(c *cli.Context) error {
				server.dbMigrate()
				return nil
			},
		},
		{
			Name: "db:seed",
			Action: func(c *cli.Context) error {
				err := seeder.DBSeed(server.DB)

				if err != nil {
					log.Fatal(err)
				}

				return nil
			},
		},
	}

	err := cmdApp.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) Run(port string) {
	fmt.Printf("Start server to port: %s", port)
	log.Fatal(http.ListenAndServe(port, server.Router))
}

func GetEnv(key, callback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return callback
}

func Run() {
	var server = Server{}
	var appConfig = AppConfig{}
	var dbConfig = DBConfig{}
	var ginConfig = GINConfig{}

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Failed to load env file")
	}

	appConfig.AppName = GetEnv("APP_NAME", "Golang")
	appConfig.AppEnv = GetEnv("APP_ENV", "local")
	appConfig.AppPort = GetEnv("APP_PORT", "8000")
	appConfig.AppTimeZone = GetEnv("APP_TIMEZONE", "Asia/Shanghai")

	dbConfig.DBHost = GetEnv("DB_HOST", "localhost")
	dbConfig.DBPort = GetEnv("DB_PORT", "5432")
	dbConfig.DBName = GetEnv("DB_DATABASE", "golang")
	dbConfig.DBUsername = GetEnv("DB_USERNAME", "postgres")
	dbConfig.DBPassword = GetEnv("DB_PASSWORD", "root")
	dbConfig.DBDriver = GetEnv("DB_DRIVER", "postgres")

	ginConfig.GINMode = GetEnv("GIN_MODE", "debug")

	flag.Parse()

	arg := flag.Arg(0)
	if arg != "" {
		server.InitializeCommands(dbConfig, appConfig)
	} else {
		if ginConfig.GINMode == "release" {
			gin.SetMode(gin.ReleaseMode)
		}
		server.Initialize(appConfig, dbConfig)
		server.Run(":" + appConfig.AppPort)
	}
}
