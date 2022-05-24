package service

import (
	"log"
	"net/http"

	"louis/pw/cache"
	"louis/pw/db"
	"louis/pw/model"

	"github.com/gin-gonic/gin"
)

type serviceSell struct {
	db    *db.Store
	cache *cache.Cache
}

// CreateSell is a handler that creates a new sell.
func (s *serviceSell) CreateSell(ctx *gin.Context) {
	var payload model.Sell
	err := ctx.BindJSON(&payload)
	if err != nil {
		log.Println("error create sell", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	se, err := s.db.Sell.CreateSell(&payload)
	if err != nil {
		log.Printf("error create sell %T %#v", err, err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, se)
}

// GetAllSell gives in JSON the sell list.
func (s *serviceSell) GetAllSell(ctx *gin.Context) {
	se, err := s.db.Sell.GetAllSell()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, se)
}

// GetSellById retrives a sell form the given id url param.
func (s *serviceSell) GetSellById(ctx *gin.Context) {
	se, err := s.db.Sell.GetSellById(ctx.Param("id"))
	if err != nil {
		log.Println("error getting sell", err)
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, se)
}
