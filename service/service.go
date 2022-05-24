package service

import (
	"louis/pw/auth"
	"louis/pw/cache"
	"louis/pw/db"

	"github.com/gin-gonic/gin"
)

func New(router *gin.Engine, db *db.Store, cacheDB *cache.Cache) error {
	su := &serviceUser{
		db:    db,
		cache: cacheDB,
	}
	sa := &serviceAuth{
		db:    db,
		cache: cacheDB,
	}
	sc := &serviceClient{
		db:    db,
		cache: cacheDB,
	}
	sp := &serviceProduct{
		db:    db,
		cache: cacheDB,
	}
	ss := &serviceSell{
		db:    db,
		cache: cacheDB,
	}

	verif := auth.VerifyJWT()
	getCache := cache.GetCache(cacheDB)

	v1 := router.Group("/v1")

	v1.POST("/login", sa.Login)

	v1.POST("/users", getCache, su.CreateUser)
	v1.GET("/users", su.GetAllUser)
	v1.GET("/users/:id", verif, su.GetUserById)
	v1.PATCH("/users/:id", su.UpdateUser)
	v1.PUT("/users/:id", su.UpdateUser)
	v1.DELETE("/users/:id", su.DeleteUser)

	v1.GET("/clients", verif, sc.GetAllClient)
	v1.GET("/clients/:id", verif, sc.GetClientById)
	v1.POST("/clients", verif, sc.CreateClient)
	v1.PATCH("/clients/:id", verif, sc.UpdateClient)
	v1.PUT("/clients/:id", verif, sc.UpdateClient)
	v1.POST("/clients/identity/:id", verif, sc.UploadIdentityCard)
	v1.POST("/clients/validate/:id", verif, sc.ValidateClient)
	v1.POST("/clients/unvalidate/:id", verif, sc.UnvalidateClient)
	v1.DELETE("/clients/:id", verif, sc.DeleteClient)

	v1.GET("/products", verif, sp.GetAllProduct)
	v1.GET("/products/:id", verif, sp.GetProductById)
	v1.POST("/products", verif, sp.CreateProduct)
	v1.PATCH("/products/:id", verif, sp.UpdateProduct)
	v1.PUT("/products/:id", verif, sp.UpdateProduct)
	v1.DELETE("/products/:id", verif, sp.DeleteProduct)

	v1.GET("/sells", verif, ss.GetAllSell)
	v1.GET("/sells/:id", verif, ss.GetSellById)
	v1.POST("/sells", verif, ss.CreateSell)

	return nil
}
