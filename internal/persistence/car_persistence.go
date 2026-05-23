package persistence

import (
	"context"
	"time"

	"github.com/v2code/autolog/internal/persistence/entities"
)

type CarPersistence interface {
	CreateCar(ctx context.Context, car *entities.Car) error
	ListCars(ctx context.Context) ([]*entities.Car, error)
	GetCarByID(ctx context.Context, id int) (*entities.Car, error)
	AddLogEntry(ctx context.Context, carID int, logEntry *entities.LogEntry) (*entities.LogEntry, error)
	EditLogEntry(ctx context.Context, carID int, logEntryID int, logEntry *entities.LogEntry) (*entities.LogEntry, error)
	EditCar(ctx context.Context, car *entities.Car) (*entities.Car, error)
	DeleteCar(ctx context.Context, id int) (*entities.Car, error)
	DeleteLogEntry(ctx context.Context, carID int, logEntryID int) (*entities.LogEntry, error)
}

type carPersistence struct {
	cars              []*entities.Car
	nextCarID         int
	nextLogEntryID int
}

func NewInMemoryCarPersistence() CarPersistence {
	cars := seedCars()

	p := &carPersistence{
		cars:              cars,
		nextCarID:         1,
		nextLogEntryID: 1,
	}

	for _, car := range cars {
		if car.ID >= p.nextCarID {
			p.nextCarID = car.ID + 1
		}
		for _, logEntry := range car.LogEntries {
			if logEntry.ID >= p.nextLogEntryID {
				p.nextLogEntryID = logEntry.ID + 1
			}
		}
	}

	return p
}

