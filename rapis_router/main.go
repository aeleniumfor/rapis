package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/rapis/rapis_router/router"
)

func handler(router *router.Routing) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		//testproject.user.localhost
		fmt.Println(r.Host)
		router.SetVRoting(r.Host)
		router.Routing()

		virtul, target := router.GetRouting()

		log.Println("[request]: " + virtul)
		log.Println("[target]: " + target)
		target = "http://" + target
		fmt.Println(target)
		remote, _ := url.Parse(target)
		w.Header().Set("X-Forwarded-For", r.Host)
		p := httputil.NewSingleHostReverseProxy(remote)
		p.ServeHTTP(w, r)
	}
}

func main() {
	r := router.Init()
	http.HandleFunc("/", handler(r))
	http.ListenAndServe(":8090", nil)
}
