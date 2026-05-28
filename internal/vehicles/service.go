package vehicles

import (
	"context"

	"github.com/v2code/autolog/internal/persistence"
	"github.com/v2code/autolog/internal/persistence/entities"
)

type Service struct {
	repo persistence.VehiclePersistence
}

func NewService(repo persistence.VehiclePersistence) *Service {
	return &Service{repo: repo}
}

func (s *Service) List(ctx context.Context) ([]*entities.Vehicle, error) {
	return s.repo.ListVehicles(ctx)
}

func (s *Service) Create(ctx context.Context, input CreateVehicleInput) (*entities.Vehicle, error) {
	vehicle := &entities.Vehicle{
		Title:      input.Title,
		Type:       input.Type,
		KM:         input.KM,
		LogEntries: []*entities.LogEntry{},
	}

	if err := s.repo.CreateVehicle(ctx, vehicle); err != nil {
		return nil, err
	}

	return vehicle, nil
}

func (s *Service) Edit(ctx context.Context, id int, input EditVehicleInput) (*entities.Vehicle, error) {
	patch := &entities.Vehicle{
		ID:    id,
		Title: input.Title,
		Type:  input.Type,
		KM:    input.KM,
	}

	return s.repo.EditVehicle(ctx, patch)
}

func (s *Service) Delete(ctx context.Context, id int) (*entities.Vehicle, error) {
	return s.repo.DeleteVehicle(ctx, id)
}

func (s *Service) AddLog(ctx context.Context, input AddLogInput) (*entities.Vehicle, *entities.LogEntry, error) {
	if !input.Type.Valid() {
		return nil, nil, persistence.ErrInvalidLogEntryType
	}

	entry := &entities.LogEntry{
		Name: input.Name,
		Type: input.Type,
		Date: input.Date,
		KM:   input.KM,
		Cost: input.Cost,
	}

	inserted, err := s.repo.AddLogEntry(ctx, input.VehicleID, entry)
	if err != nil {
		return nil, nil, err
	}

	vehicle, err := s.repo.GetVehicleByID(ctx, input.VehicleID)
	if err != nil {
		return nil, nil, err
	}

	return vehicle, inserted, nil
}

func (s *Service) EditLog(ctx context.Context, vehicleID, logEntryID int, input EditLogInput) (*entities.Vehicle, *entities.LogEntry, error) {
	if !input.Type.Valid() {
		return nil, nil, persistence.ErrInvalidLogEntryType
	}

	patch := &entities.LogEntry{
		Name: input.Name,
		Type: input.Type,
		Date: input.Date,
		KM:   input.KM,
		Cost: input.Cost,
	}

	updated, err := s.repo.EditLogEntry(ctx, vehicleID, logEntryID, patch)
	if err != nil {
		return nil, nil, err
	}

	vehicle, err := s.repo.GetVehicleByID(ctx, vehicleID)
	if err != nil {
		return nil, nil, err
	}

	return vehicle, updated, nil
}

func (s *Service) DeleteLog(ctx context.Context, vehicleID, logEntryID int) (*entities.Vehicle, *entities.LogEntry, error) {
	deleted, err := s.repo.DeleteLogEntry(ctx, vehicleID, logEntryID)
	if err != nil {
		return nil, nil, err
	}

	vehicle, err := s.repo.GetVehicleByID(ctx, vehicleID)
	if err != nil {
		return nil, nil, err
	}

	return vehicle, deleted, nil
}
