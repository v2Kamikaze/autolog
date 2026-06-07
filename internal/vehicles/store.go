package vehicles

import (
	"context"
	"log"

	"github.com/v2code/autolog/internal/database"
)

type Store struct {
	database database.Database
}

func NewStore(db database.Database) *Store {
	return &Store{database: db}
}

func (s *Store) CreateVehicle(ctx context.Context, vehicle *Vehicle) (*Vehicle, error) {
	row := s.database.Executor(ctx).QueryRowContext(ctx,
		"INSERT INTO vehicles(brand, model, year, version, engine, transmission, fuel, km) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		vehicle.Brand,
		vehicle.Model,
		vehicle.Year,
		vehicle.Version,
		vehicle.Engine,
		vehicle.Transmission,
		vehicle.Fuel,
		vehicle.KM,
	)

	var id int64

	err := row.Scan(&id)

	if err != nil {
		log.Println("Error creating vehicle:", err)
		return nil, err
	}

	return &Vehicle{
		ID:           int(id),
		Brand:        vehicle.Brand,
		Model:        vehicle.Model,
		Year:         vehicle.Year,
		Version:      vehicle.Version,
		Engine:       vehicle.Engine,
		Transmission: vehicle.Transmission,
		Fuel:         vehicle.Fuel,
		KM:           vehicle.KM,
		LogEntries:   vehicle.LogEntries,
	}, nil
}

func (s *Store) ListVehicles(ctx context.Context, page int, size int) ([]*Vehicle, error) {
	vehicles, err := s.getVehicles(ctx, page, size)
	if err != nil {
		log.Println("Error listing vehicles:", err)
		return nil, err
	}

	ids := make([]int, len(vehicles))

	for _, vehicle := range vehicles {
		ids = append(ids, vehicle.ID)
	}

	logEntries, err := s.getLogEntries(ctx, ids...)
	if err != nil {
		log.Println("Error listing log entries:", err)
		return nil, err
	}

	logMap := make(map[int][]*LogEntry)

	for _, logEntry := range logEntries {
		logMap[logEntry.VehicleID] = append(logMap[logEntry.VehicleID], logEntry)
	}

	for i := range vehicles {
		vehicles[i].LogEntries = logMap[vehicles[i].ID]
	}

	return vehicles, nil
}

func (s *Store) getVehicles(ctx context.Context, page int, size int) ([]*Vehicle, error) {
	rows, err := s.database.Executor(ctx).QueryContext(ctx,
		"SELECT id, brand, model, year, version, engine, transmission, fuel, km FROM vehicles LIMIT $1 OFFSET $2",
		size,
		(page-1)*size,
	)
	if err != nil {
		log.Println("Error getting vehicles:", err)
		return []*Vehicle{}, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("Error closing rows: %v", err)
		}
	}()

	vehicles := make([]*Vehicle, 0)

	for rows.Next() {
		var vehicle Vehicle

		if err := rows.Scan(
			&vehicle.ID,
			&vehicle.Brand,
			&vehicle.Model,
			&vehicle.Year,
			&vehicle.Version,
			&vehicle.Engine,
			&vehicle.Transmission,
			&vehicle.Fuel,
			&vehicle.KM,
		); err != nil {
			log.Println("Error scanning row:", err)
			return []*Vehicle{}, err
		}

		vehicles = append(vehicles, &vehicle)
	}

	return vehicles, nil
}

func (s *Store) getLogEntries(ctx context.Context, ids ...int) ([]*LogEntry, error) {
	rows, err := s.database.Executor(ctx).QueryContext(ctx,
		"SELECT id, vehicle_id, name, notes, type, date, km, cost FROM log_entries WHERE vehicle_id = ANY($1)",
		ids,
	)

	if err != nil {
		log.Println("Error getting log entries:", err)
		return []*LogEntry{}, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("Error closing rows: %v", err)
		}
	}()

	logEntries := make([]*LogEntry, 0)

	for rows.Next() {
		var logEntry LogEntry

		if err := rows.Scan(
			&logEntry.ID,
			&logEntry.VehicleID,
			&logEntry.Name,
			&logEntry.Notes,
			&logEntry.Type,
			&logEntry.Date,
			&logEntry.KM,
			&logEntry.Cost,
		); err != nil {
			log.Println("Error scanning row:", err)
			return []*LogEntry{}, err
		}

		logEntries = append(logEntries, &logEntry)
	}

	return logEntries, nil
}

