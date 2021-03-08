package utils

import (
	"BuildDBGo/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseErrorGin(ctx *gin.Context, message string) {
	response := dtos.Response{
		Meta: dtos.Meta{
			Code:    http.StatusBadRequest,
			Message: message,
		},
	}

	ctx.JSON(http.StatusBadRequest, response)
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	response := dtos.Response{
		Meta: dtos.Meta{
			Code:    http.StatusOK,
			Message: "success",
		},
		Data: data,
	}

	ctx.JSON(http.StatusOK, response)
}
