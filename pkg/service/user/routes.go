package user

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

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
	router.HandleFunc("/register", h.handleRegister)
}

func (h *Handler) handleLogin(rw http.ResponseWriter, r *http.Request) {
	log.Println("Received request")
	text, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	newText := strings.ReplaceAll(string(text), "this", "butt")
	rw.Write([]byte(newText))
}

func (h *Handler) handleRegister(rw http.ResponseWriter, r *http.Request) {
	// json payload for new user
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(rw, http.StatusBadRequest, err)
	}

	// check if user exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(rw, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
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
		utils.WriteError(rw, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(rw, http.StatusCreated, nil)
}
