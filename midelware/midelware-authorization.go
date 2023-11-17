package midelware

import (
	"belajar-api-goleng/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMidelware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := utils.ExtractTokenFromHeader(ctx)
		if tokenString == "" {
			ctx.JSON(401, gin.H{
				"error":   true,
				"message": "Unauthorized: Token is missing",
			})
			ctx.Abort()
			return
		}

		token, err := utils.VerifyToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(401, gin.H{
				"error":   true,
				"message": "Unauthorized: Invalid token",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func AdminRoleMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := utils.ExtractTokenFromHeader(ctx)
		if tokenString == "" {
			ctx.JSON(401, gin.H{
				"error":   true,
				"message": "Unauthorized: Token is missing",
			})
			ctx.Abort()
			return
		}

		token, err := utils.VerifyToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(401, gin.H{
				"error":   true,
				"message": "Unauthorized: Invalid token",
			})
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.JSON(401, gin.H{
				"error":   true,
				"message": "Unauthorized: Invalid token claims",
			})
			ctx.Abort()
			return
		}

		userRole, ok := claims["role"].(string)
		if !ok || userRole != "admin" {
			ctx.JSON(403, gin.H{
				"error":   true,
				"message": "Forbidden: User does not have admin access",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