func (s *Store) GetVehicleByID(ctx context.Context, id int) (*Vehicle, error) {
	row := s.database.Executor(ctx).QueryRowContext(ctx,
		"SELECT id, brand, model, year, version, engine, transmission, fuel, km FROM vehicles WHERE id = $1",
		id,
	)

	var vehicle Vehicle

	if err := row.Scan(
		&vehicle.ID,
		&vehicle.Brand,
		&vehicle.Model,
		&vehicle.Year,
		&vehicle.Version,
		&vehicle.Engine,
		&vehicle.Transmission,
		&vehicle.Fuel,
		&vehicle.KM,
	); err != nil {
		log.Println("Error getting vehicle:", err)
		return nil, err
	}

	logEntries, err := s.getLogEntries(ctx, vehicle.ID)
	if err != nil {
		log.Println("Error getting log entries:", err)
		return nil, err
	}

	vehicle.LogEntries = logEntries

	return &vehicle, nil
}

func (s *Store) AddLogEntry(ctx context.Context, logEntry *LogEntry) (*LogEntry, error) {
	row := s.database.Executor(ctx).QueryRowContext(ctx,
		"INSERT INTO log_entries(vehicle_id, name, notes, type, date, km, cost) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		logEntry.VehicleID,
		logEntry.Name,
		logEntry.Notes,
		logEntry.Type,
		logEntry.Date,
		logEntry.KM,
		logEntry.Cost,
	)

	var id int64
	if err := row.Scan(&id); err != nil {
		log.Println("Error scanning row:", err)
		return nil, err
	}

	return &LogEntry{
		ID:        int(id),
		VehicleID: logEntry.VehicleID,
		Name:      logEntry.Name,
		Type:      logEntry.Type,
		Date:      logEntry.Date,
		KM:        logEntry.KM,
		Cost:      logEntry.Cost,
		Notes:     logEntry.Notes,
	}, nil
}

func (s *Store) EditLogEntry(ctx context.Context, patch *LogEntry) (*LogEntry, error) {
	_, err := s.database.Executor(ctx).ExecContext(ctx,
		"UPDATE log_entries SET name = $1, notes = $2, type = $3, date = $4, km = $5, cost = $6 WHERE id = $7 AND vehicle_id = $8",
		patch.Name,
		patch.Notes,
		patch.Type,
		patch.Date,
		patch.KM,
		patch.Cost,
		patch.ID,
		patch.VehicleID,
	)

	if err != nil {
		log.Println("Error updating log entry:", err)
		return nil, err
	}

	return patch, nil
}

func (s *Store) EditVehicle(ctx context.Context, patch *Vehicle) (*Vehicle, error) {
	_, err := s.database.Executor(ctx).ExecContext(ctx,
		"UPDATE vehicles SET brand = $1, model = $2, year = $3, version = $4, engine = $5, transmission = $6, fuel = $7, km = $8 WHERE id = $9",
		patch.Brand,
		patch.Model,
		patch.Year,
		patch.Version,
		patch.Engine,
		patch.Transmission,
		patch.Fuel,
		patch.KM,
		patch.ID,
	)

	if err != nil {
		log.Println("Error updating vehicle:", err)
		return nil, err
	}

	logEntries, err := s.getLogEntries(ctx, patch.ID)
	if err != nil {
		log.Println("Error getting log entries:", err)
		return nil, err
	}

	patch.LogEntries = logEntries

	return patch, nil
}

func (s *Store) DeleteVehicle(ctx context.Context, id int) (*Vehicle, error) {
	row := s.database.Executor(ctx).QueryRowContext(ctx,
		"DELETE FROM vehicles WHERE id = $1 RETURNING id, brand, model, year, version, engine, transmission, fuel, km",
		id,
	)

	var vehicle Vehicle

	if err := row.Scan(
		&vehicle.ID,
		&vehicle.Brand,
		&vehicle.Model,
		&vehicle.Year,
		&vehicle.Version,
		&vehicle.Engine,
		&vehicle.Transmission,
		&vehicle.Fuel,
		&vehicle.KM,
	); err != nil {
		log.Println("Error scanning row:", err)
		return nil, err
	}

	return &vehicle, nil
}

func (s *Store) DeleteLogEntry(ctx context.Context, vehicleID int, logEntryID int) (*LogEntry, error) {
	row := s.database.Executor(ctx).QueryRowContext(ctx,
		"DELETE FROM log_entries WHERE vehicle_id = $1 AND id = $2 RETURNING id, vehicle_id, name, notes, type, date, km, cost",
		vehicleID,
		logEntryID,
	)

	var logEntry LogEntry

	if err := row.Scan(
		&logEntry.ID,
		&logEntry.VehicleID,
		&logEntry.Name,
		&logEntry.Notes,
		&logEntry.Type,
		&logEntry.Date,
		&logEntry.KM,
		&logEntry.Cost,
	); err != nil {
		log.Println("Error scanning row:", err)
		return nil, err
	}

	return &logEntry, nil
}
