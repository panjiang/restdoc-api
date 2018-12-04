package tasks

import (
	"restdoc-api/app"
	"restdoc-api/models"

	"github.com/jinzhu/gorm"
	log "github.com/panjiang/golog"
)

func must(db *gorm.DB) {
	if db.Error != nil {
		log.Fatal("migrate db", db.Error)
	}
}

// DatabaseMigrate 数据库结构更新
func DatabaseMigrate() {
	must(app.DB.AutoMigrate(
		new(models.Project),
	))

	// 自定义实现部分
	new(models.Manager).AutoMigrate()
}
