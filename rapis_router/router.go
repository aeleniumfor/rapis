package main

import (
	"github.com/rapis/rapis_router/cash_routing"
	"github.com/rapis/rapis_router/dbrouting"

	"fmt"
)

type Routing struct {
	v_urls string
	t_urls string
}

func Init() {
	dbsql := dbrouting.Init()
	redis := cash_routing.Init("localhost:6379", "tcp")
	redis.Connect()
	dbsql.Set_v_urls("locahost.com:8080")
	fmt.Println(dbsql.Get_v_urls())
}

func main() {
	//cash := cash_redis.Init("localhost:6379", "tcp")
	dbsql := dbrouting.Init()
	redis := cash_routing.Init("localhost:6379", "tcp")
	redis.Connect()
	dbsql.Set_v_urls("locahost.com:8080")
	v_urls := "testproject.user.localhost"
	s, _ := redis.GetCash(v_urls)
	if s == "" {
		dbsql.Set_v_urls(v_urls) // DBからURLを引いてくる準備
		fmt.Println(dbsql.Get_v_urls())
		dbsql.Get_Routing() //DBからURLを引く
		target := dbsql.Get_t_urls()
		redis.SetCach(v_urls, target)
	} else {
		fmt.Println("=")
		fmt.Println(s)
	}
}
