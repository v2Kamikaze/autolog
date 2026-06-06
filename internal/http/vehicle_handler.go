package http

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/schema"
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

	if err := h.ui.WriteHTML(w, vehicles.HomeView(list)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *VehicleHandler) CreateVehicle(w http.ResponseWriter, r *http.Request) {
	form, err := ParseForm[CreateVehicleForm](h, r)
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

	if err := h.ui.WriteTurbo(w, vehicles.VehicleCreatedView(vehicle)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *VehicleHandler) AddLogEntry(w http.ResponseWriter, r *http.Request) {
	vehicleID, err := pathID(r, "id", "vehicle")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	form, err := ParseForm[AddLogEntryForm](h, r)
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
		Type:      vehicles.LogEntryCategory(form.EntryType),
		Date:      date,
		KM:        form.KM,
		Cost:      form.Cost,
	})
	if err != nil {
		http.Error(w, "failed to add log entry", http.StatusInternalServerError)
		return
	}

	if err := h.ui.WriteTurbo(w, vehicles.LogEntryAddedView(vehicle, entry)); err != nil {
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

	form, err := ParseForm[EditLogEntryForm](h, r)
	if err != nil {
		http.Error(w, "invalid form data", http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", form.Date)
	if err != nil {
		http.Error(w, "invalid date", http.StatusBadRequest)
		return
	}

	vehicle, entry, err := h.vehicles.EditLog(r.Context(), vehicles.EditLogInput{
		ID:        logEntryID,
		VehicleID: vehicleID,
		Name:      form.Name,
		Type:      vehicles.LogEntryCategory(form.EntryType),
		Date:      date,
		KM:        form.KM,
		Cost:      form.Cost,
	})
	if err != nil {
		http.Error(w, "failed to edit log entry", http.StatusInternalServerError)
		return
	}

	if err := h.ui.WriteTurbo(w, vehicles.LogEntryUpdatedView(vehicle, entry)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *VehicleHandler) EditVehicle(w http.ResponseWriter, r *http.Request) {
	vehicleID, err := pathID(r, "id", "vehicle")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	form, err := ParseForm[EditVehicleForm](h, r)
	if err != nil {
		http.Error(w, "invalid form data", http.StatusBadRequest)
		return
	}

	vehicle, err := h.vehicles.Edit(r.Context(), vehicles.EditVehicleInput{
		ID:    vehicleID,
		Title: form.Title,
		Type:  form.VehicleType,
		KM:    form.KM,
	})
	if err != nil {
		http.Error(w, "failed to edit vehicle", http.StatusInternalServerError)
		return
	}

	if err := h.ui.WriteTurbo(w, vehicles.VehicleUpdatedView(vehicle)); err != nil {
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

	if err := h.ui.WriteTurbo(w, vehicles.VehicleDeletedView(vehicle)); err != nil {
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
		http.Error(w, "failed to delete log entry", http.StatusInternalServerError)
		return
	}

	if err := h.ui.WriteTurbo(w, vehicles.LogEntryDeletedView(vehicle, entry)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func pathID(r *http.Request, key string, name string) (int, error) {
	id, err := strconv.Atoi(r.PathValue(key))
	if err != nil {
		return 0, errors.New("invalid " + name + " id")
	}
	return id, nil
}
