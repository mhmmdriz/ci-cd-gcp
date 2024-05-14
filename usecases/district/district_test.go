package district

import (
	"errors"
	"lawan-tambang-liar/entities"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDistrictRepository struct {
	mock.Mock
}

func (m *MockDistrictRepository) AddDistrictsFromAPI(districts []entities.District) error {
	args := m.Called(districts)
	return args.Error(0)
}

func (m *MockDistrictRepository) GetAll(regencyID string) ([]entities.District, error) {
	args := m.Called(regencyID)
	return args.Get(0).([]entities.District), args.Error(1)
}

func (m *MockDistrictRepository) GetByID(id string) (entities.District, error) {
	args := m.Called(id)
	return args.Get(0).(entities.District), args.Error(1)
}

type MockDistrictAPI struct {
	mock.Mock
}

func (m *MockDistrictAPI) GetDistrictsDataFromAPI(regencyIDs []string) ([]entities.District, error) {
	args := m.Called(regencyIDs)
	return args.Get(0).([]entities.District), args.Error(1)
}

func TestGetByID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockDistrictRepository)
		mockAPI := new(MockDistrictAPI)
		districtUseCase := NewDistrictUseCase(mockRepo, mockAPI)

		districtID := "1"
		expectedDistrict := entities.District{ID: "1", Name: "District 1"}

		mockRepo.On("GetByID", districtID).Return(expectedDistrict, nil)

		resultDistrict, err := districtUseCase.GetByID(districtID)

		assert.NoError(t, err)
		assert.Equal(t, expectedDistrict, resultDistrict)

		mockRepo.AssertExpectations(t)
	})

	t.Run("FailedRepositoryError", func(t *testing.T) {
		mockRepo := new(MockDistrictRepository)
		mockAPI := new(MockDistrictAPI)
		districtUseCase := NewDistrictUseCase(mockRepo, mockAPI)

		districtID := "1"
		expectedError := errors.New("Repository error")

		mockRepo.On("GetByID", districtID).Return(entities.District{}, expectedError)

		resultDistrict, err := districtUseCase.GetByID(districtID)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, entities.District{}, resultDistrict)

		mockRepo.AssertExpectations(t)
	})
}
