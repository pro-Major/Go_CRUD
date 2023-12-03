package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Define a User Model
type User struct {
	ID uint `gorm:"primaryKey"`
	Name string 
	Email string 
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}




func main(){	
	 // Connect to a PostgreSQL database
	dbConfig := "user=postgres password=91 dbname=go_db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dbConfig))
	if(err != nil){
		panic(err) //The panic built-in function stops normal execution of the current goroutine. When a function F calls panic, normal execution of F stops immediately.
	}
	fmt.Println(`Port Start on 5432 and Database is Running Successfully !!!`)
	defer db.DB()

	//Auto-migrate the schema 
	db.AutoMigrate(&User{})

	//Create a Gin Router 
	router := gin.Default();

	router.GET("/users", func(c *gin.Context) {
		var users []User

		db.Find((&users))
		c.JSON(200, users)
	})
	
	router.POST("/users", func(c *gin.Context) {
		var user User
		c.BindJSON(&user)
		db.Create(&user)
		c.JSON(200, user)
	})

	router.Run("localhost:8080")
}

