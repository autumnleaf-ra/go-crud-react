package main

import (
	"github.com/autumnleaf-ra/go-crud-react/bootstrap"
	"github.com/autumnleaf-ra/go-crud-react/repository"
	"github.com/gofiber/fiber/v2"
)

type Repository repository.Repository

func main() {
	app := fiber.New()
	bootstrap.InitializeAPp(app)
}
