package main

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
)

//go:embed db/migrations
var migrationsFS embed.FS

func main() {
	// 데이터베이스 연결 설정
	dsn := "myuser:mypassword@tcp(127.0.0.1:3306)/goosedb?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}
	defer db.Close()

	// 'db/migrations' 디렉토리 내의 파일 목록을 출력
	err = printDirContents("db/migrations")
	if err != nil {
		log.Fatalf("Error reading directory contents: %v", err)
	}

	// goose에 embed.FS 사용 설정
	goose.SetDialect("mysql")
	goose.SetBaseFS(migrationsFS)

	// 마이그레이션 실행
	if err := goose.Up(db, "db/migrations"); err != nil {
		log.Fatalf("goose up failed: %v", err)
	} else {
		fmt.Println("Migration successful")
	}
}

// printDirContents 함수는 지정된 경로의 파일 목록을 읽고 출력합니다.
func printDirContents(dirPath string) error {
	// 디렉토리 내의 파일 목록을 읽음
	dirEntries, err := fs.ReadDir(migrationsFS, dirPath)
	if err != nil {
		return err
	}

	// 파일 목록을 순회하며 파일 이름을 출력
	for _, entry := range dirEntries {
		fmt.Println(entry.Name())
	}

	return nil
}
