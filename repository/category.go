package repository

import (
	"database/sql"
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

func InsertCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "INSERT INTO category (id, name, created_at) VALUES ($1,$2,$3)"

	category.Updated_at = time.Now()
	errs := db.QueryRow(sql, category.ID, category.Name, category.Updated_at)

	return errs.Err()
}

func UpdatedCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "UPDATE category SET name = $1,created_at = $2 WHERE id = $3"

	category.Updated_at = time.Now()
	errs := db.QueryRow(sql, category.Name, category.Updated_at, category.ID)

	return errs.Err()
}

func DeletedCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "DELETE FROM category WHERE id = $1"

	errs := db.QueryRow(sql, category.ID)

	return errs.Err()
}
