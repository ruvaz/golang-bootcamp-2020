/**
Student controller
*/
package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"golang-bootcamp-2020/domain/model"
)

// Usecase interface
type Usecase interface {
	ReadStudentsService() ([]model.Student, error)
	StoreURLService() ([]model.Student, error)
}

// Students Use case struct
type Students struct {
	students Usecase
}

// NewController student
func NewController(u Usecase) *Students {
	return &Students{students: u}
}

// ReadStudentsHandler 	Handler for: /readcsv
func (s *Students) ReadStudentsHandler(w http.ResponseWriter, r *http.Request) {
	students, err := s.students.ReadStudentsService()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		sJSON, err := json.Marshal(students)
		if err != nil {
			log.Fatal("Cannot encode to JSON ", err)
		}
		respondWithJSON(
			w,
			http.StatusOK,
			map[string]string{
				"ok":       "true",
				"students": string(sJSON),
			},
		)
	}
}

// StoreStudentURLHandler	 Handler for: /storedata
func (s *Students) StoreStudentURLHandler(w http.ResponseWriter, r *http.Request) {
	students, err := s.students.StoreURLService()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		total := strconv.Itoa(len(students))
		respondWithJSON(
			w,
			http.StatusOK,
			map[string]string{
				"ok":    "true",
				"msj":   "Csv created",
				"total": total,
			},
		)
	}
}

// respondWithError response with error code and message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"msj": msg, "ok": "false"})
}

// respondWithJSON  respond message in JSON
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}
