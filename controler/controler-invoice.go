package controler

import (
	"belajar-api-goleng/databases"
	"belajar-api-goleng/midelware"
	"belajar-api-goleng/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetInvoice(ctx *gin.Context) {
	invoice := new([]models.Invoice)

	errDb := databases.DB.Table("invoice").Where("user_id = ?", midelware.UserId).Preload("Carts.Products").Preload("Addresses").
		Preload("Users", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "username", "email", "role")
		}).Find(&invoice).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "get data invoice success",
		"datas":   invoice,
	})
}

func CreateInvoice(ctx *gin.Context) {
	payloadInvoice := new(models.PayloadInvoice)
	cart := new(models.Carts)
	address := new(models.Addresses)
	invoice := new(models.Invoice)

	if errReques := ctx.ShouldBind(&payloadInvoice); errReques != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": errReques.Error(),
		})
		return
	}

	errDbCart := databases.DB.Table("carts").Where("id = ?", payloadInvoice.IdCart).Find(&cart).Error
	if errDbCart != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "internal server cart error",
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

	errDbAddress := databases.DB.Table("addresses").Where("id = ?", payloadInvoice.IdAddress).Find(&address).Error
	if errDbAddress != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "internal server address error",
		})
		return
	}

	if address.Id == nil {
		ctx.JSON(404, gin.H{
			"error":   true,
			"message": "data address by id not found",
		})
		return
	}

	if payloadInvoice.Diskon != nil {
		total := uint64(*cart.PriceTotal) - uint64(*payloadInvoice.Diskon)

		invoice.IdCart = cart.Id
		invoice.IdAddress = address.Id
		invoice.Diskon = payloadInvoice.Diskon
		invoice.UserId = midelware.UserId
		invoice.PriceTotalInvoice = &total

		errCreateInvoice := databases.DB.Table("invoice").Create(&invoice).Error
		if errCreateInvoice != nil {
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": "create data invoice error 1",
			})
			return
		}
	} else {
		invoice.IdCart = cart.Id
		invoice.IdAddress = address.Id
		invoice.Diskon = nil
		invoice.UserId = midelware.UserId
		invoice.PriceTotalInvoice = cart.PriceTotal

		errCreateInvoice := databases.DB.Table("invoice").Create(&invoice).Error
		if errCreateInvoice != nil {
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": "create data invoice error 2",
			})
			return
		}
	}

	ctx.JSON(201, gin.H{
		"error":   true,
		"message": "create data success",
		"datas":   invoice,
	})
}

func UpdateInvoice(ctx *gin.Context) {
	id := ctx.Param("id")
	payloadInvoice := new(models.PayloadInvoice)
	cart := new(models.Carts)
	address := new(models.Addresses)
	invoice := new(models.Invoice)

	if errReques := ctx.ShouldBind(&payloadInvoice); errReques != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": errReques.Error(),
		})
		return
	}

	errDbCart := databases.DB.Table("carts").Where("id = ?", payloadInvoice.IdCart).Find(&cart).Error
	if errDbCart != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "internal server cart error",
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

	errDbAddress := databases.DB.Table("addresses").Where("id = ?", payloadInvoice.IdAddress).Find(&address).Error
	if errDbAddress != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "internal server address error",
		})
		return
	}

	if address.Id == nil {
		ctx.JSON(404, gin.H{
			"error":   true,
			"message": "data address by id not found",
		})
		return
	}

	errDbInvoiceId := databases.DB.Table("invoice").Where("id = ?", id).Find(&invoice).Error
	if errDbInvoiceId != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "update data invoice error 1",
		})
		return
	}

	if invoice.Id == nil {
		ctx.JSON(404, gin.H{
			"error":   true,
			"message": "data invoice by id not found",
		})
		return
	}

	if payloadInvoice.Diskon != nil {
		total := uint64(*cart.PriceTotal) - uint64(*payloadInvoice.Diskon)

		invoice.IdCart = cart.Id
		invoice.Diskon = payloadInvoice.Diskon
		invoice.IdAddress = address.Id
		invoice.PriceTotalInvoice = &total
		invoice.UserId = midelware.UserId

		errDbInvoice := databases.DB.Table("invoice").Where("id = ?", id).Updates(&invoice).Error
		if errDbInvoice != nil {
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": "update data invoice error 1",
			})
			return
		}
	} else {
		invoice.IdCart = cart.Id
		invoice.Diskon = nil
		invoice.IdAddress = address.Id
		invoice.PriceTotalInvoice = cart.PriceTotal
		invoice.UserId = midelware.UserId

		errDbInvoice := databases.DB.Table("invoice").Where("id = ?", id).Updates(&invoice).Error
		if errDbInvoice != nil {
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": "update data invoice error 2",
			})
			return
		}
	}

	ctx.JSON(201, gin.H{
		"error":   false,
		"message": "update invoice by id success",
		"datas":   invoice,
	})
}
