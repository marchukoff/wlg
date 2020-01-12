package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main()  {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/process", processHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles(
		"templates/header.html",
		"templates/index.html",
		"templates/footer.html",
		)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
    t.ExecuteTemplate(w, "index", nil)
}

func processHandler(w http.ResponseWriter, r *http.Request)  {
	login := r.FormValue("login")
	password := r.FormValue("password")
	fmt.Printf("%s:%s", login, password)
	http.Redirect(w, r, "/", http.StatusResetContent)
}