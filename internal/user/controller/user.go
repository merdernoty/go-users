package controller

import (
	"go-users/internal/user/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {
	s domain.UserService
}

func NewUserController(s domain.UserService) *UserController {
	return &UserController{s:s}
}

func (ctrl *UserController) RegisterRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	users.GET("/", ctrl.getAll)
	users.GET("/:id", ctrl.getByID)
	users.POST("/", ctrl.create)
	users.PUT("/:id", ctrl.update)
	users.DELETE("/:id", ctrl.delete)
}

func (ctrl *UserController) getAll(c *gin.Context) {
	users, err := ctrl.s.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (ctrl *UserController) getByID(c *gin.Context) {
	id , _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user , err := ctrl.s.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (ctrl *UserController) create(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.s.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (ctrl *UserController) update(c *gin.Context) {
	id , err := strconv.ParseInt(c.Param("id"), 10, 64); if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	user , err := ctrl.s.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = id
	if err := ctrl.s.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (ctrl *UserController) delete(c *gin.Context) {
	id , err := strconv.ParseInt(c.Param("id"), 10, 64); if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	if err := ctrl.s.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}