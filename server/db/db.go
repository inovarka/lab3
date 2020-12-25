package db

import (
	"database/sql"
)

type Connection struct {
	DbName         string
	User, Password string
}

func (c *Connection) ConnectionURL() string {
	str := c.User + ":" + c.Password + "@/" + c.DbName
	return str
}

func (c *Connection) Open() (*sql.DB, error) {
	return sql.Open("mysql", c.ConnectionURL())
}
