package services

import (
	"cms/v2/internal/model"
	"cms/v2/internal/repositories"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterReq struct {
	UserID   string `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
}

type RegisterRep struct {
	Message string `json:"message" binding:"required"`
}

func (c *CmsApp) Register(ctx *gin.Context) {
	var req RegisterReq
	//將傳入參數直接綁定到結構
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//密碼加密
	hashPassword, err := encryptPassword(req.Password)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{})
	}

	//帳號驗證
	accountRepo := repositories.NewAccountRepo(c.db)
	isExist, err := accountRepo.IsExist(req.UserID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	if isExist {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "帳號已存在"})

		return
	}

	//寫入DB
	now := time.Now()
	if err := accountRepo.Create(model.Account{
		UserID:     req.UserID,
		Password:   hashPassword,
		Nickname:   req.Nickname,
		Created_at: now,
		Updated_at: now,
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "創建失敗"})

		return
	}

	//由於gin.Context管理了請求的整個生命週期，並在處理結束時自動回傳回應，所以你不需要在Hello方法中明確地使用return。在Gin框架中，這種方式非常常見，可以簡化程式碼並提高可讀性。
	ctx.JSON(http.StatusOK, gin.H{ //gin.H是一个快捷方式，等價於map[string]interface{}，
		"code":    0,
		"message": "ok",
		"data": &RegisterRep{
			Message: "註冊成功",
		},
	})
}

func encryptPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashPassword), err
}
