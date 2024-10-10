package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nguyenvantuan2391996/patient-order-number/base_common/constants"
)

func APIKeyAuthentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(constants.RequestIDField, uuid.NewString())
		if constants.APIKey != ctx.Request.Header.Get(constants.ApiKeyHeader) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "You are not authorized to perform the action",
			})
			return
		}

		ctx.Next()
	}
}
