package sqls

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DB_Routing struct {
	v_urls string
	t_urls string
	urls   string
}

func (r *DB_Routing) Get_Routing() {
	db, err := sql.Open("postgres", "host=localhost port=5431 user=root password=root dbname=domain sslmode=disable")
	defer db.Close()

	if err != nil {
		fmt.Println(err)
	}

	r.v_urls = "testproject.user.localhost"
	query := fmt.Sprintf("SELECT t_urls FROM hosts WHERE v_urls = '%s'", r.v_urls)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		rows.Scan(&r.t_urls)
	}

	fmt.Println(r.v_urls + " -> " + r.t_urls)
}

func (r *DB_Routing) Set_v_urls(v_urls string) {
	r.v_urls = v_urls
}

func (r DB_Routing) Get_v_urls() string {
	return r.v_urls
}

func Init() *DB_Routing {
	//fmt.Println("Init")
	return &DB_Routing{}
}

// func main() {
// 	db := Init()
// 	db.Get_Routing()
// }
