package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

// Contact は連絡先です
type Contact struct {
	ID     int64
	Name   string
	Number string
}

func main() {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		fmt.Printf("DB接続失敗 %v", err)
		os.Exit(1)
	}
	const createSQL = `
	create table if not exists contact (
		id integer not null primary key autoincrement
		name varchar(255) not null
		number varchar(16)
	);
	`
	if _, err := db.Exec(createSQL); err != nil {
		fmt.Printf("テーブル作成失敗 %v", err)
		os.Exit(1)
	}

	contacts := []*Contact{
		{
			Name:   "Suzuki Hajime",
			Number: "0000-000-000",
		},
		{
			Name:   "Suzuki Futaba",
			Number: "0000-000-001",
		},
	}
	const insertSQL = "insert into contact(name, number) values (?, ?)"
	for i := range contacts {
		r, err := db.Exec(insertSQL, contacts[i].Name, contacts[i].Number)
		if err != nil {
			fmt.Printf("レコード挿入失敗 %v", err)
			os.Exit(1)
		}
		id, err := r.LastInsertId()
		if err != nil {
			fmt.Printf("ID取得失敗 %v", err)
			os.Exit(1)
		}
		contacts[i].ID = id
		fmt.Println("INSERT", contacts[i])
	}

	rows, err := db.Query("select * from contact")
	if err != nil {
		fmt.Printf("レコード取得失敗 %v", err)
		os.Exit(1)
	}
	for rows.Next() {
		var c Contact
		if err := rows.Scan(&c.ID, &c.Name, &c.Number); err != nil {
			fmt.Printf("レコード取得失敗 %v", err)
			os.Exit(1)
		}
		fmt.Println(c)
	}
	if err := rows.Err(); err != nil {
		fmt.Printf("レコード取得失敗 %v", err)
		os.Exit(1)
	}
	fmt.Println("hello")
}
