package handler

import (
	"encoding/json"
	"github.com/AlnurZhanibek/kazusa-server/internal/service"
	"github.com/google/uuid"
	"net/http"
)

type ActivityHandler struct {
	service *service.ActivityService
}

func NewActivityHandler(service *service.ActivityService) *ActivityHandler {
	return &ActivityHandler{service: service}
}

type NewActivity struct {
	UserID       uuid.UUID `json:"userId" validate:"required"`
	UserEmail    string    `json:"userEmail" validate:"required"`
	UserFullname string    `json:"userFullname" validate:"required"`
	CourseID     uuid.UUID `json:"courseId" validate:"required"`
	CourseName   string    `json:"courseName" validate:"required"`
	ModuleID     uuid.UUID `json:"moduleId" validate:"required"`
	IsLast       *bool     `json:"isLast"`
} // @name NewActivity

// Create activity
//
//	@Summary		Create activity
//	@Description	create activity
//	@ID				activity.create
//	@Accept			json
//	@Produce		json
//	@Param			request		body	NewActivity	true	"new activity body"
//	@Success		200			{boolean} boolean ok
//	@Failure		400			{boolean} boolean ok
//	@Router			/activity [post]
func (h *ActivityHandler) Create(w http.ResponseWriter, r *http.Request) {
	newActivity := new(NewActivity)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newActivity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	if newActivity.UserID == uuid.Nil || newActivity.CourseID == uuid.Nil || newActivity.ModuleID == uuid.Nil {
		http.Error(w, "user_id, course_id or module_id is empty!", http.StatusUnprocessableEntity)
		return
	}

	err = h.service.Create(&service.ActivityCreateBody{
		UserID:       newActivity.UserID,
		UserEmail:    newActivity.UserEmail,
		UserFullname: newActivity.UserFullname,
		CourseID:     newActivity.CourseID,
		CourseName:   newActivity.CourseName,
		ModuleID:     newActivity.ModuleID,
		IsLast:       newActivity.IsLast,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
