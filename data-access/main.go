package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "data-access",
	}

	var err error

	db, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	albums, err := albumsByArtist("John Coltrane")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums Found: %v\n", albums)
}

func albumsByArtist(name string) ([]Album, error) {
	var albums []Album

	rows, err := db.Query("Select * from album where artist = ?", name)

	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q:%v", name, err)
	}

	// defer is used to ensure that a function calls performed later in a programs's execution.
	// eg. finally in other languages.
	defer rows.Close()

	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q:%v", name, err)
		}
		albums = append(albums, alb)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q:%v", name, err)
	}

	return albums, nil
}
