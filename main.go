package main

import (
	"Hexagonal-Model/database"
	"Hexagonal-Model/handlers"
	"Hexagonal-Model/repositories"
	"Hexagonal-Model/services"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// เปิดการเชื่อมต่อกับฐานข้อมูล
	db := database.Postgresql()
	defer db.Close()

	// ตรวจสอบการเชื่อมต่อกับฐานข้อมูล
	err := db.Ping()
	if err != nil {
		log.Fatal("Database connection error: ", err)
	}

	// สร้าง instances ของ repositories, services, และ handlers
	r := repositories.NewRepositoryAdapter(db)
	s := services.NewServices(r)
	h := handlers.NewHandlers(s)

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "X-Auth-Token", "Authorization"}
	router.Use(cors.New(config))

	router.POST("/api/register", h.PostRegisterHandler)

	err = router.Run(":8062")
	if err != nil {
		log.Fatal(err.Error())
	}
}
