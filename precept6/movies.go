package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func addMovie(db *sql.DB, id int, title string, year int, genres string) {
	stmt, err := db.Prepare("INSERT or REPLACE INTO movies(movieId,title, year, genres) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(id, title, year, genres)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {

		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
}

func main() {
	// open the database
	db, err := sql.Open("sqlite3", "file:MovieLens.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	addMovie(db, 193611, "Terminator: Dark Fate", 2019, "Action|Sci-Fi|Thriller")
}
