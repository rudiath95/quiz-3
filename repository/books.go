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
	sql := "INSERT INTO books (id, name, created_at) VALUES ($1,$2,$3)"

	books.Updated_at = time.Now()
	errs := db.QueryRow(sql, books.ID, books.Title, books.Updated_at)

	return errs.Err()
}

func UpdatedBooks(db *sql.DB, books structs.Books) (err error) {
	sql := "UPDATE books SET name = $1,created_at = $2 WHERE id = $3"

	books.Updated_at = time.Now()
	errs := db.QueryRow(sql, books.Title, books.Updated_at, books.ID)

	return errs.Err()
}

func DeletedBooks(db *sql.DB, books structs.Books) (err error) {
	sql := "DELETE FROM books WHERE id = $1"

	errs := db.QueryRow(sql, books.ID)

	return errs.Err()
}
