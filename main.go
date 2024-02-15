package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
)

//go:embed db/migrations/*.sql
var migrationsFS embed.FS

func main() {
	// 데이터베이스 연결 설정
	dsn := "myuser:mypassword@tcp(localhost:3306)/goosedb?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}
	defer db.Close()

	// goose에 embed.FS 사용 설정
	goose.SetBaseFS(migrationsFS)

	// 마이그레이션 실행
	if err := goose.Up(db, "db/migrations"); err != nil {
		log.Fatalf("goose up failed: %v", err)
	} else {
		fmt.Println("Migration successful")
	}
}
