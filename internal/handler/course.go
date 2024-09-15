package handler

import (
	"context"
	"encoding/json"
	"github.com/AlnurZhanibek/kazusa-server/internal/entity"
	"github.com/AlnurZhanibek/kazusa-server/internal/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"mime/multipart"
	"net/http"
	"os"
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

type CourseCreateBody struct {
	Title       string
	Description string
	Price       int64
	Cover       service.FileWithHeader
	Attachments []service.FileWithHeader
}

// Create course
//
//	@Summary		Create course
//	@Description	create course
//	@ID				course.create
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			title formData string true "title"
//	@Param			description formData string	true "description"
//	@Param			price formData number true "price"
//	@Param			cover formData file	true "cover"
//	@Param			attachments formData file true "cover"
//	@Success		200 {boolean} boolean ok
//	@Failure		400 {boolean} boolean ok
//	@Router			/course [post]
func (h *CourseHandler) Create(w http.ResponseWriter, r *http.Request) {
	newCourse := CourseCreateBody{}

	coverFile, coverHeader, err := r.FormFile("cover")
	defer coverFile.Close()
	cover := service.FileWithHeader{
		Header: coverHeader,
		File:   coverFile,
	}

	attachmentHeaders := r.MultipartForm.File["attachments"]
	attachments := make([]service.FileWithHeader, len(attachmentHeaders))
	for i, fh := range attachmentHeaders {
		var file multipart.File
		file, err = fh.Open()
		if err != nil {
			http.Error(w, "error reading one of the attachments", http.StatusUnprocessableEntity)
			return
		}
		attachments[i] = service.FileWithHeader{
			Header: fh,
			File:   file,
		}
	}
	newCourse.Title = r.FormValue("title")
	newCourse.Description = r.FormValue("description")
	priceInt, err := strconv.Atoi(r.FormValue("price"))
	if err != nil {
		http.Error(w, "price should be a number", http.StatusUnprocessableEntity)
		return
	}
	newCourse.Price = int64(priceInt)
	newCourse.Cover = cover
	newCourse.Attachments = attachments

	if newCourse.Title == "" || newCourse.Description == "" || newCourse.Price == 0 {
		http.Error(w, "title, description or price is empty!", http.StatusUnprocessableEntity)
		return
	}

	ok, err := h.service.Create(service.CourseCreateBody{
		Title:       newCourse.Title,
		Description: newCourse.Description,
		Price:       newCourse.Price,
		Cover:       newCourse.Cover,
		Attachments: newCourse.Attachments,
	})

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
	offset, _ := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 64)
	limit, _ := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)

	id := r.URL.Query().Get("id")

	filters := entity.CourseFilters{
		ID: uuid.Nil,
	}

	if id != "" {
		filters.ID = uuid.MustParse(id)
	}

	ctx := r.Context()

	tokenCookie, err := r.Cookie("token")
	if tokenCookie != nil {
		token := new(jwt.Token)
		token, err = jwt.ParseWithClaims(tokenCookie.Value, &service.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_KEY")), nil
		})
		claims, ok := token.Claims.(*service.Claims)
		if !ok || !token.Valid {
			http.Error(w, "error parsing jwt token", http.StatusBadRequest)

		}
		ctx = context.WithValue(ctx, "user_id", claims.UserID)
	}

	courses, err := h.service.Read(ctx, entity.Pagination{
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
