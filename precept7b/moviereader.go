package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Movies struct {
	MovieId int
	Title   string
	Year    int
	Genres  string
}

func getMovie(db *sql.DB, name string) {
	query := "SELECT * FROM Movies WHERE title = ?"
	rows, err := db.Query(query, name)
	if err != nil {
		log.Fatalf("Failed to fetch rows from database for %s - %v", name, err)
	}
	for rows.Next() {
		movie := Movies{}
		rows.Scan(&movie.MovieId, &movie.Title, &movie.Year, &movie.Genres)
		fmt.Println("Got movie", movie)
	}
}

func getMovieGeneric(db *sql.DB, name string, model interface{}) {
	fmt.Println("Not implemented yet!")
}

func main() {
	fmt.Println("Let's grab a movie from the database")

	db, err := sql.Open("sqlite3", "file:MovieLens.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	getMovie(db, "Silver Spoon")
	getMovieGeneric(db, "Silver Spoon", Movies{})
}
