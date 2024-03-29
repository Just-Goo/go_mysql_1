package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Just-Goo/Go-MySql-1/cmd/config"
	"github.com/Just-Goo/Go-MySql-1/cmd/data"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	config.MyApp.Tpl.ExecuteTemplate(w, "index.html", nil)
}

func InsertHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("======== Insert Handler ==========")

	if r.Method == "GET" {
		config.MyApp.Tpl.ExecuteTemplate(w, "insert.html", nil)
		return
	}
	r.ParseForm()
	firstName := r.FormValue("firstname")
	lastName := r.FormValue("lastname")
	age := r.FormValue("age")

	if firstName == "" || lastName == "" || age == "" {
		config.MyApp.Tpl.ExecuteTemplate(w, "insert.html", "Please insert all required fields")
		return
	}

	age2, err := strconv.Atoi(age)
	if err != nil {
		config.MyApp.Tpl.ExecuteTemplate(w, "insert.html", fmt.Sprintf("%v is an invalid age format", age2))
		return
	}

	// id, err := data.InsertStudentReturnID(firstName, lastName, age2)
	err = data.InsertStudent(firstName, lastName, age2)

	if err != nil {
		fmt.Println(err)
		config.MyApp.Tpl.ExecuteTemplate(w, "insert.html", "Failed to insert student details")
		return
	}

	// config.MyApp.Tpl.ExecuteTemplate(w, "insert.html", fmt.Sprintf("Student with id of %d added successfully", id))
	config.MyApp.Tpl.ExecuteTemplate(w, "insert.html", "Student added successfully")

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("======== Home Handler ==========")

	users, err := data.GetAllStudents()
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/error", http.StatusTemporaryRedirect)
		return
	}

	config.MyApp.Tpl.ExecuteTemplate(w, "home.html", users)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("======== Update Handler ==========")

	r.ParseForm() 
	id := r.FormValue("id")

	user, err := data.GetSingleStudent(id)
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusTemporaryRedirect)
	}  
	
	config.MyApp.Tpl.ExecuteTemplate(w, "update.html", user)

}

func UpdateStudentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("======== Update Student Handler ==========")

	r.ParseForm()
	id := r.FormValue("ID")
	firstName := r.FormValue("firstname")
	lastName := r.FormValue("lastname")
	age := r.FormValue("age")
 
	studentAge, err := strconv.Atoi(age)
	if err != nil {
		config.MyApp.Tpl.ExecuteTemplate(w, "update.html", "Invalid age value")
	}

	err = data.UpdateStudent(firstName, lastName, id, studentAge)
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusTemporaryRedirect)
	}

	config.MyApp.Tpl.ExecuteTemplate(w, "result.html", "Student details updated")

}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("======== Delete Handler ==========")

	r.ParseForm() 
	id := r.FormValue("id")

	err := data.DeleteStudent(id)
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusTemporaryRedirect)
		return
	}

	config.MyApp.Tpl.ExecuteTemplate(w, "result.html", "Student deleted successfully")

}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("There was an error"))
}
