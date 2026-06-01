package persistence

import (
	"context"

	"github.com/v2code/autolog/internal/persistence/entities"
)

type VehiclePersistence interface {
	CreateVehicle(ctx context.Context, vehicle *entities.Vehicle) (*entities.Vehicle, error)
	ListVehicles(ctx context.Context) ([]*entities.Vehicle, error)
	GetVehicleByID(ctx context.Context, id int) (*entities.Vehicle, error)
	AddLogEntry(ctx context.Context, vehicleID int, logEntry *entities.LogEntry) (*entities.LogEntry, error)
	EditLogEntry(ctx context.Context, vehicleID int, logEntryID int, logEntry *entities.LogEntry) (*entities.LogEntry, error)
	EditVehicle(ctx context.Context, vehicle *entities.Vehicle) (*entities.Vehicle, error)
	DeleteVehicle(ctx context.Context, id int) (*entities.Vehicle, error)
	DeleteLogEntry(ctx context.Context, vehicleID int, logEntryID int) (*entities.LogEntry, error)
}

type vehiclePersistence struct {
	database Database
}

func NewVehiclePersistence(db Database) VehiclePersistence {
	return &vehiclePersistence{database: db}
}

func (vp *vehiclePersistence) CreateVehicle(ctx context.Context, vehicle *entities.Vehicle) (*entities.Vehicle, error) {

	res, err := vp.database.Executor(ctx).ExecContext(ctx,
		"INSERT INTO vehicles (title, type, km) VALUES ($1, $2, $3)",
		vehicle.Title, vehicle.Type, vehicle.KM,
	)

	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &entities.Vehicle{ID: int(id), Title: vehicle.Title, Type: vehicle.Type, KM: vehicle.KM}, nil
}

func (v *vehiclePersistence) AddLogEntry(ctx context.Context, vehicleID int, logEntry *entities.LogEntry) (*entities.LogEntry, error) {
	panic("unimplemented")
}

func (v *vehiclePersistence) DeleteLogEntry(ctx context.Context, vehicleID int, logEntryID int) (*entities.LogEntry, error) {
	panic("unimplemented")
}

func (v *vehiclePersistence) DeleteVehicle(ctx context.Context, id int) (*entities.Vehicle, error) {
	panic("unimplemented")
}

func (v *vehiclePersistence) EditLogEntry(ctx context.Context, vehicleID int, logEntryID int, logEntry *entities.LogEntry) (*entities.LogEntry, error) {
	panic("unimplemented")
}

func (v *vehiclePersistence) EditVehicle(ctx context.Context, vehicle *entities.Vehicle) (*entities.Vehicle, error) {
	panic("unimplemented")
}

func (v *vehiclePersistence) GetVehicleByID(ctx context.Context, id int) (*entities.Vehicle, error) {
	panic("unimplemented")
}

func (v *vehiclePersistence) ListVehicles(ctx context.Context) ([]*entities.Vehicle, error) {
	panic("unimplemented")
}
