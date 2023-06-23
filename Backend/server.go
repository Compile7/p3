package main

import (
	"log"
	"p3/handler"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Employee struct {
	ID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Email string    `json:"email"`
	Name  string    `json:"name"`
}

func main() {
	e := echo.New()
	db, err := gorm.Open(postgres.Open("postgresql://ankit:ZF4xbddmIAoPI26mqyeRAQ@people-performance-platform-3088.7s5.cockroachlabs.cloud:26257/p3?sslmode=verify-full"), &gorm.Config{})
	db.AutoMigrate(&Employee{})
	if err != nil {
		log.Fatal(err)
	}
	h := handler.Handler{
		DB: db,
	}
	employee := e.Group("/employee")
	employee.POST("/", h.CreateEmployee)
	employee.GET("/:email", h.GetEmployee)

	e.Logger.Fatal(e.Start(":1234"))
}
