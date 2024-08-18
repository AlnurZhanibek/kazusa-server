package handler

import (
	"encoding/json"
	"fmt"
	"github.com/AlnurZhanibek/kazusa-server/internal/entity"
	"github.com/AlnurZhanibek/kazusa-server/internal/service"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

type ModuleHandlerImplementation interface {
	Create(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
}

type ModuleHandler struct {
	service service.ModuleServiceImplementation
}

func NewModuleHandler(service service.ModuleServiceImplementation) *ModuleHandler {
	return &ModuleHandler{service: service}
}

// Create module
//
//	@Summary		Create module
//	@Description	create module
//	@ID				module.create
//	@Accept			json
//	@Produce		json
//	@Param			request		body		entity.NewModule	true "new module body"
//	@Success		200			{boolean} boolean ok
//	@Failure		400			{boolean} boolean ok
//	@Router			/course [post]
func (h *ModuleHandler) Create(w http.ResponseWriter, r *http.Request) {
	newModule := entity.NewModule{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newModule)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}

	if newModule.Name == "" || newModule.Content == "" || newModule.CourseID == uuid.Nil || newModule.DurationMinutes == 0 {
		http.Error(w, "name, content, content or duration is empty!", http.StatusUnprocessableEntity)
		return
	}

	ok, err := h.service.Create(newModule)
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

// Read module
//
//	@Summary		Read modules
//	@Description	read modules
//	@ID				module.read
//	@Accept			json
//	@Produce		json
//	@Param			id			query		string		false 	"id"
//	@Param			course_id	query		string		false 	"course_id"
//	@Param			offset		query		int64		false 	"offset"
//	@Param			limit		query		int64		false 	"limit"
//	@Success		200			{array}		entity.Module
//	@Failure		404			{boolean} 	boolean ok
//	@Router			/module [get]
func (h *ModuleHandler) Read(w http.ResponseWriter, r *http.Request) {
	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 64)
	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)

	id := r.URL.Query().Get("id")
	courseID := r.URL.Query().Get("course_id")
	fmt.Println(id)

	filters := entity.ModuleFilters{
		ID:       uuid.Nil,
		CourseID: uuid.Nil,
	}

	pagination := entity.Pagination{
		Offset: 0,
		Limit:  0,
	}

	if offset != 0 {
		pagination.Offset = offset
	}

	if limit != 0 {
		pagination.Limit = limit
	}

	if id != "" {
		filters.ID = uuid.MustParse(id)
	}

	if courseID != "" {
		filters.CourseID = uuid.MustParse(courseID)
	}

	modules, err := h.service.Read(filters, pagination)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(modules)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
