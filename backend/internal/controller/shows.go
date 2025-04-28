package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Gumilho/gumi-fetch/internal/types"
	"github.com/Gumilho/gumi-fetch/internal/utils"
	"github.com/go-playground/validator/v10"
)

type ShowController struct {
	showStore types.ShowStore
	logger    types.Logger
}

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}
func NewShowController(showStore types.ShowStore, logger types.Logger) *ShowController {
	return &ShowController{showStore: showStore, logger: logger}
}

func (sc *ShowController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /shows", sc.handleList)
	mux.HandleFunc("POST /shows", sc.handleRegister)
	mux.HandleFunc("DELETE /shows/{id}", sc.handleUnregister)
}

// ListShows godoc
//
//	@Summary		List registered shows
//	@Description	List registered shows
//	@Tags			shows
//	@Produce		json
//	@Success		200	{object}	types.Show
//	@Failure		500	{object}	error
//	@Router			/shows [get]
func (sc *ShowController) handleList(w http.ResponseWriter, r *http.Request) {
	shows, err := sc.showStore.List()
	if err != nil {
		utils.InternalServerError(w, r, err, sc.logger)
		return
	}
	if err := utils.JsonResponse(w, http.StatusOK, shows); err != nil {
		utils.InternalServerError(w, r, err, sc.logger)
		return
	}
}

type RegisterShowPayload struct {
	ID          int    `json:"id"`
	Title       string `json:"title" validate:"required,min=1"`
	Source      string `json:"source"`
	SourceID    string `json:"source_id"`
	MainPicture string `json:"main_picture" validate:"required,min=1,url"`
}

// RegisterShow godoc
//
//	@Summary		Register show
//	@Description	Register show
//	@Tags			shows
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		RegisterShowPayload	true	"Show payload"
//	@Success		201		{object}	types.Show
//	@Failure		500		{object}	error
//	@Router			/shows [post]
func (sc *ShowController) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload RegisterShowPayload
	if err := utils.ReadJSON(w, r, &payload); err != nil {
		utils.BadRequestResponse(w, r, err, sc.logger)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		utils.BadRequestResponse(w, r, err, sc.logger)
		return
	}

	show := types.Show{
		MALID:       payload.ID,
		Title:       payload.Title,
		Source:      payload.Source,
		SourceID:    payload.SourceID,
		MainPicture: payload.MainPicture,
	}
	err := sc.showStore.Create(show)
	if err != nil {
		utils.InternalServerError(w, r, err, sc.logger)
		return
	}
	if err := utils.JsonResponse(w, http.StatusCreated, show); err != nil {
		utils.InternalServerError(w, r, err, sc.logger)
		return
	}
}

// UnregisterShow godoc
//
//	@Summary		Unregister show
//	@Description	Unregister show
//	@Tags			shows
//	@Param			id	path	int	true	"Show ID"
//	@Produce		json
//	@Success		200
//	@Failure		500	{object}	error
//	@Router			/shows/{id} [delete]
func (sc *ShowController) handleUnregister(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	idAsInt, err := strconv.Atoi(id)
	if err != nil {
		utils.BadRequestResponse(w, r, fmt.Errorf("invalid show ID format"), sc.logger)
		return
	}
	if err := sc.showStore.Delete(idAsInt); err != nil {
		utils.InternalServerError(w, r, err, sc.logger)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
