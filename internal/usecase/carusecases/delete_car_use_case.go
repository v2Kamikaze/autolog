package carusecases

import (
	"context"

	"github.com/v2code/autolog/internal/persistence"
	"github.com/v2code/autolog/internal/persistence/entities"
)

type DeleteCarInput struct {
	ID int
}

type DeleteCarResponse struct {
	Car *entities.Car
}

type DeleteCarUseCase struct {
	carPersistence persistence.CarPersistence
}

func NewDeleteCarUseCase(carPersistence persistence.CarPersistence) *DeleteCarUseCase {
	return &DeleteCarUseCase{
		carPersistence: carPersistence,
	}
}

func (u *DeleteCarUseCase) Execute(ctx context.Context, input DeleteCarInput) (*DeleteCarResponse, error) {
	car, err := u.carPersistence.DeleteCar(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	return &DeleteCarResponse{car}, nil
}
