package router

import (
	"github.com/rapis/rapis_router/cash_routing"
	"github.com/rapis/rapis_router/dbrouting"
)

type Routing struct {
	v_urls        string
	t_urls        string
	DBResoleve    *dbrouting.DB_Routing
	RedisResoleve *cash_routing.Cach
}

func Init() *Routing {
	db := dbrouting.Init()
	redis := cash_routing.Init("localhost:6379", "tcp")
	redis.Connect()
	return &Routing{DBResoleve: db, RedisResoleve: redis}
}

func (r *Routing) Routing() {
	r.DBResoleve.Set_v_urls("localhost.com")
	s, _ := r.RedisResoleve.GetCash(r.v_urls)
	if s == "" {
		r.DBResoleve.Set_v_urls(r.v_urls)
		r.DBResoleve.Get_Routing()
		target := r.DBResoleve.Get_t_urls()
		r.RedisResoleve.SetCach(r.v_urls, target)
		r.t_urls = target
	} else {
		r.t_urls = s
	}
}

func (r *Routing) SetVRoting(v string) {
	r.v_urls = v
}
func (r Routing) GetRouting() (string, string) {
	return r.v_urls, r.t_urls
}

// func main() {
	//cash := cash_redis.Init("localhost:6379", "tcp")
	// dbsql := dbrouting.Init()
	// redis := cash_routing.Init("localhost:6379", "tcp")
	// redis.Connect()
	// dbsql.Set_v_urls("locahost.com:8080")
	// v_urls := "testproject.user.localhost"
	// s, _ := redis.GetCash(v_urls)
	// if s == "" {
	// 	dbsql.Set_v_urls(v_urls) // DBからURLを引いてくる準備
	// 	fmt.Println(dbsql.Get_v_urls())
	// 	dbsql.Get_Routing() //DBからURLを引く
	// 	target := dbsql.Get_t_urls()
	// 	redis.SetCach(v_urls, target)
	// } else {
	// 	fmt.Println("=")
	// 	fmt.Println(s)
	// }
// 	routing := Init()
// 	routing.SetVRoting("testproject.user.localhost")
// 	routing.Routing()
// 	a, b := routing.GetRouting()
// 	fmt.Println(a)
// 	fmt.Println(b)

// }
