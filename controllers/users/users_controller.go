package users

import (
	"github.com/bhanupratap1810/bookstore_users-api/controllers"
	"github.com/bhanupratap1810/bookstore_users-api/controllers/middlewares"
	"github.com/bhanupratap1810/bookstore_users-api/domain/users"
	"github.com/bhanupratap1810/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

//func getUserId(userIdParam string) (int64, *errors.RestErr) {
//	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
//	if userErr != nil {
//		return 0, errors.NewBadRequestError("user id should be a number")
//	}
//	return userId, nil
//}

func CreateUserHandler(service controllers.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user users.User
		if err := c.ShouldBindJSON(&user); err != nil {
			restErr := errors.NewBadRequestError("invalid json body")
			c.JSON(restErr.Status, restErr)
			return
		}
		u, err := service.UserServiceImpl.CreateUser(user)
		if err != nil {
			c.JSON(err.Status, err)
			return
		}

		//result, saveErr := services.CreateUser(user)
		//if saveErr != nil {
		//	c.JSON(saveErr.Status, saveErr)
		//	return
		//}
		c.JSON(http.StatusCreated, u.Marshall(false))
	}
}

func GetUserHandler(service controllers.Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		userid, role, err1 := middlewares.GetUserIdAndRoleFromContext(c)
		if err1 != nil {
			restErr:=errors.NewUnauthorizedError("unauthenticated")
			c.JSON(restErr.Status, restErr)
			return
		}
		//fmt.Println("========>", userid, role, err1 )

		userId, idErr := middlewares.GetId(c.Param("user_id"))
		if idErr != nil {
			c.JSON(idErr.Status, idErr)
			return
		}

		if userid==0 || role== "unemployed"{
			restErr:=errors.NewBadRequestError("not permitted")
			c.JSON(restErr.Status, restErr)
			return
		}

		if role!="admin" && userid!=userId{
			restErr:=errors.NewForbiddenError("only admin or the same user can get user details")
			c.JSON(restErr.Status, restErr)
			return
		}

		user, getErr := service.UserServiceImpl.GetUser(userId)
		if getErr != nil {
			c.JSON(getErr.Status, getErr)
			return
		}
		isPublic:= userid!=userId
		c.JSON(http.StatusOK, user.Marshall(isPublic))
	}
}

func UpdateUserHandler(service controllers.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		userid, role, err1 := middlewares.GetUserIdAndRoleFromContext(c)
		if err1 != nil {
			restErr:=errors.NewUnauthorizedError("unauthenticated")
			c.JSON(restErr.Status, restErr)
			return
		}
		//fmt.Println("========>", userid, role, err1)

		if userid==0 || role== "unemployed"{
			restErr:=errors.NewBadRequestError("not permitted")
			c.JSON(restErr.Status, restErr)
			return
		}

		userId, idErr := middlewares.GetId(c.Param("user_id"))
		if idErr != nil {
			c.JSON(idErr.Status, idErr)
			return
		}
		var user users.User
		if err := c.ShouldBindJSON(&user); err != nil {
			restErr := errors.NewBadRequestError("invalid json body")
			c.JSON(restErr.Status, restErr)
			return
		}

		if role!="admin" && userid!=userId{
			restErr:=errors.NewForbiddenError("only admin or the same user can update the user details")
			c.JSON(restErr.Status, restErr)
			return
		}

		user.Id = userId

		isPartial := c.Request.Method == http.MethodPatch

		result, err := service.UserServiceImpl.UpdateUser(isPartial, user)
		if err != nil {
			c.JSON(err.Status, err)
			return
		}
		c.JSON(http.StatusOK, result.Marshall(false))
	}
}

func DeleteUserHandler(service controllers.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		userid, role, err1 := middlewares.GetUserIdAndRoleFromContext(c)
		if err1 != nil {
			restErr:=errors.NewUnauthorizedError("unauthenticated")
			c.JSON(restErr.Status, restErr)
			return
		}

		if userid==0 || role== "unemployed"{
			restErr:=errors.NewBadRequestError("not permitted")
			c.JSON(restErr.Status, restErr)
			return
		}

		//fmt.Println("========>", userid, role, err1)
		userId, idErr := middlewares.GetId(c.Param("user_id"))
		if idErr != nil {
			c.JSON(idErr.Status, idErr)
			return
		}

		if role!="admin" && userid!=userId{
			restErr:=errors.NewForbiddenError("only admin or the same user can delete the user details")
			c.JSON(restErr.Status, restErr)
			return
		}

		if err := service.UserServiceImpl.DeleteUser(userId); err != nil {
			c.JSON(err.Status, err)
			return
		}
		c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
	}
}

func SearchUserHandler(service controllers.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, roleByToken, err1 := middlewares.GetUserIdAndRoleFromContext(c)
		if err1 != nil {
			restErr:=errors.NewUnauthorizedError("unauthenticated")
			c.JSON(restErr.Status, restErr)
			return
		}
		//fmt.Println("========>", userid, roleByToken, err1)

		if roleByToken!="admin" {
			restErr:=errors.NewForbiddenError("only admin can search for the user details")
			c.JSON(restErr.Status, restErr)
			return
		}

		role := c.Query("Role")

		Users, err := service.UserServiceImpl.SearchUser(role)
		if err != nil {
			c.JSON(err.Status, err)
			return
		}
		//c.JSON(http.StatusOK, Users.Marshall(c.GetHeader("X-Public") == "true"))
		c.JSON(http.StatusOK, Users.Marshall(true))
	}
}

func LoginUserHandler(service controllers.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request users.LoginRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			restErr := errors.NewBadRequestError("invalid json body")
			c.JSON(restErr.Status, restErr)
			return
		}
		token, err := service.UserServiceImpl.LoginUser(request)
		if err != nil {
			c.JSON(err.Status, err)
			return
		}
		c.JSON(http.StatusOK, token)
	}
}
