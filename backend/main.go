package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func initDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	var err error
	for i := 1; i <= 20; i++ {
		db, err = sqlx.Connect("mysql", dsn)
		if err == nil {
			log.Println("DB接続成功")
			return
		}
		log.Printf("DB接続待機中... (%d/20): %v", i, err)
		time.Sleep(3 * time.Second)
	}
	log.Fatalf("DB接続エラー: %v", err)
}

func main() {
	initDB()

	r := gin.Default()
	r.Use(cors.Default())

	api := r.Group("/api")
	{
		api.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})

		// 観光プラン
		api.GET("/plans", getPlans)
		api.POST("/plans", createPlan)
		api.GET("/plans/:id", getPlan)
	}

	r.Run(":8080")
}

type Plan struct {
	ID          int    `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Prefecture  string `db:"prefecture" json:"prefecture"`
	Days        int    `db:"days" json:"days"`
	CreatedAt   string `db:"created_at" json:"created_at"`
}

func getPlans(c *gin.Context) {
	var plans []Plan
	err := db.Select(&plans, "SELECT id, title, description, prefecture, days, created_at FROM plans ORDER BY created_at DESC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, plans)
}

func getPlan(c *gin.Context) {
	id := c.Param("id")
	var plan Plan
	err := db.Get(&plan, "SELECT id, title, description, prefecture, days, created_at FROM plans WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "プランが見つかりません"})
		return
	}
	c.JSON(http.StatusOK, plan)
}

func createPlan(c *gin.Context) {
	var plan Plan
	if err := c.ShouldBindJSON(&plan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := db.Exec(
		"INSERT INTO plans (title, description, prefecture, days) VALUES (?, ?, ?, ?)",
		plan.Title, plan.Description, plan.Prefecture, plan.Days,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	id, _ := result.LastInsertId()
	plan.ID = int(id)
	c.JSON(http.StatusCreated, plan)
}
