package fiberserver

import (
	"github.com/gofiber/fiber/v2"
)

type fiberServer struct {
	app *fiber.App
}

func New() *fiberServer {
	return &fiberServer{
		app: fiber.New(),
	}
}
