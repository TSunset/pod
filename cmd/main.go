package main

import (
	"log"
	"pod/internal/db"
	"pod/internal/handler"
	messageserver "pod/internal/messageServer"

	"github.com/labstack/echo/v4"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to DB^ %v", err)
	}

	messRepo := messageserver.NewMessageRepository(database)
	messService := messageserver.NewMessageService(messRepo)
	messHandlers := handler.NewMessageHandler(messService)

	e := echo.New()

	e.GET("/messages", messHandlers.GetHandler)
	e.POST("/messages", messHandlers.PostHandler)
	e.PATCH("/messages/:id", messHandlers.PutHandler)
	e.DELETE("/messages/:id", messHandlers.DeleteHandler)

	e.Start(":8080")
}
