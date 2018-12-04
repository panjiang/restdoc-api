package app

import (
	"github.com/go-redis/redis"
)

// RDS a global Redis client
var RDS redis.UniversalClient

// RedisConfig Redis配置
type RedisConfig struct {
	Host string `json:"host"`
	Pass string `json:"pass"`
}

// Redis Key
const (
	SPasswordCaptcha    = "PasswordCaptcha"    // 玩家密码验证码
	SResetPasswordToken = "ResetPasswordToken" // 重置密码令牌
	SRegisterCaptcha    = "RegisterCaptcha"    // 注册验证码
)

// GetRedisKey 获取key
func GetRedisKey(label string, key string) string {
	return label + "." + key
}

// InitRedisCli create a redis handler
func InitRedisCli(conf *RedisConfig) error {
	RDS = redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    []string{conf.Host},
		Password: conf.Pass,
		DB:       0,
	})

	_, err := RDS.Ping().Result()
	return err
}
