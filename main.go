package main

import (
	"database/sql"
	"fmt"
	"net/http"
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

func endApp() (w http.ResponseWriter, r *http.Request) {
	fmt.Println("End App")
	message := recover()
	fmt.Println(w, "Terjadi Error", message)

	return
}

func main() {
	defer endApp()
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
	router.Use(gin.Recovery())
	router.GET("/categories", controllers.GetAllCategory)
	router.GET("/categories/books/:id ", controllers.GetBookFromCategory)
	authorized.POST("/categories", controllers.InsertCategory)
	authorized.PUT("/categories/:id", controllers.UpdatedCategory)
	authorized.DELETE("/categories/:id", controllers.DeletedCategory)

	router.GET("/books", controllers.GetAllBooks)
	authorized.POST("/books", controllers.InsertBooks)
	authorized.PUT("/books/:id", controllers.UpdatedBook)
	authorized.DELETE("/books/:id", controllers.DeletedBook)

	router.GET("/segitiga-sama-sisi", controllers.Segitiga)
	router.GET("/persegi", controllers.Persegi)
	router.GET("/persegi-panjang", controllers.PersegiPanjang)
	router.GET("/lingkaran", controllers.Lingkaran)

	router.Run(":" + os.Getenv("PORT"))

}
