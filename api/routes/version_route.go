package routes

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

type rootResponse struct {
	Message       string `json:"message"`
	Documentation string `json:"documentation"`
	Version       string `json:"version"`
	Status        string `json:"status"`
}

// @summary Get the API infos
// @description Get the API infos
// @tags root
// @accept */*
// @produce application/json
// @success 200 {object} rootResponse
// @router / [get]
func route(c *fiber.Ctx) error {
	res := rootResponse{
		Message:       "Bienvenue sur notre API",
		Documentation: "http://localhost:8080/swagger/index.html",
		Version:       os.Getenv("VERSION"),
		Status:        "OK",
	}

	return c.JSON(res)
}

type versionResponse struct {
	Version string `json:"version"`
	IsAlive bool   `json:"isAlive"`
}

// @summary Get the API version
// @description Get the API version
// @tags version
// @accept */*
// @produce application/json
// @success 200 {object} versionResponse
// @router /version [get]
func versionInfos(c *fiber.Ctx) error {
	res := versionResponse{
		Version: os.Getenv("VERSION"),
		IsAlive: true,
	}

	return c.JSON(res)
}

func version(app *fiber.App) {
	app.Get("/", route)
	app.Get("/version", versionInfos)
}
