package recipe

import (
	"fmt"
	//"log"
	"net/http"

	//"github.com/hweepok/mixmashbackend/pkg/config"
	"github.com/google/uuid"
	"github.com/hweepok/mixmashbackend/pkg/types"
	"github.com/hweepok/mixmashbackend/pkg/utils"
)

type RecipeHandler struct {
	store types.RecipeStore
}

func NewHandler(store types.RecipeStore) *RecipeHandler {
	return &RecipeHandler{
		store: store,
	}
}

func (h *RecipeHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /recipe", h.handlePushRecipe)
}

func (h *RecipeHandler) handlePushRecipe(rw http.ResponseWriter, r *http.Request) {
	// Recipe payload
	var recipePayload types.PushRecipePayload
	if err := utils.ParseJSON(r, &recipePayload); err != nil {
		utils.WriteError(rw, http.StatusBadRequest, fmt.Errorf("send recipe just doesn't work  %s", err))
		return
	}

	// Might not be needed if UUID is implemented
	/*
		_, err := h.store.GetRecipeByName(recipePayload.Name)
		if err == nil {
			utils.WriteError(rw, http.StatusBadRequest, fmt.Errorf("recipe name already exists %s: %s", recipePayload.Name, err))
			return
		}
	*/

	err := h.store.CreateRecipe(types.Recipe{
		ID:          uuid.NewString(),
		Name:        recipePayload.Name,
		Description: recipePayload.Description,
		ImageUrl:    recipePayload.ImageUrl,
		Source:      recipePayload.Source,
		Alterations: recipePayload.Alterations,
		TotalTime:   recipePayload.TotalTime,
		Servings:    recipePayload.Servings,
	})
	if err != nil {
		utils.WriteError(rw, http.StatusInternalServerError, fmt.Errorf("couldn't create recipe %s: %s", recipePayload.Name, err))
		return
	}

	utils.WriteJSON(rw, http.StatusCreated, nil)
}
