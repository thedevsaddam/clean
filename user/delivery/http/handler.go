package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/thedevsaddam/clean/domain"
	"github.com/thedevsaddam/clean/user/presenter"
)

// UserHandler ..
type UserHandler struct {
	UserUsecase domain.UserUsecase
}

// NewUserHandler will initialize the resources endpoint
func NewUserHandler(r *chi.Mux, us domain.UserUsecase) {
	handler := &UserHandler{
		UserUsecase: us,
	}
	r.Route("/v1/users", func(r chi.Router) {
		r.Post("/", handler.Store)
		r.Get("/", handler.FetchUsers)
		r.Get("/{id}", handler.FetchUser)
	})
}

// Store ...
func (u *UserHandler) Store(w http.ResponseWriter, r *http.Request) {
	user := domain.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		(&Response{Status: http.StatusBadRequest, Message: "Bad request"}).JSON(w)
		log.Println("handler/store:", err)
		return
	}
	if user.Profile == nil {
		(&Response{Status: http.StatusBadRequest, Message: "User's profile data required"}).JSON(w)
		return
	}
	ctx := context.TODO()
	id, err := u.UserUsecase.Store(ctx, &user)
	if err != nil {
		(&Response{Status: http.StatusInternalServerError, Message: "Failed to create user"}).JSON(w)
		return
	}

	(&Response{
		Status:  http.StatusCreated,
		Message: "User created successfully",
		Data:    id,
	}).JSON(w)
}

// FetchUsers ...
func (u *UserHandler) FetchUsers(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	ctr := &domain.UserCriteria{}
	users, err := u.UserUsecase.Fetch(ctx, ctr)
	if err != nil {
		(&Response{Status: http.StatusInternalServerError, Message: "Failed to fetch users"}).JSON(w)
		return
	}

	// convert data for endpoint
	up := presenter.NewUserPresenter()
	respUsers := up.PresentUsers(users)

	(&Response{
		Status: http.StatusOK,
		Data:   respUsers,
	}).JSON(w)
}

// FetchUser ...
func (u *UserHandler) FetchUser(w http.ResponseWriter, r *http.Request) {
	uid, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		(&Response{Status: http.StatusBadRequest, Message: "Invalid user id"}).JSON(w)
		log.Println("handler/fetchuser:", err)
		return
	}
	ctx := context.TODO()
	user, err := u.UserUsecase.GetByID(ctx, uint(uid))
	if err != nil {
		(&Response{Status: http.StatusInternalServerError, Message: "Failed to fetch user"}).JSON(w)
		log.Println("handler/fetchuser", err)
		return
	}

	if user == nil {
		(&Response{Status: http.StatusNotFound, Message: "User not found"}).JSON(w)
		return
	}

	// convert data for endpoint
	up := presenter.NewUserPresenter()
	respUsers := up.PresentUser(user)

	(&Response{
		Status: http.StatusOK,
		Data:   respUsers,
	}).JSON(w)
}
