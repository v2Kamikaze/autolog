package carusecases

import (
	"context"

	"github.com/v2code/autolog/internal/persistence"
	"github.com/v2code/autolog/internal/persistence/entities"
)

type ListCarsResponse struct {
	Cars []*entities.Car
}

type ListCarsUseCase struct {
	carPersistence persistence.CarPersistence
}

func NewListCarsUseCase(carPersistence persistence.CarPersistence) *ListCarsUseCase {
	return &ListCarsUseCase{carPersistence: carPersistence}
}

func (u *ListCarsUseCase) Execute(ctx context.Context) (*ListCarsResponse, error) {
	cars, err := u.carPersistence.ListCars(ctx)
	if err != nil {
		return nil, err
	}
	return &ListCarsResponse{Cars: cars}, nil
}
