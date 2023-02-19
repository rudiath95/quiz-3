package repository

import (
	"database/sql"
	"quiz3/structs"
	"time"
)

func GetAllBooks(db *sql.DB) (err error, results []structs.Books) {
	sql := "SELECT * FROM books"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var books = structs.Books{}

		err = rows.Scan(&books.ID, &books.Title, &books.Created_at, &books.Updated_at)
		if err != nil {
			panic(err)
		}

		results = append(results, books)
	}
	return
}

func InsertBooks(db *sql.DB, books structs.Books) (err error) {
	sql := "INSERT INTO books (id, title, description, image_url, price, total_page, thickness, release_year, updated_at, category_id) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)"

	if books.TotalPage < 100 {
		books.Thickness = "tipis"
	} else if books.TotalPage < 200 {
		books.Thickness = "sedang"
	} else {
		books.Thickness = "tebal"
	}

	if books.ReleaseYear < 1980 {
		panic("ReleaseYear tidak boleh kurang dari 1980")
	} else if books.ReleaseYear > 2021 {
		panic("ReleaseYear tidak boleh lebih dari 2021")
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

	books.Updated_at = time.Now()
	errs := db.QueryRow(sql, books.Title, books.Description, books.Image, books.Price, books.TotalPage, books.Thickness, books.ReleaseYear, books.Updated_at, books.CategoryID, books.ID)

	return errs.Err()
}

func DeletedBooks(db *sql.DB, books structs.Books) (err error) {
	sql := "DELETE FROM books WHERE id = $1"

	errs := db.QueryRow(sql, books.ID)

	return errs.Err()
}
