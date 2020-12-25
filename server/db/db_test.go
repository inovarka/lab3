package db

import "testing"

func TestDbConnection_ConnectionURL(t *testing.T) {
	conn := &Connection{
		DbName:   "balancer_db",
		User:     "mysql",
		Password: "mysql",
	}
	if conn.ConnectionURL() != "mysql:mysql@/balancer_db" {
		t.Error("Unexpected connection string")
	}
}
