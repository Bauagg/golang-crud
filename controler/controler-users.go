package controler

import (
	"belajar-api-goleng/databases"
	"belajar-api-goleng/models"
	"belajar-api-goleng/utils"
	"regexp"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	users := new(models.Users)
	usersRequire := new(models.CreateRegisterUser)

	if errReq := ctx.ShouldBind(&usersRequire); errReq != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": errReq.Error(),
		})
		return
	}

	regexEmail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !regexEmail.MatchString(usersRequire.Email) {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": "Invalid email format",
		})
		return
	}

	erroValidateEmail := databases.DB.Table("users").Where("email = ?", usersRequire.Email).Find(&users).Error
	if erroValidateEmail != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	if users.Email != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": "Email has been already used",
		})
		return
	}

	if len(usersRequire.Password) <= 3 || usersRequire.Password == "" {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": "password kurang kuat",
		})
		return
	}

	bcyptHashPassword := utils.HashPassword(usersRequire.Password)

	users.Username = &usersRequire.Username
	users.Email = &usersRequire.Email
	users.Password = &bcyptHashPassword
	users.Role = &usersRequire.Role

	errCreateResgister := databases.DB.Table("users").Create(&users).Error
	if errCreateResgister != errCreateResgister {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "create register gagal",
		})
		return
	}

	ctx.JSON(201, gin.H{
		"error":   false,
		"message": "register success",
		"datas":   users,
	})
}

func Login(ctx *gin.Context) {
	users := new(models.Users)
	userRequest := new(models.CretateLoginUser)
	usersLogin := new(models.UserGetLogin)

	if userReq := ctx.ShouldBind(&userRequest); userReq != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": userReq.Error(),
		})
		return
	}

	errorRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !errorRegex.MatchString(userRequest.Email) {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": "Invalid email format",
		})
		return
	}

	errValidateEmail := databases.DB.Table("users").Where("email = ?", userRequest.Email).Find(&users).Error
	if errValidateEmail != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "internal server error",
		})
	}

	if users.Email == nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": "email and password do not match",
		})
		return
	}

	verifyPassword := utils.VerifyPassword(userRequest.Password, *users.Password)
	if verifyPassword != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": "email and password do not match",
		})
		return
	}

	createToken, err := utils.SignToken(users)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Failed to create token",
		})
		return
	}

	usersLogin.Id = users.Id
	usersLogin.Username = users.Username
	usersLogin.Email = users.Email
	usersLogin.Role = users.Role
	usersLogin.Token = &createToken

	ctx.JSON(201, gin.H{
		"error":   false,
		"message": "Login succes",
		"datas":   usersLogin,
	})
}
