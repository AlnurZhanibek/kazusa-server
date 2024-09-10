package handler

import (
	"context"
	"encoding/json"
	"github.com/AlnurZhanibek/kazusa-server/internal/entity"
	"github.com/AlnurZhanibek/kazusa-server/internal/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"net/http"
	"os"
	"strconv"
)

type ModuleHandlerImplementation interface {
	Create(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
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
//	@Router			/module [post]
func (h *ModuleHandler) Create(w http.ResponseWriter, r *http.Request) {
	newModule := entity.NewModule{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newModule)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if newModule.Name == "" || newModule.Content == "" || newModule.CourseID == uuid.Nil || newModule.DurationMinutes == 0 || newModule.Order == 0 {
		http.Error(w, "name, content, content or order duration is empty!", http.StatusUnprocessableEntity)
		return
	}

	ok, err := h.service.Create(newModule)
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
	offset, _ := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 64)
	limit, _ := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)

	id := r.URL.Query().Get("id")
	courseID := r.URL.Query().Get("course_id")

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

	modules, err := h.service.Read(ctx, pagination, filters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(modules)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Update module
//
//	@Summary		Update module
//	@Description	update module
//	@ID				module.update
//	@Accept			json
//	@Produce		json
//	@Param			request		body		entity.ModuleUpdateBody	true "update module body"
//	@Success		200			{boolean} boolean ok
//	@Failure		400			{boolean} boolean ok
//	@Router			/module [post]
func (h *ModuleHandler) Update(w http.ResponseWriter, r *http.Request) {
	body := entity.ModuleUpdateBody{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if body.ID == uuid.Nil {
		http.Error(w, "id is empty!", http.StatusUnprocessableEntity)
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

// Delete module
//
//	@Summary		Delete module
//	@Description	delete module
//	@ID				module.delete
//	@Accept			json
//	@Produce		json
//	@Param			id		    query	  string	true "module id"
//	@Success		200			{boolean} boolean ok
//	@Failure		400			{boolean} boolean ok
//	@Router			/module [delete]
func (h *ModuleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := uuid.UUID{}

	id, err := uuid.Parse(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "module handler error: error parsing id", http.StatusUnprocessableEntity)
		return
	}

	if id == uuid.Nil {
		http.Error(w, "module handler error: id is empty!", http.StatusUnprocessableEntity)
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
