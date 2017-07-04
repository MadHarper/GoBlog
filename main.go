package main

import (
	"net/http"
	"fmt"
	"html/template"
	"github.com/MadHarper/goBlog/models"
)

var Posts map[int]*models.Post
var n int

func main(){
	n = 1
	Posts = make(map[int]*models.Post, 0)
	fmt.Println(Posts)

	fmt.Println("Listening on port :3000")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/write", writeHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))


	http.ListenAndServe(":3000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil{
		fmt.Fprint(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", nil)
}

func writeHandler(w http.ResponseWriter, r *http.Request){

	if r.Method == "POST"{
		title := r.FormValue("title")
		content := r.FormValue("content")
		new_post := models.NewPost(n, title, content)
		n++
		Posts[new_post.Id] = new_post
		for _, value := range Posts {
			fmt.Println(value)
		}
	}

	t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
	if err != nil{
		fmt.Fprint(w, err.Error())
	}
	t.ExecuteTemplate(w, "write", nil)
}