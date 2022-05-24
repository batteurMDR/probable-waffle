package auth

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"louis/pw/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var hmacSampleSecret []byte = []byte("poZFZoerfTYJkoZFREfzefPAZMELAaefkoarg")

type CustomClaims struct {
	IDUser string `json:"id_user"`
	jwt.StandardClaims
}

func NewJWT(u *model.User) (string, error) {

	var cc CustomClaims
	cc.IDUser = u.ID
	cc.ExpiresAt = time.Now().Add(time.Hour * 2).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &cc)

	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		val := ctx.GetHeader("Authorization")
		if len(val) == 0 {
			log.Println("auth: verify JWT access without Authorisation vaulues")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !strings.Contains(val, "Bearer ") || len(val) < 100 {
			log.Println("auth: verify JWT access without Bearer value or too small")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := jwt.ParseWithClaims(val[7:], &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("auth: verify JWT Unexpected signing method: %v", token.Header["alg"])
			}
			return hmacSampleSecret, nil
		})
		if err != nil {
			log.Printf("auth: verify JWT error parse %v", err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		log.Printf("%t", token.Claims)

		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			log.Printf("auth: verify JWT with idUser %v expires at %v\n", claims.IDUser, claims.ExpiresAt)
			ctx.Set("session", claims)
		} else {
			log.Printf("auth: verify JWT assert not succeed")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
