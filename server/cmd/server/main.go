package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type balancer struct {
	id   int
	name string
}

func main() {
	db, err := sql.Open("mysql", "mysql:mysql@/balancer_db")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from balancer_db.Balancer")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	balancers := []balancer{}

	for rows.Next() {
		p := balancer{}
		err := rows.Scan(&p.id, &p.name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		balancers = append(balancers, p)
	}
	for _, p := range balancers {
		fmt.Println(p.id, p.name)
	}
}
