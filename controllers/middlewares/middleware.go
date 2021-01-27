package middlewares

import (
	"context"
	"fmt"
	"github.com/bhanupratap1810/bookstore_users-api/constants"
	"github.com/bhanupratap1810/bookstore_users-api/services"
	"github.com/bhanupratap1810/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"github.com/kataras/jwt"
	"net/http"
	"strconv"
)

//todo verify and search for admin

func VerifyAndServe(authService services.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("Authorization")
		if token==""{
			c.Abort()
			c.JSON(http.StatusUnauthorized, map[string]string{"status": "jwt token required"})
			return
		}
		//todo error handling
		loginJwtClaims, err := authService.ValidateToken(token)
		if err != nil {
			c.Abort()
			if err == jwt.ErrTokenSignature {
				c.JSON(http.StatusBadRequest, jwt.ErrTokenSignature)
				return
			}
			if err == jwt.ErrExpired {
				c.JSON(http.StatusBadRequest, jwt.ErrExpired)
				return
			}
			if err == jwt.ErrInvalidKey {
				c.JSON(http.StatusBadRequest, jwt.ErrInvalidKey)
				return
			}
			if err == jwt.ErrDecrypt {
				c.JSON(http.StatusBadRequest, jwt.ErrDecrypt)
				return
			}
			if err == jwt.ErrMissingKey {
				c.JSON(http.StatusBadRequest, jwt.ErrMissingKey)
				return
			}
			if err == jwt.ErrMissing {
				c.JSON(http.StatusBadRequest, jwt.ErrMissing)
				return
			}
			if err == jwt.ErrTokenSignature {
				c.JSON(http.StatusBadRequest, jwt.ErrTokenSignature)
				return
			}
			c.JSON(http.StatusBadRequest, jwt.ErrNotValidYet)
			return
		}
		ctx := context.WithValue(c.Request.Context(), constants.UserIdKey, loginJwtClaims.UserId)
		c.Request = c.Request.WithContext(ctx)
		c.Set(constants.UserIdKey, loginJwtClaims.UserId)

		ctx = context.WithValue(c.Request.Context(), constants.RoleKey, loginJwtClaims.Role)
		c.Request = c.Request.WithContext(ctx)
		c.Set(constants.RoleKey, loginJwtClaims.Role)
	}
}

//todo error handling and return role also
func GetUserIdAndRoleFromContext(ctx *gin.Context) (int64, string, error) {
	userIdInterface, ok := ctx.Get(constants.UserIdKey)
	if !ok {
		return 0, "", fmt.Errorf("userid not found")
	}

	userRoleInterface, ok := ctx.Get(constants.RoleKey)
	if !ok {
		return 0, "", fmt.Errorf("role not found")
	}

	return userIdInterface.(int64), userRoleInterface.(string), nil
}

func GetId(idParam string) (int64, *errors.RestErr) {
	id, idErr := strconv.ParseInt(idParam, 10, 64)
	if idErr != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
	}
	return id, nil
}