func seedCars() []*entities.Car {
	return []*entities.Car{
		{
			ID:    1,
			Title: "Honda Civic EXL 2020 2.0 Flex Automático",
			Type:  "sedan",
			KM:    28000,
			LogEntries: []*entities.LogEntry{
				{ID: 1, CarID: 1, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de óleo", Date: time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC), KM: 5000, Cost: 30000},
				{ID: 2, CarID: 1, Type: entities.LogEntryCategoryMaintenance, Name: "Filtro de ar", Date: time.Date(2024, time.April, 10, 0, 0, 0, 0, time.UTC), KM: 10000, Cost: 12000},
				{ID: 3, CarID: 1, Type: entities.LogEntryCategoryMaintenance, Name: "Alinhamento", Date: time.Date(2024, time.July, 10, 0, 0, 0, 0, time.UTC), KM: 15000, Cost: 15000},
				{ID: 4, CarID: 1, Type: entities.LogEntryCategoryMaintenance, Name: "Pastilhas de freio", Date: time.Date(2024, time.October, 10, 0, 0, 0, 0, time.UTC), KM: 22000, Cost: 40000},
				{ID: 5, CarID: 1, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de pneus", Date: time.Date(2025, time.January, 10, 0, 0, 0, 0, time.UTC), KM: 28000, Cost: 140000},
			},
		},
		{
			ID:    2,
			Title: "Toyota Corolla XEi 2021 2.0 Flex Automático",
			Type:  "sedan",
			KM:    22000,
			LogEntries: []*entities.LogEntry{
				{ID: 6, CarID: 2, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de óleo", Date: time.Date(2024, time.February, 5, 0, 0, 0, 0, time.UTC), KM: 4000, Cost: 28000},
				{ID: 7, CarID: 2, Type: entities.LogEntryCategoryMaintenance, Name: "Filtro de óleo", Date: time.Date(2024, time.May, 5, 0, 0, 0, 0, time.UTC), KM: 8000, Cost: 9000},
				{ID: 8, CarID: 2, Type: entities.LogEntryCategoryMaintenance, Name: "Balanceamento", Date: time.Date(2024, time.August, 5, 0, 0, 0, 0, time.UTC), KM: 12000, Cost: 12000},
				{ID: 9, CarID: 2, Type: entities.LogEntryCategoryMaintenance, Name: "Bieletas", Date: time.Date(2024, time.November, 5, 0, 0, 0, 0, time.UTC), KM: 17000, Cost: 25000},
				{ID: 10, CarID: 2, Type: entities.LogEntryCategoryMaintenance, Name: "Amortecedores", Date: time.Date(2025, time.February, 5, 0, 0, 0, 0, time.UTC), KM: 22000, Cost: 180000},
			},
		},
		{
			ID:    3,
			Title: "Nissan Kicks SV 2022 1.6 Flex Automático",
			Type:  "suv",
			KM:    15000,
			LogEntries: []*entities.LogEntry{
				{ID: 11, CarID: 3, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de óleo", Date: time.Date(2025, time.January, 15, 0, 0, 0, 0, time.UTC), KM: 3000, Cost: 26000},
				{ID: 12, CarID: 3, Type: entities.LogEntryCategoryMaintenance, Name: "Filtro de ar", Date: time.Date(2025, time.April, 15, 0, 0, 0, 0, time.UTC), KM: 6000, Cost: 11000},
				{ID: 13, CarID: 3, Type: entities.LogEntryCategoryMaintenance, Name: "Alinhamento", Date: time.Date(2025, time.July, 15, 0, 0, 0, 0, time.UTC), KM: 9000, Cost: 13000},
				{ID: 14, CarID: 3, Type: entities.LogEntryCategoryMaintenance, Name: "Pastilhas de freio", Date: time.Date(2025, time.October, 15, 0, 0, 0, 0, time.UTC), KM: 12000, Cost: 35000},
				{ID: 15, CarID: 3, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de pneus", Date: time.Date(2026, time.January, 15, 0, 0, 0, 0, time.UTC), KM: 15000, Cost: 120000},
			},
		},
		{
			ID:    4,
			Title: "Hyundai i30 2019 2.0 Automático",
			Type:  "hatch",
			KM:    35000,
			LogEntries: []*entities.LogEntry{
				{ID: 16, CarID: 4, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de óleo", Date: time.Date(2024, time.January, 20, 0, 0, 0, 0, time.UTC), KM: 8000, Cost: 27000},
				{ID: 17, CarID: 4, Type: entities.LogEntryCategoryMaintenance, Name: "Filtro de óleo", Date: time.Date(2024, time.April, 20, 0, 0, 0, 0, time.UTC), KM: 15000, Cost: 9000},
				{ID: 18, CarID: 4, Type: entities.LogEntryCategoryMaintenance, Name: "Alinhamento", Date: time.Date(2024, time.July, 20, 0, 0, 0, 0, time.UTC), KM: 20000, Cost: 14000},
				{ID: 19, CarID: 4, Type: entities.LogEntryCategoryMaintenance, Name: "Bieletas", Date: time.Date(2024, time.October, 20, 0, 0, 0, 0, time.UTC), KM: 28000, Cost: 25000},
				{ID: 20, CarID: 4, Type: entities.LogEntryCategoryMaintenance, Name: "Amortecedores", Date: time.Date(2025, time.January, 20, 0, 0, 0, 0, time.UTC), KM: 35000, Cost: 190000},
			},
		},
		{
			ID:    5,
			Title: "Hyundai Tucson 2021 1.6 Turbo Automático",
			Type:  "suv",
			KM:    20000,
			LogEntries: []*entities.LogEntry{
				{ID: 21, CarID: 5, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de óleo", Date: time.Date(2025, time.February, 25, 0, 0, 0, 0, time.UTC), KM: 4000, Cost: 30000},
				{ID: 22, CarID: 5, Type: entities.LogEntryCategoryMaintenance, Name: "Filtro de ar", Date: time.Date(2025, time.May, 25, 0, 0, 0, 0, time.UTC), KM: 8000, Cost: 12000},
				{ID: 23, CarID: 5, Type: entities.LogEntryCategoryMaintenance, Name: "Balanceamento", Date: time.Date(2025, time.August, 25, 0, 0, 0, 0, time.UTC), KM: 12000, Cost: 13000},
				{ID: 24, CarID: 5, Type: entities.LogEntryCategoryMaintenance, Name: "Pastilhas de freio", Date: time.Date(2025, time.November, 25, 0, 0, 0, 0, time.UTC), KM: 16000, Cost: 40000},
				{ID: 25, CarID: 5, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de pneus", Date: time.Date(2026, time.February, 25, 0, 0, 0, 0, time.UTC), KM: 20000, Cost: 150000},
			},
		},
		{
			ID:    6,
			Title: "Honda Fit EX 2018 1.5 Flex Automático",
			Type:  "hatch",
			KM:    45000,
			LogEntries: []*entities.LogEntry{
				{ID: 26, CarID: 6, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de óleo", Date: time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC), KM: 10000, Cost: 25000},
				{ID: 27, CarID: 6, Type: entities.LogEntryCategoryMaintenance, Name: "Filtro de ar", Date: time.Date(2024, time.April, 10, 0, 0, 0, 0, time.UTC), KM: 20000, Cost: 10000},
				{ID: 28, CarID: 6, Type: entities.LogEntryCategoryMaintenance, Name: "Alinhamento", Date: time.Date(2024, time.July, 10, 0, 0, 0, 0, time.UTC), KM: 30000, Cost: 12000},
				{ID: 29, CarID: 6, Type: entities.LogEntryCategoryMaintenance, Name: "Pastilhas de freio", Date: time.Date(2024, time.October, 10, 0, 0, 0, 0, time.UTC), KM: 38000, Cost: 30000},
				{ID: 30, CarID: 6, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de pneus", Date: time.Date(2025, time.January, 10, 0, 0, 0, 0, time.UTC), KM: 45000, Cost: 120000},
			},
		},
		{
			ID:    7,
			Title: "Toyota Yaris XS 2020 1.5 Flex Automático",
			Type:  "hatch",
			KM:    32000,
			LogEntries: []*entities.LogEntry{
				{ID: 31, CarID: 7, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de óleo", Date: time.Date(2024, time.February, 5, 0, 0, 0, 0, time.UTC), KM: 8000, Cost: 26000},
				{ID: 32, CarID: 7, Type: entities.LogEntryCategoryMaintenance, Name: "Filtro de óleo", Date: time.Date(2024, time.May, 5, 0, 0, 0, 0, time.UTC), KM: 15000, Cost: 9000},
				{ID: 33, CarID: 7, Type: entities.LogEntryCategoryMaintenance, Name: "Balanceamento", Date: time.Date(2024, time.August, 5, 0, 0, 0, 0, time.UTC), KM: 20000, Cost: 12000},
				{ID: 34, CarID: 7, Type: entities.LogEntryCategoryMaintenance, Name: "Bieletas", Date: time.Date(2024, time.November, 5, 0, 0, 0, 0, time.UTC), KM: 26000, Cost: 22000},
				{ID: 35, CarID: 7, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de pneus", Date: time.Date(2025, time.February, 5, 0, 0, 0, 0, time.UTC), KM: 32000, Cost: 130000},
			},
		},
		{
			ID:    8,
			Title: "Nissan Versa SV 2019 1.6 Flex Automático",
			Type:  "sedan",
			KM:    60000,
			LogEntries: []*entities.LogEntry{
				{ID: 36, CarID: 8, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de óleo", Date: time.Date(2023, time.January, 15, 0, 0, 0, 0, time.UTC), KM: 15000, Cost: 25000},
				{ID: 37, CarID: 8, Type: entities.LogEntryCategoryMaintenance, Name: "Filtro de ar", Date: time.Date(2023, time.June, 15, 0, 0, 0, 0, time.UTC), KM: 25000, Cost: 10000},
				{ID: 38, CarID: 8, Type: entities.LogEntryCategoryMaintenance, Name: "Alinhamento", Date: time.Date(2023, time.November, 15, 0, 0, 0, 0, time.UTC), KM: 35000, Cost: 13000},
				{ID: 39, CarID: 8, Type: entities.LogEntryCategoryMaintenance, Name: "Pastilhas de freio", Date: time.Date(2024, time.April, 15, 0, 0, 0, 0, time.UTC), KM: 45000, Cost: 32000},
				{ID: 40, CarID: 8, Type: entities.LogEntryCategoryMaintenance, Name: "Amortecedores", Date: time.Date(2024, time.December, 15, 0, 0, 0, 0, time.UTC), KM: 60000, Cost: 180000},
			},
		},
		{
			ID:    9,
			Title: "Mitsubishi Lancer HL 2016 2.0 Automático",
			Type:  "sedan",
			KM:    70000,
			LogEntries: []*entities.LogEntry{
				{ID: 41, CarID: 9, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de óleo", Date: time.Date(2023, time.February, 20, 0, 0, 0, 0, time.UTC), KM: 20000, Cost: 28000},
				{ID: 42, CarID: 9, Type: entities.LogEntryCategoryMaintenance, Name: "Filtro de óleo", Date: time.Date(2023, time.July, 20, 0, 0, 0, 0, time.UTC), KM: 30000, Cost: 9000},
				{ID: 43, CarID: 9, Type: entities.LogEntryCategoryMaintenance, Name: "Alinhamento", Date: time.Date(2024, time.January, 20, 0, 0, 0, 0, time.UTC), KM: 45000, Cost: 15000},
				{ID: 44, CarID: 9, Type: entities.LogEntryCategoryMaintenance, Name: "Pastilhas de freio", Date: time.Date(2024, time.June, 20, 0, 0, 0, 0, time.UTC), KM: 55000, Cost: 40000},
				{ID: 45, CarID: 9, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de pneus", Date: time.Date(2025, time.January, 20, 0, 0, 0, 0, time.UTC), KM: 70000, Cost: 150000},
			},
		},
		{
			ID:    10,
			Title: "Suzuki Jimny 4Sport 2018 1.3 4x4",
			Type:  "suv",
			KM:    50000,
			LogEntries: []*entities.LogEntry{
				{ID: 46, CarID: 10, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de óleo", Date: time.Date(2023, time.March, 10, 0, 0, 0, 0, time.UTC), KM: 10000, Cost: 26000},
				{ID: 47, CarID: 10, Type: entities.LogEntryCategoryMaintenance, Name: "Filtro de ar", Date: time.Date(2023, time.August, 10, 0, 0, 0, 0, time.UTC), KM: 20000, Cost: 11000},
				{ID: 48, CarID: 10, Type: entities.LogEntryCategoryMaintenance, Name: "Alinhamento", Date: time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC), KM: 30000, Cost: 14000},
				{ID: 49, CarID: 10, Type: entities.LogEntryCategoryMaintenance, Name: "Bieletas", Date: time.Date(2024, time.June, 10, 0, 0, 0, 0, time.UTC), KM: 40000, Cost: 25000},
				{ID: 50, CarID: 10, Type: entities.LogEntryCategoryMaintenance, Name: "Troca de pneus", Date: time.Date(2025, time.January, 10, 0, 0, 0, 0, time.UTC), KM: 50000, Cost: 140000},
			},
		},
	}
}

func (p *carPersistence) ListCars(_ context.Context) ([]*entities.Car, error) {
	return p.cars, nil
}

func (p *carPersistence) GetCarByID(_ context.Context, id int) (*entities.Car, error) {
	for _, c := range p.cars {
		if c.ID == id {
			return c, nil
		}
	}
	return nil, ErrCarNotFound
}

func (p *carPersistence) CreateCar(_ context.Context, car *entities.Car) error {
	if car == nil {
		return ErrNilCar
	}
	car.ID = p.nextCarID
	p.nextCarID++
	p.cars = append(p.cars, car)
	return nil
}

func (p *carPersistence) AddLogEntry(_ context.Context, carID int, logEntry *entities.LogEntry) (*entities.LogEntry, error) {
	if logEntry == nil {
		return nil, ErrNilLogEntry
	}

	var car *entities.Car
	for _, c := range p.cars {
		if c.ID == carID {
			car = c
			break
		}
	}
	if car == nil {
		return nil, ErrCarNotFound
	}

	logEntry.ID = p.nextLogEntryID
	logEntry.CarID = carID
	p.nextLogEntryID++

	car.LogEntries = append(car.LogEntries, logEntry)

	return logEntry, nil
}

func (p *carPersistence) EditLogEntry(_ context.Context, carID int, logEntryID int, patch *entities.LogEntry) (*entities.LogEntry, error) {
	if patch == nil {
		return nil, ErrNilLogEntry
	}

	var car *entities.Car
	for _, c := range p.cars {
		if c.ID == carID {
			car = c
			break
		}
	}
	if car == nil {
		return nil, ErrCarNotFound
	}

	for _, m := range car.LogEntries {
		if m.ID != logEntryID {
			continue
		}
		m.Name = patch.Name
		m.Type = patch.Type
		m.Date = patch.Date
		m.KM = patch.KM
		m.Cost = patch.Cost
		m.CarID = carID
		return m, nil
	}
	return nil, ErrLogEntryNotFound
}

func (p *carPersistence) EditCar(_ context.Context, patch *entities.Car) (*entities.Car, error) {
	if patch == nil {
		return nil, ErrNilCar
	}

	var car *entities.Car
	for _, c := range p.cars {
		if c.ID == patch.ID {
			car = c
			break
		}
	}
	if car == nil {
		return nil, ErrCarNotFound
	}

	car.Title = patch.Title
	car.Type = patch.Type
	car.KM = patch.KM

	return car, nil
}

func (p *carPersistence) DeleteCar(_ context.Context, id int) (*entities.Car, error) {
	var deletedCar *entities.Car
	for i, c := range p.cars {
		if c.ID == id {
			deletedCar = p.cars[i]
			p.cars = append(p.cars[:i], p.cars[i+1:]...)
			return deletedCar, nil
		}
	}
	return nil, ErrCarNotFound
}

func (p *carPersistence) DeleteLogEntry(_ context.Context, carID int, logEntryID int) (*entities.LogEntry, error) {
	var car *entities.Car
	for _, c := range p.cars {
		if c.ID == carID {
			car = c
			break
		}
	}
	if car == nil {
		return nil, ErrCarNotFound
	}

	var deletedLogEntry *entities.LogEntry
	for i, m := range car.LogEntries {
		if m.ID == logEntryID {
			deletedLogEntry = car.LogEntries[i]
			car.LogEntries = append(car.LogEntries[:i], car.LogEntries[i+1:]...)
			return deletedLogEntry, nil
		}
	}
	return nil, ErrLogEntryNotFound
}
