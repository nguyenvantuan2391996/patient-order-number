package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/nguyenvantuan2391996/patient-order-number/base_common/constants"
	response "github.com/nguyenvantuan2391996/patient-order-number/base_common/response"
	"github.com/spf13/viper"
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

func JWTValidationMW(role int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(constants.RequestIDField, uuid.NewString())
		resp := response.NewResponse(ctx)

		authHeader := ctx.GetHeader(constants.AuthorizationHeader)
		if authHeader == "" {
			log.Println(constants.ErrAuthorizationHeaderEmpty)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp.ToResponse(http.StatusUnauthorized,
				nil, constants.ErrAuthorizationHeaderEmpty))
			return
		}

		authHdrPart := strings.Split(authHeader, " ")
		switch len(authHdrPart) {
		case 2:
			if authHdrPart[0] != constants.AuthorizationTypeBearer {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp.ToResponse(http.StatusUnauthorized,
					nil, constants.ErrInvalidAuthorizationType))
				return
			}

			claims, err := verifyToken(authHdrPart[1])
			if err != nil {
				log.Printf("Error verifying token: %s", err.Error())
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp.ToResponse(http.StatusUnauthorized,
					nil, constants.ErrInvalidToken))
				return
			}

			userName, ok := claims["user_name"].(string)
			if !ok {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp.ToResponse(http.StatusUnauthorized,
					nil, constants.ErrNotAuthorized))
				return
			}

			ctx.Set("user_name", userName)

			userID, ok := claims["account_id"].(float64)
			if !ok {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp.ToResponse(http.StatusUnauthorized,
					nil, constants.ErrNotAuthorized))
				return
			}

			ctx.Set("account_id", userID)

			// validate role
			roleInToken, ok := claims["role"].(float64)
			if !ok {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp.ToResponse(http.StatusUnauthorized,
					nil, constants.ErrNotAuthorized))
				return
			}

			if int(roleInToken) != role {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp.ToResponse(http.StatusUnauthorized,
					nil, constants.ErrNotAuthorized))
				return
			}
		default:
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp.ToResponse(http.StatusUnauthorized,
				nil, constants.ErrInvalidToken))
			return
		}

		ctx.Next()
	}
}

func verifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("PRIVATE_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("cannot map claims from token")
	}

	return claims, nil
}
