package carusecases

import (
	"context"

	"github.com/v2code/autolog/internal/persistence"
	"github.com/v2code/autolog/internal/persistence/entities"
)

type EditCarInput struct {
	ID    int
	Title string
	Type  string
	KM    int
}

type EditCarResponse struct {
	Car *entities.Car
}

type EditCarUseCase struct {
	carPersistence persistence.CarPersistence
}

func NewEditCarUseCase(carPersistence persistence.CarPersistence) *EditCarUseCase {
	return &EditCarUseCase{
		carPersistence: carPersistence,
	}
}

func (uc *EditCarUseCase) Execute(ctx context.Context, input EditCarInput) (*EditCarResponse, error) {
	patch := &entities.Car{
		ID:    input.ID,
		Title: input.Title,
		Type:  input.Type,
		KM:    input.KM,
	}
	car, err := uc.carPersistence.EditCar(ctx, patch)
	if err != nil {
		return nil, err
	}
	return &EditCarResponse{Car: car}, nil
}
