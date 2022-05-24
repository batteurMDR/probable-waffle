package service

import (
	"log"
	"net/http"
	"os"

	"louis/pw/cache"
	"louis/pw/db"
	"louis/pw/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type serviceClient struct {
	db    *db.Store
	cache *cache.Cache
}

// CreateClient is a handler that creates a new customer.
func (s *serviceClient) CreateClient(ctx *gin.Context) {
	var payload model.Client
	err := ctx.BindJSON(&payload)
	if err != nil {
		log.Println("error create client", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c, err := s.db.Client.CreateClient(&payload)
	if err != nil {
		log.Printf("error create client %T %#v", err, err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, c)
}

// GetAllClient gives in JSON the customer list.
func (s *serviceClient) GetAllClient(ctx *gin.Context) {
	c, err := s.db.Client.GetAllClient()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, c)
}

// GetClientById retrives a customer form the given id url param.
func (s *serviceClient) GetClientById(ctx *gin.Context) {
	c, err := s.db.Client.GetClientById(ctx.Param("id"))
	if err != nil {
		log.Println("error getting client", err)
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, c)
}

// UpdateClient update an existing customer from the url parm id and the given body payload.
func (s *serviceClient) UpdateClient(ctx *gin.Context) {
	payload := make(map[string]interface{})
	err := ctx.BindJSON(&payload)
	if err != nil {
		log.Println("error update client", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c, err := s.db.Client.UpdateClient(ctx.Param("id"), payload)
	if err != nil {
		log.Println("error update client", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, c)
}

// UpdateClient update an existing customer from the url parm id and validate it.
func (s *serviceClient) ValidateClient(ctx *gin.Context) {
	c, err := s.db.Client.ValidateClient(ctx.Param("id"))
	if err != nil {
		log.Println("error update client", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, c)
}

// UpdateClient update an existing customer from the url parm id and unvalidate it.
func (s *serviceClient) UnvalidateClient(ctx *gin.Context) {
	c, err := s.db.Client.UnvalidateClient(ctx.Param("id"))
	if err != nil {
		log.Println("error update client", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, c)
}

// UploadIdentityCard update customer with new identity card
func (s *serviceClient) UploadIdentityCard(ctx *gin.Context) {
	c, err := s.db.Client.GetClientById(ctx.Param("id"))
	if err != nil {
		log.Println("error update client", err)
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		log.Println("handler: upload identity card error", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if err := os.Mkdir("./upload/"+c.ID, 0777); err != nil {
		log.Println("handler: upload identity card error", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}
	file_path := "./upload/" + c.ID + "/" + uuid.New().String() + ".jpeg"
	if err := ctx.SaveUploadedFile(file, file_path); err != nil {
		log.Println("handler: upload identity card error", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}

	payload := make(map[string]interface{})
	payload["id_card_path"] = file_path

	uc, err := s.db.Client.UpdateClient(ctx.Param("id"), payload)
	if err != nil {
		log.Println("error update client", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, uc)
}

// DeleteClient is deleting a customer from the url parm id.
func (s *serviceClient) DeleteClient(ctx *gin.Context) {
	err := s.db.Client.DeleteClient(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"client": "deleted"})
}
