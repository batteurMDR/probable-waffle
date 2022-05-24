package service

import (
	"log"
	"net/http"

	"louis/pw/auth"
	"louis/pw/cache"
	"louis/pw/db"
	"louis/pw/model"

	"github.com/gin-gonic/gin"
)

type serviceUser struct {
	db    *db.Store
	cache *cache.Cache
}

// GetAllUser gives in JSON the user list.
func (s *serviceUser) GetAllUser(ctx *gin.Context) {
	us, err := s.db.User.GetAllUser()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, us)
}

// CreateUser is a handler that creates a new user.
func (s *serviceUser) CreateUser(ctx *gin.Context) {
	var payload model.User
	err := ctx.BindJSON(&payload)
	if err != nil {
		log.Println("error create user", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	u, err := s.db.User.CreateUser(&payload)
	if err != nil {
		log.Printf("error create user %T %#v", err, err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, u)
}

// DeleteUser is deleting a user from the url parm id.
func (s *serviceUser) DeleteUser(ctx *gin.Context) {
	err := s.db.User.DeleteUser(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"user": "deleted"})
}

// UpdateUser update an existing user from the url parm id and the given body payload.
func (s *serviceUser) UpdateUser(ctx *gin.Context) {

	payload := make(map[string]interface{})
	err := ctx.BindJSON(&payload)
	if err != nil {
		log.Println("error create user", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	u, err := s.db.User.UpdateUser(ctx.Param("id"), payload)
	if err != nil {
		log.Println("error create user", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, u)
}

// GetUserById retrives a user form the given id url param.
func (s *serviceUser) GetUserById(ctx *gin.Context) {
	u, err := s.db.User.GetUserById(ctx.Param("id"))
	if err != nil {
		log.Println("error create user", err)
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	sessionValue, ok := ctx.MustGet("session").(*auth.CustomClaims)
	if !ok {
		log.Println("error assert session")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	log.Println("from session get id user", sessionValue.IDUser)

	ctx.JSON(http.StatusOK, u)
}
