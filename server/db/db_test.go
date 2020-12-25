package db

import "testing"

func TestDbConnection_ConnectionURL(t *testing.T) {
	conn := &Connection{
		DbName:   "balancer_bd",
		User:     "mysql",
		Password: "mysql",
	}
	if conn.ConnectionURL() != "mysql:mysql@/balancer_bd" {
		t.Error("Unexpected connection string")
	}
}
