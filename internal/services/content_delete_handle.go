package services

import (
	"cms/v2/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContentDeleteReq struct {
	ID int `json:"id" binding:"required"`
}

type ContentDeleteRep struct {
	Message string `json:"message" binding:"required"`
}

func (c *CmsApp) ContentDelete(ctx *gin.Context) {
	var req ContentDeleteReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contentRepo := repositories.NewContentRepo(c.db)
	ok, err := contentRepo.IsExist(req.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "內容不存在"})
		return
	}

	if err := contentRepo.Delete(req.ID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "刪除失敗"})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data": &ContentDeleteRep{
			Message: "刪除成功",
		},
	})
}
