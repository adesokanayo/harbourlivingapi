package rest

import (
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/BigListRyRy/harbourlivingapi/db/sqlc"
	"github.com/gin-gonic/gin"
)

type CreateVenueRequest struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Province    string `json:"province"`
	CountryCode string `json:"country_code" default:"CAN"`
}
type GetVenueRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (s *HTTPServer) CreateVenue(ctx *gin.Context) {
	var req CreateVenueRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	createEventReq := db.CreateVenueParams{
		Name: req.Name,
	}

	user, err := s.store.CreateVenue(ctx, createEventReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (s *HTTPServer) ListVenues(ctx *gin.Context) {
	venues, err := s.store.GetAllVenues(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, venues)
}
func (s *HTTPServer) GetVenue(ctx *gin.Context) {

	var req GetVenueRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	venue, err := s.store.GetVenue(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, venue)
}

func (s *HTTPServer) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello!")
	/*
		var req GetVenueRequest
		err := ctx.ShouldBindUri(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		venue, err := s.store.GetVenue(ctx, req.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, venue)
	*/
}
