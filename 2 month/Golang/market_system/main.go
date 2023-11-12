package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/asadbekGo/market_system/model"
	uuid "github.com/google/uuid"
	_ "github.com/lib/pq"
)

func main() {

	cfg := Load()

	connect := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	)

	db, err := sql.Open("postgres", connect)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err != nil {
		return
	}

	// Insert(db)
	// GetByID(db)
	// GetList(db)
	// Update(db)
	Delete(db)
}

func Insert(db *sql.DB) {
	var (
		categoryId = uuid.New().String()
		query      = `INSERT INTO "category"("id", "title", "parent_id", "updated_at") VALUES ($1, $2, $3, NOW())`
	)

	// exec function sql commandani ishga tushirish yani bu function value qaytmaydi
	_, err := db.Exec(
		query,
		categoryId,
		"Готовые кухни",
		"a452dcc9-8fb6-4f93-b047-722688fe5234",
	)

	if err != nil {
		log.Println("db exec function:", err.Error())
		return
	}
}
func Update(db *sql.DB) {
	var (
		query = `UPDATE "category" SET "title"=$1, "updated_at"=NOW() WHERE "id" = $2`
	)

	// exec function sql commandani ishga tushirish yani bu function value qaytmaydi
	_, err := db.Exec(
		query,
		"KAMPYUTER",
		"a452dcc9-8fb6-4f93-b047-722688fe5234",
	)

	if err != nil {
		log.Println("db exec function:", err.Error())
		return
	}
}
func Delete(db *sql.DB) {
	var (
		query = `DELETE FROM "category" WHERE id =$1`
	)

	// exec function sql commandani ishga tushirish yani bu function value qaytmaydi
	_, err := db.Exec(
		query,
		"ad0c706f-f067-4173-bf02-dfa0182a1b69",
	)

	if err != nil {
		log.Println("db exec function:", err.Error())
		return
	}
}

func GetByID(db *sql.DB) {
	var (
		category model.Category
		query    = `
			SELECT
				"id",
				"title",
				COALESCE(CAST("parent_id" AS VARCHAR), ''),
				"created_at",
				"updated_at"	
			FROM "category"
			WHERE "id" = $1
		`
	)

	// query row function bitta row olish uchun ishlatiladi
	err := db.QueryRow(query, "689dfd45-8166-402e-a490-d044d843694f").Scan(
		&category.Id,
		&category.Title,
		&category.ParentID,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if err != nil {
		log.Println("db query row function:", err.Error())
		return
	}

	fmt.Printf("Category: %+v\n", category)
}

func GetList(db *sql.DB) {
	var (
		categories []model.Category
		query      = `
			SELECT
				"id",
				"title",
				"parent_id",
				"created_at",
				"updated_at"
			FROM "category"
		`
	)

	// query function hamma listdagi ma'lumotlarni olishlik uchun ishlatiladi
	rows, err := db.Query(query)
	if err != nil {
		log.Println("db query function:", err.Error())
		return
	}

	for rows.Next() {

		var (
			category model.Category
			parentID sql.NullString
		)
		err = rows.Scan(
			&category.Id,
			&category.Title,
			&parentID,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		if err != nil {
			log.Println("error while get list scan err:", err.Error())
			return
		}

		category.ParentID = parentID.String

		categories = append(categories, category)
	}

	for _, category := range categories {
		fmt.Printf("Title: %s\n", category.Title)
	}
}
