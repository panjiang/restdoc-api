package app

import (
	"fmt"

	"github.com/jinzhu/gorm"
	log "github.com/panjiang/golog"
)

// DB global MySQL client instance
var DB *gorm.DB

// MysqlConfig 用于解析mysql配置
type MysqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"pass"`
	DB       string `json:"db"`
}

func newMysqlClient(host string, user string, pwd string, dbname string) (*gorm.DB, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local&readTimeout=10s&timeout=10s", user, pwd, host, dbname)
	log.Info("mysql:", connStr)
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	db.LogMode(!Config.Release)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetMaxIdleConns(100)
	log.Info("mysql: success")
	return db, nil
}

// InitDB init global DB instance with config
func InitDB(conf *MysqlConfig) error {
	var err error
	DB, err = newMysqlClient(conf.Host, conf.User, conf.Password, conf.DB)
	return err
}
