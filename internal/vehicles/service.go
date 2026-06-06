package vehicles

import (
	"context"
	"errors"
)

type Service struct {
	store *Store
}

func NewService(store *Store) *Service {
	return &Service{store: store}
}

func (s *Service) List(ctx context.Context) ([]*Vehicle, error) {
	return s.store.ListVehicles(ctx, 1, 10000)
}

func (s *Service) Create(ctx context.Context, input CreateVehicleInput) (*Vehicle, error) {
	vehicle := &Vehicle{
		Title:      input.Title,
		Type:       input.Type,
		KM:         input.KM,
		LogEntries: []*LogEntry{},
	}

	vehicle, err := s.store.CreateVehicle(ctx, vehicle)
	if err != nil {
		return nil, err
	}

	return vehicle, nil
}

func (s *Service) Edit(ctx context.Context, input EditVehicleInput) (*Vehicle, error) {
	patch := &Vehicle{
		ID:    input.ID,
		Title: input.Title,
		Type:  input.Type,
		KM:    input.KM,
	}

	return s.store.EditVehicle(ctx, patch)
}

func (s *Service) Delete(ctx context.Context, id int) (*Vehicle, error) {
	return s.store.DeleteVehicle(ctx, id)
}

func (s *Service) AddLog(ctx context.Context, input AddLogInput) (*Vehicle, *LogEntry, error) {
	if !input.Type.Valid() {
		return nil, nil, errors.New("invalid log entry type")
	}

	entry := &LogEntry{
		VehicleID: input.VehicleID,
		Name:      input.Name,
		Type:      input.Type,
		Date:      input.Date,
		KM:        input.KM,
		Cost:      input.Cost,
	}

	inserted, err := s.store.AddLogEntry(ctx, entry)
	if err != nil {
		return nil, nil, err
	}

	vehicle, err := s.store.GetVehicleByID(ctx, input.VehicleID)
	if err != nil {
		return nil, nil, err
	}

	return vehicle, inserted, nil
}

func (s *Service) EditLog(ctx context.Context, input EditLogInput) (*Vehicle, *LogEntry, error) {
	if !input.Type.Valid() {
		return nil, nil, errors.New("invalid log entry type")
	}

	patch := &LogEntry{
		ID:        input.ID,
		VehicleID: input.VehicleID,
		Name:      input.Name,
		Type:      input.Type,
		Date:      input.Date,
		KM:        input.KM,
		Cost:      input.Cost,
	}

	updated, err := s.store.EditLogEntry(ctx, patch)
	if err != nil {
		return nil, nil, err
	}

	vehicle, err := s.store.GetVehicleByID(ctx, patch.VehicleID)
	if err != nil {
		return nil, nil, err
	}

	return vehicle, updated, nil
}

func (s *Service) DeleteLog(ctx context.Context, vehicleID, logEntryID int) (*Vehicle, *LogEntry, error) {
	deleted, err := s.store.DeleteLogEntry(ctx, vehicleID, logEntryID)
	if err != nil {
		return nil, nil, err
	}

	vehicle, err := s.store.GetVehicleByID(ctx, vehicleID)
	if err != nil {
		return nil, nil, err
	}

	return vehicle, deleted, nil
}
