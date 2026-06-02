package mocks

import (
	"context"
	"time"

	"github.com/v2code/autolog/internal/persistence"
	"github.com/v2code/autolog/internal/persistence/entities"
)

type VehiclePersistenceMock struct {
	vehicles      []*entities.Vehicle
	nextVehicleID int
	nextLogID     int
}

func NewVehiclePersistenceMock() persistence.VehiclePersistence {
	vehicles := seedVehicles()
	return &VehiclePersistenceMock{
		vehicles:      vehicles,
		nextVehicleID: len(vehicles) + 1,
		nextLogID:     9,
	}
}

func seedVehicles() []*entities.Vehicle {
	return []*entities.Vehicle{
		{
			ID:    1,
			Title: "Etios",
			Type:  "X Plus 2019",
			KM:    82400,
			LogEntries: []*entities.LogEntry{
				{
					ID:        1,
					VehicleID: 1,
					Name:      "Troca de óleo e filtro",
					Type:      entities.LogEntryCategoryMaintenance,
					Date:      time.Date(2025, 2, 10, 0, 0, 0, 0, time.Local),
					KM:        78000,
					Cost:      34000,
				},
				{
					ID:        2,
					VehicleID: 1,
					Name:      "Abastecimento",
					Type:      entities.LogEntryCategoryExpense,
					Date:      time.Date(2025, 3, 2, 0, 0, 0, 0, time.Local),
					KM:        79250,
					Cost:      28000,
				},
			},
		},
		{
			ID:    2,
			Title: "Civic",
			Type:  "LXR 2.0 2014",
			KM:    136900,
			LogEntries: []*entities.LogEntry{
				{
					ID:        3,
					VehicleID: 2,
					Name:      "Pastilhas e fluido de freio",
					Type:      entities.LogEntryCategoryMaintenance,
					Date:      time.Date(2024, 11, 18, 0, 0, 0, 0, time.Local),
					KM:        132200,
					Cost:      69000,
				},
				{
					ID:        4,
					VehicleID: 2,
					Name:      "IPVA",
					Type:      entities.LogEntryCategoryExpense,
					Date:      time.Date(2025, 1, 8, 0, 0, 0, 0, time.Local),
					KM:        133000,
					Cost:      187000,
				},
			},
		},
		{
			ID:    3,
			Title: "Jeep Compass",
			Type:  "Limited Diesel 2021",
			KM:    58700,
			LogEntries: []*entities.LogEntry{
				{
					ID:        5,
					VehicleID: 3,
					Name:      "Revisão de 60 mil",
					Type:      entities.LogEntryCategoryMaintenance,
					Date:      time.Date(2025, 4, 12, 0, 0, 0, 0, time.Local),
					KM:        58000,
					Cost:      124000,
				},
				{
					ID:        6,
					VehicleID: 3,
					Name:      "Seguro anual",
					Type:      entities.LogEntryCategoryExpense,
					Date:      time.Date(2025, 5, 4, 0, 0, 0, 0, time.Local),
					KM:        58400,
					Cost:      352000,
				},
			},
		},
		{
			ID:    4,
			Title: "Corolla",
			Type:  "XEI 2.0 2012",
			KM:    172300,
			LogEntries: []*entities.LogEntry{
				{
					ID:        7,
					VehicleID: 4,
					Name:      "Troca de amortecedores",
					Type:      entities.LogEntryCategoryMaintenance,
					Date:      time.Date(2024, 9, 25, 0, 0, 0, 0, time.Local),
					KM:        168900,
					Cost:      158000,
				},
				{
					ID:        8,
					VehicleID: 4,
					Name:      "Combustível",
					Type:      entities.LogEntryCategoryExpense,
					Date:      time.Date(2024, 10, 3, 0, 0, 0, 0, time.Local),
					KM:        169400,
					Cost:      32000,
				},
			},
		},
	}
}

func (m *VehiclePersistenceMock) CreateVehicle(_ context.Context, vehicle *entities.Vehicle) (*entities.Vehicle, error) {
	if vehicle == nil {
		return nil, persistence.ErrNilVehicle
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
	return nil, persistence.ErrVehicleNotFound
}

func (m *VehiclePersistenceMock) AddLogEntry(ctx context.Context, vehicleID int, logEntry *entities.LogEntry) (*entities.LogEntry, error) {
	if logEntry == nil {
		return nil, persistence.ErrNilLogEntry
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
		return nil, persistence.ErrNilLogEntry
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
	return nil, persistence.ErrLogEntryNotFound
}

func (m *VehiclePersistenceMock) EditVehicle(_ context.Context, patch *entities.Vehicle) (*entities.Vehicle, error) {
	if patch == nil {
		return nil, persistence.ErrNilVehicle
	}
	for _, vehicle := range m.vehicles {
		if vehicle.ID == patch.ID {
			vehicle.Title = patch.Title
			vehicle.Type = patch.Type
			vehicle.KM = patch.KM
			return vehicle, nil
		}
	}
	return nil, persistence.ErrVehicleNotFound
}

func (m *VehiclePersistenceMock) DeleteVehicle(_ context.Context, id int) (*entities.Vehicle, error) {
	for i, vehicle := range m.vehicles {
		if vehicle.ID == id {
			removed := m.vehicles[i]
			m.vehicles = append(m.vehicles[:i], m.vehicles[i+1:]...)
			return removed, nil
		}
	}
	return nil, persistence.ErrVehicleNotFound
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
	return nil, persistence.ErrLogEntryNotFound
}
