package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"kazusa-server/internal/entity"
	"kazusa-server/internal/service"
	"net/http"
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
//	@Param			request		body		entity.ModuleReadRequest	true "module read request"
//	@Success		200			{array}	entity.Module
//	@Failure		404			{boolean} boolean ok
//	@Router			/module [get]
func (h *ModuleHandler) Read(w http.ResponseWriter, r *http.Request) {
	filters := entity.ModuleFilters{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&filters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}

	modules, err := h.service.Read(filters)
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
