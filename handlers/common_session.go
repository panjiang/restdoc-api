package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	log "github.com/panjiang/golog"
)

// Session label
const (
	labelUID   string = "uid"
	labelAdmin string = "admin"
)

// SetSession sets a value into session
func SetSession(c *gin.Context, name string, value interface{}) error {
	session := sessions.Default(c)
	session.Set(name, value)
	return session.Save()
}

// GetSessionInt64 fetchs an int64 value from session
func GetSessionInt64(c *gin.Context, name string) int64 {
	session := sessions.Default(c)
	v := session.Get(name)
	if v == nil {
		ClearSession(c)
		panic(http.StatusUnauthorized)
	}

	value, ok := v.(int64)
	if !ok {
		ClearSession(c)
		log.Errorf("Parse session %s failed: %+v", name, v)
		panic(http.StatusUnauthorized)
	}
	return value
}

// GetSessionString fetchs a string value from session
func GetSessionString(c *gin.Context, name string) string {
	session := sessions.Default(c)
	v := session.Get(name)
	if v == nil {
		ClearSession(c)
		panic(http.StatusUnauthorized)
	}

	value, ok := v.(string)
	if !ok {
		ClearSession(c)
		log.Errorf("Parse session %s failed: %+v", name, v)
		panic(http.StatusUnauthorized)
	}
	return value
}

// DelSession deletes a value from session
func DelSession(c *gin.Context, name string) {
	session := sessions.Default(c)
	session.Delete(name)
	if err := session.Save(); err != nil {
		log.Error(err)
	}
}

// GetSessionUID fetchs uid
// 1w ERC20
func GetSessionUID(c *gin.Context) int64 {
	return GetSessionInt64(c, labelUID)
}

// TryGetSessionUID fetchs uid if exist
func TryGetSessionUID(c *gin.Context) (int64, bool) {
	session := sessions.Default(c)
	v := session.Get(labelUID)
	if v == nil {
		return 0, false
	}

	value, ok := v.(int64)
	if !ok {
		return 0, false
	}
	return value, true
}

// SetSessionUID stores uid
func SetSessionUID(c *gin.Context, uid int64) error {
	return SetSession(c, labelUID, uid)
}

// GetSessionAdmin 管理员session
func GetSessionAdmin(c *gin.Context) int64 {
	return GetSessionInt64(c, labelAdmin)
}

// SetSessionAdmin 管理员session
func SetSessionAdmin(c *gin.Context, uid int64) error {
	return SetSession(c, labelAdmin, uid)
}

// DelSessionAdmin 管理员session
func DelSessionAdmin(c *gin.Context) {
	DelSession(c, labelAdmin)
}

// ClearSession clear all data stored in current session
func ClearSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}
