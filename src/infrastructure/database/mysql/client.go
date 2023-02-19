package mysql

import (
    "database/sql"
    "fmt"
    "log"
)

type Client struct {
    db *sql.DB
}

func NewClient(dataSourceName string) *Client {
    db, err := sql.Open("mysql", dataSourceName)
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
