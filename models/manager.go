package models

import (
	"log"
	"restdoc-api/app"
	"time"

	"github.com/panjiang/go-utils/cryptos"
)

// Manager 管理员
type Manager struct {
	UID       int64     `gorm:"column:uid;PRIMARY_KEY" json:"uid"` // 关联的用户ID
	Username  string    `gorm:"type:VARCHAR(50);UNIQUE" json:"username"`
	Password  string    `gorm:"type:VARCHAR(50)" json:"-"`
	CreatedAt time.Time // 创建时间
}

// AutoMigrate 自动变更
func (Manager) AutoMigrate() {
	first := !app.DB.HasTable(&Manager{})
	if err := app.DB.AutoMigrate(&Manager{}).Error; err != nil {
		log.Fatal(err)
	}
	if first {
		if err := app.DB.Create(&Manager{
			Username: "admin",
			Password: cryptos.MD5("admin"),
		}).Error; err != nil {
			log.Fatal(err)
		}
	}
}
