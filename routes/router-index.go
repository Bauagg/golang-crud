package routes

import (
	"belajar-api-goleng/controler"
	"belajar-api-goleng/midelware"

	"github.com/gin-gonic/gin"
)

func RouterIndex(app *gin.Engine) {
	router := app

	router.GET("/product", controler.GetProduct)
	router.GET("/product/:id", controler.GetProductById)
	router.POST("/product", midelware.AdminRoleMiddleware(), controler.CreateProduct)
	router.PUT("/product/:id", midelware.AdminRoleMiddleware(), controler.UpdateProduct)
	router.DELETE("/product/:id", midelware.AdminRoleMiddleware(), controler.DeleteProduct)

	// router user
	router.POST("/register", controler.Register)
	router.POST("/login", controler.Login)

	// router address
	router.GET("/address", midelware.AuthMidelware(), controler.GetAddress)
	router.POST("/address", midelware.AuthMidelware(), controler.CreateAddress)
	router.PUT("/address/:id", midelware.AuthMidelware(), controler.UpdateAddress)
	router.DELETE("/address/:id", midelware.AuthMidelware(), controler.DeleteAddress)

	// router carts
	router.GET("/carts", midelware.AuthMidelware(), controler.GetCarts)
	router.POST("/cart", midelware.AuthMidelware(), controler.CreateCart)
	router.PUT("/cart/:id", midelware.AuthMidelware(), controler.UpdateCart)
	router.DELETE("/cart/:id", controler.DeleteCart)

	// router invoice
}
