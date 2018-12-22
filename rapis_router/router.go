package main

import (
	"github.com/rapis/rapis_router/sqls"

	"fmt"
)

func main() {
	//cash := cash_redis.Init("localhost:6379", "tcp")
	dbsql := sqls.Init()
	dbsql.Set_v_urls("locahost.com:8080")
	fmt.Println(dbsql.Get_v_urls())
}
