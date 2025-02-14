package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"go-api-gateway/config"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// โหลด Config
	config.LoadConfig()

	// สร้าง Connection String
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.AppConfig.Database.Host,
		config.AppConfig.Database.Port,
		config.AppConfig.Database.User,
		config.AppConfig.Database.Password,
		config.AppConfig.Database.Name,
	)

	// เชื่อมต่อ Database
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// ตรวจสอบการเชื่อมต่อ
	err = db.Ping()
	if err != nil {
		log.Fatalf("Database is not reachable: %v", err)
	} else {
		log.Println("Connected to database successfully!")
	}

	// ตั้งค่า Gin
	r := gin.Default()

	// Health check route
	r.GET("/api/v1/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"message": "API Gateway is running",
		})
	})

	r.GET("/api/v1/", func(c *gin.Context) {
		content, err := ioutil.ReadFile("./public/html/WelcomeMessage.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "Error reading file")
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", content)	
	})

	r.Run(":8080")
	log.Println("Server is running at http://localhost:8080/api/v1")

}
