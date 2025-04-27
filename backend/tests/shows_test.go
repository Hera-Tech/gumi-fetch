package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Gumilho/gumi-fetch/internal/controller"
	"github.com/Gumilho/gumi-fetch/internal/logger"
	"github.com/Gumilho/gumi-fetch/internal/types"
	"github.com/Gumilho/gumi-fetch/mocks"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()
}

func TestShowController_handleList(t *testing.T) {
	tests := []struct {
		name           string
		mockShowStore  *mocks.MockShowStore
		expectedStatus int
		expectedBody   []types.Show
		mockListError  error
	}{
		{
			name: "successful list",
			mockShowStore: func() *mocks.MockShowStore {
				store := new(mocks.MockShowStore)
				shows := []types.Show{{MALID: 1, Title: "Show 1"}, {MALID: 2, Title: "Show 2"}}
				store.On("List").Return(shows, nil)
				return store
			}(),
			expectedStatus: http.StatusOK,
			expectedBody:   []types.Show{{MALID: 1, Title: "Show 1"}, {MALID: 2, Title: "Show 2"}},
		},
		{
			name: "error listing shows",
			mockShowStore: func() *mocks.MockShowStore {
				store := new(mocks.MockShowStore)
				store.On("List").Return([]types.Show{}, assert.AnError)
				return store
			}(),
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name: "empty list",
			mockShowStore: func() *mocks.MockShowStore {
				store := new(mocks.MockShowStore)
				store.On("List").Return([]types.Show{}, nil)
				return store
			}(),
			expectedStatus: http.StatusOK,
			expectedBody:   []types.Show{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLogger := new(logger.NoOpLogger)
			controller := controller.NewShowController(tt.mockShowStore, mockLogger)

			rr := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/shows", nil)

			mux := http.NewServeMux()
			controller.RegisterRoutes(mux)
			mux.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			if tt.expectedBody != nil {
				var envelope struct {
					Data []types.Show `json:"data"`
				}
				err := json.Unmarshal(rr.Body.Bytes(), &envelope)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, envelope.Data)
			}
		})
	}
}

func TestShowController_handleRegister(t *testing.T) {
	tests := []struct {
		name            string
		payload         controller.RegisterShowPayload
		mockShowStore   *mocks.MockShowStore
		expectedStatus  int
		mockCreateError error
		validationError bool
	}{
		{
			name: "successful registration",
			payload: controller.RegisterShowPayload{
				ID:          3,
				Title:       "New Show",
				Source:      "Some Source",
				SourceID:    "new_show_id",
				MainPicture: "https://cdn.myanimelist.net/images/anime/1770/97704.jpg",
			},
			mockShowStore: func() *mocks.MockShowStore {
				store := new(mocks.MockShowStore)
				show := types.Show{MALID: 3, Title: "New Show", Source: "Some Source", SourceID: "new_show_id", MainPicture: "https://cdn.myanimelist.net/images/anime/1770/97704.jpg"}
				store.On("Create", show).Return(nil)
				return store
			}(),
			expectedStatus: http.StatusCreated,
		},
		{
			name: "error creating show",
			payload: controller.RegisterShowPayload{
				ID:          4,
				Title:       "Failed Show",
				Source:      "Another Source",
				SourceID:    "failed_id",
				MainPicture: "https://cdn.myanimelist.net/images/anime/1770/97704.jpg",
			},
			mockShowStore: func() *mocks.MockShowStore {
				store := new(mocks.MockShowStore)
				show := types.Show{MALID: 4, Title: "Failed Show", Source: "Another Source", SourceID: "failed_id", MainPicture: "https://cdn.myanimelist.net/images/anime/1770/97704.jpg"}
				store.On("Create", show).Return(assert.AnError)
				return store
			}(),
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name: "invalid payload - missing title",
			payload: controller.RegisterShowPayload{
				ID:          5,
				Source:      "Source",
				SourceID:    "missing_title",
				MainPicture: "https://cdn.myanimelist.net/images/anime/1770/97704.jpg",
			},
			mockShowStore:   new(mocks.MockShowStore),
			expectedStatus:  http.StatusBadRequest,
			validationError: true,
		},
		{
			name: "invalid payload - missing main_picture",
			payload: controller.RegisterShowPayload{
				ID:       6,
				Title:    "Title",
				Source:   "Source",
				SourceID: "empty_title",
			},
			mockShowStore:   new(mocks.MockShowStore),
			expectedStatus:  http.StatusBadRequest,
			validationError: true,
		},
		{
			name: "invalid payload - empty main_picture",
			payload: controller.RegisterShowPayload{
				ID:          6,
				Title:       "Title",
				Source:      "Source",
				SourceID:    "empty_title",
				MainPicture: "",
			},
			mockShowStore:   new(mocks.MockShowStore),
			expectedStatus:  http.StatusBadRequest,
			validationError: true,
		},
		{
			name: "invalid payload - non-url main_picture",
			payload: controller.RegisterShowPayload{
				ID:          6,
				Title:       "Title",
				Source:      "Source",
				SourceID:    "empty_title",
				MainPicture: "random string",
			},
			mockShowStore:   new(mocks.MockShowStore),
			expectedStatus:  http.StatusBadRequest,
			validationError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLogger := new(logger.NoOpLogger)
			controller := controller.NewShowController(tt.mockShowStore, mockLogger)

			payloadBytes, _ := json.Marshal(tt.payload)
			rr := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/shows", bytes.NewReader(payloadBytes))
			req.Header.Set("Content-Type", "application/json")

			mux := http.NewServeMux()
			controller.RegisterRoutes(mux)
			mux.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			if tt.expectedStatus == http.StatusBadRequest && !tt.validationError {
				assert.Contains(t, rr.Body.String(), "invalid request payload")
			}
		})
	}
}

func TestShowController_handleUnregister(t *testing.T) {
	tests := []struct {
		name            string
		showID          string
		mockShowStore   *mocks.MockShowStore
		expectedStatus  int
		mockDeleteError error
	}{
		{
			name:   "successful unregistration",
			showID: "7",
			mockShowStore: func() *mocks.MockShowStore {
				store := new(mocks.MockShowStore)
				store.On("Delete", 7).Return(nil)
				return store
			}(),
			expectedStatus: http.StatusNoContent,
		},
		{
			name:   "error deleting show",
			showID: "8",
			mockShowStore: func() *mocks.MockShowStore {
				store := new(mocks.MockShowStore)
				store.On("Delete", 8).Return(assert.AnError)
				return store
			}(),
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "invalid show ID format",
			showID:         "abc",
			mockShowStore:  new(mocks.MockShowStore),
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLogger := new(logger.NoOpLogger)
			controller := controller.NewShowController(tt.mockShowStore, mockLogger)

			rr := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", "/shows/"+tt.showID, nil)

			mux := http.NewServeMux()
			controller.RegisterRoutes(mux)
			mux.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			if tt.expectedStatus == http.StatusBadRequest {
				assert.Contains(t, rr.Body.String(), "invalid show ID")
			}
		})
	}
}
