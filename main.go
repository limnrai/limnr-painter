package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func Index(c *fiber.Ctx) error {
	currentYear := time.Now().Year()

	return c.Render("index", fiber.Map{
		"Title":       "Limnr AI",
		"CurrentYear": currentYear,
	}, "layouts/limnr")
}

func ViewRoute(route fiber.Router) {
	route.Get("", Index)
	route.Get("/index", Index)
}

func main() {
	viewEngine := html.New("./webapp/views", ".html")

	goApp := fiber.New(fiber.Config{
		Views: viewEngine,
	})

	goApp.Static("/public", "./webapp/public",
		fiber.Static{
			MaxAge:   3600,
			Compress: true,
		},
	)

	// setup routes
	ui := goApp.Group("/")
	ViewRoute(ui.Group("/"))

	// start
	err := goApp.Listen(":3030")
	if err != nil {
		fmt.Printf("%v", err)
	}

}
