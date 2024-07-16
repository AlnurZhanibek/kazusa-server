package handler

import (
	"encoding/json"
	"kazusa-server/internal/entity"
	"kazusa-server/internal/service"
	"net/http"
)

type CourseHandlerImplementation interface {
	Create(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
}

type CourseHandler struct {
	service service.CourseServiceImplementation
}

func NewCourseHandler(service service.CourseServiceImplementation) *CourseHandler {
	return &CourseHandler{service: service}
}

// Create course
//
//	@Summary		Create course
//	@Description	create course
//	@ID				course.create
//	@Accept			json
//	@Produce		json
//	@Param			request		body		entity.NewCourse	true "new course body"
//	@Success		200			{boolean} boolean ok
//	@Failure		400			{boolean} boolean ok
//	@Router			/course [post]
func (h *CourseHandler) Create(w http.ResponseWriter, r *http.Request) {
	newCourse := entity.NewCourse{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newCourse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}

	if newCourse.Title == "" || newCourse.Description == "" || newCourse.Price == 0 {
		http.Error(w, "title, description or price is empty!", http.StatusUnprocessableEntity)
		return
	}

	ok, err := h.service.Create(newCourse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(ok)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Read course
//
//	@Summary		Read courses
//	@Description	read courses
//	@ID				course.read
//	@Accept			json
//	@Produce		json
//	@Param			request		body		entity.CourseReadRequest	true "course read request"
//	@Success		200			{array}	entity.Course
//	@Failure		404			{boolean} boolean ok
//	@Router			/course [get]
func (h *CourseHandler) Read(w http.ResponseWriter, r *http.Request) {
	pagination := entity.Pagination{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&pagination)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}

	courses, err := h.service.Read(pagination)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(courses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
