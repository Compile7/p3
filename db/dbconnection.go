package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var once sync.Once
var dbInstance *gorm.DB
var db *sql.DB

func GetInstance() *gorm.DB {
	once.Do(func() {
		dbInstance, _ = gorm.Open(postgres.Open("postgresql://ankit:ZF4xbddmIAoPI26mqyeRAQ@people-performance-platform-3088.7s5.cockroachlabs.cloud:26257/p3?sslmode=verify-full"), &gorm.Config{})

	})
	return dbInstance
}

func GetJetDB() *sql.DB {
	once.Do(func() {

	})
	//todo: why not working in once do
	db, _ = sql.Open("postgres", "postgres://ankit:ZF4xbddmIAoPI26mqyeRAQ@people-performance-platform-3088.7s5.cockroachlabs.cloud:26257/p3?sslmode=verify-full")
	return db
}
