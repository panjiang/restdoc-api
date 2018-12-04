package handlers

import (
	"fmt"
	"net/http"
	"strings"

	validator "gopkg.in/go-playground/validator.v8"

	"github.com/gin-gonic/gin"
)

// BadType 错误类型
type BadType string

// BadRequest相关错误
const (
	BadNone        BadType = ""            // 无错
	BadInvalid     BadType = "Invalid"     // 无效
	BadNotFound    BadType = "NotFound"    // DB未找到
	BadOccupied    BadType = "Occupied"    // DB已占用
	BadExceed      BadType = "Exceed"      // 超出限制
	BadUnmatch     BadType = "Unmatch"     // 不匹配
	BadUnsupported BadType = "Unsupported" // 不支持
	BadNoChange    BadType = "NoChange"    // 无变化
)

var badFormats = map[BadType]string{
	BadInvalid:     "Invalid '%s'",
	BadNotFound:    "The '%s' not found",
	BadOccupied:    "The '%s' is already occupied",
	BadExceed:      "The '%s' exceeds limit",
	BadUnmatch:     "The '%s' do not match",
	BadUnsupported: "The '%s' is unsupported",
}

type badRequestError struct {
	Field   string  `json:"field"`   // 字段名
	Type    BadType `json:"type"`    // 错误类型
	Message string  `json:"message"` // 错误信息
}

// BadRequest 400非法请求
func BadRequest(c *gin.Context, field string, bad BadType, message ...string) {
	var msg string
	if len(message) > 0 {
		msg = message[0]
	} else {
		format, ok := badFormats[bad]
		if ok {
			msg = fmt.Sprintf(format, field)
		} else {
			msg = fmt.Sprintf("%s %s", field, bad)
		}
	}
	c.JSON(400, gin.H{
		"errors": []interface{}{
			badRequestError{
				Field:   field,
				Type:    bad,
				Message: msg,
			},
		},
	})
}

// BadRequestBind 解析GinBind错误为JSON格式
func BadRequestBind(c *gin.Context, err error) {
	errs := []interface{}{}
	vErrors, ok := err.(validator.ValidationErrors)
	if ok {
		for _, v := range vErrors {
			field := strings.ToLower(v.Field)
			var message string
			if v.Tag == "min" || v.Tag == "max" {
				message = fmt.Sprintf("The '%s' exceeds '%s' limit", field, v.Tag)
			} else {
				message = fmt.Sprintf("bind '%s' failed on the tag '%s'", field, v.Tag)
			}
			errs = append(errs, &badRequestError{
				Field:   field,
				Type:    BadInvalid,
				Message: message,
			})
		}
	} else {
		errs = append(errs, &badRequestError{
			Field:   "",
			Type:    BadInvalid,
			Message: err.Error(),
		})
	}

	c.JSON(400, gin.H{
		"errors": errs,
	})
}

// Unauthorized 401待验证
func Unauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "unauthorized",
	})
}

// Success 200成功
func Success(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusOK, obj)
}
