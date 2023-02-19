package main

import (
	"database/sql"
	"fmt"
	"os"
	"quiz3/controllers"
	"quiz3/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	//ENV Configuration
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file environment")
	} else {
		fmt.Println("success load file environment")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"))

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	//Router GIN
	router := gin.Default()
	//GIN Authentication
	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
		"admin":  "password",
		"editor": "secret",
	}))
	router.GET("/category", controllers.GetAllCategory)
	authorized.POST("/category", controllers.InsertCategory)
	authorized.PUT("/category/:id", controllers.UpdatedCategory)
	authorized.DELETE("/category/:id", controllers.DeletedCategory)

	router.Run(":" + os.Getenv("PORT"))

}
