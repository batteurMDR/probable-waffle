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

type serviceAuth struct {
	db    *db.Store
	cache *cache.Cache
}

func (s *serviceAuth) Login(ctx *gin.Context) {
	var payload model.LoginPayload
	err := ctx.BindJSON(&payload)
	if err != nil {
		log.Println("service auth: login user parse payload", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	u, err := s.db.User.GetUserByEmail(payload.Email)
	if err != nil {
		log.Println("servce auth: login user db GetUserByEmail", err)
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if u.Password != payload.Password {
		log.Println("service auth: the password doesn't match")
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	jwtValue, err := auth.NewJWT(u)
	if u.Password != payload.Password {
		log.Println("service auth: create JWT", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"jwt": jwtValue})
}
