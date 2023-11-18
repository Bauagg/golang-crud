package controler

import (
	"belajar-api-goleng/databases"
	"belajar-api-goleng/models"

	"github.com/gin-gonic/gin"
)

func GetProduct(ctx *gin.Context) {
	product := new([]models.Products)
	search := databases.DB.Table("products")

	productName := ctx.Query("name")
	if productName != "" {
		search = search.Where("name_product LIKE ?", "%"+productName+"%")
	}

	err := search.Find(&product).Error

	if err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "get data product success",
		"datas":   product,
	})
}

func GetProductById(ctx *gin.Context) {
	id := ctx.Param("id")
	product := new(models.Products)

	err := databases.DB.Table("products").Where("id = ?", id).Find(&product).Error
	if err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "internal server error",
		})
		return
	}

	if product.Id == nil {
		ctx.JSON(404, gin.H{
			"error":   true,
			"message": "get data by ID NOT FOUNT",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "get data by ID success",
		"datas":   product,
	})
}

func CreateProduct(ctx *gin.Context) {
	productReques := new(models.CreateProduct)

	if errRequire := ctx.ShouldBind(&productReques); errRequire != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errRequire.Error(),
		})
		return
	}

	product := new(models.Products)
	product.NameProduct = &productReques.NameProduct
	product.Stock = &productReques.Stock
	product.Price = &productReques.Price
	product.Images = &productReques.Images

	err := databases.DB.Table("products").Create(&product).Error
	if err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "internal server error",
		})
		return
	}

	ctx.JSON(201, gin.H{
		"error":   false,
		"message": "create data product success",
		"datas":   product,
	})
}

func UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	product := new(models.Products)
	productReques := new(models.CreateProduct)

	if errReques := ctx.ShouldBind(&productReques); errReques != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": errReques.Error(),
		})
		return
	}

	errIdDb := databases.DB.Table("products").Where("id = ?", id).Find(&product).Error
	if errIdDb != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "internal server error",
		})
		return
	}

	if product.Id == nil {
		ctx.JSON(404, gin.H{
			"error":   true,
			"message": "data not Found",
		})
		return
	}

	product.NameProduct = &productReques.NameProduct
	product.Stock = &productReques.Stock
	product.Price = &productReques.Price
	product.Images = &productReques.Images

	errUpdate := databases.DB.Table("products").Where("id = ?", id).Updates(&product).Error
	if errUpdate != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": "update data gagal",
		})
		return
	}

	ctx.JSON(201, gin.H{
		"error":   false,
		"message": "update data success",
		"datas":   product,
	})
}

func DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	product := new(models.Products)
	errDb := databases.DB.Table("products").Where("id = ?", id).Find(&product).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "internal server error",
		})
	}
	if product.Id == nil {
		ctx.JSON(404, gin.H{
			"error":   true,
			"message": "delete data by id not FOUND",
		})
		return
	}

	errorDelete := databases.DB.Table("products").Where("id = ?", id).Delete(&product).Error
	if errorDelete != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": errorDelete.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "delete data by id success",
	})
}
