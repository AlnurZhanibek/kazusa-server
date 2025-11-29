package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/AlnurZhanibek/kazusa-server/internal/entity"
	"github.com/AlnurZhanibek/kazusa-server/internal/repository"
	"github.com/google/uuid"
)

type CourseServiceImplementation interface {
	Create(course CourseCreateBody) (bool, error)
	Read(ctx context.Context, pagination entity.Pagination, filters entity.CourseFilters) ([]entity.Course, error)
	Update(body entity.CourseUpdateBody) (bool, error)
	Delete(id uuid.UUID) (bool, error)
}

type CourseService struct {
	repo           repository.CourseRepositoryImplementation
	moduleService  ModuleServiceImplementation
	fileService    *FileService
	paymentService *PaymentService
}

func NewCourseService(repo repository.CourseRepositoryImplementation, moduleService ModuleServiceImplementation, fileService *FileService, paymentService *PaymentService) *CourseService {
	return &CourseService{repo: repo, moduleService: moduleService, fileService: fileService, paymentService: paymentService}
}

type FileWithHeader struct {
	Header *multipart.FileHeader
	File   multipart.File
}

type CourseCreateBody struct {
	Title       string
	Description string
	Price       int64
	Cover       FileWithHeader
	Attachments []FileWithHeader
}

func (s *CourseService) Create(course CourseCreateBody) (bool, error) {

	ctx := context.Background()
	coverFilename := strings.ToLower(strings.ReplaceAll(course.Title, " ", "-")) + "-cover" + filepath.Ext(course.Cover.Header.Filename)

	coverUrl, err := s.fileService.Put(ctx, coverFilename, course.Cover.File)
	if err != nil {
		return false, fmt.Errorf("course service create error: uploading cover image")
	}

	attachmentURLs := make([]string, len(course.Attachments))
	for i, attachment := range course.Attachments {
		filename := strings.ToLower(strings.ReplaceAll(attachment.Header.Filename, " ", "-"))
		var url *string
		url, err = s.fileService.Put(ctx, filename, attachment.File)
		if err != nil {
			return false, fmt.Errorf("course service create error: uploading attachment")
		}
		attachmentURLs[i] = *url
	}

	ok, err := s.repo.Create(repository.CourseCreateBody{
		Title:          course.Title,
		Description:    course.Description,
		Price:          course.Price,
		CoverURL:       *coverUrl,
		AttachmentURLs: attachmentURLs,
	})

	if err != nil {
		return false, fmt.Errorf("course service create error: %v", err)
	}

	return ok, nil
}

func (s *CourseService) Read(ctx context.Context, pagination entity.Pagination, filters entity.CourseFilters) ([]entity.Course, error) {
	repoCourses, err := s.repo.Read(pagination, filters)
	if err != nil {
		return nil, fmt.Errorf("course service create error: %v", err)
	}

	courses := make([]entity.Course, len(repoCourses))
	for i, repoCourse := range repoCourses {
		courses[i] = entity.Course{
			ID:             repoCourse.ID,
			CreatedAt:      repoCourse.CreatedAt,
			UpdatedAt:      repoCourse.UpdatedAt,
			Title:          repoCourse.Title,
			Description:    repoCourse.Description,
			Price:          repoCourse.Price,
			CoverURL:       repoCourse.CoverURL,
			AttachmentURLs: repoCourse.AttachmentURLs.String,
			Modules:        nil,
		}
	}

	if filters.ID != uuid.Nil || len(courses) == 1 {
		modules, err := s.moduleService.Read(ctx, entity.Pagination{}, entity.ModuleFilters{
			CourseID: courses[0].ID,
		})

		if err != nil {
			return nil, fmt.Errorf("course service get modules error: %v", err)
		}

		courses[0].Modules = &modules

		userIDCtx := ctx.Value("user_id")
		if userIDCtx != nil {
			userID := userIDCtx.(uuid.UUID)
			payment := &Payment{}
			payment, _ = s.paymentService.Read(PaymentFilters{
				UserID:   &userID,
				CourseID: &filters.ID,
			})

			if payment != nil {
				courses[0].IsPaid = true
			} else {
				courses[0].IsPaid = false
			}
		}
	}

	return courses, nil
}

func (s *CourseService) Update(course entity.CourseUpdateBody) (bool, error) {
	ok, err := s.repo.Update(course)
	if err != nil {
		return false, fmt.Errorf("course service update error: %v", err)
	}

	return ok, nil
}

func (s *CourseService) Delete(id uuid.UUID) (bool, error) {
	ok, err := s.repo.Delete(id)
	if err != nil {
		return false, fmt.Errorf("course service delete error: %v", err)
	}

	return ok, nil
}
