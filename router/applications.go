package router

import (
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"github.com/traPtitech/Jomon/model"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type ApplicationService struct {
	repo model.ApplicationRepository
}

func NewApplicationService(repo model.ApplicationRepository) *ApplicationService {
	return &ApplicationService{repo: repo}
}

type GetApplicationsListQuery struct {
	Sort           *string `query:"sort"`
	CurrentState   *string `query:"current_state"`
	FinancialYear  *int    `query:"financial_year"`
	Applicant      *string `query:"applicant"`
	Type           *string `query:"type"`
	SubmittedSince *string `query:"submitted_since"`
	SubmittedUntil *string `query:"submitted_until"`
}

type PostApplicationRequest struct {
	Type       *model.ApplicationType `json:"type"`
	Title      *string                `json:"title"`
	Remarks    *string                `json:"remarks"`
	PaidAt     *time.Time             `json:"paid_at"`
	Amount     *int                   `json:"amount"`
	RepaidToId []string               `json:"repaid_to_id"`
}

func (s *Service) GetApplicationList(c echo.Context) error {
	var query GetApplicationsListQuery
	err := c.Bind(&query)
	if err != nil {
		// TODO
		return c.NoContent(http.StatusBadRequest)
	}

	var currentState *model.StateType
	if query.CurrentState != nil {
		_currentState, err := model.GetStateType(*query.CurrentState)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		currentState = &_currentState
	}

	var typ *model.ApplicationType
	if query.Type != nil {
		_typ, err := model.GetApplicationType(*query.Type)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		typ = &_typ
	}

	var submittedSince *time.Time
	if query.SubmittedSince != nil {
		_submittedSince, err := StrToDate(*query.SubmittedSince)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		submittedSince = &_submittedSince
	}

	var submittedUntil *time.Time
	if query.SubmittedUntil != nil {
		_submittedUntil, err := StrToDate(*query.SubmittedUntil)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		_submittedUntil = _submittedUntil.AddDate(0, 0, 1)
		submittedUntil = &_submittedUntil
	}

	applications, err := s.Applications.repo.GetApplicationList(query.Sort, currentState, query.FinancialYear, query.Applicant, typ, submittedSince, submittedUntil, true)
	return c.JSON(http.StatusOK, applications)
}

func (s *Service) GetApplication(c echo.Context) error {
	applicationId := uuid.FromStringOrNil(c.Param("applicationId"))
	if applicationId == uuid.Nil {
		return c.NoContent(http.StatusBadRequest)
	}
	application, err := s.Applications.repo.GetApplication(applicationId, true, true)
	if gorm.IsRecordNotFoundError(err) {
		return c.NoContent(http.StatusNotFound)
	} else if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	fmt.Printf("%+v\n", application)

	return c.JSON(http.StatusOK, application)
}

func (s *Service) PostApplication(c echo.Context) error {
	details := c.FormValue("details")
	var req PostApplicationRequest
	if err := json.Unmarshal([]byte(details), &req); err != nil {
		// TODO more information
		return c.NoContent(http.StatusBadRequest)
	}

	if req.Type == nil || req.Title == nil || req.Remarks == nil || req.Amount == nil || req.PaidAt == nil {
		return c.NoContent(http.StatusBadRequest)
	}

	userId := ""

	id, err := s.Applications.repo.BuildApplication(userId, *req.Type, *req.Title, *req.Remarks, *req.Amount, *req.PaidAt)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	app, err := s.Applications.repo.GetApplication(id, true, true)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	// TODO Handle image files

	return c.JSON(http.StatusCreated, app)
}

func (s *Service) PatchApplication(c echo.Context) error {
	applicationId := uuid.FromStringOrNil(c.Param("applicationId"))
	userId := ""
	if applicationId == uuid.Nil {
		return c.NoContent(http.StatusBadRequest)
	}

	details := c.FormValue("details")
	var req PostApplicationRequest
	if err := json.Unmarshal([]byte(details), &req); err != nil {
		// TODO
		return c.NoContent(http.StatusBadRequest)
	}

	err := s.Applications.repo.PatchApplication(applicationId, userId, req.Type, req.Title, req.Remarks, req.Amount, req.PaidAt)
	if gorm.IsRecordNotFoundError(err) {
		return c.NoContent(http.StatusNotFound)
	} else if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	app, err := s.Applications.repo.GetApplication(applicationId, true, true)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	// TODO Handle image files

	return c.JSON(http.StatusOK, &app)
}
