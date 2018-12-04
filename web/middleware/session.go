package middleware

import (
	"log"
	"restdoc-api/app"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

// Session 会话中间件
func Session() gin.HandlerFunc {
	redisConf := app.Config.Redis
	store, err := redis.NewStore(10, "tcp", redisConf.Host, redisConf.Pass, []byte("secret"))
	if err != nil {
		log.Fatal(err)
	}

	// 配置会话
	store.Options(sessions.Options{
		Path:   "/", // 必须填写，否则会出现多个cookie的bug，导致会话混乱
		MaxAge: 86400,
	})

	return sessions.Sessions(app.Config.SiteName, store)
}
