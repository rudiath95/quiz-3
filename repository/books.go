package repository

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"quiz3/structs"

	"regexp"
	"time"
)

func endApp() (w http.ResponseWriter, r *http.Request) {
	fmt.Println("End App")
	message := recover()
	fmt.Println(w, "Terjadi Error", message)

	return
}

func GetAllBooks(db *sql.DB) (err error, results []structs.Books) {
	sql := "SELECT * FROM books"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var books = structs.Books{}

		err = rows.Scan(&books.ID, &books.Title, &books.Description, &books.Image, &books.ReleaseYear, &books.Price, &books.TotalPage, &books.Thickness, &books.Created_at, &books.Updated_at, &books.CategoryID)
		if err != nil {
			panic(err)
		}

		results = append(results, books)
	}
	return
}

func InsertBooks(db *sql.DB, books structs.Books /*, c *gin.Context*/) (err error) {
	defer endApp()
	sql := "INSERT INTO books (id, title, description, image_url, price, total_page, thickness, release_year, updated_at, category_id) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)"

	var count int

	err2 := db.QueryRow("SELECT COUNT(*) FROM books").Scan(&count)
	switch {
	case err2 != nil:
		log.Fatal(err2)
	default:
		count = count + 1
	}

	books.ID = int64(count)

	if books.TotalPage < 100 {
		books.Thickness = "tipis"
	} else if books.TotalPage < 200 {
		books.Thickness = "sedang"
	} else {
		books.Thickness = "tebal"
	}

	var regex, _ = regexp.Compile(`([^\s]+(\.(?i)(jpe?g|png|gif|bmp|webp))$)`)
	if books.ReleaseYear < 1980 && regex.MatchString(books.Image) {
		panic("ReleaseYear tidak boleh kurang dari 1980")
	} else if books.ReleaseYear > 2021 && regex.MatchString(books.Image) {
		panic("ReleaseYear tidak boleh lebih dari 2021")
	} else if books.ReleaseYear < 1980 && !regex.MatchString(books.Image) {
		panic("ReleaseYear tidak boleh kurang dari 1980 dan image_url bukan gambar")
	} else if books.ReleaseYear > 2021 && !regex.MatchString(books.Image) {
		panic("ReleaseYear tidak boleh lebih dari 2021 dan image_url bukan gambar")
	}

	books.Updated_at = time.Now()
	errs := db.QueryRow(sql, books.ID, books.Title, books.Description, books.Image, books.Price, books.TotalPage, books.Thickness, books.ReleaseYear, books.Updated_at, books.CategoryID)

	return errs.Err()
}

func UpdatedBooks(db *sql.DB, books structs.Books) (err error) {
	sql := "UPDATE books SET title = $1, description = $2, image_url = $3, price = $4, total_page = $5, thickness = $6, release_year = $7, updated_at = $8, category_id = $9 WHERE id = $10"

	if books.TotalPage < 100 {
		books.Thickness = "tipis"
	} else if books.TotalPage < 200 {
		books.Thickness = "sedang"
	} else {
		books.Thickness = "tebal"
	}

	var regex, _ = regexp.Compile(`([^\s]+(\.(?i)(jpe?g|png|gif|bmp|webp))$)`)

	if books.ReleaseYear < 1980 && regex.MatchString(books.Image) {
		panic("ReleaseYear tidak boleh kurang dari 1980")
	} else if books.ReleaseYear > 2021 && regex.MatchString(books.Image) {
		panic("ReleaseYear tidak boleh lebih dari 2021")
	} else if books.ReleaseYear < 1980 && !regex.MatchString(books.Image) {
		panic("ReleaseYear tidak boleh kurang dari 1980 dan image_url bukan gambar")
	} else if books.ReleaseYear > 2021 && !regex.MatchString(books.Image) {
		panic("ReleaseYear tidak boleh lebih dari 2021 dan image_url bukan gambar")
	}

	books.Updated_at = time.Now()
	errs := db.QueryRow(sql, books.Title, books.Description, books.Image, books.Price, books.TotalPage, books.Thickness, books.ReleaseYear, books.Updated_at, books.CategoryID, books.ID)

	return errs.Err()
}

func DeletedBooks(db *sql.DB, books structs.Books) (err error) {
	sql := "DELETE FROM books WHERE id = $1"

	errs := db.QueryRow(sql, books.ID)

	return errs.Err()
}
