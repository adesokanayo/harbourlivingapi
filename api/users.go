package api

import (
	"database/sql"
	"errors"
	db "github.com/BigListRyRy/harbourlivingapi/db/sqlc"
	"github.com/BigListRyRy/harbourlivingapi/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type CreateUserRequest struct {
	Title       string `json:"title" `
	Username    string `json:"username" binding:"required,alphanum"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required,min=6"`
	Usertype    int32  `json:"usertype" binding:"required"`
	DateOfBirth string `json:"date_of_birth" time_format:"2006-01-02"`
}

type GetUserRequest struct {
	ID int32 `uri:"id",binding:"required,min=1"`
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type userResponse struct {
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	PasswordChangedAt string    `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

type createUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

type GetUsersRequest struct {
	PageID int32 `form:"page_id" binding:"required,min=1"`
	Size   int32 `form:"size" binding:"required,min=5,max=100"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username: user.Username,
		Email:    user.Email,
		//PasswordChangedAt: string(user.PasswordChangedAt),
		CreatedAt: user.CreatedAt,
	}
}

func newLoginResponse(user db.User, token string) loginUserResponse {
	return loginUserResponse{
		AccessToken: token,
		User:        newUserResponse(user),
	}
}

// CreateUser godoc
// @Summary Create a user
// @Description Create a user and returns a token
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {object} userResponse
// @Router /users/ [post]
func (s *Server) CreateUser(ctx *gin.Context) {
	var req CreateUserRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	dateOrTime, err := util.ProcessDateTime(req.DateOfBirth)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	createUserReq := db.CreateUserParams{
		Title:       req.Title,
		FirstName:   req.FirstName,
		Username:    req.Username,
		LastName:    req.LastName,
		Email:       req.Email,
		Password:    hashedPassword,
		Usertype:    req.Usertype,
		DateOfBirth: *dateOrTime,
	}

	user, err := s.store.CreateUser(ctx, createUserReq)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("Username already exist")))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, newUserResponse(user))
}

func (s *Server) ListUsers(ctx *gin.Context) {
	users, err := s.store.GetAllUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, users)
}

// Login godoc
// @Summary Authenticate a user with username and password and generate token
// @Description Authenticate a user returns a token
// @Tags Login
// @Accept  json
// @Produce  json
// @Success 200 {object} loginUserResponse
// @Router /login/ [post]
func (s *Server) Login(ctx *gin.Context) {
	var req loginUserRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, err := s.store.GetUsername(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("invalid username & password combination")))
		return
	}

	token, err := s.tokenMaker.CreateToken(user.Username, time.Minute)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, newLoginResponse(user, token))
}

func (s *Server) GetUser(ctx *gin.Context) {
	var req GetUserRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, err := s.store.GetUser(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}
