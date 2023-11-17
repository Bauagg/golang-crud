package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func ExtractTokenFromHeader(ctx *gin.Context) string {
	header := ctx.GetHeader("Authorization")
	parts := strings.Split(header, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}

	return parts[1]
}
