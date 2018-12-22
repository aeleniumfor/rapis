package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Routing struct {
	v_urls string
	t_urls string
	urls   string
}

func (r *Routing) Get_Routing() {
	db, err := sql.Open("postgres", "host=localhost port=5431 user=root password=root dbname=domain sslmode=disable")
	defer db.Close()

	if err != nil {
		fmt.Println(err)
	}

	r.v_urls = "web.user.localhost"
	query := fmt.Sprintf("SELECT r_urls FROM hosts WHERE p_urls = '%s'", r.v_urls)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		rows.Scan(r.t_urls)
	}

	fmt.Println(r.v_urls + " -> " + r.t_urls)
}

func Init() *Routing {
	fmt.Println("Init")
	return &Routing{}
}

func main() {
	db := Init()
	db.Get_Routing()
}
