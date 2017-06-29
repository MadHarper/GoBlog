package main

import (
	"net/http"
	"fmt"
	"html/template"
)

func main(){
	http.HandleFunc("/", indexHandler)

	http.ListenAndServe(":3000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/index.html")
	if err != nil{
		fmt.Fprint(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}
