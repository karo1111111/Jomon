package router

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traPtitech/Jomon/ent"
)

type Group struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Budget      *int   `json:"budget"`
}

type Owner struct {
	ID uuid.UUID `json:"id"`
}

type OwnerResponse struct {
	Owners []uuid.UUID `json:"owners"`
}
type GroupOverview struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Budget      *int      `json:"budget"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GroupDetail struct {
	ID          uuid.UUID    `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Budget      *int         `json:"budget"`
	Owners      []*uuid.UUID `json:"owners"`
	Users       []*uuid.UUID `json:"users"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

type MemberResponse struct {
	ID []uuid.UUID `json:"members"`
}

type Member struct {
	ID uuid.UUID `json:"id"`
}

func (h *Handlers) GetGroups(c echo.Context) error {
	ctx := context.Background()
	groups, err := h.Repository.GetGroups(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	res := []*GroupOverview{}
	for _, group := range groups {
		res = append(res, &GroupOverview{
			ID:          group.ID,
			Name:        group.Name,
			Description: group.Description,
			Budget:      group.Budget,
			CreatedAt:   group.CreatedAt,
			UpdatedAt:   group.UpdatedAt,
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Handlers) PostGroup(c echo.Context) error {
	var group Group
	if err := c.Bind(&group); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	ctx := context.Background()
	created, err := h.Repository.CreateGroup(ctx, group.Name, group.Description, group.Budget)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	res := GroupOverview{
		ID:          created.ID,
		Name:        created.Name,
		Description: created.Description,
		Budget:      created.Budget,
		CreatedAt:   created.CreatedAt,
		UpdatedAt:   created.UpdatedAt,
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Handlers) PutGroup(c echo.Context) error {
	var group Group
	if err := c.Bind(&group); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	groupID, err := uuid.Parse(c.Param("groupID"))
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if groupID == uuid.Nil {
		c.Logger().Error(errors.New("invalid UUID"))
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("invalid UUID"))
	}

	ctx := context.Background()
	updated, err := h.Repository.UpdateGroup(ctx, groupID, group.Name, group.Description, group.Budget)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	res := GroupOverview{
		ID:          updated.ID,
		Name:        updated.Name,
		Description: updated.Description,
		Budget:      updated.Budget,
		CreatedAt:   updated.CreatedAt,
		UpdatedAt:   updated.UpdatedAt,
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Handlers) DeleteGroup(c echo.Context) error {
	groupID, err := uuid.Parse(c.Param("groupID"))
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if groupID == uuid.Nil {
		c.Logger().Error(errors.New("invalid UUID"))
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("invalid UUID"))
	}

	ctx := context.Background()
	err = h.Repository.DeleteGroup(ctx, groupID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}

func (h *Handlers) GetMembers(c echo.Context) error {
	groupID, err := uuid.Parse(c.Param("groupID"))
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if groupID == uuid.Nil {
		c.Logger().Error(errors.New("invalid UUID"))
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("invalid UUID"))
	}

	ctx := context.Background()
	members, err := h.Repository.GetMembers(ctx, groupID)
	if err != nil {
		if ent.IsNotFound(err) {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	var res []uuid.UUID
	for _, member := range members {
		res = append(res, member.ID)
	}

	return c.JSON(http.StatusOK, &MemberResponse{res})
}

func (h *Handlers) PostMember(c echo.Context) error {
	var member Member
	if err := c.Bind(&member); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	groupID, err := uuid.Parse(c.Param("groupID"))
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if groupID == uuid.Nil {
		c.Logger().Error(errors.New("invalid UUID"))
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("invalid UUID"))
	}

	ctx := context.Background()
	created, err := h.Repository.CreateMember(ctx, groupID, member.ID)
	if err != nil {
		if ent.IsConstraintError(err) {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	res := created.ID

	return c.JSON(http.StatusOK, &Member{res})
}

func (h *Handlers) DeleteMember(c echo.Context) error {
	var member Member
	if err := c.Bind(&member); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	groupID, err := uuid.Parse(c.Param("groupID"))
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if groupID == uuid.Nil {
		c.Logger().Error(errors.New("invalid UUID"))
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("invalid UUID"))
	}

	ctx := context.Background()
	err = h.Repository.DeleteMember(ctx, groupID, member.ID)
	if err != nil {
		if ent.IsNotFound(err) {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}

func (h *Handlers) GetOwners(c echo.Context) error {
	ctx := context.Background()
	groupID, err := uuid.Parse(c.Param("groupID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if groupID == uuid.Nil {
		c.Logger().Error(errors.New("invalid UUID"))
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("invalid UUID"))
	}
	owners, err := h.Repository.GetOwners(ctx, groupID)
	if err != nil {
		if ent.IsNotFound(err) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	var res []uuid.UUID
	for _, owner := range owners {
		res = append(res, owner.ID)
	}

	return c.JSON(http.StatusOK, &OwnerResponse{res})
}

func (h *Handlers) PostOwner(c echo.Context) error {
	ctx := context.Background()
	var owner Owner
	groupID, err := uuid.Parse(c.Param("groupID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if groupID == uuid.Nil {
		c.Logger().Error(errors.New("invalid UUID"))
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("invalid UUID"))
	}
	if err := c.Bind(&owner); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	createdOwner, err := h.Repository.CreateOwner(ctx, groupID, owner.ID)
	if err != nil {
		if ent.IsConstraintError(err) {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	res := &Owner{
		ID: createdOwner.ID,
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Handlers) DeleteOwner(c echo.Context) error {
	ctx := context.Background()
	groupID, err := uuid.Parse(c.Param("groupID"))
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if groupID == uuid.Nil {
		c.Logger().Error(errors.New("invalid UUID"))
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("invalid UUID"))
	}
	var owner Owner
	if err := c.Bind(&owner); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err = h.Repository.DeleteOwner(ctx, groupID, owner.ID)
	if err != nil {
		if ent.IsNotFound(err) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}
