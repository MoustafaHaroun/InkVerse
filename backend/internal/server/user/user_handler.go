package user

import (
	"fmt"
	"net/http"

	"github.com/MoustafaHaroun/InkVerse/pkg/auth"
	"github.com/MoustafaHaroun/InkVerse/pkg/util"
	"github.com/google/uuid"
)

type UserHandler struct {
	UserRepository SQLUserRepository
}

func NewUserHandler(UserRepository SQLUserRepository) *UserHandler {
	return &UserHandler{
		UserRepository: UserRepository,
	}
}

func (h *UserHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /login", h.Login)
	router.HandleFunc("POST /register", h.Register)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var payload LoginUserPayload

	err := util.ParseJsonPayload(r, &payload)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.UserRepository.GetByEmail(payload.Email)

	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	ok := auth.CheckPassword(payload.Password, user.Password)

	if !ok {
		util.WriteError(w, http.StatusUnauthorized, fmt.Errorf("password is incorrect"))
		return
	}

	token, err := auth.CreateToken(user.Username)

	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var payload RegisterUserPayload

	err := util.ParseJsonPayload(r, &payload)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.UserRepository.GetByEmail(payload.Email)

	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if user.ID != uuid.Nil {
		util.WriteError(w, http.StatusBadRequest, fmt.Errorf("user already exists with this email")) //TODO: move this to the repo
		return
	}

	payload.Password, err = auth.HashPassword(payload.Password)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.UserRepository.Add(User{
		Username: payload.UserName,
		Email:    payload.Email,
		Password: payload.Password,
	})

	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteJSON(w, http.StatusCreated, nil)
}
