package handlers

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	log "github.com/panjiang/golog"
)

// CheckError 如有异常抛出，外层捕获，并最终返回500
func CheckError(err error) {
	if err != nil && err != gorm.ErrRecordNotFound && err != redis.Nil {
		log.Error(err)
		panic(err)
	}
}

// RecordNotFound 记录不存在
func RecordNotFound(err error) bool {
	if err == nil {
		return false
	}
	if err == gorm.ErrRecordNotFound {
		return true
	}
	panic(err)
}
