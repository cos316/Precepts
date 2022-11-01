package main

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Movies struct {
	MovieId int    // represents the movie id
	Title   string // title of movie
	Year    int    // movie release year
	Genres  string // pipe-separated list of genres}
}

type Links struct {
	MovieId int // represents the movie id
	ImdbId  int // IMDB code
	TmdbId  int // TMDB code
}

func addMovie(db *sql.DB, movie Movies) {
	stmt, err := db.Prepare("INSERT or REPLACE INTO Movies(MovieId,Title, Year, Genres) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(movie.MovieId, movie.Title, movie.Year, movie.Genres)
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

func addLinks(db *sql.DB, links Links) {
	stmt, err := db.Prepare("INSERT or REPLACE INTO links(MovieId, ImdbId, TmdbId) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(links.MovieId, links.ImdbId, links.TmdbId)
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

// Task: Implement addRow which takes an interface as a parameter and
// automatically adds a row to the table based on the structure. The goal is to
// avoid having two separate functions (one for each table/struct)
func addRow(db *sql.DB, q interface{}) {
	// fmt.Println("Not implemented.")

	v := reflect.ValueOf(q)

	table_name := fmt.Sprintf("%v", v.Type().Name())

	field_names := []string{}
	field_values := []any{}
	wildcards := []string{}

	for i := 0; i < v.NumField(); i++ {
		field_names = append(field_names, v.Type().Field(i).Name)
		field_values = append(field_values, v.Field(i).Interface())
		wildcards = append(wildcards, "?")
	}

	field_str := strings.Join(field_names, ",")
	wildcard_str := strings.Join(wildcards, ",")

	prepare_statement := fmt.Sprintf("INSERT or REPLACE INTO %s(%s) VALUES(%s)", table_name, field_str, wildcard_str)

	fmt.Println(prepare_statement)

	stmt, err := db.Prepare(prepare_statement)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(field_values...)
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

	jokerMovie := Movies{MovieId: 193612, Title: "Joker", Year: 2019, Genres: "action|thriller"}
	jokerLinks := Links{MovieId: 193612, ImdbId: 7286456, TmdbId: 475557}

	addMovie(db, jokerMovie)
	addLinks(db, jokerLinks)

	addRow(db, jokerMovie)
	addRow(db, jokerLinks)

}
