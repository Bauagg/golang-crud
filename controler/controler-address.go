package controler

import (
	"belajar-api-goleng/databases"
	"belajar-api-goleng/models"
	"log"

	"github.com/gin-gonic/gin"
)

func GetAddress(ctx *gin.Context) {
	address := new([]models.Address)

	err := databases.DB.Table("address").Find(&address).Error
	if err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "get data address success",
		"datas":   address,
	})

}

func CreateAddress(ctx *gin.Context) {
	address := new(models.Address)
	payloadAddress := new(models.PayloadAddress)

	if errDb := ctx.ShouldBind(&payloadAddress); errDb != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": errDb.Error(),
		})
		return
	}

	log.Println(payloadAddress)

	address.Village = &payloadAddress.Village
	address.Subdistrict = &payloadAddress.Subdistrict
	address.City = &payloadAddress.City
	address.Province = &payloadAddress.Province
	address.Country = &payloadAddress.Country

	err := databases.DB.Table("address").Create(&address).Error
	if err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "create data address success",
		"datas":   address,
	})

}

func UpdateAddress(ctx *gin.Context) {
	id := ctx.Param("id")
	payloadAddress := new(models.PayloadAddress)
	address := new(models.Address)

	if errRequesAddress := ctx.ShouldBind(&payloadAddress); errRequesAddress != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": errRequesAddress.Error(),
		})
		return
	}

	errId := databases.DB.Table("address").Where("id = ?", id).Find(&address).Error
	if errId != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "internal server error",
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

	address.Village = &payloadAddress.Village
	address.Subdistrict = &payloadAddress.Subdistrict
	address.City = &payloadAddress.City
	address.Province = &payloadAddress.Province
	address.Country = &payloadAddress.Country

	errDb := databases.DB.Table("address").Where("id = ?", id).Updates(&address).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "update data by id gagal",
		})
	}

	ctx.JSON(201, gin.H{
		"error":   false,
		"message": "update data address success",
		"datas":   address,
	})
}

func DeleteAddress(ctx *gin.Context) {
	id := ctx.Param("id")
	address := new(models.Address)

	errId := databases.DB.Table("address").Where("id = ?", id).Find(&address).Error
	if errId != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "internal server error",
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

	errDb := databases.DB.Table("address").Where("id = ?", id).Delete(&address).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "delete data address by id gagal",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "delete adrres by id success",
	})
}
