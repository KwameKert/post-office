package app

import (
	//"encoding/json"

	// "strconv"
	//	"gorm.io/gorm"

	"go.mongodb.org/mongo-driver/mongo"

	"postoffice/app/core"
	"postoffice/app/core/database"

	//	"postoffice/app/models"
	"postoffice/app/repository"
	"postoffice/app/routes"
	"postoffice/app/services"

	"github.com/joho/godotenv"

	log "github.com/sirupsen/logrus"
)

type App struct{}

func init() {
	loadEnvironmentVariables()
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}

func loadEnvironmentVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func (app *App) Start(conf *core.Config) {
	log.Info("Starting Post Office server")
	connection := setupDatabase(conf)
	repo := repository.NewRepository(connection.Database("agerp-post-office"))
	services := services.NewService(repo, conf)

	server := core.NewHTTPServer(conf)
	router := routes.NewRouter(server.Engine, conf, services)

	router.RegisterRoutes()
	server.Start()

}

func setupDatabase(conf *core.Config) *mongo.Client {
	mg, err := database.GetMongoClient(conf)
	if err != nil {
		log.Fatal("failed to initialize postgres database. err:", err)
		panic(err.Error())
	}
	if err != nil {
		log.Fatal("failed to run migrations. err:", err)
	}

	return mg
}
