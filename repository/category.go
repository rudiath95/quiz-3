package repository

import (
	"database/sql"
	"log"
	"quiz3/structs"
	"time"
)

func GetAllCategory(db *sql.DB) (err error, results []structs.Category) {
	sql := "SELECT * FROM category"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var category = structs.Category{}

		err = rows.Scan(&category.ID, &category.Name, &category.Created_at, &category.Updated_at)
		if err != nil {
			panic(err)
		}

		results = append(results, category)
	}
	return
}

func GetBookFromCategory(db *sql.DB) (err error, results []structs.Books) {
	sql := "SELECT * FROM books WHERE category_id = $1"
	var books = structs.Books{}

	rows, err := db.Query(sql, books.CategoryID)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&books.ID, &books.Title, &books.Created_at, &books.Updated_at)
		if err != nil {
			panic(err)
		}

		results = append(results, books)
	}
	return
}

var CatDatas = []structs.Category{}

func InsertCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "INSERT INTO category (id, name, updated_at) VALUES ($1,$2,$3)"

	var count int

	err2 := db.QueryRow("SELECT COUNT(*) FROM category").Scan(&count)
	switch {
	case err2 != nil:
		log.Fatal(err2)
	default:
		count = count + 1
	}

	category.ID = int64(count)

	category.Updated_at = time.Now()
	errs := db.QueryRow(sql, category.ID, category.Name, category.Updated_at)

	return errs.Err()
}

func UpdatedCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "UPDATE category SET name = $1,updated_at = $2 WHERE id = $3"

	category.Updated_at = time.Now()
	errs := db.QueryRow(sql, category.Name, category.Updated_at, category.ID)

	return errs.Err()
}

func DeletedCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "DELETE FROM category WHERE id = $1"

	errs := db.QueryRow(sql, category.ID)

	return errs.Err()
}
