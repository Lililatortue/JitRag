package main

import (
	"context"
	"database/sql"
	"errors"
	"log"

	_ "modernc.org/sqlite"
)

// wrapper type of sql.DB
type Database struct {
	conn *sql.DB
}

// 
// 
func NewDatabase(path string)(*Database, error) {
	db, err := sql.Open("sqlite", path)	
	if err != nil {
		log.Fatalf("couldn't initialise db, %v", err)
		return nil, errors.New("couldn't initialise the db")
	}

    	query := `
    		CREATE TABLE IF NOT EXISTS pages (
        		id INTEGER PRIMARY KEY AUTOINCREMENT,
        		checksum BLOB NOT NULL,
        		url TEXT NOT NULL
    		);`

    	if _, err := db.Exec(query); err != nil {
      	return nil, err
    	}

	return &Database{conn: db},nil
}

// adds pages inside db
func (db *Database)InsertPage(ctx context.Context, page *Page) (int64, error) {
	sql   := db.conn
	query := "INSERT INTO pages(checksum, url) VALUES (?, ?)"

	res, err := sql.ExecContext(ctx, query, page.Checksum[:], page.URL)

	if err != nil {
		return 0, err 
	}
	return res.LastInsertId()
}

//fetches from group of id
//if it finds nothing returns an empty array
func (db *Database)FetchPage(ctx context.Context, ids []int64) []Page {
	sql   := db.conn
	query := "SELECT url, checksum FROM pages WHERE id = ?"
	pages:= make([]Page,0,len(ids))

	for _, id :=range ids {
		var p Page
		var tmp []byte

		res := sql.QueryRowContext(ctx, query, id)

		if err := res.Scan(&p.URL, &tmp); err != nil {
			log.Printf("couldn't scan row nb %d", id)
			continue
		}

		copy(p.Checksum[:],tmp)
		pages = append(pages, p)
	}
	return pages
}


