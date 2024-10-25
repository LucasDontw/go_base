package services

import (
	"cms/v2/internal/repositories"
	"cms/v2/internal/utils"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type LoginReq struct {
	UserID   string `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRep struct {
	SessionID string `json:"Session_id"`
	UserID    string `json:"user_id"`
	Nickname  string `json:"nickname"`
}

func (c *CmsApp) Login(ctx *gin.Context) {
	var req LoginReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var (
		userID   = req.UserID
		password = req.Password
	)

	//取得account
	accountRepo := repositories.NewAccountRepo(c.db)
	account, err := accountRepo.FirstByUserID(userID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "請輸入正確帳號ID"})
		return
	}

	//比對密碼
	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "請輸入正確密碼"})
		return
	}

	//產生session
	sessionID, err := c.generateSessionID(context.Background(), userID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "系統錯誤"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data": &LoginRep{
			SessionID: sessionID,
			UserID:    account.UserID,
			Nickname:  account.Nickname,
		},
	})
}

func (c *CmsApp) generateSessionID(ctx context.Context, userID string) (string, error) {
	sessionID := uuid.New().String()
	sessionKey := utils.GetSessionKey(userID)

	err := c.rdb.Set(ctx, sessionKey, sessionID, time.Hour*8).Err()

	if err != nil {
		fmt.Printf("rdb set error = %v \n", err)
		return "", err
	}

	// 第二個session主要是用來記錄用戶登入時間，後續像是帳號只可以登入一個等其他功能會用到
	authKey := utils.GetAuthKey(sessionID)
	err = c.rdb.Set(ctx, authKey, time.Now().Unix(), time.Hour*8).Err()

	if err != nil {
		fmt.Printf("rdb set error = %v \n", err)
		return "", err
	}

	return sessionID, nil
}
