package handler

import (
	"encoding/json"
	"github.com/AlnurZhanibek/kazusa-server/internal/entity"
	"github.com/AlnurZhanibek/kazusa-server/internal/service"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

type UserHandlerImplementation interface {
	Create(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
	service service.UserServiceImplementation
}

func NewUserHandler(service service.UserServiceImplementation) *UserHandler {
	return &UserHandler{service: service}
}

// Create user
//
//	@Summary		Create user
//	@Description	create user
//	@ID				user.create
//	@Accept			json
//	@Produce		json
//	@Param			request		body		entity.NewUser	true "new user body"
//	@Success		200			{boolean} boolean ok
//	@Failure		400			{boolean} boolean ok
//	@Router			/user [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	newUser := entity.NewUser{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if newUser.Name == "" || newUser.Email == "" {
		http.Error(w, "name or email is empty!", http.StatusUnprocessableEntity)
		return
	}

	ok, err := h.service.Create(newUser)
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

// Read User
//
//	@Summary		Read users
//	@Description	read users
//	@ID				user.read
//	@Accept			json
//	@Produce		json
//	@Param			offset		query		int64	true "offset"
//	@Param			limit		query		int64	true "limit"
//	@Param			id			query		string	false "id"
//	@Success		200			{array}	entity.User
//	@Failure		404			{boolean} boolean ok
//	@Router			/user [get]
func (h *UserHandler) Read(w http.ResponseWriter, r *http.Request) {
	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}
	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}

	id := r.URL.Query().Get("id")

	filters := entity.UserFilters{
		ID: &uuid.Nil,
	}

	if id != "" {
		parsedID := uuid.MustParse(id)
		filters.ID = &parsedID
	}

	users, err := h.service.Read(entity.Pagination{
		Offset: offset,
		Limit:  limit,
	}, filters)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Update user
//
//	@Summary		Update user
//	@Description	update user
//	@ID				user.update
//	@Accept			json
//	@Produce		json
//	@Param			request		body		entity.UserUpdateBody	true "update user body"
//	@Success		200			{boolean} boolean ok
//	@Failure		400			{boolean} boolean ok
//	@Router			/user [put]
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	body := entity.UserUpdateBody{}

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

// Delete user
//
//	@Summary		Delete user
//	@Description	delete user
//	@ID				user.delete
//	@Accept			json
//	@Produce		json
//	@Param			id		    query	  string	true "user id"
//	@Success		200			{boolean} boolean ok
//	@Failure		400			{boolean} boolean ok
//	@Router			/user [delete]
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := uuid.UUID{}

	id, err := uuid.Parse(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "user handler error: error parsing id", http.StatusUnprocessableEntity)
		return
	}

	if id == uuid.Nil {
		http.Error(w, "user handler error: id is empty!", http.StatusUnprocessableEntity)
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
