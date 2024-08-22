package handler

import (
	"encoding/json"
	"github.com/AlnurZhanibek/kazusa-server/internal/entity"
	"github.com/AlnurZhanibek/kazusa-server/internal/service"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

type CourseHandlerImplementation interface {
	Create(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
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
		return
	}

	if newCourse.Title == "" || newCourse.Description == "" || newCourse.Price == 0 {
		http.Error(w, "title, description or price is empty!", http.StatusUnprocessableEntity)
		return
	}

	ok, err := h.service.Create(newCourse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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
//	@Param			offset		query		int64	true "offset"
//	@Param			limit		query		int64	true "limit"
//	@Param			id			query		string	false "id"
//	@Success		200			{array}	entity.Course
//	@Failure		404			{boolean} boolean ok
//	@Router			/course [get]
func (h *CourseHandler) Read(w http.ResponseWriter, r *http.Request) {
	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	id := r.URL.Query().Get("id")

	filters := entity.CourseFilters{
		ID: uuid.Nil,
	}

	if id != "" {
		filters.ID = uuid.MustParse(id)
	}

	courses, err := h.service.Read(entity.Pagination{
		Offset: offset,
		Limit:  limit,
	}, filters)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(courses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Update course
//
//	@Summary		Update course
//	@Description	update course
//	@ID				course.update
//	@Accept			json
//	@Produce		json
//	@Param			request		body		entity.CourseUpdateBody	true "update course body"
//	@Success		200			{boolean} boolean ok
//	@Failure		400			{boolean} boolean ok
//	@Router			/course [put]
func (h *CourseHandler) Update(w http.ResponseWriter, r *http.Request) {
	body := entity.CourseUpdateBody{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if body.ID == uuid.Nil {
		http.Error(w, "course handler error: id is empty!", http.StatusUnprocessableEntity)
		return
	}

	ok, err := h.service.Update(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(ok)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Delete course
//
//	@Summary		Delete course
//	@Description	delete course
//	@ID				course.delete
//	@Accept			json
//	@Produce		json
//	@Param			id			query	  string	true "course id"
//	@Success		200			{boolean} boolean ok
//	@Failure		400			{boolean} boolean ok
//	@Router			/course [delete]
func (h *CourseHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := uuid.UUID{}

	id, err := uuid.Parse(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "course handler error: error parsing id", http.StatusUnprocessableEntity)
		return
	}

	if id == uuid.Nil {
		http.Error(w, "course handler error: id is empty!", http.StatusUnprocessableEntity)
		return
	}

	ok, err := h.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(ok)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
