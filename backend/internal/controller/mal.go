package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Gumilho/gumi-fetch/internal/types"
	"github.com/Gumilho/gumi-fetch/internal/utils"
)

type MALController struct {
	logger   types.Logger
	ClientID string
}

func NewMALController(logger types.Logger, clientID string) *MALController {
	return &MALController{logger: logger, ClientID: clientID}
}

func (mc *MALController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /search", mc.handleSearch)
}

type MALSearchResponse struct {
	Data []struct {
		Node struct {
			ID          int    `json:"id"`
			Title       string `json:"title"`
			MainPicture struct {
				Large  string `json:"large"`
				Medium string `json:"medium"`
			} `json:"main_picture"`
			AlternativeTitles struct {
				Synonyms []string `json:"synonyms"`
				En       string   `json:"en"`
				Ja       string   `json:"ja"`
			} `json:"alternative_titles"`
			MediaType types.MALMediaTypes `json:"media_type"`
			Status    types.MALStatus     `json:"status"`
		} `json:"node"`
	} `json:"data"`
	Paging struct {
		Previous string `json:"previous"`
		Next     string `json:"next"`
	} `json:"paging"`
}

// SearchShows godoc
//
//	@Summary		Search shows
//	@Description	Search shows
//	@Tags			search
//	@Produce		json
//	@Success		200
//	@Failure		500	{object}	error
//	@Router			/search [get]
func (mc *MALController) handleSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.RawQuery
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://api.myanimelist.net/v2/anime?"+query, nil)
	if err != nil {
		utils.InternalServerError(w, r, err, mc.logger)
		return
	}

	req.Header.Set("X-MAL-CLIENT-ID", mc.ClientID)
	mc.logger.Infow("request", "client id", mc.ClientID, "url", req.URL)
	resp, err := client.Do(req)
	if err != nil {
		utils.InternalServerError(w, r, err, mc.logger)
		return
	}

	var response MALSearchResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {

		utils.InternalServerError(w, r, err, mc.logger)
		return
	}
	pagination := utils.Pagination{
		Previous: response.Paging.Previous,
		Next:     response.Paging.Next,
	}
	if err := utils.JsonResponseWithPagination(w, http.StatusOK, response.Data, pagination); err != nil {
		utils.InternalServerError(w, r, err, mc.logger)
		return
	}
}
