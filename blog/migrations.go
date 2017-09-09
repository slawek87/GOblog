package blog

import "github.com/slawek87/GOblog/storage"


func InitMigrations() {
	db, _ := storage.InitDB()

	db.LogMode(true)
	db.AutoMigrate(&Blog{})
}
