package controler

import (
	"belajar-api-goleng/databases"
	"belajar-api-goleng/models"

	"github.com/gin-gonic/gin"
)

func GetCarts(ctx *gin.Context) {
	carts := new([]models.CartsTabel)

	errDb := databases.DB.Table("carts").Find(&carts).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "get data carts success",
		"datas":   carts,
	})
}

func CreateCart(ctx *gin.Context) {
	cart := new(models.CartsTabel)
	payloadCart := new(models.PayloadCart)
	product := new(models.Product)

	if errReq := ctx.ShouldBind(&payloadCart); errReq != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": errReq.Error(),
		})
		return
	}

	errProduct := databases.DB.Table("products").Where("id = ?", payloadCart.IdProduct).Find(&product).Error
	if errProduct != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "internal server product error",
		})
		return
	}

	if product.Id == nil {
		ctx.JSON(404, gin.H{
			"error":   true,
			"message": "data product by id not found",
		})
		return
	}

	total := uint64(*product.Price) * uint64(payloadCart.Quantity)

	cart.IdProduct = &payloadCart.IdProduct
	cart.Quantity = &payloadCart.Quantity
	cart.PriceTotal = &total

	errDB := databases.DB.Table("carts").Create(&cart).Error
	if errDB != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "create data cart error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "create data cart success",
		"datas":   cart,
	})
}

func UpdateCart(ctx *gin.Context) {
	id := ctx.Param("id")
	payloadCart := new(models.PayloadCart)
	product := new(models.Product)
	cart := new(models.CartsTabel)

	if errReques := ctx.ShouldBind(&payloadCart); errReques != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": errReques.Error(),
		})
		return
	}

	errIdProduct := databases.DB.Table("products").Where("id = ?", payloadCart.IdProduct).Find(&product).Error
	if errIdProduct != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "internal product by id server error",
		})
		return
	}

	if product.Id == nil {
		ctx.JSON(404, gin.H{
			"error":   true,
			"message": "get data product by id not found",
		})
		return
	}

	total := uint64(*product.Price) * uint64(payloadCart.Quantity)

	cart.IdProduct = &payloadCart.IdProduct
	cart.Quantity = &payloadCart.Quantity
	cart.PriceTotal = &total

	errDb := databases.DB.Table("carts").Where("id = ?", id).Updates(&cart).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "update data cart error",
		})
		return
	}

	if cart.Id == nil {
		ctx.JSON(404, gin.H{
			"error":   true,
			"message": "get data cart by id not found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "update data cart by id success",
		"datas":   cart,
	})
}

func DeleteCart(ctx *gin.Context) {
	id := ctx.Param("id")
	cart := new(models.CartsTabel)

	errIdcart := databases.DB.Table("carts").Where("id = ?", id).Find(&cart).Error
	if errIdcart != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "internal sever error",
		})
		return
	}

	if cart.Id == nil {
		ctx.JSON(404, gin.H{
			"error":   true,
			"message": "data cart by id not found",
		})
		return
	}

	errDb := databases.DB.Table("carts").Where("id = ?", id).Delete(&cart).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "internal sever error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "delete data cart by id success",
	})
}
