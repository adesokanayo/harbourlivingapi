package api

import (
	"database/sql"
	db "github.com/BigListRyRy/harbourlivingapi/db/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateEventRequest struct {
	Title       string         `json:"title"`
	Description string         `json:"description"`
	BannerImage string         `json:"banner_image"`
	StartDate   string         `json:"start_date"`
	EndDate     string         `json:"end_date"`
	Venue       int32          `json:"venue"`
	Type        int32          `json:"type"`
	UserID      int32          `json:"user_id"`
	Category    int32          `json:"category"`
	Subcategory int32          `json:"subcategory"`
	Status      sql.NullString `json:"status"`
	Image1      sql.NullString `json:"image1"`
	Image2      sql.NullString `json:"image2"`
	Image3      sql.NullString `json:"image3"`
	Video1      sql.NullString `json:"video1"`
	Video2      sql.NullString `json:"video2"`
}

type GetEventRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type GetEventsRequest struct {
	PageID      int32 `form:"page_id" binding:"required,min=1"`
	PageSize    int32 `form:"page_size" binding:"required,min=1,max=100"`
	Category    int32 `form:"category"`
	SubCategory int32 `form:"subcategory"`
}

func (s *Server) CreateEvent(ctx *gin.Context) {
	var req CreateEventRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	startDate, err := ProcessDateTime(req.StartDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	endDate, err := ProcessDateTime(req.StartDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateEventParams{
		Title:       req.Title,
		Description: req.Description,
		StartDate:   *startDate,
		EndDate:     *endDate,
		Venue:       req.Venue,
		Type:        req.Type,
		UserID:      req.UserID,
		Category:    req.Category,
		Subcategory: req.Subcategory,
	}
	event, err := s.store.CreateEvent(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, event)
}

func (s *Server) ListEvents(ctx *gin.Context) {
	var req GetEventsRequest

	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetAllEventsParams{
		Category:    req.Category,
		Subcategory: req.SubCategory,
		Limit:       req.PageSize,
		Offset:      (req.PageID - 1) * req.PageSize,
	}
	events, err := s.store.GetAllEvents(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func (s *Server) GetEvent(ctx *gin.Context) {

	var req GetEventRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	event, err := s.store.GetEvent(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, event)
}
