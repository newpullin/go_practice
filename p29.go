package main

import (
	"fmt"
	"net/http"
)

type router struct {
	// 키 : http 메서드
	// 값 : URL 패턴별로 실행할 handle Func
	handlers map[string]map[string]http.HandlerFunc
}
type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func (r *router) HandleFunc(method, pattern string, h http.HandlerFunc) {
	m, ok := r.handlers[method]
	if !ok {
		m = make(map[string]http.HandlerFunc)
		r.handlers[method] = m
	}

	m[pattern] = h
}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "welcome")
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "about")
	})

	http.ListenAndServe(":8000", nil)
}
