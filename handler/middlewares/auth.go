package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nguyenvantuan2391996/patient-order-number/base_common/constants"
)

func APIKeyAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		if constants.APIKey != c.Request.Header.Get(constants.ApiKeyHeader) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "You are not authorized to perform the action",
			})
			return
		}

		c.Next()
	}
}
