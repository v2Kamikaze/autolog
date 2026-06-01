package mocks

import (
	"context"

	"github.com/v2code/autolog/internal/persistence"
	"github.com/v2code/autolog/internal/persistence/entities"
)

type VehiclePersistenceMock struct {
	vehicles      []*entities.Vehicle
	nextVehicleID int
	nextLogID     int
}

func NewVehiclePersistenceMock() persistence.VehiclePersistence {
	return &VehiclePersistenceMock{
		vehicles:      []*entities.Vehicle{},
		nextVehicleID: 1,
		nextLogID:     1,
	}
}

func (m *VehiclePersistenceMock) CreateVehicle(_ context.Context, vehicle *entities.Vehicle) (*entities.Vehicle, error) {
	if vehicle == nil {
		return nil, ErrNilVehicle
	}
	vehicle.ID = m.nextVehicleID
	m.nextVehicleID++
	m.vehicles = append(m.vehicles, vehicle)
	return vehicle, nil
}

func (m *VehiclePersistenceMock) ListVehicles(_ context.Context) ([]*entities.Vehicle, error) {
	return m.vehicles, nil
}

func (m *VehiclePersistenceMock) GetVehicleByID(_ context.Context, id int) (*entities.Vehicle, error) {
	for _, vehicle := range m.vehicles {
		if vehicle.ID == id {
			return vehicle, nil
		}
	}
	return nil, ErrVehicleNotFound
}

func (m *VehiclePersistenceMock) AddLogEntry(ctx context.Context, vehicleID int, logEntry *entities.LogEntry) (*entities.LogEntry, error) {
	if logEntry == nil {
		return nil, ErrNilLogEntry
	}
	vehicle, err := m.GetVehicleByID(ctx, vehicleID)
	if err != nil {
		return nil, err
	}

	logEntry.ID = m.nextLogID
	logEntry.VehicleID = vehicleID
	m.nextLogID++

	vehicle.LogEntries = append(vehicle.LogEntries, logEntry)
	return logEntry, nil
}

func (m *VehiclePersistenceMock) EditLogEntry(ctx context.Context, vehicleID int, logEntryID int, patch *entities.LogEntry) (*entities.LogEntry, error) {
	if patch == nil {
		return nil, ErrNilLogEntry
	}
	vehicle, err := m.GetVehicleByID(ctx, vehicleID)
	if err != nil {
		return nil, err
	}
	for _, entry := range vehicle.LogEntries {
		if entry.ID == logEntryID {
			entry.Name = patch.Name
			entry.Type = patch.Type
			entry.Date = patch.Date
			entry.KM = patch.KM
			entry.Cost = patch.Cost
			entry.VehicleID = vehicleID
			return entry, nil
		}
	}
	return nil, ErrLogEntryNotFound
}

func (m *VehiclePersistenceMock) EditVehicle(_ context.Context, patch *entities.Vehicle) (*entities.Vehicle, error) {
	if patch == nil {
		return nil, ErrNilVehicle
	}
	for _, vehicle := range m.vehicles {
		if vehicle.ID == patch.ID {
			vehicle.Title = patch.Title
			vehicle.Type = patch.Type
			vehicle.KM = patch.KM
			return vehicle, nil
		}
	}
	return nil, ErrVehicleNotFound
}

func (m *VehiclePersistenceMock) DeleteVehicle(_ context.Context, id int) (*entities.Vehicle, error) {
	for i, vehicle := range m.vehicles {
		if vehicle.ID == id {
			removed := m.vehicles[i]
			m.vehicles = append(m.vehicles[:i], m.vehicles[i+1:]...)
			return removed, nil
		}
	}
	return nil, ErrVehicleNotFound
}

func (m *VehiclePersistenceMock) DeleteLogEntry(ctx context.Context, vehicleID int, logEntryID int) (*entities.LogEntry, error) {
	vehicle, err := m.GetVehicleByID(ctx, vehicleID)
	if err != nil {
		return nil, err
	}
	for i, entry := range vehicle.LogEntries {
		if entry.ID == logEntryID {
			removed := vehicle.LogEntries[i]
			vehicle.LogEntries = append(vehicle.LogEntries[:i], vehicle.LogEntries[i+1:]...)
			return removed, nil
		}
	}
	return nil, ErrLogEntryNotFound
}
