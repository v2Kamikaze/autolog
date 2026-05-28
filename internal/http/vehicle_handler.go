package http

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/schema"
	"github.com/v2code/autolog/internal/persistence"
	"github.com/v2code/autolog/internal/persistence/entities"
	"github.com/v2code/autolog/internal/ui"
	"github.com/v2code/autolog/internal/vehicles"
)

type VehicleHandler struct {
	vehicles    *vehicles.Service
	ui          *ui.Engine
	formDecoder *schema.Decoder
}

func NewVehicleHandler(vehiclesService *vehicles.Service, engine *ui.Engine) *VehicleHandler {
	formDecoder := schema.NewDecoder()
	formDecoder.IgnoreUnknownKeys(true)

	return &VehicleHandler{
		vehicles:    vehiclesService,
		ui:          engine,
		formDecoder: formDecoder,
	}
}

func (h *VehicleHandler) Home(w http.ResponseWriter, r *http.Request) {
	list, err := h.vehicles.List(r.Context())
	if err != nil {
		http.Error(w, "failed to list vehicles", http.StatusInternalServerError)
		return
	}

	if err := h.ui.WriteHTML(w, vehicles.Home(list)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *VehicleHandler) CreateVehicle(w http.ResponseWriter, r *http.Request) {
	form, err := decodeForm[CreateVehicleForm](h, r)
	if err != nil {
		http.Error(w, "invalid form data", http.StatusBadRequest)
		return
	}

	vehicle, err := h.vehicles.Create(r.Context(), vehicles.CreateVehicleInput{
		Title: form.Title,
		Type:  form.VehicleType,
		KM:    form.KM,
	})
	if err != nil {
		http.Error(w, "failed to create vehicle", http.StatusInternalServerError)
		return
	}

	if err := h.ui.WriteTurbo(w, vehicles.VehicleCreated(vehicle)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *VehicleHandler) AddLogEntry(w http.ResponseWriter, r *http.Request) {
	vehicleID, err := pathID(r, "id", "vehicle")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	form, err := decodeForm[AddLogEntryForm](h, r)
	if err != nil {
		http.Error(w, "invalid form data", http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", form.Date)
	if err != nil {
		http.Error(w, "invalid date", http.StatusBadRequest)
		return
	}

	vehicle, entry, err := h.vehicles.AddLog(r.Context(), vehicles.AddLogInput{
		VehicleID: vehicleID,
		Name:      form.Name,
		Type:      entities.LogEntryCategory(form.EntryType),
		Date:      date,
		KM:        form.KM,
		Cost:      form.Cost,
	})
	if err != nil {
		writePersistenceError(w, err)
		return
	}

	if err := h.ui.WriteTurbo(w, vehicles.LogEntryAdded(vehicle, entry)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *VehicleHandler) EditLogEntry(w http.ResponseWriter, r *http.Request) {
	vehicleID, err := pathID(r, "vehicleId", "vehicle")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logEntryID, err := pathID(r, "logEntryId", "log entry")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	form, err := decodeForm[EditLogEntryForm](h, r)
	if err != nil {
		http.Error(w, "invalid form data", http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", form.Date)
	if err != nil {
		http.Error(w, "invalid date", http.StatusBadRequest)
		return
	}

	vehicle, entry, err := h.vehicles.EditLog(r.Context(), vehicleID, logEntryID, vehicles.EditLogInput{
		Name: form.Name,
		Type: entities.LogEntryCategory(form.EntryType),
		Date: date,
		KM:   form.KM,
		Cost: form.Cost,
	})
	if err != nil {
		writePersistenceError(w, err)
		return
	}

	if err := h.ui.WriteTurbo(w, vehicles.LogEntryUpdated(vehicle, entry)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *VehicleHandler) EditVehicle(w http.ResponseWriter, r *http.Request) {
	vehicleID, err := pathID(r, "id", "vehicle")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	form, err := decodeForm[EditVehicleForm](h, r)
	if err != nil {
		http.Error(w, "invalid form data", http.StatusBadRequest)
		return
	}

	vehicle, err := h.vehicles.Edit(r.Context(), vehicleID, vehicles.EditVehicleInput{
		Title: form.Title,
		Type:  form.VehicleType,
		KM:    form.KM,
	})
	if err != nil {
		if errors.Is(err, persistence.ErrVehicleNotFound) {
			http.Error(w, "vehicle not found", http.StatusNotFound)
			return
		}
		http.Error(w, "failed to edit vehicle", http.StatusInternalServerError)
		return
	}

	if err := h.ui.WriteTurbo(w, vehicles.VehicleUpdated(vehicle)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *VehicleHandler) DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	vehicleID, err := pathID(r, "id", "vehicle")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vehicle, err := h.vehicles.Delete(r.Context(), vehicleID)
	if err != nil {
		http.Error(w, "failed to delete vehicle", http.StatusInternalServerError)
		return
	}

	if err := h.ui.WriteTurbo(w, vehicles.VehicleDeleted(vehicle)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *VehicleHandler) DeleteLogEntry(w http.ResponseWriter, r *http.Request) {
	vehicleID, err := pathID(r, "vehicleId", "vehicle")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logEntryID, err := pathID(r, "logEntryId", "log entry")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vehicle, entry, err := h.vehicles.DeleteLog(r.Context(), vehicleID, logEntryID)
	if err != nil {
		writePersistenceError(w, err)
		return
	}

	if err := h.ui.WriteTurbo(w, vehicles.LogEntryDeleted(vehicle, entry)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func decodeForm[T any](h *VehicleHandler, r *http.Request) (T, error) {
	var form T
	if err := r.ParseForm(); err != nil {
		return form, err
	}
	if err := h.formDecoder.Decode(&form, r.PostForm); err != nil {
		return form, err
	}
	return form, nil
}

func pathID(r *http.Request, key string, name string) (int, error) {
	id, err := strconv.Atoi(r.PathValue(key))
	if err != nil {
		return 0, errors.New("invalid " + name + " id")
	}
	return id, nil
}

func writePersistenceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, persistence.ErrVehicleNotFound):
		http.Error(w, "vehicle not found", http.StatusNotFound)
	case errors.Is(err, persistence.ErrLogEntryNotFound):
		http.Error(w, "log entry not found", http.StatusNotFound)
	case errors.Is(err, persistence.ErrInvalidLogEntryType):
		http.Error(w, "invalid entry type", http.StatusBadRequest)
	default:
		http.Error(w, "request failed", http.StatusInternalServerError)
	}
}
