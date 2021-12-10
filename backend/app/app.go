package app

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"skema/app/controller"
	"skema/config"
	"skema/database"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type App struct {
	Router        *mux.Router
	DatabaseMongo *mongo.Database
	DatabaseGorm  *gorm.DB
}

func InitAndRun(config *config.Config) {
	app := new(App)

	app.Init(config)
	app.Run(config.ServerHost)
}

func (app *App) Init(config *config.Config) {
	// app.DatabaseMongo = database.InitialConnect("golang", config.MongoHost)
	db, err := database.StubConnection()

	if err != nil {
		log.Printf("Error database init")
	}

	app.DatabaseGorm = db

	app.Router = mux.NewRouter()
	app.SetRouters()
}

func (app *App) SetRouters() {
	controller.InitUserController(app.Router, app.DatabaseGorm)
	controller.InitSchemaController(app.Router, app.DatabaseGorm)
}

func (app *App) Run(host string) {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt, os.Kill)

	go func() {
		log.Fatal(http.ListenAndServe(host, app.Router))
	}()

	log.Printf("Server is listening on %s\n", host)
	sig := <-sigs
	log.Println("Signal:", sig)

	// log.Println("Stoping MongoDB Connection...")
	// app.DatabaseMongo.Client().Disconnect(context.Background())
}

func (app *App) createIndexes() {
	// username and email will be unique.
	keys := bsonx.Doc{
		{Key: "username", Value: bsonx.Int32(1)},
		{Key: "email", Value: bsonx.Int32(1)},
	}

	user := app.DatabaseMongo.Collection("user")
	database.SetIndexes(user, keys)
}
