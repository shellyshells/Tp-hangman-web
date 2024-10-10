package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Student struct {
	Name   string
	Age    int
	Gender string
}

type UserData struct {
	Name      string
	FirstName string
	BirthDate time.Time
	Gender    string
}

var (
	students     []Student
	viewCount    int
	viewCountMux sync.Mutex
	templates    = template.Must(template.ParseGlob("templates/*.html"))
)

func init() {
	students = []Student{
		{Name: "Alice", Age: 22, Gender: "female"},
		{Name: "Bob", Age: 25, Gender: "male"},
		{Name: "Charlie", Age: 23, Gender: "male"},
		{Name: "Diana", Age: 24, Gender: "female"},
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/promo", promoHandler)
	http.HandleFunc("/change", changeHandler)
	http.HandleFunc("/user/form", userFormHandler)
	http.HandleFunc("/user/treatment", userTreatmentHandler)
	http.HandleFunc("/user/display", userDisplayHandler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "home.html", nil)
}

func promoHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Students []Student
	}{
		Students: students,
	}
	templates.ExecuteTemplate(w, "promo.html", data)
}

func changeHandler(w http.ResponseWriter, r *http.Request) {
	viewCountMux.Lock()
	viewCount++
	currentCount := viewCount
	viewCountMux.Unlock()

	var message string
	if currentCount%2 == 0 {
		message = "This page has been viewed an even number of times."
	} else {
		message = "This page has been viewed an odd number of times."
	}

	data := struct {
		Message string
		Count   int
	}{
		Message: message,
		Count:   currentCount,
	}
	templates.ExecuteTemplate(w, "change.html", data)
}

func userFormHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "user_form.html", nil)
}

func userTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	firstName := r.FormValue("firstName")
	birthDateStr := r.FormValue("birthDate")
	gender := r.FormValue("gender")

	if name == "" || firstName == "" || birthDateStr == "" || gender == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	birthDate, err := time.Parse("2006-01-02", birthDateStr)
	if err != nil {
		http.Error(w, "Invalid birth date format", http.StatusBadRequest)
		return
	}

	_ = birthDate

	if gender != "male" && gender != "female" && gender != "other" {
		http.Error(w, "Invalid gender", http.StatusBadRequest)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "userData",
		Value: fmt.Sprintf("%s|%s|%s|%s", name, firstName, birthDateStr, gender),
		Path:  "/",
	})

	http.Redirect(w, r, "/user/display", http.StatusSeeOther)
}

func userDisplayHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("userData")
	if err != nil {
		http.Error(w, "User data not found", http.StatusBadRequest)
		return
	}

	parts := strings.Split(cookie.Value, "|")
	if len(parts) != 4 {
		http.Error(w, "Invalid user data", http.StatusBadRequest)
		return
	}

	birthDate, _ := time.Parse("2006-01-02", parts[2])
	userData := UserData{
		Name:      parts[0],
		FirstName: parts[1],
		BirthDate: birthDate,
		Gender:    parts[3],
	}

	templates.ExecuteTemplate(w, "user_display.html", userData)
}
