package web

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

type Result struct {
	Post   bool
	Text   string
	Banner string
	Output string
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles("./ui/templates/base.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func AsciiPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		http.NotFound(w, r)
		return
	}
	var res Result
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error())
		return
	}
	// проверка на ключевые слова
	if r.Form == nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	for k := range r.Form {
		if !(k == "text" || k == "banner") {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	}
	// считывание с формы
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	res.Banner = banner
	res.Text = text
	// проверка на существование файла
	_, err = os.Stat("banners/" + banner + ".txt")
	if err != nil {
		// Используем os.IsNotExist(err) так os.Stat может возвращать и другие виды ошибок ввода/вывода, таймаут и ещё букет ошибок
		if os.IsNotExist(err) {
			fmt.Println("file does not exist") // os.IsNotExist это проверка ошибки об отсутствия файла
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
	res.Output, err = Converter(res)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	tmpl, err := template.ParseFiles("./ui/templates/base.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	res.Post = true
	err = tmpl.Execute(w, res)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
