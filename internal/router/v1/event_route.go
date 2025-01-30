package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	dto "github.com/richieieie/event-booking/internal/DTO"
	database "github.com/richieieie/event-booking/internal/db"
	middlewares "github.com/richieieie/event-booking/internal/middleware"
	"github.com/richieieie/event-booking/internal/repository"
	"github.com/richieieie/event-booking/internal/service"
)

type EventHandler struct {
	iEventService service.IEventService
}

func NewEventHandler(iEventService service.IEventService) *EventHandler {
	return &EventHandler{iEventService: iEventService}
}

func (handler EventHandler) GetEvents(c *gin.Context) {
	events, err := handler.iEventService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch data", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok", "events": events})
}

func (handler *EventHandler) GetEventById(c *gin.Context) {
	idParsed, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid id", "message": err.Error()})
		return
	}

	event, err := handler.iEventService.GetById(idParsed)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Could not find your event", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "event": event})
}

func (handler EventHandler) CreateEvent(c *gin.Context) {
	userId := c.GetInt64("userId")
	var eventDTO dto.CreateEventDTO
	err := c.ShouldBindJSON(&eventDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request does not contain enough fields", "message": "Please ensure that your request body contains name, description, location and dateTime fields"})
		return
	}
	eventDTO.UserId = userId
	id, err := handler.iEventService.CreateOne(eventDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create data", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "ok", "id": id})
}

func (handler EventHandler) UpdateEventById(c *gin.Context) {
	userId := c.GetInt64("userId")
	var eventDTO dto.EventUpdateDTO
	err := c.ShouldBindJSON(&eventDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request does not contain enough fields", "message": "Please ensure that your request body contains name, description, location and date_time fields"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id", "message": "Please ensure that the event id is correct"})
		return
	}

	err = handler.iEventService.UpdateOne(id, eventDTO, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot update", "message": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (handler EventHandler) DeleteEventById(c *gin.Context) {
	userId := c.GetInt64("userId")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id", "message": "Please ensure that the event id is correct"})
		return
	}

	err = handler.iEventService.DeleteOne(id, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot delete", "message": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func InitEventHandler(r *gin.RouterGroup) {
	// Repository
	var iEventRepository repository.IEventRepository = repository.NewEventRepository(database.Db)

	// Services
	var iEventService service.IEventService = service.NewEventService(iEventRepository)

	handler := NewEventHandler(iEventService)

	eventV1 := r.Group("/events")
	{
		eventV1.GET("/", handler.GetEvents)
		eventV1.GET("/:id", handler.GetEventById)
		authenticate := eventV1.Group("/")
		authenticate.Use(middlewares.Authenticate)
		authenticate.POST("/", middlewares.Authenticate, handler.CreateEvent)
		authenticate.PUT("/:id", handler.UpdateEventById)
		authenticate.DELETE("/:id", handler.DeleteEventById)
	}
}
