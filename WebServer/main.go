package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type LoginForm struct{
	Name string;
	Password string
}

func main(){
	http.HandleFunc("/login",loginHandler)
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080",nil)	
}

func loginHandler(w http.ResponseWriter, r *http.Request){
	if r.Method==http.MethodGet{
		tmpl := template.Must(template.ParseFiles("login.html"))
		tmpl.Execute(w,nil)
	}else if r.Method==http.MethodPost{
		r.ParseForm()
		form := LoginForm{
			Name : r.FormValue("name"),
			Password : r.FormValue("password"),
		}


	fmt.Printf("Login Attempt: Username=%s, Password=%s\n", form.Name, form.Password)
	
	if form.Name == "admin" && form.Password == "admin" {
		
		fmt.Fprintf(w, "Login successful for user: %s", form.Name)
	}else{
		err := http.StatusNotFound
		fmt.Fprintf(w, "Login failed for user: %s", form.Name)
		fmt.Fprintf(w,"Error: %d",err)
	}

	}
}