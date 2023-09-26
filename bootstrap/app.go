package bootstrap

import (
	"log"
	"os"

	"github.com/autumnleaf-ra/go-crud-react/database/migrations"
	"github.com/autumnleaf-ra/go-crud-react/database/storage"
	"github.com/autumnleaf-ra/go-crud-react/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func InitializeAPp(app *fiber.App) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("Could not load the database !")
	}

	err = migrations.MigrateUsers(db)

	if err != nil {
		log.Fatal("Could not migrates db")
	}

	repo := repository.Repository{
		DB: db,
	}

	app.Use(cors.New(cors.Config{AllowCredentials: true}))
	repo.SetupRoute(app)
	app.Listen(":4000")
}
