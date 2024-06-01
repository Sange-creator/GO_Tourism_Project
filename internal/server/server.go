package server

import (
	"github.com/gofiber/fiber/v2"

	"JAMJAMNEPAL/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "JAMJAMNEPAL",
			AppName:      "JAMJAMNEPAL",
		}),

		db: database.New(),
	}

	return server
}
