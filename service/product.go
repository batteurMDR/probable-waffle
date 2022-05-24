package service

import (
	"log"
	"net/http"

	"louis/pw/cache"
	"louis/pw/db"
	"louis/pw/model"

	"github.com/gin-gonic/gin"
)

type serviceProduct struct {
	db    *db.Store
	cache *cache.Cache
}

// CreateProduct is a handler that creates a new product.
func (s *serviceProduct) CreateProduct(ctx *gin.Context) {
	var payload model.Product
	err := ctx.BindJSON(&payload)
	if err != nil {
		log.Println("error create product", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	p, err := s.db.Product.CreateProduct(&payload)
	if err != nil {
		log.Printf("error create product %T %#v", err, err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, p)
}

// GetAllProduct gives in JSON the products list.
func (s *serviceProduct) GetAllProduct(ctx *gin.Context) {
	p, err := s.db.Product.GetAllProduct()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, p)
}

// GetProductById retrives a product form the given id url param.
func (s *serviceProduct) GetProductById(ctx *gin.Context) {
	p, err := s.db.Product.GetProductById(ctx.Param("id"))
	if err != nil {
		log.Println("error getting product", err)
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, p)
}

// UpdateProduct update an existing product from the url parm id and the given body payload.
func (s *serviceProduct) UpdateProduct(ctx *gin.Context) {
	payload := make(map[string]interface{})
	err := ctx.BindJSON(&payload)
	if err != nil {
		log.Println("error update product", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	p, err := s.db.Product.UpdateProduct(ctx.Param("id"), payload)
	if err != nil {
		log.Println("error update product", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, p)
}

// DeleteProduct is deleting a product from the url parm id.
func (s *serviceProduct) DeleteProduct(ctx *gin.Context) {
	err := s.db.Product.DeleteProduct(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"product": "deleted"})
}
