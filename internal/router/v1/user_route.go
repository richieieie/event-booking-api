package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dto "github.com/richieieie/event-booking/internal/DTO"
	database "github.com/richieieie/event-booking/internal/db"
	"github.com/richieieie/event-booking/internal/repository"
	"github.com/richieieie/event-booking/internal/service"
)

type UserHandler struct {
	iUserService service.IUserService
}

func NewUserHandler(iUserService service.IUserService) *UserHandler {
	return &UserHandler{iUserService: iUserService}
}

func (h UserHandler) Login(c *gin.Context) {
	var userDto dto.LoginUserDTO
	err := c.ShouldBindJSON(&userDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request does not contain enough fields", "message": "Please ensure that your request body contains email, password fields"})
		return
	}

	accessToken, err := h.iUserService.Login(userDto)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Could not log in", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok", "accessToken": accessToken, "refreshToken": "refreshToken"})
}

func (handler UserHandler) SignUp(c *gin.Context) {
	var userDTO dto.SignUpUserDTO
	err := c.ShouldBindJSON(&userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request does not contain enough fields", "message": "Please ensure that your request body contains email, password fields"})
		return
	}

	id, err := handler.iUserService.SignUp(userDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create data", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "ok", "id": id})
}

func InitUserHandler(r *gin.RouterGroup) {
	// Repository
	var iUserRepository repository.IUserRepository = repository.NewUserRepository(database.Db)

	// Services
	var iUserService service.IUserService = service.NewUserService(iUserRepository)

	handler := NewUserHandler(iUserService)

	userV1 := r.Group("/users")
	{
		userV1.POST("/login", handler.Login)
		userV1.POST("/signup", handler.SignUp)
	}
}
