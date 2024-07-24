package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloReq struct {
	Name string `json:"name" binding:"required"`
}

type HelloRep struct {
	Message string `json:"message" binding:"required"`
}

func (c *CmsApp) Hello(ctx *gin.Context) {
	var req HelloReq
	//將傳入參數直接綁定到結構
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	//由於gin.Context管理了請求的整個生命週期，並在處理結束時自動回傳回應，所以你不需要在Hello方法中明確地使用return。在Gin框架中，這種方式非常常見，可以簡化程式碼並提高可讀性。
	ctx.JSON(http.StatusOK, gin.H{ //gin.H是一个快捷方式，等價於map[string]interface{}，
		"code":    0,
		"message": "ok",
		"data": &HelloRep{
			Message: "hello",
		},
	})
}
