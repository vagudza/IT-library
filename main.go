package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"www/mydatabase"

	"github.com/gorilla/mux"
)

// some commands:
// go run main.go - start programm
// go build - build project !

type User struct {
	Username            string
	Firstname, Lastname string
	Password            string
	Hobbies             []string
}

/*
func home_page(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(page, "Go is super easy!")
	//var someText string = "Go lang is super easy"
	//testUser := User{"tester", "Bob", "Smile", "f37ae12fx98", []string{"Scala", "Ruby", "Python"}}
	var projects []mydatabase.Project = mydatabase.DBGetProjects()
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, projects)
}
*/

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html", "templates/project_card.html")

	if err != nil {
		// TODO: узнать про пакет log
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// загрузка проектов для главной страницы
	var projects []mydatabase.Project = mydatabase.DBGetProjects()

	// загрузка блока "index" на страницу
	t.ExecuteTemplate(w, "index", projects)
}

func show_project(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ID: %v\n", vars["id"])
}

func handleFunc() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", index).Methods("GET")
	rtr.HandleFunc("/projects/{id:[0-9]+}", show_project).Methods("GET")

	http.Handle("/", rtr)
	// при обращении к физическим файлам по ссылке static/* ищем файл в папке проекта
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("start")
	handleFunc()
}
