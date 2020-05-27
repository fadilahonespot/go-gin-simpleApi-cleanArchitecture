package utils

import (
	"simpleApic/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleSucces(c *gin.Context, data interface{}) {
	responData := model.Respon{
		Status: "200",
		Message: "Success",
		Data: data,
	}
	c.JSON(http.StatusOK, responData)
}

func HandleError(c *gin.Context, status int, message string) {
	responData := model.Respon{
		Status: strconv.Itoa(status),
		Message: message,
	}
	c.JSON(status, responData)
}