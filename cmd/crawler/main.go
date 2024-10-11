package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

const USER_AGENT = "Mozilla/5.0 (compatible; cookrisan2; +https://github.com/eniehack/cookrisan2)"

func main() {
	target_url := flag.String("url", "", "")
	db_url := flag.String("db", "", "")
	flag.Parse()
	if len(*target_url) < 1 || len(*db_url) < 1 {
		os.Exit(1)
	}

	db, err := sql.Open("sqlite3", *db_url)
	if err != nil {
		log.Fatalln(err)
	}

	client := new(http.Client)
	req, err := http.NewRequest("GET", *target_url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", USER_AGENT)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("bad status code returned: %s\n", resp.Status)
	}
	tx, err := db.Begin()
	if err != nil {
		log.Fatalln(err)
	}
	now := time.Now().Format(time.RFC3339)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	title := doc.Find(".heading1__title").First().Text()
	recipe_id := uuid.New().String()
	if _, err := tx.Exec(
		"INSERT INTO recipe (id, name, url, created_at, updated_at) VALUES (?, ?, ?, ?, ?);",
		recipe_id,
		title,
		target_url,
		now,
		now,
	); err != nil {
		log.Fatalln(err)
	}
	doc.Find(".recipe-detail-ingredient__name").Each(func(i int, s *goquery.Selection) {
		ingredient := s.Find(".recipe-detail-ingredient__name__type--nolink, .recipe-detail-ingredient__name__type").Text()
		if ingredient == "" {
			return
		}
		if err := insertIngredient(tx, ingredient, recipe_id); err != nil {
			log.Println(err)
			return
		}
	})
	tx.Commit()
}

func insertIngredient(db *sql.Tx, ingredient string, recipe_id string) error {
	var ingredient_id string
	if err := db.QueryRow("SELECT id FROM crawled_ingredient WHERE name = ?;", ingredient).Scan(&ingredient_id); err != nil && err != sql.ErrNoRows {
		return err
	} else if err == sql.ErrNoRows {
		ingredient_id = uuid.NewString()
		if _, err := db.Exec("INSERT INTO crawled_ingredient (id, name) VALUES (?, ?);", ingredient_id, ingredient); err != nil {
			return err
		}
	}
	if _, err := db.Exec(
		"INSERT INTO recipe_ingredient (id, recipe_id, ingredient_id) VALUES (?, ?, ?);",
		uuid.NewString(),
		recipe_id,
		ingredient_id,
	); err != nil {
		return err
	}
	return nil
}
