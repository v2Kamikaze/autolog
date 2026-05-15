package carusecases

import (
	"context"

	"github.com/v2code/autolog/internal/persistence"
	"github.com/v2code/autolog/internal/persistence/entities"
)

type CreateCarInput struct {
	Title string
	KM    int
}

type CreateCarResponse struct {
	Car *entities.Car
}

type CreateCarUseCase struct {
	carPersistence persistence.CarPersistence
}

func NewCreateCarUseCase(carPersistence persistence.CarPersistence) *CreateCarUseCase {
	return &CreateCarUseCase{carPersistence: carPersistence}
}

func (u *CreateCarUseCase) Execute(ctx context.Context, input CreateCarInput) (*CreateCarResponse, error) {
	car := &entities.Car{
		Title:        input.Title,
		KM:           input.KM,
		Maintenances: []*entities.Maintenance{},
	}

	if err := u.carPersistence.CreateCar(ctx, car); err != nil {
		return nil, err
	}
	return &CreateCarResponse{Car: car}, nil
}
