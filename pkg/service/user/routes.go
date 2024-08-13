package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hweepok/mixmashbackend/pkg/config"
	"github.com/hweepok/mixmashbackend/pkg/service/auth"
	"github.com/hweepok/mixmashbackend/pkg/types"
	"github.com/hweepok/mixmashbackend/pkg/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/login", h.handleLogin)
	router.HandleFunc("POST /register", h.handleRegister)
}

func (h *Handler) handleLogin(rw http.ResponseWriter, r *http.Request) {
	var payload types.LoginUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(rw, http.StatusBadRequest, fmt.Errorf("just doesn't work %s: %s", err, payload))
		utils.WriteError(rw, http.StatusBadRequest, err)
		return
	}

	// Validate payload TODO

	// check if user exists
	user, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(rw, http.StatusBadRequest, fmt.Errorf("Invalid email or password"))
		return
	}

	if !auth.ComparePasswords(user.Password, []byte(payload.Password)) {
		utils.WriteError(rw, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, user.ID)
	if err != nil {
		utils.WriteError(rw, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(rw, http.StatusOK, map[string]string{"token": token})
}

func (h *Handler) handleRegister(rw http.ResponseWriter, r *http.Request) {
	// json payload for new user
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(rw, http.StatusBadRequest, fmt.Errorf("just doesn't work %s: %s", err, payload))
		utils.WriteError(rw, http.StatusBadRequest, err)
		return
	}

	// Validate payload
	// TODO

	// check if user exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(rw, http.StatusBadRequest, fmt.Errorf("user with email %s already exists %s", payload.Email, err))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = h.store.CreateUser(types.User{
		UserName: payload.UserName,
		Email:    payload.Email,
		Password: hashedPassword,
	})
	if err != nil {
		utils.WriteError(rw, http.StatusInternalServerError, fmt.Errorf("couldn't create user"))
		utils.WriteError(rw, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(rw, http.StatusCreated, nil)
}
