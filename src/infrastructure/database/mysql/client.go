package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

type Client struct {
	db *sql.DB
}

func NewClient() *Client {
	dataSource := os.Getenv("DATA_SOURCE")
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	fmt.Println("successfully connected to database")

	return &Client{db}
}

func (c *Client) GetDb() *sql.DB {
	return c.db
}

func (c *Client) Close() {
	err := c.db.Close()
	if err != nil {
		log.Fatalf("failed to close database: %v", err)
	}
}
