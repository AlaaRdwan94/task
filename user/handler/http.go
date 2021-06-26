package handler

import (
	"errors"
	"github.com/InnoSoft/task/middleware"
	"github.com/InnoSoft/task/middleware/auth"
	"github.com/InnoSoft/task/model"
	"github.com/InnoSoft/task/user"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
)

type UserHandler struct {
	Uusecase user.Usecase
}

func (h UserHandler) RegisterUser(c *gin.Context) {

	c.Writer.Header().Set("Content-Type", "application/json")
	user := model.UserData{}
	if err := c.ShouldBindBodyWith(&user, binding.JSON); err != nil {
		log.Printf("%+v", err)
	}
	id , err := h.Uusecase.CreateAccount(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError,err)
		return
	}
	byId, err := h.Uusecase.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError,err)
		return
	}
	c.JSON(http.StatusCreated,byId)
}
var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
func (h UserHandler) Login(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	login := model.LoginData{}
	if err := c.ShouldBindBodyWith(&login, binding.JSON); err != nil {
		log.Printf("%+v", err)
	}
	user , err := h.Uusecase.GetUserByEmail(login.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError,err)
		return
	}

	if login.Password != user.Password {
		c.JSON(http.StatusUnauthorized,errors.New("wrong password").Error())
		return
	}

	token, err := auth.CreateToken(uint32(user.ID))
	if err != nil {
		c.JSON(http.StatusUnauthorized,err)
		return
	}
	user.Token = token
	c.JSON(http.StatusCreated,user)
	return
}

func (h UserHandler) UpdateUserProfilePicture(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	updateProfilePic := model.UpdateProfilePicture{}
	if err := c.ShouldBindBodyWith(&updateProfilePic, binding.JSON); err != nil {
		log.Printf("%+v", err)
	}
	key := c.Writer.Header().Get("uid")
	atoi, err := strconv.Atoi(key)
	if err != nil {
		c.JSON(http.StatusUnauthorized,err)
		return
	}
	id := uint(atoi)
	user, err := h.Uusecase.UpdateProfilePic(updateProfilePic.Url, id)
	if err != nil {
		c.JSON(http.StatusUnauthorized,err)
		return
	}
	user.Token = c.Request.Header.Get("Authorization")
	c.JSON(http.StatusCreated,user)
}

func NewUserHandler(e *gin.RouterGroup, uus user.Usecase)  {
	handler := &UserHandler{Uusecase: uus}
	e.POST("/register-user",handler.RegisterUser)
	e.POST("/login",handler.Login)
	e.PUT("/update-profile-pic",middleware.SetMiddlewareAuthentication(handler.UpdateUserProfilePicture))
	e.GET("/send-notification",handler.Send)
}