package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dileepkushwaha/sre-bootcamp/internal/model"
	"github.com/dileepkushwaha/sre-bootcamp/internal/service"
	"github.com/gorilla/mux"
)

type StudentController struct {
	studentService service.StudentService
}

func NewStudentController(studentService service.StudentService) *StudentController {
	return &StudentController{studentService: studentService}
}

func (c *StudentController) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student model.Student
	json.NewDecoder(r.Body).Decode(&student)
	err := c.studentService.CreateStudent(&student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *StudentController) GetAllStudents(w http.ResponseWriter, r *http.Request) {
	students, err := c.studentService.GetAllStudents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(students)
}

func (c *StudentController) GetStudentByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	student, err := c.studentService.GetStudentByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(student)
}

func (c *StudentController) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var student model.Student
	json.NewDecoder(r.Body).Decode(&student)
	student.ID = uint(id)
	err := c.studentService.UpdateStudent(&student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *StudentController) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	err := c.studentService.DeleteStudent(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (c *StudentController) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
