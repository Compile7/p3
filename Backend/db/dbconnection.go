package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var once sync.Once
var dbInstance *gorm.DB

func GetInstance() *gorm.DB {
	once.Do(func() {
		dbInstance, _ = gorm.Open(postgres.Open("postgresql://ankit:ZF4xbddmIAoPI26mqyeRAQ@people-performance-platform-3088.7s5.cockroachlabs.cloud:26257/p3?sslmode=verify-full"), &gorm.Config{})
	})
	return dbInstance
}
