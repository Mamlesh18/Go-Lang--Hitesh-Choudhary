package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseName string  `json:"course_name"`
	CourseID   string  `json:"course_id"`
	Author     *Author `json:"author"`
}

type Author struct {
	AuthorName string `json:"fullname"`
}

var courses []Course

func main() {
	fmt.Println("API DEVELOPING")
	r := mux.NewRouter()
	courses = append(courses, Course{
		CourseName: "Introduction to Go",
		CourseID:   "1",
		Author:     &Author{AuthorName: "John Doe"},
	})

	courses = append(courses, Course{
		CourseName: "Advanced Go Programming",
		CourseID:   "2",
		Author:     &Author{AuthorName: "Jane Smith"},
	})

	// Route Handlers
	r.HandleFunc("/", serverHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourse).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", AddCourse).Methods("POST")
	r.HandleFunc("/course/{id}", UpdateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", DeleteCourse).Methods("DELETE")

	// Start the server
	http.ListenAndServe(":8080", r)
}

func (c *Course) IsEmpty() bool {
	return c.CourseID == "" && c.CourseName == ""
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hey, I am JSON</h1>"))
}

func getAllCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one Course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found")
}

func AddCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add Course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please add some data")
		return
	}

	rand.Seed(time.Now().UnixNano())
	course.CourseID = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	courseID := params["id"]

	for index, course := range courses {
		if course.CourseID == courseID {
			var updatedCourse Course
			_ = json.NewDecoder(r.Body).Decode(&updatedCourse)

			if updatedCourse.IsEmpty() {
				json.NewEncoder(w).Encode("Please provide updated course data")
				return
			}

			updatedCourse.CourseID = courseID // Maintain the original course ID
			courses[index] = updatedCourse
			json.NewEncoder(w).Encode(updatedCourse)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with the given ID")
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	courseID := params["id"]

	for index, course := range courses {
		if course.CourseID == courseID {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted successfully")
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with the given ID")
}
