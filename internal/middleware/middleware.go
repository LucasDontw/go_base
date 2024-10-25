package middleware

import (
	"cms/v2/internal/utils"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

const SessionKey = "session_id"

type SessionAuth struct {
	rdb *redis.Client
}

func NewSessionAuth() *SessionAuth {
	s := &SessionAuth{}
	connRdb(s)

	return s
}

func (s *SessionAuth) Auth(ctx *gin.Context) {
	sessionID := ctx.GetHeader(SessionKey)

	if sessionID == "" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, "session is null")
	}

	authKey := utils.GetAuthKey(sessionID)

	loginTime, err := s.rdb.Get(ctx, authKey).Result()

	if err != nil && err != redis.Nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "session auth error")
	}

	if loginTime == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, "session auth fail")
	}

	ctx.Next()
}

func connRdb(s *SessionAuth) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rdb.Ping(context.Background()).Result()

	if err != nil {
		panic(err)
	}

	s.rdb = rdb
}
