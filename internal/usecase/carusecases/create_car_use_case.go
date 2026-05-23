package carusecases

import (
	"context"

	"github.com/v2code/autolog/internal/persistence"
	"github.com/v2code/autolog/internal/persistence/entities"
)

type CreateCarInput struct {
	Title string
	Type  string
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
		Type:         input.Type,
		KM:           input.KM,
		LogEntries: []*entities.LogEntry{},
	}

	if err := u.carPersistence.CreateCar(ctx, car); err != nil {
		return nil, err
	}
	return &CreateCarResponse{Car: car}, nil
}
