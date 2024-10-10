package middlewares

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/nguyenvantuan2391996/patient-order-number/base_common/constants"
	"github.com/sirupsen/logrus"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logrus.Error(string(debug.Stack()))
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": constants.SomethingWentWrong,
				})
			}
		}()

		c.Next()
	}
}
