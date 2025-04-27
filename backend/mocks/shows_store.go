package mocks

import (
	"github.com/Gumilho/gumi-fetch/internal/types"
	"github.com/stretchr/testify/mock"
)

// MockShowStore is a mock implementation of the ShowStore interface.
type MockShowStore struct {
	mock.Mock
}

func (m *MockShowStore) Create(show types.Show) error {
	args := m.Called(show)
	return args.Error(0)
}

func (m *MockShowStore) List() ([]types.Show, error) {
	args := m.Called()
	return args.Get(0).([]types.Show), args.Error(1)
}

func (m *MockShowStore) Get(id int) (*types.Show, error) {
	args := m.Called(id)
	return args.Get(0).(*types.Show), args.Error(1)
}

func (m *MockShowStore) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
