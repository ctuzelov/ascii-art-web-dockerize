package main

import (
	s "ascii-art-web/cmd/web"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.Home)
	mux.HandleFunc("/ascii-art", s.AsciiPage)
	mux.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("./ui/style/"))))
	log.Println("Запуск сервера на http://127.0.0.1:7777")
	err := http.ListenAndServe(":7777", mux)
	log.Fatal(err)
}

// vfdbednbfsd bfhsrbje ghnty jmgkmfyg m
