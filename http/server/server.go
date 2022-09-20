package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type HttpServer struct {
	Addr      string
	router    map[string]http.HandlerFunc
	method    map[string]int
	reqMethod map[string]int
}

func NewHttpServer(addr string) *HttpServer {
	return &HttpServer{
		Addr:   addr,
		router: make(map[string]http.HandlerFunc),
		method: make(map[string]int),
		reqMethod: map[string]int{
			http.MethodGet:    1 << 0,
			http.MethodPost:   1 << 1,
			http.MethodPut:    1 << 2,
			http.MethodDelete: 1 << 3,
		},
	}
}

func (s *HttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	path := r.URL.Path
	if code, ok := s.method[path]; ok {
		method := r.Method
		if s.reqMethod[method]&code == 0 {
			http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
			end := time.Now()
			log.Println("[YuanCheng] | 405 | ", s.Addr, " | ", end.Sub(start).String(), " | ", method, path)
			return
		}
		route := fmt.Sprint(method, ":", path)
		if handler, ok := s.router[route]; ok {
			w.Header().Set("Content-Type", "application/json")
			handler.ServeHTTP(w, r)
			end := time.Now()
			log.Println("[YuanCheng] | 200 | ", s.Addr, " | ", end.Sub(start).String(), " | ", method, path)
		}
	} else {
		http.NotFound(w, r)
		end := time.Now()
		log.Println("[YuanCheng] | 404 | ", s.Addr, " | ", end.Sub(start).String(), " | ", path)
	}
}

func (s *HttpServer) addRoute(method string, route string, f http.HandlerFunc) {
	s.method[route] = s.method[route] | s.reqMethod[method]
	sprint := fmt.Sprint(method, ":", route)
	if _, ok := s.router[sprint]; ok {
		log.Println(method, " Method : '", route, "' is exist")
		return
	}
	s.router[sprint] = f
}

func (s *HttpServer) GET(route string, f http.HandlerFunc) {
	s.addRoute("GET", route, f)
}

func (s *HttpServer) POST(route string, f http.HandlerFunc) {
	s.addRoute("POST", route, f)
}

func (s *HttpServer) PUT(route string, f http.HandlerFunc) {
	s.addRoute("PUT", route, f)
}

func (s *HttpServer) DELETE(route string, f http.HandlerFunc) {
	s.addRoute("DELETE", route, f)
}

func (s *HttpServer) Run() {
	log.Println("Http Server start at", s.Addr)
	err := http.ListenAndServe(s.Addr, s)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
