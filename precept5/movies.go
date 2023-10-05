package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func listMoviesGenresByYear(db *sql.DB, year int) {
	// access title and genres fields from Movies table
	var (
		title  string
		genres string
	)

	// form the SQL query - store result in rows
	rows, err := db.Query("select title, genres from Movies where year = ?;", year)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// interate over each row
	for rows.Next() {
		err := rows.Scan(&title, &genres)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(title, genres)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func listMoviesByYear(db *sql.DB, year int) {
	// create a query statement
	stmt, err := db.Prepare("select title from Movies where year = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// execute the query
	rows, err := stmt.Query(year)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var title string
	for rows.Next() {
		err := rows.Scan(&title)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(title)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func listMoviesById(db *sql.DB, id int) {
	var title string

	// fetch a single row
	err := db.QueryRow("select title from Movies where movieId = ?", id).Scan(&title)
	if err != nil {
		log.Fatal(err)
	}

	// print the title of the movie
	fmt.Println(title)
}

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

	listMoviesGenresByYear(db, 1933)

	listMoviesByYear(db, 1995)

	listMoviesById(db, 1)

	addMovie(db, 193611, "Terminator: Dark Fate", 2019, "Action|Sci-Fi|Thriller")
}
