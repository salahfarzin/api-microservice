package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/salahfarzin/api-microservice/domain/errors"
	"github.com/salahfarzin/api-microservice/domain/users"
	"github.com/salahfarzin/api-microservice/services"
)

var UsersHandler = usersHandler{}

type usersHandler struct{}

func (controller usersHandler) Create(ctx *gin.Context) {
	var user users.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		httpErr := errors.NewBadRequestError("Invlaid json body")
		response(ctx, httpErr.Code, httpErr)
		return
	}

	createdUser, err := services.UsersService.Create(user)
	if err != nil {
		response(ctx, err.Code, err)
		return
	}

	response(ctx, http.StatusOK, createdUser)
}

func (controller usersHandler) Get(ctx *gin.Context) {
	userId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		httpErr := errors.NewBadRequestError("Invlaid user id")
		response(ctx, httpErr.Code, httpErr)
		return
	}

	user, getErr := services.UsersService.Get(userId)
	if getErr != nil {
		response(ctx, getErr.Code, getErr)
		return
	}

	response(ctx, http.StatusOK, user)
}

func response(ctx *gin.Context, code int, body interface{}) {
	if ctx.GetHeader("Accept") == "application/xml" {
		ctx.XML(code, body)
		return
	}

	ctx.JSON(code, body)
}
