package main

import (
	"log"
	"os"

	"github.com/GuilhermeW1/backend-suino/api/handler"
	"github.com/GuilhermeW1/backend-suino/db"
	"github.com/GuilhermeW1/backend-suino/middleware"
	"github.com/GuilhermeW1/backend-suino/repository"
	"github.com/GuilhermeW1/backend-suino/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found")
	}

	db := db.Init()

	r := gin.Default()
	r.Use(middleware.APIKeyMiddleware(os.Getenv("API_KEY")))
	r.Use(middleware.RateLimitMiddleware())

	sowRepo := &repository.SowRepository{DB: db}
	cycleRepo := &repository.CycleRepository{DB: db}
	eventRepo := &repository.EventRepository{DB: db}

	sowService := &service.SowService{R: sowRepo}
	cycleService := &service.CycleService{R: cycleRepo}

	eventService := &service.EventService{R: eventRepo, CycleService: cycleService}

	sowHandler := &handler.SowHandler{Service: sowService}
	eventHandler := &handler.EventHandler{S: eventService}

	// sowHandler := &handler.

	//TODO: delete sow isnt changing sow status, and not closing her cycles
	sows := r.Group("/sows")
	{
		sows.POST("/", sowHandler.Create)
		sows.GET("/:id", sowHandler.GetById)
		sows.GET("/", sowHandler.GetAllActive)
		sows.DELETE("/:id", sowHandler.Delete)
		sows.GET(("/earTag/:earTag"), sowHandler.GetByEarTag)
	}

	events := r.Group("/events")
	{
		events.POST("/", eventHandler.AddEvent)
		events.GET("/", eventHandler.GetEvents)
		events.GET("/:sowId", eventHandler.GetEventsBySowId)
	}

	r.Run(":8000")
}
