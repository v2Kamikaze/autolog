package persistence

import (
	"context"
	"time"

	"github.com/v2code/autolog/internal/persistence/entities"
)

type VehiclePersistence interface {
	CreateVehicle(ctx context.Context, vehicle *entities.Vehicle) error
	ListVehicles(ctx context.Context) ([]*entities.Vehicle, error)
	GetVehicleByID(ctx context.Context, id int) (*entities.Vehicle, error)
	AddLogEntry(ctx context.Context, vehicleID int, logEntry *entities.LogEntry) (*entities.LogEntry, error)
	EditLogEntry(ctx context.Context, vehicleID int, logEntryID int, logEntry *entities.LogEntry) (*entities.LogEntry, error)
	EditVehicle(ctx context.Context, vehicle *entities.Vehicle) (*entities.Vehicle, error)
	DeleteVehicle(ctx context.Context, id int) (*entities.Vehicle, error)
	DeleteLogEntry(ctx context.Context, vehicleID int, logEntryID int) (*entities.LogEntry, error)
}

type vehiclePersistence struct {
	vehicles         []*entities.Vehicle
	nextVehicleID         int
	nextLogEntryID int
}

func NewInMemoryVehiclePersistence() VehiclePersistence {
	vehicles := seedVehicles()

	p := &vehiclePersistence{
		vehicles:         vehicles,
		nextVehicleID:    1,
		nextLogEntryID:   1,
	}

	for _, vehicle := range vehicles {
		if vehicle.ID >= p.nextVehicleID {
			p.nextVehicleID = vehicle.ID + 1
		}
		for _, logEntry := range vehicle.LogEntries {
			if logEntry.ID >= p.nextLogEntryID {
				p.nextLogEntryID = logEntry.ID + 1
			}
		}
	}

	return p
}

func seedVehicles() []*entities.Vehicle {
	return []*entities.Vehicle{
		{
			ID:    1,
			Title: "Honda Civic EXL 2020 2.0 Flex Automático",
			Type:  "sedan",
			KM:    28000,
			LogEntries: []*entities.LogEntry{
				{ID: 1, VehicleID: 1, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de óleo", Date: time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC), KM: 5000, Cost: 30000},
				{ID: 2, VehicleID: 1, Type: entities.LogEntryCategoryMaintenance, Name: "Filtro de ar", Date: time.Date(2024, time.April, 10, 0, 0, 0, 0, time.UTC), KM: 10000, Cost: 12000},
				{ID: 3, VehicleID: 1, Type: entities.LogEntryCategoryMaintenance, Name: "Alinhamento", Date: time.Date(2024, time.July, 10, 0, 0, 0, 0, time.UTC), KM: 15000, Cost: 15000},
				{ID: 4, VehicleID: 1, Type: entities.LogEntryCategoryMaintenance, Name: "Pastilhas de freio", Date: time.Date(2024, time.October, 10, 0, 0, 0, 0, time.UTC), KM: 22000, Cost: 40000},
				{ID: 5, VehicleID: 1, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de pneus", Date: time.Date(2025, time.January, 10, 0, 0, 0, 0, time.UTC), KM: 28000, Cost: 140000},
			},
		},
		{
			ID:    2,
			Title: "Toyota Corolla XEi 2021 2.0 Flex Automático",
			Type:  "sedan",
			KM:    22000,
			LogEntries: []*entities.LogEntry{
				{ID: 6, VehicleID: 2, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de óleo", Date: time.Date(2024, time.February, 5, 0, 0, 0, 0, time.UTC), KM: 4000, Cost: 28000},
				{ID: 7, VehicleID: 2, Type: entities.LogEntryCategoryMaintenance, Name: "Filtro de óleo", Date: time.Date(2024, time.May, 5, 0, 0, 0, 0, time.UTC), KM: 8000, Cost: 9000},
				{ID: 8, VehicleID: 2, Type: entities.LogEntryCategoryMaintenance, Name: "Balanceamento", Date: time.Date(2024, time.August, 5, 0, 0, 0, 0, time.UTC), KM: 12000, Cost: 12000},
				{ID: 9, VehicleID: 2, Type: entities.LogEntryCategoryMaintenance, Name: "Bieletas", Date: time.Date(2024, time.November, 5, 0, 0, 0, 0, time.UTC), KM: 17000, Cost: 25000},
				{ID: 10, VehicleID: 2, Type: entities.LogEntryCategoryMaintenance, Name: "Amortecedores", Date: time.Date(2025, time.February, 5, 0, 0, 0, 0, time.UTC), KM: 22000, Cost: 180000},
			},
		},
		{
			ID:    3,
			Title: "Nissan Kicks SV 2022 1.6 Flex Automático",
			Type:  "suv",
			KM:    15000,
			LogEntries: []*entities.LogEntry{
				{ID: 11, VehicleID: 3, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de óleo", Date: time.Date(2025, time.January, 15, 0, 0, 0, 0, time.UTC), KM: 3000, Cost: 26000},
				{ID: 12, VehicleID: 3, Type: entities.LogEntryCategoryMaintenance, Name: "Filtro de ar", Date: time.Date(2025, time.April, 15, 0, 0, 0, 0, time.UTC), KM: 6000, Cost: 11000},
				{ID: 13, VehicleID: 3, Type: entities.LogEntryCategoryMaintenance, Name: "Alinhamento", Date: time.Date(2025, time.July, 15, 0, 0, 0, 0, time.UTC), KM: 9000, Cost: 13000},
				{ID: 14, VehicleID: 3, Type: entities.LogEntryCategoryMaintenance, Name: "Pastilhas de freio", Date: time.Date(2025, time.October, 15, 0, 0, 0, 0, time.UTC), KM: 12000, Cost: 35000},
				{ID: 15, VehicleID: 3, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de pneus", Date: time.Date(2026, time.January, 15, 0, 0, 0, 0, time.UTC), KM: 15000, Cost: 120000},
			},
		},
		{
			ID:    4,
			Title: "Hyundai i30 2019 2.0 Automático",
			Type:  "hatch",
			KM:    35000,
			LogEntries: []*entities.LogEntry{
				{ID: 16, VehicleID: 4, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de óleo", Date: time.Date(2024, time.January, 20, 0, 0, 0, 0, time.UTC), KM: 8000, Cost: 27000},
				{ID: 17, VehicleID: 4, Type: entities.LogEntryCategoryMaintenance, Name: "Filtro de óleo", Date: time.Date(2024, time.April, 20, 0, 0, 0, 0, time.UTC), KM: 15000, Cost: 9000},
				{ID: 18, VehicleID: 4, Type: entities.LogEntryCategoryMaintenance, Name: "Alinhamento", Date: time.Date(2024, time.July, 20, 0, 0, 0, 0, time.UTC), KM: 20000, Cost: 14000},
				{ID: 19, VehicleID: 4, Type: entities.LogEntryCategoryMaintenance, Name: "Bieletas", Date: time.Date(2024, time.October, 20, 0, 0, 0, 0, time.UTC), KM: 28000, Cost: 25000},
				{ID: 20, VehicleID: 4, Type: entities.LogEntryCategoryMaintenance, Name: "Amortecedores", Date: time.Date(2025, time.January, 20, 0, 0, 0, 0, time.UTC), KM: 35000, Cost: 190000},
			},
		},
		{
			ID:    5,
			Title: "Hyundai Tucson 2021 1.6 Turbo Automático",
			Type:  "suv",
			KM:    20000,
			LogEntries: []*entities.LogEntry{
				{ID: 21, VehicleID: 5, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de óleo", Date: time.Date(2025, time.February, 25, 0, 0, 0, 0, time.UTC), KM: 4000, Cost: 30000},
				{ID: 22, VehicleID: 5, Type: entities.LogEntryCategoryMaintenance, Name: "Filtro de ar", Date: time.Date(2025, time.May, 25, 0, 0, 0, 0, time.UTC), KM: 8000, Cost: 12000},
				{ID: 23, VehicleID: 5, Type: entities.LogEntryCategoryMaintenance, Name: "Balanceamento", Date: time.Date(2025, time.August, 25, 0, 0, 0, 0, time.UTC), KM: 12000, Cost: 13000},
				{ID: 24, VehicleID: 5, Type: entities.LogEntryCategoryMaintenance, Name: "Pastilhas de freio", Date: time.Date(2025, time.November, 25, 0, 0, 0, 0, time.UTC), KM: 16000, Cost: 40000},
				{ID: 25, VehicleID: 5, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de pneus", Date: time.Date(2026, time.February, 25, 0, 0, 0, 0, time.UTC), KM: 20000, Cost: 150000},
			},
		},
		{
			ID:    6,
			Title: "Honda Fit EX 2018 1.5 Flex Automático",
			Type:  "hatch",
			KM:    45000,
			LogEntries: []*entities.LogEntry{
				{ID: 26, VehicleID: 6, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de óleo", Date: time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC), KM: 10000, Cost: 25000},
				{ID: 27, VehicleID: 6, Type: entities.LogEntryCategoryMaintenance, Name: "Filtro de ar", Date: time.Date(2024, time.April, 10, 0, 0, 0, 0, time.UTC), KM: 20000, Cost: 10000},
				{ID: 28, VehicleID: 6, Type: entities.LogEntryCategoryMaintenance, Name: "Alinhamento", Date: time.Date(2024, time.July, 10, 0, 0, 0, 0, time.UTC), KM: 30000, Cost: 12000},
				{ID: 29, VehicleID: 6, Type: entities.LogEntryCategoryMaintenance, Name: "Pastilhas de freio", Date: time.Date(2024, time.October, 10, 0, 0, 0, 0, time.UTC), KM: 38000, Cost: 30000},
				{ID: 30, VehicleID: 6, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de pneus", Date: time.Date(2025, time.January, 10, 0, 0, 0, 0, time.UTC), KM: 45000, Cost: 120000},
			},
		},
		{
			ID:    7,
			Title: "Toyota Yaris XS 2020 1.5 Flex Automático",
			Type:  "hatch",
			KM:    32000,
			LogEntries: []*entities.LogEntry{
				{ID: 31, VehicleID: 7, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de óleo", Date: time.Date(2024, time.February, 5, 0, 0, 0, 0, time.UTC), KM: 8000, Cost: 26000},
				{ID: 32, VehicleID: 7, Type: entities.LogEntryCategoryMaintenance, Name: "Filtro de óleo", Date: time.Date(2024, time.May, 5, 0, 0, 0, 0, time.UTC), KM: 15000, Cost: 9000},
				{ID: 33, VehicleID: 7, Type: entities.LogEntryCategoryMaintenance, Name: "Balanceamento", Date: time.Date(2024, time.August, 5, 0, 0, 0, 0, time.UTC), KM: 20000, Cost: 12000},
				{ID: 34, VehicleID: 7, Type: entities.LogEntryCategoryMaintenance, Name: "Bieletas", Date: time.Date(2024, time.November, 5, 0, 0, 0, 0, time.UTC), KM: 26000, Cost: 22000},
				{ID: 35, VehicleID: 7, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de pneus", Date: time.Date(2025, time.February, 5, 0, 0, 0, 0, time.UTC), KM: 32000, Cost: 130000},
			},
		},
		{
			ID:    8,
			Title: "Nissan Versa SV 2019 1.6 Flex Automático",
			Type:  "sedan",
			KM:    60000,
			LogEntries: []*entities.LogEntry{
				{ID: 36, VehicleID: 8, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de óleo", Date: time.Date(2023, time.January, 15, 0, 0, 0, 0, time.UTC), KM: 15000, Cost: 25000},
				{ID: 37, VehicleID: 8, Type: entities.LogEntryCategoryMaintenance, Name: "Filtro de ar", Date: time.Date(2023, time.June, 15, 0, 0, 0, 0, time.UTC), KM: 25000, Cost: 10000},
				{ID: 38, VehicleID: 8, Type: entities.LogEntryCategoryMaintenance, Name: "Alinhamento", Date: time.Date(2023, time.November, 15, 0, 0, 0, 0, time.UTC), KM: 35000, Cost: 13000},
				{ID: 39, VehicleID: 8, Type: entities.LogEntryCategoryMaintenance, Name: "Pastilhas de freio", Date: time.Date(2024, time.April, 15, 0, 0, 0, 0, time.UTC), KM: 45000, Cost: 32000},
				{ID: 40, VehicleID: 8, Type: entities.LogEntryCategoryMaintenance, Name: "Amortecedores", Date: time.Date(2024, time.December, 15, 0, 0, 0, 0, time.UTC), KM: 60000, Cost: 180000},
			},
		},
		{
			ID:    9,
			Title: "Mitsubishi Lancer HL 2016 2.0 Automático",
			Type:  "sedan",
			KM:    70000,
			LogEntries: []*entities.LogEntry{
				{ID: 41, VehicleID: 9, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de óleo", Date: time.Date(2023, time.February, 20, 0, 0, 0, 0, time.UTC), KM: 20000, Cost: 28000},
				{ID: 42, VehicleID: 9, Type: entities.LogEntryCategoryMaintenance, Name: "Filtro de óleo", Date: time.Date(2023, time.July, 20, 0, 0, 0, 0, time.UTC), KM: 30000, Cost: 9000},
				{ID: 43, VehicleID: 9, Type: entities.LogEntryCategoryMaintenance, Name: "Alinhamento", Date: time.Date(2024, time.January, 20, 0, 0, 0, 0, time.UTC), KM: 45000, Cost: 15000},
				{ID: 44, VehicleID: 9, Type: entities.LogEntryCategoryMaintenance, Name: "Pastilhas de freio", Date: time.Date(2024, time.June, 20, 0, 0, 0, 0, time.UTC), KM: 55000, Cost: 40000},
				{ID: 45, VehicleID: 9, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de pneus", Date: time.Date(2025, time.January, 20, 0, 0, 0, 0, time.UTC), KM: 70000, Cost: 150000},
			},
		},
		{
			ID:    10,
			Title: "Suzuki Jimny 4Sport 2018 1.3 4x4",
			Type:  "suv",
			KM:    50000,
			LogEntries: []*entities.LogEntry{
				{ID: 46, VehicleID: 10, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de óleo", Date: time.Date(2023, time.March, 10, 0, 0, 0, 0, time.UTC), KM: 10000, Cost: 26000},
				{ID: 47, VehicleID: 10, Type: entities.LogEntryCategoryMaintenance, Name: "Filtro de ar", Date: time.Date(2023, time.August, 10, 0, 0, 0, 0, time.UTC), KM: 20000, Cost: 11000},
				{ID: 48, VehicleID: 10, Type: entities.LogEntryCategoryMaintenance, Name: "Alinhamento", Date: time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC), KM: 30000, Cost: 14000},
				{ID: 49, VehicleID: 10, Type: entities.LogEntryCategoryMaintenance, Name: "Bieletas", Date: time.Date(2024, time.June, 10, 0, 0, 0, 0, time.UTC), KM: 40000, Cost: 25000},
				{ID: 50, VehicleID: 10, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de pneus", Date: time.Date(2025, time.January, 10, 0, 0, 0, 0, time.UTC), KM: 50000, Cost: 140000},
			},
		},
	}
}

func (p *vehiclePersistence) ListVehicles(_ context.Context) ([]*entities.Vehicle, error) {
	return p.vehicles, nil
}

func (p *vehiclePersistence) GetVehicleByID(_ context.Context, id int) (*entities.Vehicle, error) {
	for _, c := range p.vehicles {
		if c.ID == id {
			return c, nil
		}
	}
	return nil, ErrVehicleNotFound
}

func (p *vehiclePersistence) CreateVehicle(_ context.Context, vehicle *entities.Vehicle) error {
	if vehicle == nil {
		return ErrNilVehicle
	}
	vehicle.ID = p.nextVehicleID
	p.nextVehicleID++
	p.vehicles = append(p.vehicles, vehicle)
	return nil
}

func (p *vehiclePersistence) AddLogEntry(_ context.Context, vehicleID int, logEntry *entities.LogEntry) (*entities.LogEntry, error) {
	if logEntry == nil {
		return nil, ErrNilLogEntry
	}

	var vehicle *entities.Vehicle
	for _, c := range p.vehicles {
		if c.ID == vehicleID {
			vehicle = c
			break
		}
	}
	if vehicle == nil {
		return nil, ErrVehicleNotFound
	}

	logEntry.ID = p.nextLogEntryID
	logEntry.VehicleID = vehicleID
	p.nextLogEntryID++

	vehicle.LogEntries = append(vehicle.LogEntries, logEntry)

	return logEntry, nil
}

func (p *vehiclePersistence) EditLogEntry(_ context.Context, vehicleID int, logEntryID int, patch *entities.LogEntry) (*entities.LogEntry, error) {
	if patch == nil {
		return nil, ErrNilLogEntry
	}

	var vehicle *entities.Vehicle
	for _, c := range p.vehicles {
		if c.ID == vehicleID {
			vehicle = c
			break
		}
	}
	if vehicle == nil {
		return nil, ErrVehicleNotFound
	}

	for _, m := range vehicle.LogEntries {
		if m.ID != logEntryID {
			continue
		}
		m.Name = patch.Name
		m.Type = patch.Type
		m.Date = patch.Date
		m.KM = patch.KM
		m.Cost = patch.Cost
		m.VehicleID = vehicleID
		return m, nil
	}
	return nil, ErrLogEntryNotFound
}

func (p *vehiclePersistence) EditVehicle(_ context.Context, patch *entities.Vehicle) (*entities.Vehicle, error) {
	if patch == nil {
		return nil, ErrNilVehicle
	}

	var vehicle *entities.Vehicle
	for _, c := range p.vehicles {
		if c.ID == patch.ID {
			vehicle = c
			break
		}
	}
	if vehicle == nil {
		return nil, ErrVehicleNotFound
	}

	vehicle.Title = patch.Title
	vehicle.Type = patch.Type
	vehicle.KM = patch.KM

	return vehicle, nil
}

func (p *vehiclePersistence) DeleteVehicle(_ context.Context, id int) (*entities.Vehicle, error) {
	var deletedVehicle *entities.Vehicle
	for i, c := range p.vehicles {
		if c.ID == id {
			deletedVehicle = p.vehicles[i]
			p.vehicles = append(p.vehicles[:i], p.vehicles[i+1:]...)
			return deletedVehicle, nil
		}
	}
	return nil, ErrVehicleNotFound
}

func (p *vehiclePersistence) DeleteLogEntry(_ context.Context, vehicleID int, logEntryID int) (*entities.LogEntry, error) {
	var vehicle *entities.Vehicle
	for _, c := range p.vehicles {
		if c.ID == vehicleID {
			vehicle = c
			break
		}
	}
	if vehicle == nil {
		return nil, ErrVehicleNotFound
	}

	var deletedLogEntry *entities.LogEntry
	for i, m := range vehicle.LogEntries {
		if m.ID == logEntryID {
			deletedLogEntry = vehicle.LogEntries[i]
			vehicle.LogEntries = append(vehicle.LogEntries[:i], vehicle.LogEntries[i+1:]...)
			return deletedLogEntry, nil
		}
	}
	return nil, ErrLogEntryNotFound
}
